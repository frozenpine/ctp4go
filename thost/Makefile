VERSION ?= v6.7.13
PLATFORM ?= future

default:
	go run ../gen/main.go -dep ../dependencies -plat $(PLATFORM) -sdk hdr=ThostFtdcUserApiStruct.h,ver=$(VERSION) -output thost
	cd types && $(RM) thost*.go && gofmt -w . && go generate .