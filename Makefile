IMAGE ?= berqnet/updater
TAG   ?= latest

.PHONY: test build release digest clean

test:
	go test ./...

build:
	go build -o bin/updater .

release:
	docker build -t $(IMAGE):$(TAG) .

digest:
	@docker inspect --format='{{.Id}}' $(IMAGE):$(TAG)

size:
	@docker images $(IMAGE):$(TAG) --format '{{.Size}}'

clean:
	rm -rf bin
