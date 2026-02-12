package main

import (
	"fmt"
	"context"
	"time"
	"github.com/google/uuid"
	"github.com/benjaminafoster/gator/internal/database"
)

func handlerRegister(state *State, cmd Command) error {
	/* Usage: gator register <name> */
	if len(cmd.args) != 1 {
		return fmt.Errorf("Usage: gator register <name>\n")
	}
	
	newUser := database.CreateUserParams{
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name: cmd.args[0],
		
	}
	
	dbUser, err := state.db.CreateUser(context.Background(), newUser)
	if err != nil {
		return fmt.Errorf("failed to create user: %w\n", err)
	}
	
	err = state.cfg.SetUser(dbUser.Name)
	if err != nil {
		return fmt.Errorf("failed to set user: %w\n", err)
	}
	
	fmt.Printf("User %s registered successfully\n", dbUser.Name)
	fmt.Printf("User data:\n")
	fmt.Printf(" - ID: %s\n", dbUser.ID)
	fmt.Printf(" - Name: %s\n", dbUser.Name)
	fmt.Printf(" - Created At: %s\n", dbUser.CreatedAt)
	fmt.Printf(" - Updated At: %s\n", dbUser.UpdatedAt)
	
	return nil
}