package cluster

import "github.com/jpellizzari/fake-wego/pkg/application"

type ClusterService interface {
	ApplyApplication(c Cluster, a application.Application) error
}

func NewClusterService() ClusterService {
	return clusterService{}
}

type clusterService struct{}

func (cs clusterService) ApplyApplication(c Cluster, a application.Application) error {
	return nil
}
