pipeline {
    // install golang 1.14 on Jenkins node
    agent any
    tools {
        go 'go1.20'
    }
    // environment {
    //     GO114MODULE = 'on'
    //     CGO_ENABLED = 0 
    //     GOPATH = "${JENKINS_HOME}/jobs/${JOB_NAME}/builds/${BUILD_ID}"
    // }
    stages {
    
        stage("build") {
            steps {
                echo 'BUILD EXECUTION STARTED'
                sh 'docker build -t localhost:5000/panicscript/test-jenkins-qrcode:v1.1 .'
            }
        }
        // stage('deliver') {
        //     agent any
        //     steps {
        //         withCredentials([usernamePassword(credentialsId: 'dockerhub', 
        //                            passwordVariable: 'dockerhubPassword', 
        //                            usernameVariable: 'dockerhubUser')]) {
        //         sh "docker login -u ${env.dockerhubUser} -p ${env.dockerhubPassword}"
        //         sh 'docker push shadowshotx/product-go-micro'
        //         }
        //     }
        // }
    }
}