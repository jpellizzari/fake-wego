package cluster

import "github.com/jpellizzari/fake-wego/pkg/application"

type Service interface {
	ApplyApplication(c Cluster, a application.Application) error
}

func NewService() Service {
	return svc{}
}

type svc struct{}

func (cs svc) ApplyApplication(c Cluster, a application.Application) error {
	return nil
}
