package main

import (
	"net/http"

	"github.com/jpellizzari/fake-wego/pkg/add"
	"github.com/jpellizzari/fake-wego/pkg/cluster"
	commits "github.com/jpellizzari/fake-wego/pkg/commit"
	"github.com/jpellizzari/fake-wego/pkg/deploykey"
	"github.com/jpellizzari/fake-wego/pkg/get"
	"github.com/jpellizzari/fake-wego/pkg/gitrepo"
	"github.com/jpellizzari/fake-wego/pkg/pullrequest"
	"github.com/jpellizzari/fake-wego/pkg/server"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

func main() {
	mux := http.NewServeMux()
	gs := gitrepo.NewService()
	prs := pullrequest.NewPullRequestService()
	cs := cluster.NewClusterService()
	k := fake.NewFakeClient()
	dks := deploykey.NewService(cs, k)
	as := add.NewAddService(gs, prs, cs, dks)
	getSvc := get.NewService(k)
	commitsSvc := commits.NewService()

	mux.Handle("/applications/new", server.AddApp(as))
	mux.Handle("/applications", server.GetApp(getSvc))
	mux.Handle("/applications/commits", server.ListCommits(getSvc, commitsSvc))

	if err := http.ListenAndServe("0.0.0.0:8080", mux); err != nil {
		panic(err)
	}
}
