pipeline {
    agent any

    environment {
        REGISTRY_USER = "berqnet-ci"
        REGISTRY_PASS = "Rg!2024_pushToken" // SORUN: Registry password plaintext olarak girilmiş. Repoya erişebilen herkes görebilir.
        IMAGE         = "berqnet/updater"
    }

    stages { 
        // SORUN: Build, Test'ten önce geliyor. Önce test çalışmalı; test geçmezse build/push olmamalı.
        stage('Build') {
            steps {
                sh "docker build -t ${IMAGE}:latest ."
            }
        }

        stage('Test') {
            steps {
                sh "go test ./..."
            }
        }

        stage('Push') {
            steps {
                sh "docker login -u ${REGISTRY_USER} -p ${REGISTRY_PASS}"
                sh "docker push ${IMAGE}:latest"
            }
        }

        stage('Deploy to prod') {
            steps {
                sh "kubectl set image deployment/updater updater=${IMAGE}:latest -n prod"
            }
        }
    }
}
