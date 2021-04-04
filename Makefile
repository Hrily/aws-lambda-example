TMPDIR=./localstack-temp
GOOS=linux
BINARY=main
ZIP=lambda-example.zip
LOCALSTACK_ENDPOINT=http://localhost:4566 

aws-profile:
ifeq (,$(wildcard ~/.aws))
	echo "test\ntest\nus-east-1\njson\n" | aws configure --profile default
endif

localstack: aws-profile
	mkdir -p ${TMPDIR}
	TMPDIR=${TMPDIR} docker-compose up -d

build: main.go
	GOOS=${GOOS} go build -o ${BINARY} main.go

zip:
	zip ${ZIP} ${BINARY}

create-lambda: build zip
	aws lambda create-function \
		--endpoint ${LOCALSTACK_ENDPOINT} \
		--function-name lambda-example \
		--zip-file fileb://${ZIP} \
		--handler main \
		--runtime go1.x \
		--role admin

run: localstack create-lambda

stop:
	$(eval CONTAINER_ID := $(shell docker ps | grep localstack | cut -d " " -f 1))
	docker stop $(CONTAINER_ID)
	docker rm   $(CONTAINER_ID)
