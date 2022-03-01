package main

import (
	"fmt"
	"github.com/gabe565/matrimony/internal/config"
	"github.com/gabe565/matrimony/internal/database"
	"github.com/gabe565/matrimony/internal/datadir"
	"github.com/gabe565/matrimony/internal/server"
	flag "github.com/spf13/pflag"
	"io/fs"
	"log"
	"net/http"
	"os"
	"strings"
)

const EnvPrefix = "MATRIMONY_"

func main() {
	var err error

	address := flag.String("address", ":3000", "Override listen address.")
	frontendDir := flag.String("frontend", defaultFrontend, "Override frontend asset directory."+frontendHelpExt)
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

	var frontendFs fs.FS
	if *frontendDir != "" {
		frontendFs = os.DirFS(*frontendDir)
	} else {
		frontendFs, err = fs.Sub(embeddedFrontend, "frontend/dist")
		if err != nil {
			panic(err)
		}
	}
	router := server.Router(db, frontendFs, datadir.PublicFS())
	log.Println("Listening on " + *address)
	err = http.ListenAndServe(*address, router)
	if err != nil {
		panic(err)
	}
	fmt.Println("Exiting")
}
