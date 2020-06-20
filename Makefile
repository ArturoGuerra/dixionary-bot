.PHONY: all build clean docker docker-build docker-push

GOBUILD = go build
DOCKER = docker
APPNAME = dixionarybot

all: clean build

clean:
	rm -rf ./bin

build: clean
	$(GOBUILD) -o bin/$(APPNAME) main.go
	cp -f ./assets/* ./bin

docker-test:
	test $(DOCKERREPO)

docker-build: docker-test
	$(DOCKER) build . -t $(DOCKERREPO)

docker-push: docker-test
	$(DOCKER) push $(DOCKERREPO)

docker: docker-build docker-push