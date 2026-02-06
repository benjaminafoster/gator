package main

import (
	"fmt"
	"context"
	"github.com/benjaminafoster/gator/internal/database"
)

func middlewareLoggedIn(handler func(s *State, cmd Command, user database.User) error) func(s *State, cmd Command) error {
	return func(s *State, cmd Command) error {
		currentUsername := s.cfg.CurrentUser
		currentUser, err := s.db.GetUser(context.Background(), currentUsername)
		if err != nil {
			return fmt.Errorf("Failed to get current user: %v", err)
		}
		
		return handler(s, cmd, currentUser)
	}
}