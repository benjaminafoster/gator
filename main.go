package main

import (
	"fmt"
	"os"
	"log"
	"github.com/benjaminafoster/gator/internal/config"
)

func main() {
	cfg, _ := config.Read()
	state := &State{CfgPtr: &cfg}

	cmds := commands{
		registeredCommands: make(map[string]func(*State, command) error),
	}

	cmds.register("login", handlerLogin)

	if len(os.Args) < 3 {
		fmt.Println("a minimum of 2 command line arguments must be passed with gator")
		os.Exit(1)
	}

	cmdName := os.Args[1]
	cmdArgs := os.Args[2:]


	err := cmds.run(state, command{name: cmdName, args: cmdArgs})
	if err != nil {
		log.Fatal(err)
	}

}