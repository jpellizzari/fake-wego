package gitrepo

import (
	"bytes"
	"context"
	"io"

	gogit "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/jpellizzari/fake-wego/pkg/models"
)

type Service interface {
	CommitApplication(repo models.GitRepo, dk models.DeployKey, a models.Application) error
}

func NewService() Service {
	return gitService{}
}

type gitService struct {
}

func (gs gitService) CommitApplication(repo models.GitRepo, dk models.DeployKey, a models.Application) error {
	auth, err := dk.PublicKeys()
	if err != nil {
		return err
	}
	branchRef := plumbing.NewBranchReferenceName("")
	repoClient, err := gogit.PlainCloneContext(context.Background(), "", false, &gogit.CloneOptions{
		URL:           repo.URL.String(),
		Auth:          auth,
		RemoteName:    gogit.DefaultRemoteName,
		ReferenceName: branchRef,
		Tags:          gogit.NoTags,
	})

	if err != nil {
		return err
	}

	wt, err := repoClient.Worktree()
	if err != nil {
		return err
	}

	f, err := wt.Filesystem.Create("")
	if err != nil {
		return err
	}
	defer f.Close()

	yml := a.ManifestYaml()

	src := yml["source.yaml"]
	kust := yml["kustomization.yaml"]

	_, err = io.Copy(f, bytes.NewReader(src))
	if err != nil {
		return err
	}
	_, err = io.Copy(f, bytes.NewReader(kust))
	if err != nil {
		return err
	}

	return nil
}
