package main

import (
	"context"
	"fmt"
	"strconv"
	"github.com/benjaminafoster/gator/internal/database"
)

func handlerBrowse(s *State, cmd Command, user database.User) error {
	if len(cmd.args) > 1 {
		return fmt.Errorf("Usage: gator browse [limit: int]")
	}
	
	var limit int32
	limit = 5
	
	if len(cmd.args) == 1 {
		limitStr := cmd.args[0]
		parsed, err := strconv.ParseInt(limitStr, 10, 32)
		if err != nil {
			return fmt.Errorf("Invalid limit -- must be an integer: %s", limitStr)
		}
		
		limit = int32(parsed)
	}
	
	
	postByUserIDParams := database.GetPostsByUserIDParams{
		ID: user.ID,
		Limit: limit,
	}
	
	posts, err := s.db.GetPostsByUserID(context.Background(), postByUserIDParams)
	if err != nil {
		return fmt.Errorf("Failed to retrieve posts: %v", err)
	}
	
	for _, post := range posts {
		fmt.Printf("\n###############################################\n")
		fmt.Printf("%s\n", post.Title)
		fmt.Printf("%s\n", post.Description)
	}
	
	return nil
}