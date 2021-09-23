package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/jpellizzari/fake-wego/pkg/add"
	"github.com/jpellizzari/fake-wego/pkg/application"
	"github.com/jpellizzari/fake-wego/pkg/cluster"
	"github.com/jpellizzari/fake-wego/pkg/gitrepo"
	"github.com/jpellizzari/fake-wego/pkg/pullrequest"
)

func main() {
	var name string
	var sourceUrl string
	params := add.AddParams{}

	flag.String(name, "name", "")
	flag.String(sourceUrl, "sourceurl", "")
	flag.Bool("automerge", params.AutoMerge, "")
	flag.String("destinationurl", params.ConfigDestinationRepoURL, "")

	flag.Parse()

	app := application.New(name, sourceUrl)
	if err := app.Validate(); err != nil {
		panic(err.Error())
	}

	gs := gitrepo.NewService()
	prs := pullrequest.NewPullRequestService()
	cs := cluster.NewClusterService()
	addSvc := add.NewAddService(gs, prs, cs)

	params.Token = os.Getenv("GITHUB_TOKEN")

	if err := addSvc.Add(app, params); err != nil {
		panic(err.Error())
	}

	fmt.Println("Success!!")
}
