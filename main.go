package main

import (
	"fmt"
	"os"
	"database/sql"

	"github.com/benjaminafoster/gator/internal/config"
	"github.com/benjaminafoster/gator/internal/database"

	_ "github.com/lib/pq"
)

func main() {
	cfg, _ := config.Read()
	db, err := sql.Open("postgres", cfg.DB_URL)
	if err != nil {
		fmt.Printf("failed to open connection to postgres server at %s: %s", cfg.DB_URL, err)
	}
	dbQueries := database.New(db)

	state := &State{Db: dbQueries, CfgPtr: &cfg}

	cmds := commands{
		registeredCommands: make(map[string]func(*State, command) error),
	}

	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handlerReset)
	cmds.register("users", handlerUsers)

	if len(os.Args) < 2 {
		fmt.Println("a minimum of 2 command line arguments must be passed with gator")
		os.Exit(1)
	}

	cmdName := os.Args[1]
	cmdArgs := os.Args[2:]


	err = cmds.run(state, command{name: cmdName, args: cmdArgs})
	if err != nil {
		fmt.Printf("error: %s\n", err)
		os.Exit(1)
	}

}