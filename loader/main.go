package main

import (
	"fmt"
	"io"
	"os"
	"user-api/entities"
	"user-api/exception"

	"ariga.io/atlas-provider-gorm/gormschema"
)

func main() {
	stmts, err := gormschema.New("postgres").Load(&entities.User{})
	exception.FatalLogging(err, fmt.Sprintf("failed to load gorm schema: %v\n", err))
	io.WriteString(os.Stdout, stmts)
}
