package main

import (
	"github.com/AlfredoPrograma/diployable/bootstrap"
)

func main() {
	env := bootstrap.LoadEnv()
	bootstrap.ConnectDB(env.DbConnectionString)
}
