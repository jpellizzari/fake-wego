package application

import (
	"github.com/jpellizzari/fake-wego/pkg/services/cluster"
	"github.com/jpellizzari/fake-wego/pkg/services/deploykey"

	"github.com/jpellizzari/fake-wego/pkg/models"
	"github.com/jpellizzari/fake-wego/pkg/services/gitrepo"
	"github.com/jpellizzari/fake-wego/pkg/services/pullrequest"
)

type Adder interface {
	Add(app models.Application, cl models.Cluster, params AddParams) error
}

func NewAdder(gs gitrepo.AppComitter, prs pullrequest.Manager, cs cluster.Applier, dks deploykey.Manager) Adder {
	return addService{
		gs:  gs,
		prs: prs,
		cs:  cs,
		dks: dks,
	}
}

type addService struct {
	gs  gitrepo.AppComitter
	prs pullrequest.Manager
	cs  cluster.Applier
	dks deploykey.Manager
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

	pr, err := a.prs.Create(destRepo, params.Token, "main")
	if err != nil {
		return err
	}

	if params.AutoMerge {
		if err := a.prs.Merge(pr, params.Token); err != nil {
			return err
		}
	}

	if err := a.cs.ApplyApplication(cl, app); err != nil {
		return err
	}

	return nil
}
