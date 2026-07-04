VERSION ?= v6.7.13
PLATFORM ?= future

mduser:
	go run ../gen/main.go -dep ../dependencies -plat $(PLATFORM) -sdk hdr=ThostFtdcUserApiStruct.h,ver=$(VERSION) -output thost
	cd $(PLATFORM) && gofmt -w .
	cd $(PLATFORM)/types && gofmt -w . && go generate .

clean:
	$(RM) -r $(PLATFORM)/types
	$(RM) $(PLATFORM)/ctp_structs.go