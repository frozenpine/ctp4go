package v6_7_13

/*
#cgo CFLAGS: -I. -I${SRCDIR} -I${SRCDIR}/../../dependencies/future/v6.7.13/
#cgo LDFLAGS: -ldl

#include "md_spi_helper.h"
*/
import "C"

import (
	"context"
	"log/slog"
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
func CgoOnFrontConnected(this unsafe.Pointer) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnFrontConnected called",
		slog.Any("this", this),
	)

	(*ThostFtdcMduserSpi)(
		(*C.CThostFtdcMduserSpiExt)(this).spi,
	).callback.OnFrontConnected()
}

//export CgoOnFrontDisconnected
func CgoOnFrontDisconnected(this unsafe.Pointer, nReason C.int) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnFrontDisconnected called",
		slog.Any("this", this),
	)

	(*ThostFtdcMduserSpi)(
		(*C.CThostFtdcMduserSpiExt)(this).spi,
	).callback.OnFrontDisconnected(int(nReason))
}

//export CgoOnHeartBeatWarning
func CgoOnHeartBeatWarning(this unsafe.Pointer, nTimeLapse C.int) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnHeartBeatWarning called",
		slog.Any("this", this),
	)

	(*ThostFtdcMduserSpi)(
		(*C.CThostFtdcMduserSpiExt)(this).spi,
	).callback.OnHeartBeatWarning(int(nTimeLapse))
}

//export CgoOnRspUserLogin
func CgoOnRspUserLogin(
	this unsafe.Pointer,
	pRspUserLogin *C.struct_CThostFtdcRspUserLoginField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int, bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspUserLogin called",
		slog.Any("this", this),
	)

	(*ThostFtdcMduserSpi)(
		(*C.CThostFtdcMduserSpiExt)(this).spi,
	).callback.OnRspUserLogin(
		(*thost.CThostFtdcRspUserLoginField)(unsafe.Pointer(pRspUserLogin)),
		(*thost.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID), bool(bIsLast),
	)
}

//export CgoOnRspUserLogout
func CgoOnRspUserLogout(
	this unsafe.Pointer,
	pUserLogout *C.struct_CThostFtdcUserLogoutField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int, bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspUserLogout called",
		slog.Any("this", this),
	)

	(*ThostFtdcMduserSpi)(
		(*C.CThostFtdcMduserSpiExt)(this).spi,
	).callback.OnRspUserLogout(
		(*thost.CThostFtdcUserLogoutField)(unsafe.Pointer(pUserLogout)),
		(*thost.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID), bool(bIsLast),
	)
}

//export CgoOnRspQryMulticastInstrument
func CgoOnRspQryMulticastInstrument(
	this unsafe.Pointer,
	pMulticastInstrument *C.struct_CThostFtdcMulticastInstrumentField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int, bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspQryMulticastInstrument called",
		slog.Any("this", this),
	)

	(*ThostFtdcMduserSpi)(
		(*C.CThostFtdcMduserSpiExt)(this).spi,
	).callback.OnRspQryMulticastInstrument(
		(*thost.CThostFtdcMulticastInstrumentField)(
			unsafe.Pointer(pMulticastInstrument)),
		(*thost.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID), bool(bIsLast),
	)
}

//export CgoOnRspError
func CgoOnRspError(
	this unsafe.Pointer,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int, bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspError called",
		slog.Any("this", this),
	)

	(*ThostFtdcMduserSpi)(
		(*C.CThostFtdcMduserSpiExt)(this).spi,
	).callback.OnRspError(
		(*thost.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID), bool(bIsLast),
	)
}

//export CgoOnRspSubMarketData
func CgoOnRspSubMarketData(
	this unsafe.Pointer,
	pSpecificInstrument *C.struct_CThostFtdcSpecificInstrumentField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int, bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspSubMarketData called",
		slog.Any("this", this),
	)

	(*ThostFtdcMduserSpi)(
		(*C.CThostFtdcMduserSpiExt)(this).spi,
	).callback.OnRspSubMarketData(
		(*thost.CThostFtdcSpecificInstrumentField)(
			unsafe.Pointer(pSpecificInstrument)),
		(*thost.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID), bool(bIsLast),
	)
}

//export CgoOnRspUnSubMarketData
func CgoOnRspUnSubMarketData(
	this unsafe.Pointer,
	pSpecificInstrument *C.struct_CThostFtdcSpecificInstrumentField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int, bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspUnSubMarketData called",
		slog.Any("this", this),
	)

	(*ThostFtdcMduserSpi)(
		(*C.CThostFtdcMduserSpiExt)(this).spi,
	).callback.OnRspUnSubMarketData(
		(*thost.CThostFtdcSpecificInstrumentField)(
			unsafe.Pointer(pSpecificInstrument)),
		(*thost.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID), bool(bIsLast),
	)
}

//export CgoOnRspSubForQuoteRsp
func CgoOnRspSubForQuoteRsp(
	this unsafe.Pointer,
	pSpecificInstrument *C.struct_CThostFtdcSpecificInstrumentField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int, bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspSubForQuoteRsp called",
		slog.Any("this", this),
	)

	(*ThostFtdcMduserSpi)(
		(*C.CThostFtdcMduserSpiExt)(this).spi,
	).callback.OnRspSubForQuoteRsp(
		(*thost.CThostFtdcSpecificInstrumentField)(
			unsafe.Pointer(pSpecificInstrument)),
		(*thost.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID), bool(bIsLast),
	)
}

//export CgoOnRspUnSubForQuoteRsp
func CgoOnRspUnSubForQuoteRsp(
	this unsafe.Pointer,
	pSpecificInstrument *C.struct_CThostFtdcSpecificInstrumentField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int, bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspUnSubForQuoteRsp called",
		slog.Any("this", this),
	)

	(*ThostFtdcMduserSpi)(
		(*C.CThostFtdcMduserSpiExt)(this).spi,
	).callback.OnRspUnSubForQuoteRsp(
		(*thost.CThostFtdcSpecificInstrumentField)(
			unsafe.Pointer(pSpecificInstrument)),
		(*thost.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID), bool(bIsLast),
	)
}

//export CgoOnRtnDepthMarketData
func CgoOnRtnDepthMarketData(
	this unsafe.Pointer,
	pDepthMarketData *C.struct_CThostFtdcDepthMarketDataField,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRtnDepthMarketData called",
		slog.Any("this", this),
	)

	(*ThostFtdcMduserSpi)(
		(*C.CThostFtdcMduserSpiExt)(this).spi,
	).callback.OnRtnDepthMarketData(
		(*thost.CThostFtdcDepthMarketDataField)(
			unsafe.Pointer(pDepthMarketData)),
	)
}

//export CgoOnRtnForQuoteRsp
func CgoOnRtnForQuoteRsp(
	this unsafe.Pointer,
	pForQuoteRsp *C.struct_CThostFtdcForQuoteRspField,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRtnForQuoteRsp called",
		slog.Any("this", this),
	)

	(*ThostFtdcMduserSpi)(
		(*C.CThostFtdcMduserSpiExt)(this).spi,
	).callback.OnRtnForQuoteRsp(
		(*thost.CThostFtdcForQuoteRspField)(
			unsafe.Pointer(pForQuoteRsp)),
	)
}
