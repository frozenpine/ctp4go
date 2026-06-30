
#pragma once
#ifndef MDUSER_API_HELPER_H
#define MDUSER_API_HELPER_H

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
// converted from ..\..\dependencies\future\v6.5.1\ThostFtdcMdApi.h
// function call ptr
typedef void (*Release)(void *this);

typedef void (*Init)(void *this);

typedef int (*Join)(void *this);

typedef char* (*GetTradingDay)(void *this);

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


// api methods all vtable
typedef struct
{
    Release CThostFtdcMdApiVTable_Release;

    Init CThostFtdcMdApiVTable_Init;

    Join CThostFtdcMdApiVTable_Join;

    GetTradingDay CThostFtdcMdApiVTable_GetTradingDay;

    RegisterFront CThostFtdcMdApiVTable_RegisterFront;

    RegisterNameServer CThostFtdcMdApiVTable_RegisterNameServer;

    RegisterFensUserInfo CThostFtdcMdApiVTable_RegisterFensUserInfo;

    RegisterSpi CThostFtdcMdApiVTable_RegisterSpi;

    SubscribeMarketData CThostFtdcMdApiVTable_SubscribeMarketData;

    UnSubscribeMarketData CThostFtdcMdApiVTable_UnSubscribeMarketData;

    SubscribeForQuoteRsp CThostFtdcMdApiVTable_SubscribeForQuoteRsp;

    UnSubscribeForQuoteRsp CThostFtdcMdApiVTable_UnSubscribeForQuoteRsp;

    ReqUserLogin CThostFtdcMdApiVTable_ReqUserLogin;

    ReqUserLogout CThostFtdcMdApiVTable_ReqUserLogout;

    ReqQryMulticastInstrument CThostFtdcMdApiVTable_ReqQryMulticastInstrument;

} CThostFtdcMdApiVTable;

// cpp api instance C wrapper
typedef struct
{
    CThostFtdcMdApiVTable *vtable;
} CThostFtdcMdApiExt;

// static functions to create cpp api instance & get version

typedef void* (*CreateFtdcMdApi)(char *pszFlowPath, const bool bIsUsingUdp, const bool bIsMulticast);
void* CallCreateFtdcMdApi(CreateFtdcMdApi fn, char *pszFlowPath, const bool bIsUsingUdp, const bool bIsMulticast);

typedef char* (*GetApiVersion)();
char* CallGetApiVersion(GetApiVersion fn);


// helper function call for cgo
void CallRelease(Release fn, void *this);

void CallInit(Init fn, void *this);

int CallJoin(Join fn, void *this);

char* CallGetTradingDay(GetTradingDay fn, void *this);

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