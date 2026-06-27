#pragma once

#ifndef MDUSER_V6713_API_HELPER_H
#define MDUSER_V6713_API_HELPER_H

#ifdef __cplusplus
extern "C"
{
#endif

#include <stdbool.h>
#include <stdint.h>
#include <stdlib.h>
#include <stdio.h>
#include <string.h>
#include <dlfcn.h>

#include "ThostFtdcUserApiStruct.h"

    typedef void (*Release)(void *this);

    typedef void (*Init)(void *this);

    typedef int (*Join)(void *this);

    typedef const char *(*GetTradingDay)(void *this);

    typedef void (*RegisterFront)(void *this, char *pszFrontAddress);

    typedef void (*RegisterNameServer)(void *this, char *pszNsAddress);

    typedef void (*RegisterFensUserInfo)(void *this, struct CThostFtdcFensUserInfoField *pFensUserInfo);

    typedef void (*RegisterSpi)(void *this, void *pSpi);

    typedef int (*SubscribeMarketData)(void *this, char *ppInstrumentID[], int nCount);

    typedef int (*UnSubscribeMarketData)(void *this, char *ppInstrumentID[], int nCount);

    typedef int (*SubscribeForQuoteRsp)(void *this, char *ppInstrumentID[], int nCount);

    typedef int (*UnSubscribeForQuoteRsp)(void *this, char *ppInstrumentID[], int nCount);

    typedef int (*ReqUserLogin)(void *this, struct CThostFtdcReqUserLoginField *pReqUserLoginField, int nRequestID);

    typedef int (*ReqUserLogout)(void *this, struct CThostFtdcUserLogoutField *pUserLogout, int nRequestID);

    typedef int (*ReqQryMulticastInstrument)(void *this, struct CThostFtdcQryMulticastInstrumentField *pQryMulticastInstrument, int nRequestID);

    typedef struct
    {

        Release CThostFtdcMduserApiVTable_Release;
        Init CThostFtdcMduserApiVTable_Init;
        Join CThostFtdcMduserApiVTable_Join;
        GetTradingDay CThostFtdcMduserApiVTable_GetTradingDay;
        RegisterFront CThostFtdcMduserApiVTable_RegisterFront;
        RegisterNameServer CThostFtdcMduserApiVTable_RegisterNameServer;
        RegisterFensUserInfo CThostFtdcMduserApiVTable_RegisterFensUserInfo;
        RegisterSpi CThostFtdcMduserApiVTable_RegisterSpi;
        SubscribeMarketData CThostFtdcMduserApiVTable_SubscribeMarketData;
        UnSubscribeMarketData CThostFtdcMduserApiVTable_UnSubscribeMarketData;
        SubscribeForQuoteRsp CThostFtdcMduserApiVTable_SubscribeForQuoteRsp;
        UnSubscribeForQuoteRsp CThostFtdcMduserApiVTable_UnSubscribeForQuoteRsp;
        ReqUserLogin CThostFtdcMduserApiVTable_ReqUserLogin;
        ReqUserLogout CThostFtdcMduserApiVTable_ReqUserLogout;
        ReqQryMulticastInstrument CThostFtdcMduserApiVTable_ReqQryMulticastInstrument;
    } CThostFtdcMduserApiVTable;

    typedef struct
    {
        CThostFtdcMduserApiVTable *vtable;
    } CThostFtdcMduserApiExt;

    typedef void *(*CreateFtdcMdApi)(const char *pszFlowPath, const bool bIsUsingUdp, const bool bIsMulticast, bool bIsProductionMode);

    typedef const char *(*GetApiVersion)();

    void *CallCreateFtdcMdApi(CreateFtdcMdApi fn, const char *pszFlowPath, const bool bIsUsingUdp, const bool bIsMulticast, bool bIsProductionMode);

    const char *CallGetApiVersion(GetApiVersion fn);

    void CallRelease(Release fn, void *this);

    void CallInit(Init fn, void *this);

    int CallJoin(Join fn, void *this);

    const char *CallGetTradingDay(GetTradingDay fn, void *this);

    void CallRegisterFront(RegisterFront fn, void *this, char *pszFrontAddress);

    void CallRegisterNameServer(RegisterNameServer fn, void *this, char *pszNsAddress);

    void CallRegisterFensUserInfo(RegisterFensUserInfo fn, void *this, struct CThostFtdcFensUserInfoField *pFensUserInfo);

    void CallRegisterSpi(RegisterSpi fn, void *this, void *pSpi);

    int CallSubscribeMarketData(SubscribeMarketData fn, void *this, char *ppInstrumentID[], int nCount);

    int CallUnSubscribeMarketData(UnSubscribeMarketData fn, void *this, char *ppInstrumentID[], int nCount);

    int CallSubscribeForQuoteRsp(SubscribeForQuoteRsp fn, void *this, char *ppInstrumentID[], int nCount);

    int CallUnSubscribeForQuoteRsp(UnSubscribeForQuoteRsp fn, void *this, char *ppInstrumentID[], int nCount);

    int CallReqUserLogin(ReqUserLogin fn, void *this, struct CThostFtdcReqUserLoginField *pReqUserLoginField, int nRequestID);

    int CallReqUserLogout(ReqUserLogout fn, void *this, struct CThostFtdcUserLogoutField *pUserLogout, int nRequestID);

    int CallReqQryMulticastInstrument(ReqQryMulticastInstrument fn, void *this, struct CThostFtdcQryMulticastInstrumentField *pQryMulticastInstrument, int nRequestID);

#ifdef __cplusplus
}
#endif
#endif