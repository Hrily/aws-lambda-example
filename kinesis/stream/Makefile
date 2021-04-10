create-stream:
	${AWS} kinesis create-stream \
		--stream-name ${STREAM_NAME} \
		--shard-count 1

create-target: create-stream
	$(eval ARN := $(shell ${AWS} kinesis describe-stream --stream-name ${STREAM_NAME} | jq .StreamDescription.StreamARN | tr -d '"'))
	${AWS} events put-targets \
		--rule ${RULE_NAME} \
		--event-bus-name "${BUS_NAME}" \
		--targets '{ "Id": "9", "Arn": "${ARN}" }'

all: create-target
