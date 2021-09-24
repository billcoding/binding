# binding

A model binding written in Golang


[![Go Report Card](https://goreportcard.com/badge/github.com/billcoding/binding)](https://goreportcard.com/report/github.com/billcoding/binding)
[![GoDoc](https://pkg.go.dev/badge/github.com/billcoding/binding?status.svg)](https://pkg.go.dev/github.com/billcoding/binding?tab=doc)

## quickstart
```go
package main

import (
	"fmt"
	"github.com/billcoding/binding"
)

func main() {
	type model struct {
		ID string `binding:"name(id) default(100) trim(T) prefix(PREFIX-) suffix(-SUFFIX)"`
	}
	m := model{}
	binding.New(&m).BindMap(map[string]interface{}{"ID": "hello world"})
	fmt.Println(m.ID)
	// outputs: PREFIX-100-SUFFIX
}
```