package pullrequest

type ProviderName string

const ProviderNameGitHub ProviderName = "github"

type PullRequest struct {
	Name       string
	Provider   ProviderName
	BranchName string
}
