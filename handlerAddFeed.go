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

	err := addFeed(s, name, url)
	if err != nil {
		return fmt.Errorf("failed to add feed: %w\n", err)
	}

	return nil
}