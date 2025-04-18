package main

import (
	// "fmt"
	"github.com/benjaminafoster/gator/internal/config"
	"github.com/benjaminafoster/gator/internal/database"
)

type State struct {
	Db           *database.Queries
	CfgPtr       *config.Config
}