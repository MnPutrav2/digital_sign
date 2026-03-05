package main

import (
	"fmt"
	"os"
	"strings"

	generate "github.com/MnPutrav2/go_architecture/cmd/_cli/_generator"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	if len(os.Args) < 2 {
		Help()
		return
	}

	switch os.Args[1] {
	case "help":
		Help()
		return
	case "make:template":
		_, name, _ := strings.Cut(os.Args[2], "=")
		_, ty, _ := strings.Cut(os.Args[3], "=")

		generate.Tempate(name, ty)
	default:
		fmt.Println("command not found, use 'make help' for see available commands")
	}

}

func Help() {
	fmt.Print(`Available commands:
		
make template name=<file_name> type=<type>		 create template, example : make template name=user type=all
make run 						 running server
make build 						 build project for production
	`)
}
