package main

import (
	"io"
	"io/ioutil"
	"os"
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	bs, err := ioutil.ReadFile("/content/index.html")

	if err != nil {
		fmt.Printf("Couldn't read index.html: %v", err);
		os.Exit(1)
	}

	io.WriteString(w, string(bs[:]))
}

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}
