# Gator CLI - Blog Aggregator

Gator is a command-line interface (CLI) application written in Go that acts as a blog aggregator. It allows you to follow RSS feeds, fetch posts automatically, and browse them directly from your terminal.

## Prerequisites

To run this program, you will need to have the following installed on your machine:
*   [Go](https://go.dev/doc/install) (1.20+ recommended)
*   [PostgreSQL](https://www.postgresql.org/download/)

You must also have a PostgreSQL server running and accessible.

## Installation

You can install the `gator` CLI directly using Go's `install` command. Run the following command in your terminal:

```bash
go install github.com/vohrr/blog_aggregator@latest
```

*Note: Make sure your `$(go env GOPATH)/bin` is added to your system's `$PATH` so you can run the `gator` command from anywhere.*

## Configuration

Before running the program, you need to set up your configuration file. Create a file named `.gatorconfig.json` in your home directory (`~/.gatorconfig.json`).

Add the following JSON structure to the file, replacing the `db_url` with your actual PostgreSQL connection string:

```json
{
  "db_url": "postgres://username:password@localhost:5432/gator?sslmode=disable",
  "current_user_name": ""
}
```

* Make sure to create the `gator` database in your local Postgres instance before running the commands.

## Running the Program

Once your configuration is set up and your database is running, you can start using the `gator` CLI. 

The basic syntax is:
```bash
gator <command> [arguments...]
```

### Available Commands

Here are a few commands you can run to get started:

*   **`register <username>`**: Registers a new user and automatically logs them in.
    ```bash
    gator register alice
    ```
*   **`login <username>`**: Logs in as an existing user.
    ```bash
    gator login alice
    ```
*   **`addfeed <name> <url>`**: Adds a new RSS feed to the aggregator and follows it as the current user.
    ```bash
    gator addfeed "Boot.dev Blog" "https://blog.boot.dev/index.xml"
    ```
*   **`agg <time_between_requests>`**: Starts the background worker to aggregate feeds (e.g., `1m` for 1 minute).
    ```bash
    gator agg 1m
    ```
*   **`browse [limit]`**: Browses the latest posts from the feeds you follow. You can optionally provide a limit (default is 2).
    ```bash
    gator browse 5
    ```
*   **`users`**: Lists all registered users in the database.
