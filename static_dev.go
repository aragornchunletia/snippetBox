//+build dev
//go:build dev
// +build dev

package main

import (
	"fmt"
	"net/http"
)

func public() http.Handler {
	fmt.Println("building static files for dev...")
	return http.StripPrefix("/public/", http.FileServer(http.Dir("public")))
}
