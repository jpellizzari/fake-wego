package server

import (
	"encoding/json"
	"net/http"

	"github.com/jpellizzari/fake-wego/pkg/models"
	"github.com/jpellizzari/fake-wego/pkg/services/application"
	"github.com/jpellizzari/fake-wego/pkg/services/commit"
)

type newAppRequest struct {
	Name             string
	SourceRepoURL    string
	SourceRepoBranch string
	SourceRepoPath   string
	ConfigRepoURL    string
	AutoMerge        bool
}

func GetApp(gs application.Getter) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")

		app, err := gs.Get(name)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		b, err := json.Marshal(app)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Write(b)
	})
}

func ListCommits(getSvc application.Getter, cs commit.Lister) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		token := r.Header.Get("Authorization")

		c, err := cs.List(name, token)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		b, err := json.Marshal(c)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Write(b)
	})
}

func AddApp(as application.Adder) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		req := &newAppRequest{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		a := models.Application{
			Name:          req.Name,
			SourceURL:     req.SourceRepoURL,
			ConfigRepoURL: req.ConfigRepoURL,
		}

		if err := a.Validate(); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		params := application.AddParams{
			Token:     token,
			AutoMerge: req.AutoMerge,
		}

		if err := as.Add(a, models.DetectDefaultCluster(), params); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	})
}
