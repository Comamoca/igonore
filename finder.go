package main

import (
	// "fmt"
	// "fmt"
	"log"
	"os"
	"strings"

	"github.com/cuonglm/gogi"
	fuzzyfinder "github.com/ktr0731/go-fuzzyfinder"
)

func finder() string {
	gogiClient, _ := gogi.NewHTTPClient()
	data, err := gogiClient.List()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	langList := strings.Split(data, ",")

	data, err = gogiClient.Create("go")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	// fmt.Println(data)

	idx, err := fuzzyfinder.Find(langList, func(i int) string {
		return langList[i]
	})

	return langList[idx]
}
