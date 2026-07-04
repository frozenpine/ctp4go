VERSION ?= v6.7.13
PLATFORM ?= future

default:
	go run ../gen/main.go -dep ../dependencies -plat $(PLATFORM) -sdk name=mduser,ver=$(VERSION) -output api,spi
	cd $(VERSION) && go generate . && gofmt -w . && go build .