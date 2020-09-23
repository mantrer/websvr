// main
package main

import (
	"fmt"
	"net/http"
)

const root string = "/site/sunsea.ru/public"

func main() {

	handler := http.FileServer(http.Dir(root))

	http.Handle("/", handler)

	fmt.Println("starting server at :8080")
	http.ListenAndServe(":8080", nil)
}
