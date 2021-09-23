package deploykey

import (
	"context"

	"github.com/jpellizzari/fake-wego/pkg/application"
	"github.com/jpellizzari/fake-wego/pkg/cluster"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type Service interface {
	Provision(c cluster.Cluster, a application.Application, token string) (DeployKey, error)
	Fetch(c cluster.Cluster, a application.Application) (DeployKey, error)
}

type svc struct {
	cs cluster.Service
	k  client.Client
}

func NewService(cs cluster.Service, k client.Client) Service {
	return svc{cs: cs, k: k}
}

func (s svc) Provision(c cluster.Cluster, a application.Application, token string) (DeployKey, error) {
	secret, err := doFluxThingsToMakeASecret(a.ConfigRepoURL)
	if err != nil {
		return DeployKey{}, err
	}

	if err := s.k.Create(context.Background(), secret); err != nil {
		return DeployKey{}, err
	}

	dk := DeployKey{pem: []byte(secret.Data["id_rsa"])}

	if err := dk.Validate(); err != nil {
		return DeployKey{}, err
	}

	return dk, nil

}

func (s svc) Fetch(c cluster.Cluster, a application.Application) (DeployKey, error) {
	secret := &corev1.Secret{}

	if err := s.k.Get(context.Background(), types.NamespacedName{Name: a.DeployKeyName(c.Name)}, secret); err != nil {
		return DeployKey{}, nil
	}

	dk := DeployKey{pem: []byte(secret.Data["id_rsa"])}
	if err := dk.Validate(); err != nil {
		return DeployKey{}, err
	}

	return dk, nil
}

func doFluxThingsToMakeASecret(name string) (*corev1.Secret, error) {
	return &corev1.Secret{}, nil
}
