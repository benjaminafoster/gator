package main

import (
	"github.com/benjaminafoster/gator/internal/config"
	"github.com/benjaminafoster/gator/internal/database"
)

type State struct {
	db *database.Queries
	cfg *config.Config
}
