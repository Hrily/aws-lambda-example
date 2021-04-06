GOOS=linux
BINARY=main
ZIP=lambda-example.zip

clean:
	rm ${ZIP} ${BINARY}

build: *.go
	GOOS=${GOOS} go build -o ${BINARY} .

create-zip:
	zip -j ${ZIP} ${BINARY}

create-lambda: build create-zip
	${AWS} lambda create-function \
		--function-name scraper \
		--zip-file fileb://${ZIP} \
		--timeout 900 \
		--handler main \
		--runtime go1.x \
		--role admin

all: create-lambda
