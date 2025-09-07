package main

import (
	"fmt"

	"github.com/AlfredoPrograma/diployable/bootstrap"
)

func main() {
	env := bootstrap.LoadEnv()
	fmt.Printf("%#v", env)
}
