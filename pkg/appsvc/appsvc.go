package appsvc

import (
	"github.com/jpellizzari/fake-wego/pkg/add"
	"github.com/jpellizzari/fake-wego/pkg/cluster"
	"github.com/jpellizzari/fake-wego/pkg/deploykey"
	"github.com/jpellizzari/fake-wego/pkg/get"
	"github.com/jpellizzari/fake-wego/pkg/gitrepo"
	"github.com/jpellizzari/fake-wego/pkg/pullrequest"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type Service interface {
	add.AddService
	get.Service
}

type svc struct {
	add.AddService
	get.Service
}

func NewService(gs gitrepo.Service, prs pullrequest.Service, cs cluster.Service, dks deploykey.Service, k client.Client) Service {
	return svc{
		add.NewAddService(gs, prs, cs, dks),
		get.NewService(k),
	}
}
