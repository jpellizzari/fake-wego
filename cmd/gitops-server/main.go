package main

import (
	"net/http"

	"github.com/jpellizzari/fake-wego/pkg/add"
	"github.com/jpellizzari/fake-wego/pkg/cluster"
	"github.com/jpellizzari/fake-wego/pkg/commits"
	"github.com/jpellizzari/fake-wego/pkg/deploykey"
	"github.com/jpellizzari/fake-wego/pkg/get"
	"github.com/jpellizzari/fake-wego/pkg/gitrepo"
	"github.com/jpellizzari/fake-wego/pkg/list"
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
	getSvc := get.NewGetService(cs)
	commitsSvc := commits.NewService()
	ls := list.NewService(fake.NewFakeClient())

	mux.Handle("/applications/new", server.AddApp(as))
	mux.Handle("/applications", server.ListApp(ls))
	mux.Handle("/applications/commits", server.ListCommits(getSvc, commitsSvc))

	if err := http.ListenAndServe("0.0.0.0:8080", mux); err != nil {
		panic(err)
	}
}
