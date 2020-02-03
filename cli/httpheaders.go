package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/binaryfigments/httpheaders"
)

func main() {
	webapp := os.Args[1]

	data := httpheaders.Get(webapp, false)
	fmt.Println(data.Headers.Get("Server"))

	json, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%s\n", json)
}
