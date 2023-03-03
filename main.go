package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		requestDump, err := httputil.DumpRequest(r, true)
		if err != nil {
			fmt.Fprint(w, err.Error())
		} else {
			fmt.Fprint(w, string(requestDump))
		}
	})

	log.Fatal(http.ListenAndServe(":8080", nil))

}
