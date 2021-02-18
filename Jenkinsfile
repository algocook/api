pipeline {
  agent any
  stages {
    stage('build') {
      steps {
        sh 'docker build'
      }
    }

    stage('') {
      steps {
        sh 'kubectl applyt -k .'
      }
    }

  }
}