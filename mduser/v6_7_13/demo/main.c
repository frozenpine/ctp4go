#include "../md_api_helper.h"
#include "../md_spi_helper.h"

#ifdef _WIN32
const static char *CREATOR_FN_NAME = "?CreateFtdcMdApi@CThostFtdcMdApi@@SAPEAV1@PEBD_N1_N@Z";
const static char *VERSION_FN_NAME = "?GetApiVersion@CThostFtdcMdApi@@SAPEBDXZ";
#else
const static char *CREATOR_FN_NAME = "_ZN15CThostFtdcMdApi15CreateFtdcMdApiEPKcbbb";
const static char *VERSION_FN_NAME = "_ZN15CThostFtdcMdApi13GetApiVersionEv";
#endif

void DemoOnFrontConnected(void *this)
{
    fprintf(stdout, "%p front connected\n", this);

    struct CThostFtdcReqUserLoginField login = {0};
    memcpy(login.BrokerID, "5100", sizeof(TThostFtdcBrokerIDType) - 1);
    memcpy(login.UserID, "000008", sizeof(TThostFtdcUserIDType) - 1);
    memcpy(login.Password, "xxx", sizeof(TThostFtdcPasswordType) - 1);

    CThostFtdcMduserApiExt *api = ((CThostFtdcMduserSpiExt *)this)->spi;

    api->vtable->CThostFtdcMduserApiVTable_ReqUserLogin(api, &login, 1);
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

    CreateFtdcMdApi creator = dlsym(handler, CREATOR_FN_NAME);
    if (creator == NULL)
    {
        fprintf(stderr, "find creator failed: %s", dlerror());
        return -2;
    }

    GetApiVersion apiGetter = dlsym(handler, VERSION_FN_NAME);
    if (apiGetter != NULL)
    {
        const char *version = apiGetter();
        fprintf(stdout, "%s\n", version);
    }

    CThostFtdcMduserApiExt *instance = creator(argv[2], false, false, true);

    CThostFtdcMduserSpiVTable *spiVt = malloc(sizeof(CThostFtdcMduserSpiVTable));
    memset(spiVt, 0, sizeof(CThostFtdcMduserSpiVTable));
    spiVt->CThostFtdcMduserSpi_OnFrontConnected = DemoOnFrontConnected;
    spiVt->CThostFtdcMduserSpi_OnRspUserLogin = DemoOnRspUserLogin;

    CThostFtdcMduserSpiExt *spi = malloc(sizeof(CThostFtdcMduserSpiExt));
    spi->vtable = spiVt;
    spi->spi = instance;

    fprintf(stderr, "api[%p], spi[%p]\n", instance, spi);

    instance->vtable->CThostFtdcMduserApiVTable_RegisterSpi(instance, spi);

    instance->vtable->CThostFtdcMduserApiVTable_RegisterFront(instance, "tcp://180.166.6.245:51205");

    instance->vtable->CThostFtdcMduserApiVTable_Init(instance);

    fprintf(stderr, "press any key to EXIT\n");
    getchar();

    instance->vtable->CThostFtdcMduserApiVTable_Release(instance);
    dlclose(handler);
}