package pullrequest

import "github.com/jpellizzari/fake-wego/pkg/gitrepo"

type Service interface {
	CreatePullRequest(repo gitrepo.GitRepo, token string, branchName string) (PullRequest, error)
	MergePullRequest(pr PullRequest, token string) error
}

type prService struct {
}

func NewPullRequestService() Service {
	return prService{}
}

func (prs prService) CreatePullRequest(repo gitrepo.GitRepo, token string, branchName string) (PullRequest, error) {
	return PullRequest{}, nil
}

func (prs prService) MergePullRequest(pr PullRequest, token string) error {
	return nil
}
