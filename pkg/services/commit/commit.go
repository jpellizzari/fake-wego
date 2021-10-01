package commit

import (
	"context"

	"github.com/fluxcd/go-git-providers/gitprovider"

	"github.com/jpellizzari/fake-wego/pkg/models"
	"github.com/jpellizzari/fake-wego/pkg/services/application"
)

type Lister interface {
	List(appName string, token string) ([]models.Commit, error)
}

func NewService(gs application.Getter) Lister {
	return svc{
		getApp:         gs,
		providerClient: defaultProviderClient,
	}
}

type svc struct {
	getApp         application.Getter
	providerClient func(providerName string, token string) gitprovider.Client
}

func (s svc) List(appName string, token string) ([]models.Commit, error) {
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

	// If we want to change this to do a gogit clone (instead of go-git-providers),
	// its no problem because how we get the commits is abstracted away.
	repo, err := client.OrgRepositories().Get(ctx, gitprovider.OrgRepositoryRef{})
	if err != nil {
		return nil, err
	}

	com, err := repo.Commits().ListPage(ctx, a.Branch, 10, 1)
	if err != nil {
		return nil, err
	}

	l := []models.Commit{}
	for _, commit := range com {
		l = append(l, models.Commit{Hash: commit.Get().Sha})
	}

	return l, nil
}

func defaultProviderClient(providerName string, token string) gitprovider.Client {
	return nil
}

func findProviderName(s string) (string, error) {
	return "", nil
}
