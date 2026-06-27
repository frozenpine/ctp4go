#include "mduser_api_helper.h"

void *CallCreateFtdcMdApi(CreateFtdcMdApi fn, const char *pszFlowPath, const bool bIsUsingUdp, const bool bIsMulticast, bool bIsProductionMode)
{
    return fn(pszFlowPath, bIsUsingUdp, bIsMulticast, bIsProductionMode);
}

const char *CallGetApiVersion(GetApiVersion fn)
{
    return fn();
}

void CallRelease(Release fn, void *this)
{
    return fn(this);
}

void CallInit(Init fn, void *this)
{
    return fn(this);
}

int CallJoin(Join fn, void *this)
{
    return fn(this);
}

const char *CallGetTradingDay(GetTradingDay fn, void *this)
{
    return fn(this);
}

void CallRegisterFront(RegisterFront fn, void *this, char *pszFrontAddress)
{
    return fn(this, pszFrontAddress);
}

void CallRegisterNameServer(RegisterNameServer fn, void *this, char *pszNsAddress)
{
    return fn(this, pszNsAddress);
}

void CallRegisterFensUserInfo(RegisterFensUserInfo fn, void *this, struct CThostFtdcFensUserInfoField *pFensUserInfo)
{
    return fn(this, pFensUserInfo);
}

void CallRegisterSpi(RegisterSpi fn, void *this, void *pSpi)
{
    return fn(this, pSpi);
}

int CallSubscribeMarketData(SubscribeMarketData fn, void *this, char *ppInstrumentID[], int nCount)
{
    return fn(this, ppInstrumentID, nCount);
}

int CallUnSubscribeMarketData(UnSubscribeMarketData fn, void *this, char *ppInstrumentID[], int nCount)
{
    return fn(this, ppInstrumentID, nCount);
}

int CallSubscribeForQuoteRsp(SubscribeForQuoteRsp fn, void *this, char *ppInstrumentID[], int nCount)
{
    return fn(this, ppInstrumentID, nCount);
}

int CallUnSubscribeForQuoteRsp(UnSubscribeForQuoteRsp fn, void *this, char *ppInstrumentID[], int nCount)
{
    return fn(this, ppInstrumentID, nCount);
}

int CallReqUserLogin(ReqUserLogin fn, void *this, struct CThostFtdcReqUserLoginField *pReqUserLoginField, int nRequestID)
{
    return fn(this, pReqUserLoginField, nRequestID);
}

int CallReqUserLogout(ReqUserLogout fn, void *this, struct CThostFtdcUserLogoutField *pUserLogout, int nRequestID)
{
    return fn(this, pUserLogout, nRequestID);
}

int CallReqQryMulticastInstrument(ReqQryMulticastInstrument fn, void *this, struct CThostFtdcQryMulticastInstrumentField *pQryMulticastInstrument, int nRequestID)
{
    return fn(this, pQryMulticastInstrument, nRequestID);
}