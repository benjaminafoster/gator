package main

import (
	"fmt"
	"os"
	"github.com/benjaminafoster/gator/internal/config"
)

func main() {
	
	cfg, err := config.Read()
	if err != nil {
		fmt.Println(err)
	}
	
	state := &State{
		cfg: &cfg,
	}
	
	commands := GetCommands()
	
	commandName := os.Args[1]
	commandArgs := os.Args[2:]
	
	err = commands.run(state, Command{
		name: commandName,
		args: commandArgs,
	})
	if err != nil {
		fmt.Println(err)
	}
	
}