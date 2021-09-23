package server

import (
	"encoding/json"
	"net/http"

	"github.com/jpellizzari/fake-wego/pkg/add"
	"github.com/jpellizzari/fake-wego/pkg/application"
	"github.com/jpellizzari/fake-wego/pkg/cluster"
	"github.com/jpellizzari/fake-wego/pkg/commits"
	"github.com/jpellizzari/fake-wego/pkg/get"
)

type newAppRequest struct {
	Name                     string
	SourceRepoURL            string
	SourceRepoBranch         string
	SourceRepoPath           string
	ConfigDestinationRepoURL string
	AutoMerge                bool
}

func AddApp(as add.AddService) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		req := &newAppRequest{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		a := application.New(req.Name, req.SourceRepoURL)

		if err := a.Validate(); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		params := add.AddParams{
			Token:                    token,
			ConfigDestinationRepoURL: req.ConfigDestinationRepoURL,
			AutoMerge:                req.AutoMerge,
		}

		if err := as.Add(a, params); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	})
}

func ListApp(cs cluster.ClusterService) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c := cluster.DetectDefaultCluster()

		apps, err := cs.ListApplications(c)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		b, err := json.Marshal(apps)
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
