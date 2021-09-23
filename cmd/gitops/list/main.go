package main

import (
	"flag"
	"fmt"

	"github.com/jpellizzari/fake-wego/pkg/cluster"
)

func main() {
	var name string

	flag.String(name, "name", "")

	flag.Parse()

	cs := cluster.NewClusterService()

	apps, err := cs.ListApplications(cluster.DetectDefaultCluster())
	if err != nil {
		panic(err)
	}

	for _, app := range apps {
		fmt.Println(app.Name())
	}
}
