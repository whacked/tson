// example.go
package main

import (
	"fmt"

	tson "github.com/whacked/tson/lib"
)

func main() {
	j := []byte(`{"name":"gorilla"}`)

	// tson.Edit([]byte) will return []byte, error
	res, err := tson.Edit(j)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(res))
}
