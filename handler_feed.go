package main

import (
	"blogaggregator/internal/database"
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
)

func handlerAddFeed(s *state, cmd command) error {
	if len(cmd.Args) != 2 {
		return fmt.Errorf("usage %v <name> <url>", cmd.Name)
	}

	row, err := s.db.GetUser(context.Background() ,s.cfg.CurrentUserName)
	if err != nil {
		return err
	}
	
	feed, err := s.db.CreateFeed(context.Background(), 
		database.CreateFeedParams{
			ID: 		uuid.New(),
			CreatedAt: 	time.Now().UTC(),
			UpdatedAt: 	time.Now().UTC(),
			Name: 		cmd.Args[0],
			Url: 		cmd.Args[1],
			UserID: 	row.ID,
	})
	if err != nil {
		return fmt.Errorf("couldn't create feed: %w", err)
	}
	
	fmt.Println("Feed created sucessfully:")
	printFeed(feed)
	return nil
}

func printFeed(feed database.Feed) {
	fmt.Printf("* ID: 			%s\n", feed.ID)
	fmt.Printf("* CreatedAt:	%v\n", feed.CreatedAt)
	fmt.Printf("* UpdatedAt:	%v\n", feed.UpdatedAt)
	fmt.Printf("* Name:			%s\n", feed.Name)
	fmt.Printf("* URL:			%s\n", feed.Url)
}