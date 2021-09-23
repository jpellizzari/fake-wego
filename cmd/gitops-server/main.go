package main

import (
	"net/http"

	"github.com/jpellizzari/fake-wego/pkg/add"
	"github.com/jpellizzari/fake-wego/pkg/cluster"
	"github.com/jpellizzari/fake-wego/pkg/commits"
	"github.com/jpellizzari/fake-wego/pkg/get"
	"github.com/jpellizzari/fake-wego/pkg/gitrepo"
	"github.com/jpellizzari/fake-wego/pkg/pullrequest"
	"github.com/jpellizzari/fake-wego/pkg/server"
)

func main() {
	mux := http.NewServeMux()
	gs := gitrepo.NewService()
	prs := pullrequest.NewPullRequestService()
	cs := cluster.NewClusterService()
	as := add.NewAddService(gs, prs, cs)
	getSvc := get.NewGetService(cs)
	commitsSvc := commits.NewService()

	mux.Handle("/applications/new", server.AddApp(as))
	mux.Handle("/applications", server.ListApp(cs))
	mux.Handle("/applications/commits", server.ListCommits(getSvc, commitsSvc))
}
