package add

import (
	"github.com/jpellizzari/fake-wego/pkg/application"
	"github.com/jpellizzari/fake-wego/pkg/cluster"
	"github.com/jpellizzari/fake-wego/pkg/deploykey"
	"github.com/jpellizzari/fake-wego/pkg/gitrepo"
	"github.com/jpellizzari/fake-wego/pkg/pullrequest"
)

type AddService interface {
	Add(app application.Application, params AddParams) error
}

func NewAddService(gs gitrepo.Service, prs pullrequest.Service, cs cluster.Service, dks deploykey.Service) AddService {
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
	cs  cluster.Service
	dks deploykey.Service
}

type AddParams struct {
	AutoMerge bool
	Token     string
}

func (a addService) Add(app application.Application, params AddParams) error {
	destRepo := gitrepo.NewFromURL(app.ConfigRepoURL)

	cl := cluster.DetectDefaultCluster()

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

	c := cluster.DetectDefaultCluster()

	if err := a.cs.ApplyApplication(c, app); err != nil {
		return err
	}

	return nil
}
