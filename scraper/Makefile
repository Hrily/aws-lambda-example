GOOS=linux
BINARY=main
ZIP=lambda-example.zip

clean:
	rm ${ZIP} ${BINARY}

build-scraper: *.go
	GOOS=${GOOS} go build -o ${BINARY} .

create-zip:
	zip -j ${ZIP} ${BINARY}

create-lambda: build-scraper create-zip
	${AWS} lambda create-function \
		--function-name lambda-example \
		--zip-file fileb://${ZIP} \
		--timeout 900 \
		--handler main \
		--runtime go1.x \
		--role admin
