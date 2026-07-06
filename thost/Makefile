VERSION ?= v6.7.13
PLATFORM ?= future
GEN_TYPES ?= 0

mduser:
	go run ../gen/main.go -dep ../dependencies -plat $(PLATFORM) -sdk hdr=ThostFtdcUserApiStruct.h,ver=$(VERSION) -output thost
	cd $(PLATFORM) && gofmt -w . && go build .
	cd $(PLATFORM)/types && gofmt -w . && go build .
	if [ $(GEN_TYPES) -eq 1 ]; then cd $(PLATFORM)/types && go generate .; fi

clean:
	$(RM) -r $(PLATFORM)/types
	$(RM) $(PLATFORM)/ctp_structs.go