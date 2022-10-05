package celeritas

import (
	"github.com/AmKreta/laravel_clone/src/helper"
	"github.com/joho/godotenv"
)

//const version = "1.0.0"

func (c *Celeritas) New(rootPath string) error {
	pathConfig := initPaths{
		rootPath:    rootPath,
		folderNames: []string{"handlers", "migrations", "views", "data", "public", "tmp", "logs", "middlewares"},
	}
	err := c.Init(pathConfig)
	if err != nil {
		return err
	}

	err = godotenv.Load(rootPath + "/.env")

	if err != nil {
		return err
	}
	return nil
}

func (c *Celeritas) Init(p initPaths) error {
	root := p.rootPath
	err := helper.Create_dontenv(root, c.AppName, c.Debug)
	if err != nil {
		return err
	}
	for _, path := range p.folderNames {
		err := helper.CreateFolder(root + "/" + path)
		if err != nil {
			return err
		}
	}
	return nil
}
