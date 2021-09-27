package application

import (
	"context"

	"github.com/jpellizzari/fake-wego/pkg/models"
	wego "github.com/weaveworks/weave-gitops/api/v1alpha1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type Getter interface {
	Get(name string) (models.Application, error)
}

func NewGetter(k client.Client) Getter {
	return svc{k: k}
}

type svc struct {
	k client.Client
}

func (s svc) Get(name string) (models.Application, error) {
	k8sApp := &wego.Application{}

	if err := s.k.Get(context.Background(), types.NamespacedName{Name: name}, k8sApp); err != nil {
		return models.Application{}, nil
	}

	app := models.Application{Name: k8sApp.Name}

	if err := app.Validate(); err != nil {
		return models.Application{}, nil
	}

	return app, nil
}
