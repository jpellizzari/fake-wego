package main

import (
	"flag"
	"fmt"

	"github.com/jpellizzari/fake-wego/pkg/get"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

func main() {
	var name string
	flag.String("name", name, "")
	c := fake.NewFakeClient()
	ls := get.NewService(c)

	app, err := ls.Get(name)
	if err != nil {
		panic(err)
	}

	fmt.Println(app.Name)
	fmt.Println(app.SourceURL)
}
