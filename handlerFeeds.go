package main

import (
	"context"
	"fmt"
)

func handlerFeeds(s *State, cmd Command) error {
	if len(cmd.args) != 0 {
		return fmt.Errorf("Usage: gator feeds\n")
	}
	
	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("error retrieving feeds: %w\n", err)
	}
	
	for _, feed := range feeds {
		user, err := s.db.GetUserByID(context.Background(), feed.UserID)
		if err != nil {
			return fmt.Errorf("error retrieving user by ID: %w\n", err)
		}
		fmt.Printf("Feed title: %s\n", feed.Name)
		fmt.Printf("Feed URL: %s\n", feed.Url)
		fmt.Printf("Feed creator: %s\n\n", user.Name)
	}
	
	return nil
}