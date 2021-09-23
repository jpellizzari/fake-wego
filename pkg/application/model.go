package application

type Application struct {
	name          string
	sourceURL     string
	configRepoURL string
	manifests     map[string][]byte
	branch        string
}

func doFluxThingsToGenerateYaml(name string, u string) ([]byte, []byte) {
	// Imagine we do some flux calls here to populate all the manifests
	return []byte("source yaml"), []byte("kustomization yaml")
}

func New(name string, source string) Application {
	sYaml, kYaml := doFluxThingsToGenerateYaml(name, source)

	return Application{
		name:      name,
		sourceURL: source,
		manifests: map[string][]byte{
			"kustomization.yaml": kYaml,
			"source.yaml":        sYaml,
		},
	}
}

func (a Application) Validate() error {
	return nil
}

func (a Application) ManifestYaml() map[string][]byte {
	return a.manifests
}

func (a Application) ConfigRepo() string {
	return a.configRepoURL
}

func (a Application) Branch() string {
	return a.branch
}

func (a Application) Name() string {
	return a.name
}
