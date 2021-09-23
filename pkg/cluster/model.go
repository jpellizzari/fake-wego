package cluster

type Cluster struct {
	Name string
}

func DetectDefaultCluster() Cluster {
	// Does some sort of environment lookup to figure out what cluster we should target.
	// Do this by looking at a local kubeconfig or something.
	return Cluster{}
}
