package gitrepo

import (
	"github.com/jpellizzari/fake-wego/pkg/application"
)

type Service interface {
	CommitApplication(repo GitRepo, branch string, a application.Application) error
}

func NewService() Service {
	return gitService{}
}

type gitService struct {
}

func (gs gitService) CommitApplication(repo GitRepo, branch string, a application.Application) error {
	return nil
}
