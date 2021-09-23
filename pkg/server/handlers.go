package server

import (
	"encoding/json"
	"net/http"

	"github.com/jpellizzari/fake-wego/pkg/add"
	"github.com/jpellizzari/fake-wego/pkg/application"
	commits "github.com/jpellizzari/fake-wego/pkg/commit"
	"github.com/jpellizzari/fake-wego/pkg/get"
)

type newAppRequest struct {
	Name             string
	SourceRepoURL    string
	SourceRepoBranch string
	SourceRepoPath   string
	ConfigRepoURL    string
	AutoMerge        bool
}

func AddApp(as add.AddService) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		req := &newAppRequest{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		a := application.Application{
			Name:          req.Name,
			SourceURL:     req.SourceRepoURL,
			ConfigRepoURL: req.ConfigRepoURL,
		}

		if err := a.Validate(); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		params := add.AddParams{
			Token:     token,
			AutoMerge: req.AutoMerge,
		}

		if err := as.Add(a, params); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	})
}

func GetApp(gs get.Service) http.Handler {
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

func ListCommits(getSvc get.Service, cs commits.Service) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		token := r.Header.Get("Authorization")

		a, err := getSvc.Get(name)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		c, err := cs.List(a, token)
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
