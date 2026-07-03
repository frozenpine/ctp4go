VERSION ?= v6.7.13
PLATFORM ?= future

mduser:
	@go run ../gen/main.go -dep ../dependencies -plat $(PLATFORM) -sdk name=mduser,ver=$(VERSION) -output thost
	@cd .. && cd mduser && cd $(VERSION) && go generate ./...

trader:
	@go run ../gen/main.go -dep ../dependencies -plat $(PLATFORM) -sdk name=trader,ver=$(VERSION) -output thost
	@cd .. && cd trader && cd $(VERSION) && go generate ./...