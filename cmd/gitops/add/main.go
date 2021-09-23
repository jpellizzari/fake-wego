package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/jpellizzari/fake-wego/pkg/add"
	"github.com/jpellizzari/fake-wego/pkg/application"
	"github.com/jpellizzari/fake-wego/pkg/cluster"
	"github.com/jpellizzari/fake-wego/pkg/deploykey"
	"github.com/jpellizzari/fake-wego/pkg/gitrepo"
	"github.com/jpellizzari/fake-wego/pkg/pullrequest"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

func main() {
	var name string
	var sourceUrl string
	var configRepoURL string
	params := add.AddParams{}

	flag.String(name, "name", "")
	flag.String(sourceUrl, "sourceurl", "")
	flag.Bool("automerge", params.AutoMerge, "")
	flag.String("configrepourl", configRepoURL, "")

	flag.Parse()

	app := application.Application{
		Name:          name,
		SourceURL:     sourceUrl,
		ConfigRepoURL: configRepoURL,
	}
	if err := app.Validate(); err != nil {
		panic(err.Error())
	}

	k := fake.NewFakeClient()
	gs := gitrepo.NewService()
	prs := pullrequest.NewPullRequestService()
	cs := cluster.NewClusterService()
	dks := deploykey.NewService(cs, k)
	addSvc := add.NewAddService(gs, prs, cs, dks)

	params.Token = os.Getenv("GITHUB_TOKEN")

	if err := addSvc.Add(app, params); err != nil {
		panic(err.Error())
	}

	fmt.Println("Success!!")
}
