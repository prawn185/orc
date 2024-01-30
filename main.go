/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"fmt"
	"llama-orc/cmd"
)

var version string

func main() {
	//version number
	fmt.Println("Llama Orc v" + version)
	cmd.Execute()
}
