package pullrequest

import (
	"github.com/jpellizzari/fake-wego/pkg/models"
)

type Service interface {
	CreatePullRequest(repo models.GitRepo, token string, branchName string) (models.PullRequest, error)
	MergePullRequest(pr models.PullRequest, token string) error
}

type prService struct {
}

func NewPullRequestService() Service {
	return prService{}
}

func (prs prService) CreatePullRequest(repo models.GitRepo, token string, branchName string) (models.PullRequest, error) {
	return models.PullRequest{}, nil
}

func (prs prService) MergePullRequest(pr models.PullRequest, token string) error {
	return nil
}
