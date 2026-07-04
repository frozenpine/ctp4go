VERSION ?= v6.7.13
PLATFORM ?= future

mduser:
	go run ../gen/main.go -dep ../dependencies -plat $(PLATFORM) -sdk hdr=ThostFtdcUserApiStruct.h,ver=$(VERSION) -output thost
	gofmt -w .
	cd types && gofmt -w . && go generate .

clean:
	$(RM) -r types
	$(RM) ctp_structs.go