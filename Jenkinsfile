pipeline {
    agent any
    
    tools {
        go 'GO'
    }

    environment {
        AWS_REGION = "us-east-1"
    }

    stages {
        stage('Checkout Code') {
            steps {
                checkout scm
            }
        }

        stage('Setup Go') {
            steps {
                sh '''
                go version
                go env
                '''
            }
        }

        stage('Build Go Program') {
            steps {
                sh '''

                go mod init s3sync || true
                go mod tidy
                go build -o s3sync
                '''
            }
        }

        stage('Deploy to S3') {
            steps {
                sh '''
            
                ./s3sync
                '''
            }
        }
    }

    post {
        success {
            echo 'Deployment to S3 completed successfully'
        }
        failure {
            echo 'Deployment failed'
        }
    }
}
