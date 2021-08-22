package repository

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

func GetPlayers(ctx context.Context, client *firestore.Client) (response []map[string]interface{}, err error) {
	// fmt.Printf("MASUK %+v", client)
	iter := client.Collection("players").Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate: %v", err)
		}

		// for k := range doc.Data() {
		// 	response += k + "\n"
		// }
		response = append(response, doc.Data())
	}
	return response, nil
}

func PostPlayers(ctx context.Context, client *firestore.Client, body map[string]interface{}) (response map[string]interface{}, err error) {
	doc, res, err := client.Collection("players").Add(ctx, body)
	if err != nil {
		log.Fatalf("Failed adding data: %v", err)
	}
	response = make(map[string]interface{})
	response["id"] = doc.ID
	response["timestamp"] = res.UpdateTime
	return response, nil
}

func PutPlayers(ctx context.Context, client *firestore.Client, id string, body map[string]interface{}) (response map[string]interface{}, err error) {
	res, err := client.Collection("players").Doc(id).Set(ctx, body)
	if err != nil {
		log.Fatalf("Failed adding data: %v", err)
	}
	response = make(map[string]interface{})
	response["id"] = id
	response["timestamp"] = res.UpdateTime
	return response, nil
}

func DeletePlayers(ctx context.Context, client *firestore.Client, id string, body map[string]interface{}) (response map[string]interface{}, err error) {
	res, err := client.Collection("players").Doc(id).Delete(ctx)
	if err != nil {
		log.Fatalf("Failed adding data: %v", err)
	}
	response = make(map[string]interface{})
	response["id"] = id
	response["timestamp"] = res.UpdateTime
	return response, nil
}
