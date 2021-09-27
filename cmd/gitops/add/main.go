package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/jpellizzari/fake-wego/pkg/models"
	"github.com/jpellizzari/fake-wego/pkg/services/application"
	"github.com/jpellizzari/fake-wego/pkg/services/cluster"
	"github.com/jpellizzari/fake-wego/pkg/services/deploykey"
	"github.com/jpellizzari/fake-wego/pkg/services/gitrepo"
	"github.com/jpellizzari/fake-wego/pkg/services/pullrequest"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

func main() {
	var name string
	var sourceUrl string
	var configRepoURL string
	params := application.AddParams{}

	flag.String(name, "name", "")
	flag.String(sourceUrl, "sourceurl", "")
	flag.Bool("automerge", params.AutoMerge, "")
	flag.String("configrepourl", configRepoURL, "")

	flag.Parse()

	app := models.Application{
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
	cs := cluster.NewApplier()
	dks := deploykey.NewService(cs, k)
	addSvc := application.NewAdder(gs, prs, cs, dks)

	params.Token = os.Getenv("GITHUB_TOKEN")

	if err := addSvc.Add(app, models.DetectDefaultCluster(), params); err != nil {
		panic(err.Error())
	}

	fmt.Println("Success!!")
}
