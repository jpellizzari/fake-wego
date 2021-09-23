package get

import (
	"github.com/jpellizzari/fake-wego/pkg/application"
	"github.com/jpellizzari/fake-wego/pkg/cluster"
)

type Service interface {
	Get(name string) (application.Application, error)
}

func NewGetService(cs cluster.ClusterService) Service {
	return svc{cs: cs}
}

type svc struct {
	cs cluster.ClusterService
}

func (s svc) Get(name string) (application.Application, error) {
	return application.Application{}, nil
}
