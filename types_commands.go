package main

import (
	"fmt"
)

type Commands struct {
	cmds map[string](func(s *State, cmd Command) error)
}

func GetCommands() *Commands {
	cmds := &Commands{
		cmds: map[string]func(s *State, cmd Command) error{
			"login": handlerLogin,
			"register": handlerRegister,
			"reset": handlerReset,
			"users": handlerUsers,
			"agg": handlerAggregate,
			"addfeed": middlewareLoggedIn(handlerAddFeedv2),
			"feeds": handlerFeeds,
			"follow": middlewareLoggedIn(handlerFollowv2),
			"following": middlewareLoggedIn(handlerFollowingv2),
			"unfollow": middlewareLoggedIn(handlerUnfollow),
		},
	}
	return cmds
}

func (c *Commands) run(s *State, cmd Command) error {
	callback, ok := c.cmds[cmd.name]
	if !ok {
		return fmt.Errorf("unknown command: %s", cmd.name)
	}
	
	err := callback(s, cmd)
	if err != nil {
		return fmt.Errorf("error running command '%s': %w", cmd.name, err)
	}
	
	return err
}
