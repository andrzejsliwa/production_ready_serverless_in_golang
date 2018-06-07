# Production Ready Serverless in GoLang

Examples from [Production Ready Serverless](https://www.manning.com/livevideo/production-ready-serverless), made in Golang (Video Course is based on Node.js)

## Setup

```shell
[~]$ cd $GOPATH/src/github.com/andrzejsliwa
[~/go/src/github.com/andrzejsliwa]$ git clone git@github.com:andrzejsliwa/production_ready_serverless_in_golang.git
Cloning into 'production_ready_serverless_in_golang'...
remote: Counting objects: 10, done.
remote: Compressing objects: 100% (9/9), done.
remote: Total 10 (delta 0), reused 10 (delta 0), pack-reused 0
Receiving objects: 100% (10/10), done.
[~/go/src/github.com/andrzejsliwa]$ cd production_ready_serverless_in_golang
[master][~/go/src/github.com/andrzejsliwa/production_ready_serverless_in_golang]$
```

## Deployment

```shell
[master][~/go/src/github.com/andrzejsliwa/production_ready_serverless_in_golang]$ cd production_ready_serverless_in_golang
[master][~/go/src/github.com/andrzejsliwa/production_ready_serverless_in_golang/hello_manning]$ make
dep ensure
env GOOS=linux go build -ldflags="-s -w" -o bin/hello hello/main.go

[master][~/go/src/github.com/andrzejsliwa/production_ready_serverless_in_golang/hello_manning]$ sls deploy
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

## List of Examples

- hello_manning - simple api-gateway example with golang and serverless
- sam_local_with_localstack_and_dynamodb - simple example with using localstack and running on AWS SAM Local (with example test with mocking dynamodb) 

## Notes

- Serverless is using CloudFormation underhood for provisioning
- Currently Serverless is not supporting invoke on local for Golang (you can make it with `AWS SAM Local`)
- AWS Sam Local is running locally lambdas in docker images which are close to production in behaviour
- Together with AWS SAM Local you are able to run LocalStack to emulate resources locally
