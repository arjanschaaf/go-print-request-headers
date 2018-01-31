pipeline {
  agent {
    node {
      label 'ACI-container'
    }
    
  }
  stages {
    stage('build') {
      steps {
        sh 'docker build .'
      }
    }
  }
}