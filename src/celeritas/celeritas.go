package celeritas

import (
	"log"
	"os"
	"strconv"

	"github.com/AmKreta/laravel_clone/src/helper"
	"github.com/joho/godotenv"
)

const version = "1.0.0"

func (c *Celeritas) New(rootPath string) error {

	// initializing path config
	pathConfig := initPaths{
		rootPath:    rootPath,
		folderNames: []string{"handlers", "migrations", "views", "data", "public", "tmp", "logs", "middlewares"},
	}
	err := c.Init(pathConfig)
	if err != nil {
		return err
	}

	// loading env file
	err = godotenv.Load(rootPath + "/.env")
	if err != nil {
		return err
	}

	// creaing loggers
	c.startLoggers()
	return nil
}

func (c *Celeritas) Init(p initPaths) error {
	root := p.rootPath

	// creatinh env file
	err := helper.Create_dontenv(root, c.AppName)
	if err != nil {
		return err
	}

	// creating folder structure
	for _, path := range p.folderNames {
		err := helper.CreateFolder(root + "/" + path)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Celeritas) startLoggers() {
	c.InfoLog = log.New(os.Stdout, "Info:\t", log.Ldate|log.Ltime)
	c.ErrorLog = log.New(os.Stdout, "Error\t", log.Ldate|log.Ltime|log.Lshortfile)
	c.Debug, _ = strconv.ParseBool(os.Getenv("DEBUG"))
	c.Version = version
}
