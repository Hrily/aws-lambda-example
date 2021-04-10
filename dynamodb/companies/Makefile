COMPANIES_TABLE=companies

define ATTRIBUTE_DEF
'[
	{
			"AttributeName": "Symbol",
			"AttributeType": "S"
	}
]'
endef

define KEY_SCHEMA
'[
	{
			"AttributeName": "Symbol",
			"KeyType": "HASH"
	}
]'
endef

create-companies-table:
	${AWS} dynamodb create-table \
		--table-name ${COMPANIES_TABLE} \
		--attribute-definitions $(shell echo \'${ATTRIBUTE_DEF}\' | tr -d "\n") \
		--key-schema $(shell echo \'${KEY_SCHEMA}\' | tr -d "\n") \
		--provisioned-throughput '{ "ReadCapacityUnits": 1, "WriteCapacityUnits": 1 }'

all: create-companies-table
