pipeline {
    agent any

    environment {
        AWS_DEFAULT_REGION = "us-east-1"
    }

    stages {
        stage('Clone Repo') {
            steps {
                git branch: 'main',
                    url: 'https://github.com/panwarharsh15/s3-store.git'
            }
        }

        stage('Build Go Script') {
            steps {
                sh '''
                go mod init s3uploader || true
                go get github.com/aws/aws-sdk-go-v2
                go get github.com/aws/aws-sdk-go-v2/service/s3
                go build -o uploader upload_to_s3.go
                '''
            }
        }

        stage('Upload to S3') {
            steps {
                sh '''
                mkdir -p repo
                cp -r * repo/
                ./uploader
                '''
            }
        }
    }
}
