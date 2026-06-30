#pragma once
#ifndef MDUSER_SPI_HELPER_H
#define MDUSER_SPI_HELPER_H

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
typedef void (*OnFrontConnected)(void *this);

typedef void (*OnFrontDisconnected)(void *this, int nReason);

typedef void (*OnHeartBeatWarning)(void *this, int nTimeLapse);

typedef void (*OnRspUserLogin)(void *this, struct CThostFtdcRspUserLoginField *pRspUserLogin, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

typedef void (*OnRspUserLogout)(void *this, struct CThostFtdcUserLogoutField *pUserLogout, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

typedef void (*OnRspQryMulticastInstrument)(void *this, struct CThostFtdcMulticastInstrumentField *pMulticastInstrument, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

typedef void (*OnRspError)(void *this, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

typedef void (*OnRspSubMarketData)(void *this, struct CThostFtdcSpecificInstrumentField *pSpecificInstrument, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

typedef void (*OnRspUnSubMarketData)(void *this, struct CThostFtdcSpecificInstrumentField *pSpecificInstrument, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

typedef void (*OnRspSubForQuoteRsp)(void *this, struct CThostFtdcSpecificInstrumentField *pSpecificInstrument, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

typedef void (*OnRspUnSubForQuoteRsp)(void *this, struct CThostFtdcSpecificInstrumentField *pSpecificInstrument, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

typedef void (*OnRtnDepthMarketData)(void *this, struct CThostFtdcDepthMarketDataField *pDepthMarketData);

typedef void (*OnRtnForQuoteRsp)(void *this, struct CThostFtdcForQuoteRspField *pForQuoteRsp);


// helper function callback for cgo
void COnFrontConnected(void *this);

void COnFrontDisconnected(void *this, int nReason);

void COnHeartBeatWarning(void *this, int nTimeLapse);

void COnRspUserLogin(void *this, struct CThostFtdcRspUserLoginField *pRspUserLogin, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

void COnRspUserLogout(void *this, struct CThostFtdcUserLogoutField *pUserLogout, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

void COnRspQryMulticastInstrument(void *this, struct CThostFtdcMulticastInstrumentField *pMulticastInstrument, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

void COnRspError(void *this, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

void COnRspSubMarketData(void *this, struct CThostFtdcSpecificInstrumentField *pSpecificInstrument, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

void COnRspUnSubMarketData(void *this, struct CThostFtdcSpecificInstrumentField *pSpecificInstrument, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

void COnRspSubForQuoteRsp(void *this, struct CThostFtdcSpecificInstrumentField *pSpecificInstrument, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

void COnRspUnSubForQuoteRsp(void *this, struct CThostFtdcSpecificInstrumentField *pSpecificInstrument, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

void COnRtnDepthMarketData(void *this, struct CThostFtdcDepthMarketDataField *pDepthMarketData);

void COnRtnForQuoteRsp(void *this, struct CThostFtdcForQuoteRspField *pForQuoteRsp);


// spi callback vtable
typedef struct
{
    OnFrontConnected CThostFtdcMdSpiVTable_OnFrontConnected;

    OnFrontDisconnected CThostFtdcMdSpiVTable_OnFrontDisconnected;

    OnHeartBeatWarning CThostFtdcMdSpiVTable_OnHeartBeatWarning;

    OnRspUserLogin CThostFtdcMdSpiVTable_OnRspUserLogin;

    OnRspUserLogout CThostFtdcMdSpiVTable_OnRspUserLogout;

    OnRspQryMulticastInstrument CThostFtdcMdSpiVTable_OnRspQryMulticastInstrument;

    OnRspError CThostFtdcMdSpiVTable_OnRspError;

    OnRspSubMarketData CThostFtdcMdSpiVTable_OnRspSubMarketData;

    OnRspUnSubMarketData CThostFtdcMdSpiVTable_OnRspUnSubMarketData;

    OnRspSubForQuoteRsp CThostFtdcMdSpiVTable_OnRspSubForQuoteRsp;

    OnRspUnSubForQuoteRsp CThostFtdcMdSpiVTable_OnRspUnSubForQuoteRsp;

    OnRtnDepthMarketData CThostFtdcMdSpiVTable_OnRtnDepthMarketData;

    OnRtnForQuoteRsp CThostFtdcMdSpiVTable_OnRtnForQuoteRsp;

} CThostFtdcMdSpiVTable;

// cpp spi instance C wrapper
typedef struct
{
    CThostFtdcMdSpiVTable *vtable;
    // spi instance ptr in go, ptr must be Pinned
    void *spi;
} CThostFtdcMdSpiExt;

#ifdef __cplusplus
}
#endif

#endif