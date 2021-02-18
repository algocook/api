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
        sh 'kubernetes/cluster/kubectl.sh apply -k '
      }
    }

  }
}