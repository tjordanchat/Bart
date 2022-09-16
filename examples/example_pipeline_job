pipeline {
    agent any
    stages {
        stage('Initialization') {
            environment { 
                   JOB_TIME = sh (returnStdout: true, script: "date '+%A %W %Y %X'").trim()
            }
            steps {
                sh 'echo $JOB_TIME'
            }
        }
    }
}
