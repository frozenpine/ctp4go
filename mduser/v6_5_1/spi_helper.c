#include "spi_helper.h"

extern void CgoOnFrontConnected(void *this);

extern void CgoOnFrontDisconnected(void *this, int nReason);

extern void CgoOnHeartBeatWarning(void *this, int nTimeLapse);

extern void CgoOnRspUserLogin(void *this, struct CThostFtdcRspUserLoginField *pRspUserLogin, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

extern void CgoOnRspUserLogout(void *this, struct CThostFtdcUserLogoutField *pUserLogout, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

extern void CgoOnRspQryMulticastInstrument(void *this, struct CThostFtdcMulticastInstrumentField *pMulticastInstrument, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

extern void CgoOnRspError(void *this, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

extern void CgoOnRspSubMarketData(void *this, struct CThostFtdcSpecificInstrumentField *pSpecificInstrument, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

extern void CgoOnRspUnSubMarketData(void *this, struct CThostFtdcSpecificInstrumentField *pSpecificInstrument, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

extern void CgoOnRspSubForQuoteRsp(void *this, struct CThostFtdcSpecificInstrumentField *pSpecificInstrument, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

extern void CgoOnRspUnSubForQuoteRsp(void *this, struct CThostFtdcSpecificInstrumentField *pSpecificInstrument, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

extern void CgoOnRtnDepthMarketData(void *this, struct CThostFtdcDepthMarketDataField *pDepthMarketData);

extern void CgoOnRtnForQuoteRsp(void *this, struct CThostFtdcForQuoteRspField *pForQuoteRsp);

void COnFrontConnected(void *this)
{
    return CgoOnFrontConnected(this);
}

void COnFrontDisconnected(void *this, int nReason)
{
    return CgoOnFrontDisconnected(this, nReason);
}

void COnHeartBeatWarning(void *this, int nTimeLapse)
{
    return CgoOnHeartBeatWarning(this, nTimeLapse);
}

void COnRspUserLogin(void *this, struct CThostFtdcRspUserLoginField *pRspUserLogin, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
{
    return CgoOnRspUserLogin(this, pRspUserLogin, pRspInfo, nRequestID, bIsLast);
}

void COnRspUserLogout(void *this, struct CThostFtdcUserLogoutField *pUserLogout, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
{
    return CgoOnRspUserLogout(this, pUserLogout, pRspInfo, nRequestID, bIsLast);
}

void COnRspQryMulticastInstrument(void *this, struct CThostFtdcMulticastInstrumentField *pMulticastInstrument, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
{
    return CgoOnRspQryMulticastInstrument(this, pMulticastInstrument, pRspInfo, nRequestID, bIsLast);
}

void COnRspError(void *this, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
{
    return CgoOnRspError(this, pRspInfo, nRequestID, bIsLast);
}

void COnRspSubMarketData(void *this, struct CThostFtdcSpecificInstrumentField *pSpecificInstrument, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
{
    return CgoOnRspSubMarketData(this, pSpecificInstrument, pRspInfo, nRequestID, bIsLast);
}

void COnRspUnSubMarketData(void *this, struct CThostFtdcSpecificInstrumentField *pSpecificInstrument, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
{
    return CgoOnRspUnSubMarketData(this, pSpecificInstrument, pRspInfo, nRequestID, bIsLast);
}

void COnRspSubForQuoteRsp(void *this, struct CThostFtdcSpecificInstrumentField *pSpecificInstrument, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
{
    return CgoOnRspSubForQuoteRsp(this, pSpecificInstrument, pRspInfo, nRequestID, bIsLast);
}

void COnRspUnSubForQuoteRsp(void *this, struct CThostFtdcSpecificInstrumentField *pSpecificInstrument, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
{
    return CgoOnRspUnSubForQuoteRsp(this, pSpecificInstrument, pRspInfo, nRequestID, bIsLast);
}

void COnRtnDepthMarketData(void *this, struct CThostFtdcDepthMarketDataField *pDepthMarketData)
{
    return CgoOnRtnDepthMarketData(this, pDepthMarketData);
}

void COnRtnForQuoteRsp(void *this, struct CThostFtdcForQuoteRspField *pForQuoteRsp)
{
    return CgoOnRtnForQuoteRsp(this, pForQuoteRsp);
}
