create-event-bus:
	${AWS} events create-event-bus \
		--name "${BUS_NAME}"

create-source-rule: create-event-bus
	${AWS} events put-rule \
		--name "${RULE_NAME}" \
		--event-bus-name "${BUS_NAME}"

all: create-source-rule
