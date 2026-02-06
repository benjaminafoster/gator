package main

import (
	"context"
	"fmt"
)

func handlerFollowing(s *State, cmd Command) error {
	if len(cmd.args) != 0 {
		return fmt.Errorf("Usage: gator following\n")
	}
	
	currentUsername := s.cfg.CurrentUser
	currentUser, err := s.db.GetUser(context.Background(), currentUsername)
	if err != nil {
		return fmt.Errorf("Failed to get current user: %v", err)
	}
	
	following, err := s.db.GetFeedFollowsForUser(context.Background(), currentUser.ID)
	if err != nil {
		return fmt.Errorf("Failed to get following list: %v", err)
	}
	
	fmt.Println("Feeds you follow:")
	for _, followItem := range following {
		fmt.Printf("  - %s\n",followItem.FeedName)
	}
	
	return nil
}