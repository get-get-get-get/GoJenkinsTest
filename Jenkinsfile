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
                withEnv(['XDG_CACHE_HOME=/tmp']) {
                    sh 'go test .'
                }
            }
        }
    }
}
