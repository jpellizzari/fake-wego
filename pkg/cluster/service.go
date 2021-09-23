package cluster

import "github.com/jpellizzari/fake-wego/pkg/application"

type ClusterService interface {
	ApplyApplication(c Cluster, a application.Application) error
	ListApplications(c Cluster) ([]application.Application, error)
}

func NewClusterService() ClusterService {
	return clusterService{}
}

type clusterService struct{}

func (cs clusterService) ApplyApplication(c Cluster, a application.Application) error {
	return nil
}

func (cs clusterService) ListApplications(c Cluster) ([]application.Application, error) {
	return []application.Application{}, nil
}
