package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/benjaminafoster/gator/internal/database"
	"github.com/google/uuid"
)



type command struct {
	name           string
	args           []string
}

func handlerLogin(s *State, cmd command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.name)
	}
	
	
	userName := cmd.args[0]

	
	_, err := s.Db.GetUser(context.Background(), userName)
    if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			// User does not exist, provide error. You can't log in with a user that doesn't exist
			os.Exit(1)
			//return fmt.Errorf("user does not exist")
		} else {
			// Some other error occurred
			return fmt.Errorf("error checking if user exists: %w", err)
		}
	} else {
		// User already exists
		err := s.CfgPtr.SetUser(userName)
		if err != nil {
			return fmt.Errorf("couldn't set current user: %w", err)
		}
	}
			
	fmt.Printf("user switched successfully!\n")
	
	return nil
	
}

func handlerRegister(s *State, cmd command) error {
    if len(cmd.args) != 1 {
        return fmt.Errorf("usage: %s <name>", cmd.name)
    }

    userName := cmd.args[0]

	userParams := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      userName,
	}
	_, err := s.Db.GetUser(context.Background(), userName)
    if err != nil {
        if errors.Is(err, sql.ErrNoRows) {
            // User does not exist, proceed with registration
			_, err = s.Db.CreateUser(context.Background(), userParams)
			if err != nil {
				return fmt.Errorf("error creating user in database: %w", err)
			}
        } else {
            // Some other error occurred
            return fmt.Errorf("error checking if user exists: %w", err)
        }
    } else {
        // User already exists
		os.Exit(1)
    }

	s.CfgPtr.SetUser(userName)

	fmt.Printf("user '%s' successfully registered\n",s.CfgPtr.CurrentUser)

    return nil
}

type commands struct {
	registeredCommands         map[string]func(*State, command) error
}

func (c *commands) register(name string, f func(*State, command) error) {
	c.registeredCommands[name] = f
}

func (c *commands) run(s *State, cmd command) error {
	if f, exists := c.registeredCommands[cmd.name]; exists {
		f(s, cmd)
		return nil
	}

	return fmt.Errorf("error executing command: %s", cmd.name)
}