package repository

import (
	"context"
	"io/ioutil"
	"log"
	"os"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

func DownloadData(b []byte) ([]byte, error) {
	config := &firebase.Config{
		StorageBucket: os.Getenv("FIREBASE_MABAR_PROJECT_ID") + ".appspot.com",
	}
	opt := option.WithCredentialsJSON([]byte(b))

	app, err := firebase.NewApp(context.Background(), config, opt)
	if err != nil {
		log.Fatalln(err)
	}

	client, err := app.Storage(context.Background())
	if err != nil {
		log.Fatalln(err)
	}

	bucket, err := client.DefaultBucket()
	if err != nil {
		log.Fatalln(err)
	}

	// 'bucket' is an object defined in the cloud.google.com/go/storage package.
	// See https://godoc.org/cloud.google.com/go/storage#BucketHandle
	// for more details.storage.go

	rc, err := bucket.Object("2021-08-26T12:00:59_72896/all_namespaces/all_kinds/output-0").NewReader(context.Background())
	if err != nil {
		return nil, err
	}
	defer rc.Close()

	data, err := ioutil.ReadAll(rc)
	if err != nil {
		return nil, err
	}
	return data, nil
}
