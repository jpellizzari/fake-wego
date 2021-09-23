package gitrepo

import (
	"net/url"
)

type GitRepo struct {
	URL url.URL
}

func NewFromURL(uri string) GitRepo {
	return GitRepo{}
}
