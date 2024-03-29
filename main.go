package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gabe565/matrimony/internal/config"
	"github.com/gabe565/matrimony/internal/database"
	"github.com/gabe565/matrimony/internal/datadir"
	"github.com/gabe565/matrimony/internal/server"
	flag "github.com/spf13/pflag"
)

const EnvPrefix = "MATRIMONY_"

func main() {
	var err error

	address := flag.String("address", ":3000", "Override listen address.")
	frontendDir := flag.String("frontend", "frontend", "Override frontend asset directory.")
	flag.Parse()

	// Load flags from envs
	flag.CommandLine.VisitAll(func(f *flag.Flag) {
		optName := strings.ToUpper(f.Name)
		optName = strings.ReplaceAll(optName, "-", "_")
		varName := EnvPrefix + optName
		if val, ok := os.LookupEnv(varName); !f.Changed && ok {
			err = f.Value.Set(val)
			if err != nil {
				log.Fatalln(err)
			}
		}
	})

	err = config.Setup()
	if err != nil {
		panic(err)
	}

	db, err := database.Setup()
	if err != nil {
		panic(err)
	}

	frontendFs := os.DirFS(*frontendDir)
	router := server.Router(db, frontendFs, datadir.PublicFS())
	log.Println("Listening on " + *address)
	err = http.ListenAndServe(*address, router)
	if err != nil {
		panic(err)
	}
	fmt.Println("Exiting")
}
