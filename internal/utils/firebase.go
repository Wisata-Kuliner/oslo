package utils

import (
	"context"
	"fmt"
	"os"
	"strings"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"

	"google.golang.org/api/option"

	"encoding/json"
)

type Config struct {
	ConfigType   string `json:"type"`
	ProjectId    string `json:"project_id"`
	PrivateKeyId string `json:"private_key_id"`
	PrivateKey   string `json:"private_key"`
	ClientEmail  string `json:"client_email"`
	ClientId     string `json:"client_id"`
	AuthUrl      string `json:"auth_url"`
	ClientUrl    string `json:"client_url"`
	AuthProvider string `json:"auth_provider_x509_cert_url"`
	CertUrl      string `json:"client_x509_cert_url"`
}

func NewConfig(ctx context.Context) (*firebase.App, error) {
	serviceAccount := &Config{
		ConfigType:   os.Getenv("FIREBASE_MABAR_TYPE"),
		ProjectId:    os.Getenv("FIREBASE_MABAR_PROJECT_ID"),
		PrivateKeyId: os.Getenv("FIREBASE_MABAR_PRIVATE_KEY_ID"),
		PrivateKey: strings.Replace(
			os.Getenv("FIREBASE_MABAR_PRIVATE_KEY"), "\\n", "\n", -1),
		ClientEmail:  os.Getenv("FIREBASE_MABAR_CLIENT_EMAIL"),
		ClientId:     os.Getenv("FIREBASE_MABAR_CLIENT_ID"),
		AuthUrl:      os.Getenv("FIREBASE_MABAR_AUTH_URL"),
		ClientUrl:    os.Getenv("FIREBASE_MABAR_CLIENT_URL"),
		AuthProvider: os.Getenv("FIREBASE_MABAR_AUTH_PROVIDER"),
		CertUrl:      os.Getenv("FIREBASE_MABAR_CERT_URL"),
	}
	b, err := json.Marshal(serviceAccount)
	if err != nil {
		return nil, fmt.Errorf("error initializing app: %v", err)
	}
	// fmt.Println(string(b))
	opt := option.WithServiceAccountFile(string(b))
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		return nil, fmt.Errorf("error initializing app: %v", err)
	}
	return app, nil
}

func CreateClient(ctx context.Context) (*firestore.Client, error) {
	// Sets your Google Cloud Platform project ID.
	projectID := os.Getenv("FIREBASE_MABAR_PROJECT_ID")

	b, err := ReadServiceAccount()
	if err != nil {
		return nil, err
	}

	// fmt.Println(string(b))
	opt := option.WithCredentialsJSON([]byte(b))

	client, err := firestore.NewClient(ctx, projectID, opt)
	if err != nil {
		// log.Fatalf("Failed to create client: %v", err)
		return nil, fmt.Errorf("error initializing app: %v", err)
	}
	// Close client when done with
	// defer client.Close()
	return client, nil
}

func ReadServiceAccount() ([]byte, error) {
	serviceAccount := &Config{
		ConfigType:   os.Getenv("FIREBASE_MABAR_TYPE"),
		ProjectId:    os.Getenv("FIREBASE_MABAR_PROJECT_ID"),
		PrivateKeyId: os.Getenv("FIREBASE_MABAR_PRIVATE_KEY_ID"),
		PrivateKey: strings.Replace(
			os.Getenv("FIREBASE_MABAR_PRIVATE_KEY"), "\\n", "\n", -1),
		ClientEmail:  os.Getenv("FIREBASE_MABAR_CLIENT_EMAIL"),
		ClientId:     os.Getenv("FIREBASE_MABAR_CLIENT_ID"),
		AuthUrl:      os.Getenv("FIREBASE_MABAR_AUTH_URL"),
		ClientUrl:    os.Getenv("FIREBASE_MABAR_CLIENT_URL"),
		AuthProvider: os.Getenv("FIREBASE_MABAR_AUTH_PROVIDER"),
		CertUrl:      os.Getenv("FIREBASE_MABAR_CERT_URL"),
	}
	b, err := json.Marshal(serviceAccount)
	if err != nil {
		return nil, fmt.Errorf("error initializing app: %v", err)
	}
	return b, nil
}
