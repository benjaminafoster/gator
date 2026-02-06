package main

import (
	"fmt"
)

func handlerAddFeed(s *State, cmd Command) error {
	if len(cmd.args) != 2 {
		return fmt.Errorf("Usage: addFeed <name> <url>\n")
	}

	name := cmd.args[0]
	url := cmd.args[1]

	feed, err := addFeed(s, name, url)
	if err != nil {
		return fmt.Errorf("failed to add feed: %w\n", err)
	}
	
	fmt.Printf("New feed added: %s\n", feed.Name)
	
	followFeed, err := followFeed(s, url)
	if err != nil {
		return err
	}
	
	fmt.Printf("Now following feed: %s\n", followFeed.FeedName)

	return nil
}