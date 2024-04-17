package main

import (
	"io"
	"log"
	"os"
	"user-api/entities"

	"ariga.io/atlas-provider-gorm/gormschema"
)

func main() {
	stmts, err := gormschema.New("postgres").Load(&entities.User{})
	if err != nil {
		log.Fatalf("Failed to load gorm schema: %v\n", err)
	}
	io.WriteString(os.Stdout, stmts)
}
