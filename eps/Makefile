GOOS=linux
BINARY=main
ZIP=lambda-example.zip

clean:
	rm ${ZIP} ${BINARY}

build: *.go
	GOOS=${GOOS} CGO_ENABLED=0 go build -o ${BINARY} .

create-zip:
	zip -j ${ZIP} ${BINARY}

create-lambda: build create-zip
	${AWS} lambda create-function \
		--function-name eps \
		--zip-file fileb://${ZIP} \
		--timeout 900 \
		--handler main \
		--runtime go1.x \
		--role admin

create-event-source-mapping:
	$(eval ARN := $(shell ${AWS} kinesis describe-stream --stream-name ${STREAM_NAME} | jq .StreamDescription.StreamARN | tr -d '"'))
	${AWS} lambda create-event-source-mapping \
		--function-name eps \
		--batch-size 1 \
		--starting-position TRIM_HORIZON \
		--event-source-arn "$(ARN)"

all: create-lambda create-event-source-mapping
