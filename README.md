# Gator -- A Boot.dev RSS aggregator written in Go

## Prerequisites
In order to run this project, you will need to have the following installed:

- PostgreSQL (on localhost or remote)
- Go (version 1.16 or higher)

## Installation
To install Gator, run the following command in your terminal:
```shell
go install github.com/benjaminafoster/gator
```

## Setup
Gator requires a configuration file named `.gatorconfig.json` in your home directory. It is very important that the name of the file is exactly as shown, otherwise Gator will not be able to read it and run the program.
Within your gator file, you need to specify the following:
- `db_url`: The PostgreSQL database connection string for your preconfigured gator database.
- `current_user`: An array of RSS feed URLs to aggregate.

### Example Configuration File
{
    "db_url": "postgres://user:password@localhost:5432/gator",
    "current_user": "user1"
}

## Usage
Once installed, you can run Gator by executing the gator command with supported subcommands. General usage is as follows:
```shell
gator <subcommand> [arguments]
```

### Available subcommands:
See a list of available subcommands below with their respective positional arguments in order below:
- 'register': Register a new user with Gator.
  - username (required): the username of the new user.
- 'login': Log in to an existing user account.
  - username (required): the username of the user to log in.
- 'users': Get a list of registered Gator users.
- 'feeds': Get a list of all feeds registered in the Gator database.
- 'addfeed': Add a new RSS feed to the Gator database.
  - name (required): The name of the new feed.
  - url (required): The URL of the new feed.
- 'follow': Follow a specified RSS feed.
  - url (required): The URL of the feed to follow.
- 'unfollow': Unfollow a specified RSS feed.
  - url (required): The URL of the feed to unfollow.
- 'following': Get a list of all feeds followed by the current user.
- 'browse': Browse the current user's saved feeds and posts.
  - limit (optional): The maximum number of posts to display (default: 5)
- 'agg': Aggregate posts from all feeds in the Gator database.
  - time_between_requests (required): The time interval between requests to the RSS feed (i.e. '1m' for 1 minute, '1h' for 1 hour, '1d' for 1 day)
- 'reset': Reset the Gator database to an empty state.

If a command expects arguments, you will be provided with a usage statement when running the subcommand without arguments.