package pullrequest

import (
	"github.com/jpellizzari/fake-wego/pkg/models"
)

type Creator interface {
	Create(repo models.GitRepo, token string, branchName string) (models.PullRequest, error)
}

type Merger interface {
	Merge(pr models.PullRequest, token string) error
}

// Composition example; it may not be a good choice to split this up in such a granular way.
type Manager interface {
	Creator
	Merger
}

type mgr struct {
	creator
	merger
}

func NewManager() Manager {
	return mgr{creator{}, merger{}}
}

type creator struct {
}

func NewPullRequestCreator() Creator {
	return creator{}
}

func (prs creator) Create(repo models.GitRepo, token string, branchName string) (models.PullRequest, error) {
	return models.PullRequest{}, nil
}

type merger struct{}

func NewPullRequestMerger() Merger {
	return merger{}
}

func (m merger) Merge(pr models.PullRequest, token string) error {
	return nil
}
