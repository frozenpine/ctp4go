package v6_7_13

/*
#cgo CFLAGS: -I. -I${SRCDIR} -I${SRCDIR}/../../dependencies/future/v6.7.13/
#cgo LDFLAGS: -ldl

#include "mduser_spi_helper.h"
*/
import "C"

import (
	"runtime"
	"unsafe"

	"github.com/frozenpine/ctp4go/thost"
)

var (
	spiCVtablePtr *C.CThostFtdcMduserSpiVTable
)

func init() {
	spiCVtablePtr = (*C.CThostFtdcMduserSpiVTable)(C.malloc(
		C.sizeof_CThostFtdcMduserSpiVTable))

	spiCVtablePtr.CThostFtdcMduserSpi_OnFrontConnected = (C.OnFrontConnected)(
		unsafe.Pointer(C.COnFrontConnected))
	spiCVtablePtr.CThostFtdcMduserSpi_OnFrontDisconnected = (C.OnFrontDisconnected)(
		unsafe.Pointer(C.COnFrontDisconnected))
	spiCVtablePtr.CThostFtdcMduserSpi_OnHeartBeatWarning = (C.OnHeartBeatWarning)(
		unsafe.Pointer(C.COnHeartBeatWarning))
	spiCVtablePtr.CThostFtdcMduserSpi_OnRspUserLogin = (C.OnRspUserLogin)(
		unsafe.Pointer(C.COnRspUserLogin))
	spiCVtablePtr.CThostFtdcMduserSpi_OnRspUserLogout = (C.OnRspUserLogout)(
		unsafe.Pointer(C.COnRspUserLogout))
	spiCVtablePtr.CThostFtdcMduserSpi_OnRspQryMulticastInstrument = (C.OnRspQryMulticastInstrument)(
		unsafe.Pointer(C.COnRspQryMulticastInstrument))
	spiCVtablePtr.CThostFtdcMduserSpi_OnRspError = (C.OnRspError)(
		unsafe.Pointer(C.COnRspError))
	spiCVtablePtr.CThostFtdcMduserSpi_OnRspSubMarketData = (C.OnRspSubMarketData)(
		unsafe.Pointer(C.COnRspSubMarketData))
	spiCVtablePtr.CThostFtdcMduserSpi_OnRspUnSubMarketData = (C.OnRspUnSubMarketData)(
		unsafe.Pointer(C.COnRspUnSubMarketData))
	spiCVtablePtr.CThostFtdcMduserSpi_OnRspSubForQuoteRsp = (C.OnRspSubForQuoteRsp)(
		unsafe.Pointer(C.COnRspSubForQuoteRsp))
	spiCVtablePtr.CThostFtdcMduserSpi_OnRspUnSubForQuoteRsp = (C.OnRspUnSubForQuoteRsp)(
		unsafe.Pointer(C.COnRspUnSubForQuoteRsp))
	spiCVtablePtr.CThostFtdcMduserSpi_OnRtnDepthMarketData = (C.OnRtnDepthMarketData)(
		unsafe.Pointer(C.COnRtnDepthMarketData))
	spiCVtablePtr.CThostFtdcMduserSpi_OnRtnForQuoteRsp = (C.OnRtnForQuoteRsp)(
		unsafe.Pointer(C.COnRtnForQuoteRsp))
}

type ThostFtdcMduserSpi struct {
	runtime.Pinner
	callback thost.MdSpi
}

//export CgoOnFrontConnected
func CgoOnFrontConnected(this unsafe.Pointer) {}

//export CgoOnFrontDisconnected
func CgoOnFrontDisconnected(this unsafe.Pointer, nReason C.int) {}

//export CgoOnHeartBeatWarning
func CgoOnHeartBeatWarning(this unsafe.Pointer, nTimeLapse C.int) {}

//export CgoOnRspUserLogin
func CgoOnRspUserLogin(
	this unsafe.Pointer,
	pRspUserLogin *C.struct_CThostFtdcRspUserLoginField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int, bIsLast C.bool,
) {
}

//export CgoOnRspUserLogout
func CgoOnRspUserLogout(
	this unsafe.Pointer,
	pUserLogout *C.struct_CThostFtdcUserLogoutField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int, bIsLast C.bool,
) {
}

//export CgoOnRspQryMulticastInstrument
func CgoOnRspQryMulticastInstrument(
	this unsafe.Pointer,
	pMulticastInstrument *C.struct_CThostFtdcMulticastInstrumentField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int, bIsLast C.bool,
) {
}

//export CgoOnRspError
func CgoOnRspError(
	this unsafe.Pointer,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int, bIsLast C.bool,
) {
}

//export CgoOnRspSubMarketData
func CgoOnRspSubMarketData(
	this unsafe.Pointer,
	pSpecificInstrument *C.struct_CThostFtdcSpecificInstrumentField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int, bIsLast C.bool,
) {
}

//export CgoOnRspUnSubMarketData
func CgoOnRspUnSubMarketData(
	this unsafe.Pointer,
	pSpecificInstrument *C.struct_CThostFtdcSpecificInstrumentField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int, bIsLast C.bool,
) {
}

//export CgoOnRspSubForQuoteRsp
func CgoOnRspSubForQuoteRsp(
	this unsafe.Pointer,
	pSpecificInstrument *C.struct_CThostFtdcSpecificInstrumentField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int, bIsLast C.bool,
) {
}

//export CgoOnRspUnSubForQuoteRsp
func CgoOnRspUnSubForQuoteRsp(
	this unsafe.Pointer,
	pSpecificInstrument *C.struct_CThostFtdcSpecificInstrumentField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int, bIsLast C.bool,
) {
}

//export CgoOnRtnDepthMarketData
func CgoOnRtnDepthMarketData(
	this unsafe.Pointer,
	pDepthMarketData *C.struct_CThostFtdcDepthMarketDataField,
) {
}

//export CgoOnRtnForQuoteRsp
func CgoOnRtnForQuoteRsp(
	this unsafe.Pointer,
	pForQuoteRsp *C.struct_CThostFtdcForQuoteRspField,
) {
}
