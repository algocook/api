pipeline {
  agent any
  stages {
    stage('build') {
      steps {
        sh 'docker build .'
      }
    }

    stage('error') {
      steps {
        sh 'kubectl applyt -k .'
      }
    }

  }
}