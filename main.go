// echo simply returns the body of an HTTP request.
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {

	var port string

	if len(os.Args) > 2 {
		if os.Args[1] == "--help" || os.Args[1] == "-h" {
			fmt.Printf("Usage: %s [PORT]\n", os.Args[0])
			fmt.Println()
			fmt.Println("An HTTP echo server.")
			fmt.Println("Returns the request body along some request information.")
			os.Exit(0)
		}
		port = os.Args[1]
	} else {
		port = "8080"
	}

	log.Printf("Serving on port %s\n", port)

	http.HandleFunc("/", echo)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func echo(w http.ResponseWriter, r *http.Request) {

	var out string
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprint(w, err)
		return
	}
	defer r.Body.Close()

	out += fmt.Sprintf("Method: %s\n", r.Method)
	out += fmt.Sprintf("Proto: %s\n", r.Proto)
	out += fmt.Sprintf("URL: %s\n", r.URL)
	out += fmt.Sprintf("ContentLength: %d\n\n", r.ContentLength)
	out += fmt.Sprintf("Body: %s\n", string(body))

	fmt.Fprint(w, out)
}
