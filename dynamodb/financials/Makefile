FINANCIALS_TABLE=financials

define ATTRIBUTE_DEF
'[
	{
			"AttributeName": "Symbol",
			"AttributeType": "S"
	},
	{
			"AttributeName": "Year",
			"AttributeType": "N"
	}
]'
endef

define KEY_SCHEMA
'[
	{
			"AttributeName": "Symbol",
			"KeyType": "HASH"
	},
	{
			"AttributeName": "Year",
			"KeyType": "RANGE"
	}
]'
endef

create-financials-table:
	${AWS} dynamodb create-table \
		--table-name ${FINANCIALS_TABLE} \
		--attribute-definitions $(shell echo \'${ATTRIBUTE_DEF}\' | tr -d "\n") \
		--key-schema $(shell echo \'${KEY_SCHEMA}\' | tr -d "\n") \
		--provisioned-throughput '{ "ReadCapacityUnits": 1, "WriteCapacityUnits": 1 }'

all: create-financials-table
