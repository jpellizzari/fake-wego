package main

import (
	"fmt"

	"github.com/jpellizzari/fake-wego/pkg/list"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

func main() {
	c := fake.NewFakeClient()
	ls := list.NewService(c)

	apps, err := ls.List()
	if err != nil {
		panic(err)
	}

	for _, app := range apps {
		fmt.Println(app.Name())
	}
}
