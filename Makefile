MAIN_CMD = ./main.go
MAIN_BIN = plock
MAIN_BIN_PATH = ./bin/$(MAIN_BIN)

GO_ENV_DARWIN_64 = GOARCH=amd64 GOOS=darwin
GO_ENV_LINUX_64 = GOARCH=amd64 GOOS=linux
GO_ENV_WIN_64 = GOARCH=amd64 GOOS=windows
GO_ENV_ARM_64 = GOARCH=arm64 GOOS=linux GOARM=7

GO_PROD_FLAGS = -ldflags "-s -w"

.DEFAULT_GOAL := build
.PHONY: clean

build:
	$(GO_ENV_DARWIN_64) go build -o ./bin/darwin/$(MAIN_BIN) $(MAIN_CMD)
	$(GO_ENV_LINUX_64) go build -o ./bin/linux/$(MAIN_BIN) $(MAIN_CMD)
	$(GO_ENV_WIN_64) go build -o ./bin/win/$(MAIN_BIN).exe $(MAIN_CMD)
	$(GO_ENV_ARM_64) go build $(GO_PROD_FLAGS) -o ./bin/arm/$(MAIN_BIN) $(MAIN_CMD)

clean:
	go clean
	rm ./bin/darwin/$(MAIN_BIN)
	rm ./bin/linux/$(MAIN_BIN)
	rm ./bin/win/$(MAIN_BIN).exe
	rm ./bin/arm/$(MAIN_BIN)
