package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

const (
	bucketName = "my-s3-bucket-name"
	region     = "ap-south-1"
)

func uploadFile(client *s3.Client, filePath string, key string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(key),
		Body:   file,
	})
	return err
}

func main() {
	dir := "./repo" // Jenkins workspace path

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(region))
	if err != nil {
		log.Fatal(err)
	}

	client := s3.NewFromConfig(cfg)

	err = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		s3Key := path[len(dir)+1:]
		fmt.Println("Uploading:", s3Key)

		return uploadFile(client, path, s3Key)
	})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Upload completed successfully")
}
