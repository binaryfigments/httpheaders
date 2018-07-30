package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/binaryfigments/httpheaders"
)

func main() {
	webapp := os.Args[1]

	data := httpheaders.Get(webapp)

	json, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%s\n", json)
}
