package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type TwitterPost struct {
	UserID    string `bson:"user_id"`
	Message   string `bson:"message,omitempty"`
	Timestamp string `bson:"timestamp,omitempty"`
	Retweets  int    `bson:"retweets,omitempty"`
	Favorites int    `bson:"favorites,omitempty"`
}

type InstagramPost struct {
	UserID     string `bson:"user_id"`
	Caption    string `bson:"caption,omitempty"`
	ImageURL   string `bson:"image_url,omitempty"`
	Likes      int    `bson:"likes,omitempty"`
	Comments   int    `bson:"comments,omitempty"`
	PostedDate string `bson:"posted_date,omitempty"`
}

type FacebookPost struct {
	UserID      string `bson:"user_id"`
	Content     string `bson:"content,omitempty"`
	Reactions   int    `bson:"reactions,omitempty"`
	Comments    int    `bson:"comments,omitempty"`
	Shares      int    `bson:"shares,omitempty"`
	PublishedAt string `bson:"published_at,omitempty"`
}

func main() {
	ctx := context.Background()
	// Connect to MongoDB.
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	// Access the "socialmedia" database.
	database := client.Database("socialmedia")

	// Access a common collection for all types of posts.
	collection := database.Collection("posts")

	// Insert a Twitter post.
	twitterPost := TwitterPost{
		UserID:    "twitter_user1",
		Message:   "Exciting news on Twitter!",
		Timestamp: "2023-01-20T14:45:00",
		Retweets:  25,
	}
	_, err = collection.InsertOne(ctx, twitterPost)
	if err != nil {
		log.Fatal(err)
	}

	// Insert an Instagram post.
	instagramPost := InstagramPost{
		UserID:     "instagram_user1",
		Caption:    "Beautiful sunset on Instagram!",
		ImageURL:   "https://example.com/sunset.jpg",
		Likes:      100,
		Comments:   15,
		PostedDate: "2023-01-22T18:30:00",
	}
	_, err = collection.InsertOne(ctx, instagramPost)
	if err != nil {
		log.Fatal(err)
	}

	// Insert a Facebook post.
	facebookPost := FacebookPost{
		UserID:      "facebook_user1",
		Content:     "Sharing a great article on Facebook!",
		Reactions:   50,
		Comments:    10,
		Shares:      5,
		PublishedAt: "2023-01-24T09:15:00",
	}
	_, err = collection.InsertOne(ctx, facebookPost)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted Twitter, Instagram, and Facebook posts into one collection.")
}
