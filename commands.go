package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/benjaminafoster/gator/internal/database"
	"github.com/google/uuid"
)



type command struct {
	name           string
	args           []string
}

func handlerLogin(s *State, cmd command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.name)
	}
	
	
	userName := cmd.args[0]

	
	_, err := s.Db.GetUser(context.Background(), userName)
    if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			// User does not exist, provide error. You can't log in with a user that doesn't exist
			return fmt.Errorf("couldn't find user: %w", err)
			//return fmt.Errorf("user does not exist")
		} else {
			// Some other error occurred
			return fmt.Errorf("error checking if user exists: %w", err)
		}
	} else {
		// User already exists
		err := s.CfgPtr.SetUser(userName)
		if err != nil {
			return fmt.Errorf("couldn't set current user: %w", err)
		}
	}
			
	fmt.Printf("user switched successfully!\n")
	
	return nil
	
}

func handlerRegister(s *State, cmd command) error {
    if len(cmd.args) != 1 {
        return fmt.Errorf("usage: %s <name>", cmd.name)
    }

    userName := cmd.args[0]

	userParams := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      userName,
	}
	_, err := s.Db.GetUser(context.Background(), userName)
    if err != nil {
        if errors.Is(err, sql.ErrNoRows) {
            // User does not exist, proceed with registration
			_, err = s.Db.CreateUser(context.Background(), userParams)
			if err != nil {
				return fmt.Errorf("error creating user in database: %w", err)
			}
        } else {
            // Some other error occurred
            return fmt.Errorf("error checking if user exists: %w", err)
        }
    } else {
        // User already exists
		return fmt.Errorf("user already exists")
    }

	s.CfgPtr.SetUser(userName)

	fmt.Printf("user '%s' successfully registered\n",s.CfgPtr.CurrentUser)

    return nil
}

func handlerReset(s *State, cmd command) error {
	if len(cmd.args) != 0 {
		return fmt.Errorf("reset command takes no additional arguments")
	}

	fmt.Println("users database successfully reset")
	s.Db.DeleteAllUsers(context.Background())

	return nil
}

func handlerUsers(s *State, cmd command) error {
	if len(cmd.args) != 0 {
		return fmt.Errorf("users command takes no additional arguments")
	}

	users, err := s.Db.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("error getting user names from users database: %w", err)
	}
	for _, user := range users {
		if user == s.CfgPtr.CurrentUser {
			fmt.Printf("%s (current)\n", user)
		} else {
			fmt.Printf("%s\n", user)
		}
	}

	return nil
}

func handlerAgg(_ *State, cmd command) error {
	if len(cmd.args) != 0 {
		return fmt.Errorf("agg command takes no additional parameters")
	}
	
	feed, err := fetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		return fmt.Errorf("error fetching the feed: %w", err)
	}


	fmt.Printf("%+v\n", feed)
	return nil
}

func handlerAddFeed(s *State, cmd command) error {
	if len(cmd.args) != 2 {
		return fmt.Errorf("addfeed usage: gator <name_of_feed> <url_of_feed>")
	}

	name := cmd.args[0]
	url := cmd.args[1]

	db_user, err := s.Db.GetUser(context.Background(),s.CfgPtr.CurrentUser)
	if err != nil {
		return fmt.Errorf("error retrieving current user name from users db: %w", err)
	}

	user_id := db_user.ID

	params := database.AddFeedParams{
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name: name,
		Url: url,
		UserID: user_id,
	}

	feed, err := s.Db.AddFeed(context.Background(), params)
	if err != nil {
		return fmt.Errorf("error adding feed to feeds database: %w", err)
	}

	fmt.Println("Successfully added a feed with the following details:")
	fmt.Printf("Name: %s\n", feed.Name)
	fmt.Printf("URL: %s\n", feed.Url)

	return nil

}

type commands struct {
	registeredCommands         map[string]func(*State, command) error
}

func (c *commands) register(name string, f func(*State, command) error) {
	c.registeredCommands[name] = f
}

func (c *commands) run(s *State, cmd command) error {
	if f, exists := c.registeredCommands[cmd.name]; exists {
		err := f(s, cmd)
		return err
	}

	return fmt.Errorf("error executing command: %s", cmd.name)
}