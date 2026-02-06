package main

import (
	"context"
	"fmt"
)

func handlerFollowing(s *State, cmd Command) error {
	if len(cmd.args) != 0 {
		return fmt.Errorf("Usage: gator following\n")
	}
	
	following, err := s.db.GetFeedFollowsForUser(context.Background())
	if err != nil {
		return fmt.Errorf("Failed to get following list: %v", err)
	}
	
	fmt.Println("Feeds you follow:")
	for _, followItem := range following {
		fmt.Printf("  - %s\n",followItem.FeedName)
	}
	
	return nil
}