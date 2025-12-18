package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

const (
	bucketName = "automated-store-obj"
	region     = "us-east-1"
	rootDir    = "." // Jenkins workspace
)

func main() {
	cfg, err := config.LoadDefaultConfig(
		context.TODO(),
		config.WithRegion(region),
	)
	if err != nil {
		log.Fatalf("AWS config error: %v", err)
	}

	s3Client := s3.NewFromConfig(cfg)

	err = filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip directories and .git
		if info.IsDir() || strings.Contains(path, ".git") {
			return nil
		}

		// S3 object key
		key := strings.TrimPrefix(path, "./")

		file, err := os.Open(path)
		if err != nil {
			return err
		}
		defer file.Close()

		_, err = s3Client.PutObject(context.TODO(), &s3.PutObjectInput{
			Bucket: &bucketName,
			Key:    &key,
			Body:   file,
		})
		if err != nil {
			return err
		}

		fmt.Printf("Uploaded: %s → s3://%s/%s\n", path, bucketName, key)
		return nil
	})

	if err != nil {
		log.Fatalf("Upload failed: %v", err)
	}

	fmt.Println("✅ All files uploaded successfully")
}
