# aws_lambda_in_motion_golang

Examples from [AWS Lambda in Motion](https://www.manning.com/livevideo/production-ready-serverless) (renamed recently to `Production-Ready Serverless`), made in Golang (Video Course is based on Node.js)

## Setup

```shell
[~]$ cd $GOPATH/src/github.com/andrzejsliwa
[~/go/src/github.com/andrzejsliwa]$ git clone git@github.com:andrzejsliwa/aws_lambda_in_motion_golang.git
Cloning into 'aws_lambda_in_motion_golang'...
remote: Counting objects: 10, done.
remote: Compressing objects: 100% (9/9), done.
remote: Total 10 (delta 0), reused 10 (delta 0), pack-reused 0
Receiving objects: 100% (10/10), done.
[~/go/src/github.com/andrzejsliwa]$ cd aws_lambda_in_motion_golang
[master][~/go/src/github.com/andrzejsliwa/aws_lambda_in_motion_golang]$
```

## Deployment

```shell
[master][~/go/src/github.com/andrzejsliwa/aws_lambda_in_motion_golang]$ cd aws_lambda_in_motion_golang
[master][~/go/src/github.com/andrzejsliwa/aws_lambda_in_motion_golang/hello_manning]$ make
dep ensure
env GOOS=linux go build -ldflags="-s -w" -o bin/hello hello/main.go

[master][~/go/src/github.com/andrzejsliwa/aws_lambda_in_motion_golang/hello_manning]$ sls deploy
Serverless: Packaging service...
Serverless: Excluding development dependencies...
Serverless: Uploading CloudFormation file to S3...
Serverless: Uploading artifacts...
Serverless: Uploading service .zip file to S3 (2.25 MB)...
Serverless: Validating template...
Serverless: Updating Stack...
Serverless: Checking Stack update progress...
..........
Serverless: Stack update finished...
Service Information
service: hello-manning
stage: dev
region: us-west-1
stack: hello-manning-dev
api keys:
  None
endpoints:
  GET - https://*********.execute-api.us-west-1.amazonaws.com/dev/
functions:
  hello: hello-manning-dev-hello
```

## Running

Open endpoint (from your console output, not from this file)
Example:
```shell
$ open https://*********.execute-api.us-west-1.amazonaws.com/dev/
```