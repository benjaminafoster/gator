package main

import (
	"fmt"
	"context"
	"github.com/benjaminafoster/gator/internal/database"
)

func handlerUnfollow(s *State, cmd Command, user database.User) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("Usage: gator unfollow <feed_url>")
	}
	
	userID := user.ID
	feedUrl := cmd.args[0]
	
	feedRow, err := s.db.GetFeedByUrl(context.Background(), feedUrl)
	if err != nil {
		return fmt.Errorf("Failed to retrieve feed: %v\n", err)
	}
	
	deleteFeedFollowParams := database.DeleteFeedFollowParams{
		UserID: userID,
		FeedID: feedRow.ID,
	}
	
	err = s.db.DeleteFeedFollow(context.Background(), deleteFeedFollowParams)
	if err != nil {
		return fmt.Errorf("Failed to unfollow feed: %v\n", err)
	}
	
	return nil
}