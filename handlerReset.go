package main

import (
	"fmt"
	"context"
)

func handlerReset(s *State, cmd Command) error {
	if len(cmd.args) != 0 {
		return fmt.Errorf("Usage: gator reset\n")
	}
	
	// run the s.db.ResetUsers() command once made
	
	err := s.db.ResetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("failed to reset users: %w\n", err)	
	}
	
	fmt.Println("successfully reset users database")
	return nil
}