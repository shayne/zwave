.PHONY: build deploy run

build:
	@echo Building
	@CXX=arm-none-linux-gnueabi-g++ CC=arm-none-linux-gnueabi-gcc GOOS=linux GOARCH=arm CGO_ENABLED=1 go build

deploy: build
	@echo Deploying
	@scp zwave raspberrypi:

run: deploy
	@echo Running
	@ssh -t raspberrypi ./zwave
