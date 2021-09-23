package application

import "fmt"

type Application struct {
	Name          string
	SourceURL     string
	ConfigRepoURL string
	Branch        string
}

func (a Application) ManifestYaml() map[string][]byte {
	appYaml, sYaml, kYaml := doFluxThingsToGenerateYaml(a.Name, a.SourceURL)
	return map[string][]byte{
		"application.yaml":   appYaml,
		"source.yaml":        sYaml,
		"kustomization.yaml": kYaml,
	}
}

func (a Application) DeployKeyName(clusterName string) string {
	return fmt.Sprintf("%s-%s-deploy-key", a.Name, clusterName)
}

func (a Application) Validate() error {
	return nil
}

func doFluxThingsToGenerateYaml(name string, u string) ([]byte, []byte, []byte) {
	// Imagine we do some flux calls here to populate all the manifests
	return []byte("application yaml"), []byte("source yaml"), []byte("kustomization yaml")
}
