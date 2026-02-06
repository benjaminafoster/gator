package main

import (
	"fmt"
)


func handlerFollow(s *State, cmd Command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("Usage: gator follow <feed_url>")
	}
	
	feedUrl := cmd.args[0]
	
	feedFollow, err := followFeed(s, feedUrl)
	if err != nil {
		return err
	}
	
	fmt.Printf("Feed name: %s\n", feedFollow.FeedName)
	fmt.Printf("Current user: %s\n", feedFollow.UserName)
	
	return nil
} 