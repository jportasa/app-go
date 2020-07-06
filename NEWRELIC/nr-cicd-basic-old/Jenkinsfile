pipeline {
    agent none

    stages {
        stage('Tests') {
            agent {
                ecs {
                   cloud 'jenkins-slave-ecs'
                   image 'jportasa/nr-alpine-go:1.6'
                   launchType 'FARGATE'
                   memory 1024
                   cpu 256
                   subnets('subnet-08086cbe2d97a1ff1')
                   securityGroups('sg-08f3f54702fb3992e')
                   taskrole 'arn:aws:iam::953835556803:role/ecsTaskExecutionRole'
                   executionRole 'arn:aws:iam::953835556803:role/ecsTaskExecutionRole'
                   assignPublicIp true
                   logDriver 'awslogs'
                   logDriverOptions([[name: 'awslogs-region', value:'us-east-1'], [name: 'awslogs-group', value: '/ecs/jenkins-slave'], [name: 'awslogs-stream-prefix', value: 'alpine']])

                }
            }
            steps {
                sh '''
                    pwd
                    env
                    cd golang-simple-app
                    make test
                    make build
                    aws s3 cp ./target/webapp s3://nr-artifacts/webapp
                   '''

            }
        }
    }
}