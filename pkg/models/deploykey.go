package models

import (
	"errors"

	"github.com/go-git/go-git/v5/plumbing/transport/ssh"
)

type DeployKey struct {
	pem []byte
}

func NewDeployKey(pem []byte) DeployKey {
	return DeployKey{
		pem: pem,
	}
}

func (dk DeployKey) Validate() error {
	if dk.pem == nil {
		return errors.New("no pem!")
	}

	return nil
}

func (dk DeployKey) PublicKeys() (*ssh.PublicKeys, error) {
	return ssh.NewPublicKeys("git", dk.pem, "")
}
