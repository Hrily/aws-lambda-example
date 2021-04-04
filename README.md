# Lambda Example

## Run on Local(stack)

```bash
make run
```

## Test

```bash
aws --endpoint-url=http://localhost:4566 lambda invoke --function-name lambda-example --payload '{
    "What is your name?": "Jim",
    "How old are you?": 33
}' /dev/stdout
```

## Stop Local(stack)

```bash
make stop
```
