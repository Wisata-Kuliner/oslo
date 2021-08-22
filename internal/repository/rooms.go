package repository

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

func GetRooms(ctx context.Context, client *firestore.Client) (response []map[string]interface{}, err error) {
	iter := client.Collection("room_chats").Documents(ctx)
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

func PostRooms(ctx context.Context, client *firestore.Client, body map[string]interface{}) (response map[string]interface{}, err error) {
	doc, res, err := client.Collection("room_chats").Add(ctx, body)
	if err != nil {
		log.Fatalf("Failed adding data: %v", err)
	}
	response = make(map[string]interface{})
	response["id"] = doc.ID
	response["timestamp"] = res.UpdateTime
	return response, nil
}

func PutRooms(ctx context.Context, client *firestore.Client, id string, body map[string]interface{}) (response map[string]interface{}, err error) {
	res, err := client.Collection("room_chats").Doc(id).Set(ctx, body)
	if err != nil {
		log.Fatalf("Failed adding data: %v", err)
	}
	response = make(map[string]interface{})
	response["id"] = id
	response["timestamp"] = res.UpdateTime
	return response, nil
}

func DeleteRooms(ctx context.Context, client *firestore.Client, id string, body map[string]interface{}) (response map[string]interface{}, err error) {
	res, err := client.Collection("room_chats").Doc(id).Delete(ctx)
	if err != nil {
		log.Fatalf("Failed adding data: %v", err)
	}
	response = make(map[string]interface{})
	response["id"] = id
	response["timestamp"] = res.UpdateTime
	return response, nil
}
