package application

import (
	"github.com/jpellizzari/fake-wego/pkg/cluster"
	"github.com/jpellizzari/fake-wego/pkg/deploykey"

	"github.com/jpellizzari/fake-wego/pkg/gitrepo"
	"github.com/jpellizzari/fake-wego/pkg/models"
	"github.com/jpellizzari/fake-wego/pkg/pullrequest"
)

type Adder interface {
	Add(app models.Application, cl models.Cluster, params AddParams) error
}

func NewAdder(gs gitrepo.Service, prs pullrequest.Service, cs cluster.Applier, dks deploykey.Service) Adder {
	return addService{
		gs:  gs,
		prs: prs,
		cs:  cs,
		dks: dks,
	}
}

type addService struct {
	gs  gitrepo.Service
	prs pullrequest.Service
	cs  cluster.Applier
	dks deploykey.Service
}

type AddParams struct {
	AutoMerge bool
	Token     string
}

func (a addService) Add(app models.Application, cl models.Cluster, params AddParams) error {
	destRepo := models.NewGitRepoFromURL(app.ConfigRepoURL)

	dk, err := a.dks.Fetch(cl, app)
	if err != nil {
		return err
	}

	if err := a.gs.CommitApplication(destRepo, dk, app); err != nil {
		return err
	}

	pr, err := a.prs.CreatePullRequest(destRepo, params.Token, "main")
	if err != nil {
		return err
	}

	if params.AutoMerge {
		if err := a.prs.MergePullRequest(pr, params.Token); err != nil {
			return err
		}
	}

	if err := a.cs.ApplyApplication(cl, app); err != nil {
		return err
	}

	return nil
}
