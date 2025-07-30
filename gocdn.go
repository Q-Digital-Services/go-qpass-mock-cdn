package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	s3 "github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/joho/godotenv"
)

var (
	s3Client *s3.Client
	bucket   string
)

func main() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, relying on environment vars.")
	}

	bucket = os.Getenv("MINIO_BUCKET")
	if bucket == "" {
		log.Fatal("MINIO_BUCKET is required")
	}

	endpoint := os.Getenv("MINIO_ENDPOINT")

	region := os.Getenv("MINIO_REGION")
	if(region == ""){
		region="us-east-1"
	}

	accessKey := os.Getenv("MINIO_ACCESS_KEY")
	secretKey := os.Getenv("MINIO_SECRET_KEY")

	// AWS config with static creds for MinIO
	customResolver := aws.EndpointResolverWithOptionsFunc(
		func(service, region string, _ ...interface{}) (aws.Endpoint, error) {
			return aws.Endpoint{
				PartitionID:   "aws",
				URL:           endpoint,
				SigningRegion: region,
			}, nil
		},
	)

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(region),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(accessKey, secretKey, "")),
		config.WithEndpointResolverWithOptions(customResolver),
	)
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	fmt.Println(endpoint,region, accessKey, secretKey,bucket )

	s3Client = s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.UsePathStyle = true	
	})

	// Start HTTP server
	http.HandleFunc("/", handleRequest)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Println("📦 CDN running at http://localhost:" + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	key := strings.TrimPrefix(r.URL.Path, "/")
	if key == "" {
		http.Error(w, "Missing object path", http.StatusBadRequest)
		return
	}

	obj, err := s3Client.GetObject(context.TODO(), &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		http.Error(w, "Object not found", http.StatusNotFound)
		log.Printf("S3 error: %v", err)
		return
	}
	defer obj.Body.Close()

	// Set content type and cache headers
	if obj.ContentType != nil {
		w.Header().Set("Content-Type", *obj.ContentType)
	}
	w.Header().Set("Cache-Control", "public, max-age=31536000")

	_, err = io.Copy(w, obj.Body)
	if err != nil {
		log.Printf("Streaming error: %v", err)
	}
}
