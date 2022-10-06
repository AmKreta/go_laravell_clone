package main

import (
	"log"
	"os"

	"github.com/AmKreta/laravel_clone/src/celeritas"
	"github.com/AmKreta/laravel_clone/src/helper"
)

func initApplication(AppName string) *applications {
	path, err := os.Getwd()

	if err != nil {
		log.Fatal(err)
	}

	new_path := path + "/app_" + AppName
	err = helper.CreateFolder(new_path)

	if err != nil {
		log.Fatal(err)
	}

	// init celeratas
	cel := &celeritas.Celeritas{
		AppName: AppName,
	}
	err = cel.New(new_path)

	if err != nil {
		log.Fatal(err)
	}

	app := &applications{
		App: cel,
	}

	return app
}
