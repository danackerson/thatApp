package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/index", hello)
	http.HandleFunc("/this", thisapp)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello That!")
}

func thisapp(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello Dan!")
	response, err := http.Get("http://thisapp:8080/")
	checkError(err)
	contents := readBody(response.Body)
	fmt.Fprint(w, contents)
}

func readBody(body io.ReadCloser) []byte {
	defer body.Close()
	contents, err := ioutil.ReadAll(body)
	checkError(err)
	return contents
}
func checkError(err error) {
	if err != nil {
		fmt.Printf("%s", err)
		//os.Exit(1)
	}
}
