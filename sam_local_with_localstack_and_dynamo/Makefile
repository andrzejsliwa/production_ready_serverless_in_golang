build:
	 GOOS=linux go build -o bin/main
	 chmod +x bin/main

start: build
	AWS_DYNAMODB_ENDPOINT=http://172.17.0.4:4569 sam local start-api --docker-network localstack_default

localstack:
	localstack start --docker