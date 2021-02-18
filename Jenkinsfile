pipeline {
  agent any
  stages {
    stage('build') {
      steps {
        sh 'docker build .'
      }
    }

    stage('kuber') {
      steps {
        sh 'cd /root'
      }
    }

  }
}