#pragma once

#ifndef MDUSER_V6713_SPI_HELPER_H
#define MDUSER_V6713_SPI_HELPER_H

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

    typedef struct
    {
        OnFrontConnected CThostFtdcMduserSpi_OnFrontConnected;
        OnFrontDisconnected CThostFtdcMduserSpi_OnFrontDisconnected;
        OnHeartBeatWarning CThostFtdcMduserSpi_OnHeartBeatWarning;
        OnRspUserLogin CThostFtdcMduserSpi_OnRspUserLogin;
        OnRspUserLogout CThostFtdcMduserSpi_OnRspUserLogout;
        OnRspQryMulticastInstrument CThostFtdcMduserSpi_OnRspQryMulticastInstrument;
        OnRspError CThostFtdcMduserSpi_OnRspError;
        OnRspSubMarketData CThostFtdcMduserSpi_OnRspSubMarketData;
        OnRspUnSubMarketData CThostFtdcMduserSpi_OnRspUnSubMarketData;
        OnRspSubForQuoteRsp CThostFtdcMduserSpi_OnRspSubForQuoteRsp;
        OnRspUnSubForQuoteRsp CThostFtdcMduserSpi_OnRspUnSubForQuoteRsp;
        OnRtnDepthMarketData CThostFtdcMduserSpi_OnRtnDepthMarketData;
        OnRtnForQuoteRsp CThostFtdcMduserSpi_OnRtnForQuoteRsp;
    } CThostFtdcMduserSpiVTable;

    typedef struct
    {
        CThostFtdcMduserSpiVTable *vtable;
        void *spi;
    } CThostFtdcMduserSpiExt;

#ifdef __cplusplus
}
#endif

#endif