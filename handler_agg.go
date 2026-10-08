package main

import (
	"blogaggregator/internal/database"
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/google/uuid"
)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage %v <time>", cmd.Name)
	}

	duration, err := time.ParseDuration(cmd.Args[0])
	if err != nil {
		return fmt.Errorf("couldn't convert into duration: %w", err)
	}

	fmt.Printf("Collection feeds every %v", cmd.Args[0])

	ticker := time.NewTicker(duration)

	for ; ; <-ticker.C {
		scrapeFeeds(s)
	}
}

func scrapeFeeds(s *state) {
	feed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		log.Println("couldn't get feed: %w", err)
	}
	log.Println("Found a feed to fetch!")
	scrapeFeed(s.db, feed)
}

func scrapeFeed(db *database.Queries, feed database.Feed) {
	_, err := db.MarkFeedFetched(context.Background(), feed.ID)
	if err != nil {
		log.Printf("couldn't mark feed %s fetched: %v", feed.Name, err)
		return
	}

	feedData, err := fetchFeed(context.Background(), feed.Url)
	if err != nil {
		log.Printf("couldn't collect feed %s: %v", feed.Name, err)
		return
	}
	for _, item := range feedData.Channel.Item {
		publishedAt := sql.NullTime{}
		if t, err := time.Parse(time.RFC1123Z, item.PubDate); err != nil {
			publishedAt = sql.NullTime{
				Time: t,
				Valid: true,
			}
		}
		_, err = db.CreatePost(context.Background(), database.CreatePostParams{
			ID: uuid.New(),
			CreatedAt: time.Now().UTC(),
			UpdatedAt: time.Now().UTC(),
			PublishedAt: publishedAt,
			Title: item.Title,
			Url: item.Link,
			Description: sql.NullString{
				String: item.Description, 
				Valid: true,
			},
			FeedID: feed.ID,
		})
		if err != nil {
			if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
				continue
			}
			log.Printf("couldn't create post: %v", err)
			continue
		}
	}
	log.Printf("Feed %s collected, %v posts found", feed.Name, len(feedData.Channel.Item))
}