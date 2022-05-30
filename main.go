package main

import (
	"fmt"
	"github.com/mvrdgs/bootcamp-go-dh/dia7"
	"net/http"
)

func helloHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Olá")
}

func main() {
	dia7.Tarde2()
}
