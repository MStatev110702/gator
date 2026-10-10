# Gator

Gator is a command-line application written in Go that allows users to manage and browse RSS feeds.

## Prerequisites

Before running Gator, make sure you have the following installed:

Go — Install Go from the official [Go website](https://go.dev/doc/install).

PostgreSQL — Install PostgreSQL from the official [PostgreSQL website](https://www.postgresql.org/download/).

Verify your installations:

```bash
go version
psql --version
```

## Installation

Install the Gator CLI using go install:

```bash
git clone https://github.com/MStatev110702/gator.git
cd gator
go install .
```

## Configuration

Gator requires a configuration file to store your username and database connection string.

Create a configuration file at:

```bash
~/.gatorconfig.json
```

Example configuration:

```json
{
  "db_url": "postgres://postgres:password@localhost:5432/gator?sslmode=disable",
  "current_user_name": ""
}
```

Update the database connection string with your PostgreSQL username, password, host, port, and database name.

Create the database in PostgreSQL if it doesn't already exist:

```bash
CREATE DATABASE gator;
```

The application reads the configuration file to determine which database to connect to and which user is currently active.

## Database Setup

Apply the database schema and migrations before running the application. Inside the root of your project run:

```bash
cd sql/schema
goose postgres "postgres://username@localhost:5432/gator" up
```

## Available Commands

The following are examples of common Gator commands. Replace or adjust them to match the commands implemented in your project.

| command                    | Descrption                                        |
| -------------------------- | ------------------------------------------------- |
| gator register <username>  | Register a new user                               |
| gator login <username>     | Login as an existing user                         |
| gator users                | List registered users                             |
| gator addfeed <name> <url> | Add an RSS feed                                   |
| gator feeds                | List available feeds                              |
| gator follow <url>         | Follow an RSS feed                                |
| gator following            | List feeds followed by the current user           |
| gator agg <duration>       | Fetch and aggregate posts at a specified interval |
| gator browse <limit>       | Browse recent posts                               |

### Example

```bash
gator register alice
gator login alice
gator addfeed "Example Blog" "https://example.com/rss.xml"
gator feeds
gator follow https://example.com/rss.xml
gator following
gator agg 1m
gator browse 10
```
