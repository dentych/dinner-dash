pipeline {
    agent any

    environment {
        IMAGENAME = "dinnerdash/backend"
    }

    stages {
        stage("Build backend docker image") {
            steps {
                dir("backend") {
                    sh "docker build -t $IMAGENAME ."
                }
            }
        }

        stage("Push backend docker image to Docker Hub") {
            steps {
                withDockerRegistry([url: "", credentialsId: "docker-hub-dentych"]) {
                    sh "docker push $IMAGENAME"
                }
            }
        }
    }
}
