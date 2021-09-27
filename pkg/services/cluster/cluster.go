package cluster

import (
	"github.com/jpellizzari/fake-wego/pkg/models"
)

type Applier interface {
	ApplyApplication(c models.Cluster, a models.Application) error
}

func NewApplier() Applier {
	return svc{}
}

type svc struct{}

func (cs svc) ApplyApplication(c models.Cluster, a models.Application) error {
	return nil
}
