docker:
	docker build -t kesshabot .
linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/kessha-amd64 *.go

