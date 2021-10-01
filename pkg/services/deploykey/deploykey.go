package deploykey

import (
	"context"

	"github.com/jpellizzari/fake-wego/pkg/models"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type Manager interface {
	Create(c models.Cluster, a models.Application, token string) (models.DeployKey, error)
	Fetch(c models.Cluster, a models.Application) (models.DeployKey, error)
}

type svc struct {
	k client.Client
}

func NewService(k client.Client) Manager {
	return svc{k: k}
}

func (s svc) Create(c models.Cluster, a models.Application, token string) (models.DeployKey, error) {
	secret, err := doFluxThingsToMakeASecret(a.ConfigRepoURL)
	if err != nil {
		return models.DeployKey{}, err
	}

	if err := s.k.Create(context.Background(), secret); err != nil {
		return models.DeployKey{}, err
	}

	dk := models.NewDeployKey([]byte(secret.Data["id_rsa"]))

	return dk, nil

}

func (s svc) Fetch(c models.Cluster, a models.Application) (models.DeployKey, error) {
	secret := &corev1.Secret{}

	if err := s.k.Get(context.Background(), types.NamespacedName{Name: a.DeployKeyName(c.Name)}, secret); err != nil {
		return models.DeployKey{}, nil
	}

	dk := models.NewDeployKey([]byte(secret.Data["id_rsa"]))
	if err := dk.Validate(); err != nil {
		return models.DeployKey{}, err
	}

	return dk, nil
}

func doFluxThingsToMakeASecret(name string) (*corev1.Secret, error) {
	return &corev1.Secret{}, nil
}
