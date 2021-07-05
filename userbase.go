package main

import (
	"context"
	"log"

	firebase "firebase.google.com/go/v4"
	"google.golang.org/api/iterator"
)

func main() {

	ctx := context.Background()

	app, err := firebase.NewApp(ctx, nil)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	client, err := app.Auth(ctx)
	if err != nil {
		log.Fatalf("error getting Auth client: %v\n", err)
	}

	uidList := []string{}

	iter := client.Users(ctx, "")
	for {
		user, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("error listing users: %s\n", err)
		}
		uidList = append(uidList, user.UID)
	}

	res, err := client.DeleteUsers(ctx, uidList)
	if err != nil {
		log.Fatalf("error deleting users: %v\n", err)
	}

	log.Printf("Successfully deleted %d users", res.SuccessCount)
	log.Printf("Failed to delete %d users", res.FailureCount)

	for _, err := range res.Errors {
		log.Printf("%v", err)
	}
}
