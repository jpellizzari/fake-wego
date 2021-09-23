package get

import (
	"context"

	"github.com/jpellizzari/fake-wego/pkg/application"
	wego "github.com/weaveworks/weave-gitops/api/v1alpha1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type Service interface {
	Get(name string) (application.Application, error)
}

func NewService(k client.Client) Service {
	return svc{k: k}
}

type svc struct {
	k client.Client
}

func (s svc) Get(name string) (application.Application, error) {
	k8sApp := &wego.Application{}

	if err := s.k.Get(context.Background(), types.NamespacedName{Name: name}, k8sApp); err != nil {
		return application.Application{}, nil
	}

	app := application.Application{Name: k8sApp.Name}

	if err := app.Validate(); err != nil {
		return application.Application{}, nil
	}

	return app, nil
}
