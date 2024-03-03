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
                powershell 'go build ./main.go'
            }
        }

        stage('Run') {
            steps {
                powershell './main.exe'
            }
        }
    }

    post {
        always {
            cleanWs()
        }
    }
}
