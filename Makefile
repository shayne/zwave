.PHONY: build deploy run
.DEFAULT: default

default: build deploy run

build:
	@echo Building
	@CXX=arm-none-linux-gnueabi-g++ CC=arm-none-linux-gnueabi-gcc GOOS=linux GOARCH=arm CGO_ENABLED=1 go build

deploy:
	@echo Deploying
	@scp zwave raspberrypi:

run:
	@echo Running
	@ssh -t raspberrypi ./zwave
