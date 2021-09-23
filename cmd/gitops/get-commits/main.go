package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/jpellizzari/fake-wego/pkg/cluster"
	"github.com/jpellizzari/fake-wego/pkg/commits"
	"github.com/jpellizzari/fake-wego/pkg/get"
)

func main() {
	var name string

	flag.String(name, "name", "")

	flag.Parse()

	cs := cluster.NewClusterService()
	get := get.NewGetService(cs)
	commitsSvc := commits.NewService()

	a, err := get.Get(name)
	if err != nil {
		panic(err)
	}

	token := os.Getenv("GITHUB_TOKEN")

	c, err := commitsSvc.List(a, token)
	if err != nil {
		panic(err)
	}

	for _, commit := range c {
		fmt.Println(commit.Hash)
	}
}
