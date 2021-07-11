build:
	docker build -t grabity:latest .

dev:
	docker build -t grabity:latest . && docker run --env-file .env -it --rm -v ${CURDIR}:/usr/src/app/go/src/github.com/mrauer/grabity grabity:latest && docker exec -it grabity:latest

binary:
	env GOOS=linux GOARCH=amd64 go build -i -o grabity
