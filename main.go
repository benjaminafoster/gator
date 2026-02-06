package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/benjaminafoster/gator/internal/config"
	"github.com/benjaminafoster/gator/internal/database"
	_ "github.com/lib/pq"
)

func main() {
	
	cfg, err := config.Read()
	if err != nil {
		fmt.Println(err)
	}
	
	dbURL := cfg.DBURL
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer db.Close()
	
	dbQueries := database.New(db)
	
	state := &State{
		cfg: &cfg,
		db: dbQueries,
	}
	
	
	commands := GetCommands()
	
	if len(os.Args) == 1 {
		fmt.Println("Usage: gator <command> [args]")
		os.Exit(1)
	}
	
	commandName := os.Args[1]
	commandArgs := os.Args[2:]
	
	err = commands.run(state, Command{
		name: commandName,
		args: commandArgs,
	})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	
}