package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
)

func getSecretValue(secretName string) (string, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return "", fmt.Errorf("unable to load SDK config, %v", err)
	}

	client := secretsmanager.NewFromConfig(cfg)

	resp, err := client.GetSecretValue(context.TODO(), &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(secretName),
	})
	if err != nil {
		return "", fmt.Errorf("unable to get secret value: %v", err)
	}

	return aws.ToString(resp.SecretString), nil
}

func main() {
	http.HandleFunc("/secret", func(w http.ResponseWriter, r *http.Request) {
		secretName := os.Getenv("AWS_SECRET_NAME")
		if secretName == "" {
			http.Error(w, "AWS_SECRET_NAME not set", http.StatusBadRequest)
			return
		}

		secret, err := getSecretValue(secretName)
		if err != nil {
			log.Printf("error retrieving secret: %v", err)
			http.Error(w, "internal error", http.StatusInternalServerError)
			return
		}

		fmt.Fprintf(w, "Secret value: %s\n", secret)
	})

	log.Println("Server started on :8080")
	http.ListenAndServe(":8080", nil)
}
