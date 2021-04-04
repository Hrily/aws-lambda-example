# Lambda Example

## Run on Local(stack)

```bash
make run
```

## Test

```bash
aws --endpoint-url=http://localhost:4566 lambda invoke --function-name lambda-example --payload '{
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

## Stop Local(stack)

```bash
make stop
```
