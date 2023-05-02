package main

import (
	"fmt"
	"net/http"
)

const (
	defaultAddress string = "0.0.0.0"
	defaultPort    int    = 8080
)

func main() {
	addressPort := fmt.Sprintf("%s:%d", defaultAddress, defaultPort)
	fmt.Printf("Lisen on %s\n", addressPort)
	http.ListenAndServe(addressPort, http.FileServer(http.Dir("./html")))

}
