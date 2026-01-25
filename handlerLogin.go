package main

import (
	"fmt"
	"errors"
	"context"
)

func handlerLogin(s *State, cmd Command) error {
	if len(cmd.args) == 0 {
		return errors.New("username is required")
	}
	
	if len(cmd.args) > 1 {
		return errors.New("too many arguments")
	}
	
	username := cmd.args[0]
	
	dbUser, err := s.db.GetUser(context.Background(), username)
	if err != nil {
		return fmt.Errorf("failed to get user while logging in: %w", err)
	}
	
	err = s.cfg.SetUser(dbUser.Name)
	if err != nil {
		return fmt.Errorf("failed to setting user after logging in: %w", err)
	}
	
	return err
}