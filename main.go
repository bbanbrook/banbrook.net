// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 20.
//!+

// Server2 is a minimal "echo" and counter server.
package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"sync"
)

var mu sync.Mutex

var count float64

func main() {
	//static file handler
	fs := http.FileServer(http.Dir("asset"))
	http.Handle("/js/", http.StripPrefix("/js/", fs))

	//http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("0.0.0.0:80", nil))
}

// handler echoes the Path component of the requested URL.
func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++

	mu.Unlock()
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

// counter echoes the number of calls so far.
func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	//count++
	count += 1
	if r.Method == "GET" {
		//fmt.Fprintf(w, "Count %d\n", count)

		t, _ := template.ParseFiles("index.gtpl")
		err := t.Execute(w, count)
		if err != nil {
			log.Fatalf("execution failed: %s", err)
		}
	}
	mu.Unlock()
}

//!-
