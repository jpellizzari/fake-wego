package models

import (
	"net/url"
)

type GitRepo struct {
	URL url.URL
}

func NewGitRepoFromURL(uri string) GitRepo {
	return GitRepo{}
}
