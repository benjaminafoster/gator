package main

import (
	"fmt"
	"errors"
)

func handlerLogin(s *State, cmd Command) error {
	if len(cmd.args) == 0 {
		return errors.New("username is required")
	}
	
	if len(cmd.args) > 1 {
		return errors.New("too many arguments")
	}
	
	username := cmd.args[0]
	err := s.cfg.SetUser(username)
	if err != nil {
		return fmt.Errorf("failed to login: %w", err)
	}
	
	//fmt.Printf("Logged in as %s\n", username)
	return err
}