package main

import (
	"fmt"
	"errors"
	"context"
)

func handlerUsers(s *State, cmd Command) error {
	if len(cmd.args) != 0 {
		return errors.New("users command takes no arguments")
	}
	
	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("failed to get users: %w", err)
	}
	
	for _, user := range users {
		if user.Name == s.cfg.CurrentUser {
			fmt.Printf("* %s (current)\n", user.Name)
		} else {
			fmt.Printf("* %s\n", user.Name)
		}
	}
	
	return nil
}