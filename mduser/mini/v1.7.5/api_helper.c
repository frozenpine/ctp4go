#include "api_helper.h"

void* CallCreateFtdcMdApi(CreateFtdcMdApi fn, char* pszFlowPath, const bool bIsUsingUdp, const bool bIsMulticast)
{
    return fn(pszFlowPath, bIsUsingUdp, bIsMulticast);
}

char* CallGetApiVersion(GetApiVersion fn)
{
    return fn();
}

void CallRelease(Release fn, void *this)
{
    return fn(this);
}

void CallInit(Init fn, void *this, bool bContinuousm)
{
    return fn(this, bContinuousm);
}

int CallJoin(Join fn, void *this)
{
    return fn(this);
}

char* CallGetTradingDay(GetTradingDay fn, void *this)
{
    return fn(this);
}

void CallRegisterFront(RegisterFront fn, void *this, char* pszFrontAddress)
{
    return fn(this, pszFrontAddress);
}

void CallRegisterSpi(RegisterSpi fn, void *this, void *pSpi)
{
    return fn(this, pSpi);
}

int CallSubscribeMarketData(SubscribeMarketData fn, void *this, char* ppInstrumentID[], int nCount)
{
    return fn(this, ppInstrumentID, nCount);
}

int CallUnSubscribeMarketData(UnSubscribeMarketData fn, void *this, char* ppInstrumentID[], int nCount)
{
    return fn(this, ppInstrumentID, nCount);
}

int CallSubscribeForQuoteRsp(SubscribeForQuoteRsp fn, void *this, char* ppInstrumentID[], int nCount)
{
    return fn(this, ppInstrumentID, nCount);
}

int CallUnSubscribeForQuoteRsp(UnSubscribeForQuoteRsp fn, void *this, char* ppInstrumentID[], int nCount)
{
    return fn(this, ppInstrumentID, nCount);
}

int CallReqUserLogin(ReqUserLogin fn, void *this, struct CThostFtdcReqUserLoginField* pReqUserLoginField, int nRequestID)
{
    return fn(this, pReqUserLoginField, nRequestID);
}

int CallReqUserLoginEncrypt(ReqUserLoginEncrypt fn, void *this, struct CThostFtdcReqUserLoginField* pReqUserLoginField, int nRequestID)
{
    return fn(this, pReqUserLoginField, nRequestID);
}

int CallReqUserLogout(ReqUserLogout fn, void *this, struct CThostFtdcUserLogoutField* pUserLogout, int nRequestID)
{
    return fn(this, pUserLogout, nRequestID);
}
