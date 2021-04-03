# Lambda Example

## Build

```bash
docker build -t lambda-example .
```

## Run on Local

```bash
docker run -p 9000:8080 lambda-example:latest /main
```

## Test

```bash
curl -XPOST "http://localhost:9000/2015-03-31/functions/function/invocations" -d '
{
	"What is your name?": "Jim",
	"How old are you?": 33
}'
```
