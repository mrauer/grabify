build:
	docker build -t grabify:latest .

dev:
	docker build -t grabify:latest . && docker run --env-file .env -it --rm -v ${CURDIR}:/usr/src/app/go/src/github.com/mrauer/grabify grabify:latest && docker exec -it grabify:latest

binary:
	env GOOS=linux GOARCH=amd64 go build -i -o grabify
