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
                powershell 'C:\\Users\\ruslan\\sdk\\go1.22.0\\bin\\go.exe build .\\main.go'
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
