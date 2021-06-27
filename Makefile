### Bootstrapping with Booty ;-)
APP=booty
VERSION=$(shell cat VERSION)

all: clean build

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-X booty/booty.BuildVersion=$(VERSION)" .

run:
	go run main.go

clean:
	rm -rf booty
