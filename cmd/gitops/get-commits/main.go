package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/jpellizzari/fake-wego/pkg/services/application"
	commits "github.com/jpellizzari/fake-wego/pkg/services/commit"

	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

func main() {
	var name string

	flag.String(name, "name", "")

	flag.Parse()

	k := fake.NewFakeClient()
	get := application.NewGetter(k)
	commitsSvc := commits.NewService(get)

	token := os.Getenv("GITHUB_TOKEN")

	c, err := commitsSvc.List(name, token)
	if err != nil {
		panic(err)
	}

	for _, commit := range c {
		fmt.Println(commit.Hash)
	}
}
