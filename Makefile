LOCALSTACK_TMPDIR=./localstack-temp
LOCALSTACK_ENDPOINT=http://localhost:4566 
AWS=aws --endpoint ${LOCALSTACK_ENDPOINT}
BUCKET=webpages

export # variables

create-aws-profile:
ifeq (,$(wildcard ~/.aws))
	echo "test\ntest\nus-east-1\njson\n" | aws configure --profile default
endif

start-localstack: create-aws-profile
	mkdir -p ${LOCALSTACK_TMPDIR}
	TMPDIR=${LOCALSTACK_TMPDIR} docker-compose up -d

create-bucket:
	${AWS} s3api create-bucket --bucket ${BUCKET}

create-lambda:
	make -C scraper create-lambda

stop:
	$(eval CONTAINER_ID := $(shell docker ps | grep localstack | cut -d " " -f 1))
	docker stop $(CONTAINER_ID)
	docker rm   $(CONTAINER_ID)

run: stop start-localstack create-bucket create-lambda
