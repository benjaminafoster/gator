package main

import (
	"fmt"
)



type command struct {
	name           string
	args           []string
}

func handlerLogin(s *State, cmd command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.name)
	}
	
	err := s.CfgPtr.SetUser(cmd.args[0])
	if err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}
	
	fmt.Printf("user switched successfully!")
	
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
		f(s, cmd)
		return nil
	}

	return fmt.Errorf("error executing command: %s", cmd.name)
}