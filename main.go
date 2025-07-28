/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"ptt/cmd"
	"ptt/config"
)

func main() {
	_ = config.InitConfig()
	cmd.Execute()
}
