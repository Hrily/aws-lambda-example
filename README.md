# Lambda Example

## Run on Local(stack)

```bash
make run
```

## Test

```bash
aws --endpoint-url=http://localhost:4566 lambda invoke --function-name scraper --payload '{
        "company": {
                "name": "relianceindustries",
                "symbol": "RI"
        }
}' /dev/stdout
```

Check webpage uploaded on s3:

```bash
aws --endpoint-url=http://localhost:4566 s3 cp s3://webpages/relianceindustries/RI/1 /dev/stdout
```

Check kinesis records:

```bash
# Get shard iterator
aws --endpoint-url=http://localhost:4566 kinesis get-shard-iterator --shard-id shardId-000000000000 --shard-iterator-type TRIM_HORIZON --stream-name webpages-upload-stream
# Get records
aws --endpoint-url=http://localhost:4566 kinesis get-records --shard-iterator "<shard-iterator-from-above-command-output>"
```

Check DynamoDB records:

```bash
aws --endpoint-url=http://localhost:4566 dynamodb scan --table-name companies
aws --endpoint-url=http://localhost:4566 dynamodb scan --table-name financials
```


## Stop Local(stack)

```bash
make stop
```
