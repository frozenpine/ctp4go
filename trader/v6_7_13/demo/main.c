#include "../td_api_helper.h"
#include "../td_spi_helper.h"

void DemoOnFrontConnected(void *this)
{
    fprintf(stdout, "%p front connected", this);

    struct CThostFtdcReqAuthenticateField auth = {0};
    memcpy(auth.BrokerID, "5100", sizeof(TThostFtdcBrokerIDType) - 1);
    memcpy(auth.UserID, "000008", sizeof(TThostFtdcUserIDType) - 1);
    memcpy(auth.AppID, "client_dtprobe_1.0.0", sizeof(TThostFtdcAppIDType) - 1);
    memcpy(auth.AuthCode, "xxx", sizeof(TThostFtdcAuthCodeType) - 1);

    CThostFtdcTraderApiExt *api = ((CThostFtdcTraderSpiExt *)this)->spi;

    api->vtable->CThostFtdcTraderApiVTable_ReqAuthenticate(api, &auth, 1);
}

void DemoOnRspAuthenticate(
    void *this,
    struct CThostFtdcRspAuthenticateField *pRspAuthenticateField,
    struct CThostFtdcRspInfoField *pRspInfo,
    int nRequestID, bool bIsLast)
{
    fprintf(stdout, "[%d] %s.%s authenticate: %s\n",
            pRspInfo->ErrorID,
            pRspAuthenticateField->BrokerID,
            pRspAuthenticateField->UserID,
            pRspAuthenticateField->AppID);

    if (pRspInfo->ErrorID != 0)
    {
        return;
    }

    struct CThostFtdcReqUserLoginField login = {0};
    memcpy(login.BrokerID, "5100", sizeof(TThostFtdcBrokerIDType) - 1);
    memcpy(login.UserID, "000008", sizeof(TThostFtdcUserIDType) - 1);
    memcpy(login.Password, "xxx", sizeof(TThostFtdcPasswordType) - 1);

    CThostFtdcTraderApiExt *api = ((CThostFtdcTraderSpiExt *)this)->spi;

    api->vtable->CThostFtdcTraderApiVTable_ReqUserLogin(api, &login, nRequestID + 1);
}

void DemoOnRspUserLogin(
    void *this,
    struct CThostFtdcRspUserLoginField *pRspUserLogin,
    struct CThostFtdcRspInfoField *pRspInfo,
    int nRequestID, bool bIsLast)
{
    fprintf(stdout, "[%d] %s.%s login: %d\n",
            pRspInfo->ErrorID,
            pRspUserLogin->BrokerID,
            pRspUserLogin->UserID,
            pRspUserLogin->FrontID);
}

int main(int argc, char *argv[])
{
    if (argc < 2)
    {
        fprintf(stderr, "not enough args: %d\n", argc);
        return -1;
    }

    void *handler = dlopen(argv[1], RTLD_LAZY);
    if (handler == NULL)
    {
        fprintf(stderr, "open lib failed: %s\n", dlerror());
        return -1;
    }

    CreateFtdcTraderApi creator = dlsym(handler, "_ZN19CThostFtdcTraderApi19CreateFtdcTraderApiEPKcb");
    if (creator == NULL)
    {
        fprintf(stderr, "find creator failed: %s", dlerror());
        return -2;
    }

    GetApiVersion apiGetter = dlsym(handler, "_ZN19CThostFtdcTraderApi13GetApiVersionEv");
    if (apiGetter != NULL)
    {
        const char *version = apiGetter();
        fprintf(stdout, "%s\n", version);
    }

    CThostFtdcTraderApiExt *instance = creator(argv[2], true);

    CThostFtdcTraderSpiVTable *spiVt = malloc(sizeof(CThostFtdcTraderSpiVTable));
    memset(spiVt, 0, sizeof(CThostFtdcTraderSpiVTable));
    spiVt->CThostFtdcTraderSpi_OnFrontConnected = DemoOnFrontConnected;
    spiVt->CThostFtdcTraderSpi_OnRspAuthenticate = DemoOnRspAuthenticate;

    CThostFtdcTraderSpiExt *spi = malloc(sizeof(CThostFtdcTraderSpiExt));
    spi->vtable = spiVt;
    spi->spi = instance;

    fprintf(stderr, "api[%p], spi[%p]\n", instance, spi);

    instance->vtable->CThostFtdcTraderApiVTable_RegisterSpi(instance, spi);

    instance->vtable->CThostFtdcTraderApiVTable_SubscribePrivateTopic(instance, 2, 1);
    instance->vtable->CThostFtdcTraderApiVTable_SubscribePublicTopic(instance, 2);
    instance->vtable->CThostFtdcTraderApiVTable_RegisterFront(instance, "tcp://180.166.6.245:51205");

    instance->vtable->CThostFtdcTraderApiVTable_Init(instance);

    fprintf(stderr, "press any key to EXIT\n");
    getchar();

    instance->vtable->CThostFtdcTraderApiVTable_Release(instance);
    dlclose(handler);
}