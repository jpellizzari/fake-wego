package main

import (
	"net/http"

	"github.com/jpellizzari/fake-wego/pkg/adapters/server"
	"github.com/jpellizzari/fake-wego/pkg/services/application"
	"github.com/jpellizzari/fake-wego/pkg/services/cluster"
	"github.com/jpellizzari/fake-wego/pkg/services/commit"
	"github.com/jpellizzari/fake-wego/pkg/services/deploykey"
	"github.com/jpellizzari/fake-wego/pkg/services/gitrepo"
	"github.com/jpellizzari/fake-wego/pkg/services/pullrequest"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

func main() {
	mux := http.NewServeMux()
	gs := gitrepo.NewService()
	prs := pullrequest.NewManager()
	cs := cluster.NewApplier()
	k := fake.NewFakeClient()
	dks := deploykey.NewService(k)
	as := application.NewAdder(gs, prs, cs, dks)
	getSvc := application.NewGetter(k)
	commitsSvc := commit.NewService(getSvc)

	mux.Handle("/application/:name", server.GetApp(getSvc))
	mux.Handle("/applications/commits", server.ListCommits(getSvc, commitsSvc))
	mux.Handle("/applications/new", server.AddApp(as))

	if err := http.ListenAndServe("0.0.0.0:8080", mux); err != nil {
		panic(err)
	}
}
