LOCALSTACK_TMPDIR=./localstack-temp
LOCALSTACK_ENDPOINT=http://localhost:4566 
AWS=aws --endpoint ${LOCALSTACK_ENDPOINT} --cli-read-timeout 0 --cli-connect-timeout 0
BUCKET_NAME=webpages
BUS_NAME=webpages-bus
RULE_NAME=webpages-event-rule
STREAM_NAME=webpages-upload-stream

export # variables

create-aws-profile:
ifeq (,$(wildcard ~/.aws))
	echo "test\ntest\nus-east-1\njson\n" | aws configure --profile default
endif

start-localstack: create-aws-profile
	mkdir -p ${LOCALSTACK_TMPDIR}
	TMPDIR=${LOCALSTACK_TMPDIR} docker-compose up -d

create-bucket:
	${AWS} s3api create-bucket --bucket ${BUCKET_NAME}

create-scraper-lambda:
	make -C scraper all

create-eps-lambda:
	make -C eps all

create-eventbridge:
	make -C eventbridge-rules all

create-kinesis-stream:
	make -C kinesis/stream all

stop:
	$(eval CONTAINER_ID := $(shell docker ps | grep localstack | cut -d " " -f 1 ))
	[[ ! -z "$(CONTAINER_ID)" ]] && docker stop $(CONTAINER_ID) || true
	[[ ! -z "$(CONTAINER_ID)" ]] && docker rm   $(CONTAINER_ID) || true

run: \
	start-localstack \
	create-bucket \
	create-scraper-lambda \
	create-eventbridge \
	create-kinesis-stream \
	create-eps-lambda
