package main

import (
	"fmt"
	"errors"
	"context"
)

func handlerReset(s *State, cmd Command) error {
	if len(cmd.args) != 0 {
		return errors.New("reset command does not take any arguments")
	}
	
	// run the s.db.ResetUsers() command once made
	
	err := s.db.ResetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("failed to reset users: %w", err)	
	}
	
	fmt.Println("finished resetting users database")
	return nil
}