package main

import (
	"log"
	"os"
	"strings"

	"github.com/cuonglm/gogi"
)

func fetch(langs []string) string {
	langStr := strings.Join(langs, ",")
	gogiClient, _ := gogi.NewHTTPClient()
	data, err := gogiClient.Create(langStr)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	return data
}
