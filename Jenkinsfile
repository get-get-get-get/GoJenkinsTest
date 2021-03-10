pipeline {
    agent { docker { image 'golang' } }
    stages {
        stage('get go version') {
            steps {
                script {
                    sh 'go version'
                }
            }
        }
        stage('test') {
            steps {
                script {
                    sh 'go test .'
                }
            }
        }
    }
}
