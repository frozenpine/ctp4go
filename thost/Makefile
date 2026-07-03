VERSION ?= v6.7.13
PLATFORM ?= future

defines:
	@cd types && $(RM) thost*.go && gofmt -w . && go generate .

mduser:
	@go run ../gen/main.go -dep ../dependencies -plat $(PLATFORM) -sdk name=mduser,ver=$(VERSION) -output thost && \
		gofmt -w . && \
		go build . && \
		cd .. && cd mduser && cd $(VERSION) && go generate ./...

trader:
	@go run ../gen/main.go -dep ../dependencies -plat $(PLATFORM) -sdk name=trader,ver=$(VERSION) -output thost && \
		gofmt -w . && \
		go build . && \
		cd .. && cd trader && cd $(VERSION) && go generate ./...