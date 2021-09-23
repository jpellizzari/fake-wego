package commit

import (
	"context"

	"github.com/fluxcd/go-git-providers/gitprovider"
	"github.com/jpellizzari/fake-wego/pkg/get"
)

type Service interface {
	List(appName string, token string) ([]Commit, error)
}

func NewService(gs get.Service) Service {
	return svc{
		getApp:         gs,
		providerClient: defaultProviderClient,
	}
}

type svc struct {
	getApp         get.Service
	providerClient func(providerName string, token string) gitprovider.Client
}

func defaultProviderClient(providerName string, token string) gitprovider.Client {
	return nil
}

func findProviderName(s string) (string, error) {
	return "", nil
}

func (s svc) List(appName string, token string) ([]Commit, error) {
	ctx := context.Background()

	a, err := s.getApp.Get(appName)
	if err != nil {
		return nil, err
	}

	provider, err := findProviderName(a.ConfigRepoURL)
	if err != nil {
		return nil, err
	}

	client := s.providerClient(provider, token)

	repo, err := client.OrgRepositories().Get(ctx, gitprovider.OrgRepositoryRef{})
	if err != nil {
		return nil, err
	}

	com, err := repo.Commits().ListPage(ctx, a.Branch, 10, 1)
	if err != nil {
		return nil, err
	}

	l := []Commit{}
	for _, commit := range com {
		l = append(l, Commit{Hash: commit.Get().Sha})
	}

	return l, nil
}
