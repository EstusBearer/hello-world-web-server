pipeline {
    agent any

    environment {
        GO111MODULE = 'on'
    }

    stages {
        stage('Checkout') {
            steps {
                checkout scm
            }
        }

        stage('Set GoPath') {
            steps {
                script {
                    env.GOPATH = "${WORKSPACE}"
                }
            }
        }

        stage('Build') {
            steps {
                sh 'go build -o myGoApp'
            }
        }

        stage('Run') {
            steps {
                sh './myGoApp'
            }
        }
    }

    post {
        always {
            cleanWs()
        }
    }
}
