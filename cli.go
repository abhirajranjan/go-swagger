// this file code will go under server cli commands
package main

import (
	"github.com/swaggo/swag"
	"github.com/swaggo/swag/gen"
)

func build() {
	gen.New().Build(&gen.Config{
		SearchDir:          "swagger,./",
		MainAPIFile:        "swagger.go",
		PropNamingStrategy: swag.CamelCase,
		OutputDir:          "swagger",
		OutputTypes:        []string{"json", "yaml"},
	})
}
