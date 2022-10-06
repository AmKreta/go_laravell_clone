package main

import (
	"fmt"

	"github.com/AmKreta/laravel_clone/src/celeritas"
)

type applications struct {
	App *celeritas.Celeritas
}

func main() {
	amk := initApplication("amk")

	fmt.Println("debug", amk.App.Debug)
}
