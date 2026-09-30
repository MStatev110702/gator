package main

import (
	"blogaggregator/internal/config"
	"blogaggregator/internal/database"
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

type state struct {
	db *database.Queries
	cfg *config.Config
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	db, err := sql.Open("postgres", cfg.DBURL)
	if err != nil {
		log.Fatalf("error connecting to db: %v", err)
	}
	defer db.Close()
	dbQueries := database.New(db)

	programState := &state{
		db: dbQueries,
		cfg: &cfg,
	}

	cmds := commands{
		registeredCommands: make(map[string]func(*state, command) error),
	}
	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handlerReset)
	
	if len(os.Args) < 2 {
		log.Fatalf("Usage: cli <command> [args...]")
	}

	err = cmds.run(programState, command{Name: os.Args[1], Args: os.Args[2:]})
	if err != nil {
		log.Fatalf("error running command: %v", err)
	}

}