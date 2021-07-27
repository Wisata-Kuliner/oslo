package repository

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

func GetUsers(ctx context.Context, client *firestore.Client) (response []map[string]interface{}, err error) {
	iter := client.Collection("users").Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate: %v", err)
		}
		response = append(response, doc.Data())
	}
	return response, nil
}
