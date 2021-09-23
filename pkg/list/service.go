package list

import (
	"context"

	"github.com/jpellizzari/fake-wego/pkg/application"
	wego "github.com/weaveworks/weave-gitops/api/v1alpha1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type Service interface {
	List() ([]application.Application, error)
}

type svc struct {
	c client.Client
}

func NewService(c client.Client) Service {
	return svc{c: c}
}

func (s svc) List() ([]application.Application, error) {
	l := &wego.ApplicationList{}

	if err := s.c.List(context.Background(), l); err != nil {
		return nil, err
	}

	out := []application.Application{}
	for _, a := range l.Items {
		out = append(out, application.Application{
			Name:          a.Name,
			SourceURL:     a.Spec.URL,
			ConfigRepoURL: a.Spec.ConfigURL,
		})
	}

	return out, nil
}
