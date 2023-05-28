/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"github.com/rngallen/mycli/cmd"
	"github.com/rngallen/mycli/data"
)

func main() {
	// Open database
	data.ConnectDb()
	cmd.Execute()
}
