package v1_7_5

/*
#cgo CFLAGS: -I. -I${SRCDIR} -I${SRCDIR}/../../../dependencies/mini/v1.7.5/
#cgo LDFLAGS: -ldl

#include "spi_helper.h"
*/
import "C"
import (
	"context"
	"log/slog"
	"runtime"
	"unsafe"

	"github.com/frozenpine/ctp4go/thost/mini"
)

var (
	// 全局回调函数虚表
	spiCVtablePtr *C.CThostFtdcMdSpiVTable
)

func init() {
	// C端为虚表分配内存
	spiCVtablePtr = (*C.CThostFtdcMdSpiVTable)(C.malloc(
		C.sizeof_CThostFtdcMdSpiVTable))

	spiCVtablePtr.CThostFtdcMdSpiVTable_OnFrontConnected = (C.OnFrontConnected)(
		unsafe.Pointer(C.COnFrontConnected))

	spiCVtablePtr.CThostFtdcMdSpiVTable_OnFrontDisconnected = (C.OnFrontDisconnected)(
		unsafe.Pointer(C.COnFrontDisconnected))

	spiCVtablePtr.CThostFtdcMdSpiVTable_OnHeartBeatWarning = (C.OnHeartBeatWarning)(
		unsafe.Pointer(C.COnHeartBeatWarning))

	spiCVtablePtr.CThostFtdcMdSpiVTable_OnRspUserLogin = (C.OnRspUserLogin)(
		unsafe.Pointer(C.COnRspUserLogin))

	spiCVtablePtr.CThostFtdcMdSpiVTable_OnRspUserLogout = (C.OnRspUserLogout)(
		unsafe.Pointer(C.COnRspUserLogout))

	spiCVtablePtr.CThostFtdcMdSpiVTable_OnRspError = (C.OnRspError)(
		unsafe.Pointer(C.COnRspError))

	spiCVtablePtr.CThostFtdcMdSpiVTable_OnRspSubMarketData = (C.OnRspSubMarketData)(
		unsafe.Pointer(C.COnRspSubMarketData))

	spiCVtablePtr.CThostFtdcMdSpiVTable_OnRspUnSubMarketData = (C.OnRspUnSubMarketData)(
		unsafe.Pointer(C.COnRspUnSubMarketData))

	spiCVtablePtr.CThostFtdcMdSpiVTable_OnRspSubForQuoteRsp = (C.OnRspSubForQuoteRsp)(
		unsafe.Pointer(C.COnRspSubForQuoteRsp))

	spiCVtablePtr.CThostFtdcMdSpiVTable_OnRspUnSubForQuoteRsp = (C.OnRspUnSubForQuoteRsp)(
		unsafe.Pointer(C.COnRspUnSubForQuoteRsp))

	spiCVtablePtr.CThostFtdcMdSpiVTable_OnRtnDepthMarketData = (C.OnRtnDepthMarketData)(
		unsafe.Pointer(C.COnRtnDepthMarketData))

	spiCVtablePtr.CThostFtdcMdSpiVTable_OnRtnMBLMarketData = (C.OnRtnMBLMarketData)(
		unsafe.Pointer(C.COnRtnMBLMarketData))

	spiCVtablePtr.CThostFtdcMdSpiVTable_OnRtnForQuoteRsp = (C.OnRtnForQuoteRsp)(
		unsafe.Pointer(C.COnRtnForQuoteRsp))

}

type ThostFtdcMdSpi struct {
	runtime.Pinner
	callback mini.MdSpi
}

//export CgoOnFrontConnected
func CgoOnFrontConnected(
	this unsafe.Pointer,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnFrontConnected called",
		slog.Any("this", this),
	)

	(*ThostFtdcMdSpi)(
		(*C.CThostFtdcMdSpiExt)(this).spi,
	).callback.OnFrontConnected()
}

//export CgoOnFrontDisconnected
func CgoOnFrontDisconnected(
	this unsafe.Pointer,
	nReason C.int,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnFrontDisconnected called",
		slog.Any("this", this),
	)

	(*ThostFtdcMdSpi)(
		(*C.CThostFtdcMdSpiExt)(this).spi,
	).callback.OnFrontDisconnected(
		int(nReason),
	)
}

//export CgoOnHeartBeatWarning
func CgoOnHeartBeatWarning(
	this unsafe.Pointer,
	nTimeLapse C.int,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnHeartBeatWarning called",
		slog.Any("this", this),
	)

	(*ThostFtdcMdSpi)(
		(*C.CThostFtdcMdSpiExt)(this).spi,
	).callback.OnHeartBeatWarning(
		int(nTimeLapse),
	)
}

//export CgoOnRspUserLogin
func CgoOnRspUserLogin(
	this unsafe.Pointer,
	pRspUserLogin *C.struct_CThostFtdcRspUserLoginField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int,
	bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspUserLogin called",
		slog.Any("this", this),
	)

	(*ThostFtdcMdSpi)(
		(*C.CThostFtdcMdSpiExt)(this).spi,
	).callback.OnRspUserLogin(
		(*mini.CThostFtdcRspUserLoginField)(unsafe.Pointer(pRspUserLogin)),
		(*mini.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID),
		bool(bIsLast),
	)
}

//export CgoOnRspUserLogout
func CgoOnRspUserLogout(
	this unsafe.Pointer,
	pUserLogout *C.struct_CThostFtdcUserLogoutField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int,
	bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspUserLogout called",
		slog.Any("this", this),
	)

	(*ThostFtdcMdSpi)(
		(*C.CThostFtdcMdSpiExt)(this).spi,
	).callback.OnRspUserLogout(
		(*mini.CThostFtdcUserLogoutField)(unsafe.Pointer(pUserLogout)),
		(*mini.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID),
		bool(bIsLast),
	)
}

//export CgoOnRspError
func CgoOnRspError(
	this unsafe.Pointer,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int,
	bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspError called",
		slog.Any("this", this),
	)

	(*ThostFtdcMdSpi)(
		(*C.CThostFtdcMdSpiExt)(this).spi,
	).callback.OnRspError(
		(*mini.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID),
		bool(bIsLast),
	)
}

//export CgoOnRspSubMarketData
func CgoOnRspSubMarketData(
	this unsafe.Pointer,
	pSpecificInstrument *C.struct_CThostFtdcSpecificInstrumentField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int,
	bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspSubMarketData called",
		slog.Any("this", this),
	)

	(*ThostFtdcMdSpi)(
		(*C.CThostFtdcMdSpiExt)(this).spi,
	).callback.OnRspSubMarketData(
		(*mini.CThostFtdcSpecificInstrumentField)(unsafe.Pointer(pSpecificInstrument)),
		(*mini.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID),
		bool(bIsLast),
	)
}

//export CgoOnRspUnSubMarketData
func CgoOnRspUnSubMarketData(
	this unsafe.Pointer,
	pSpecificInstrument *C.struct_CThostFtdcSpecificInstrumentField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int,
	bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspUnSubMarketData called",
		slog.Any("this", this),
	)

	(*ThostFtdcMdSpi)(
		(*C.CThostFtdcMdSpiExt)(this).spi,
	).callback.OnRspUnSubMarketData(
		(*mini.CThostFtdcSpecificInstrumentField)(unsafe.Pointer(pSpecificInstrument)),
		(*mini.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID),
		bool(bIsLast),
	)
}

//export CgoOnRspSubForQuoteRsp
func CgoOnRspSubForQuoteRsp(
	this unsafe.Pointer,
	pSpecificInstrument *C.struct_CThostFtdcSpecificInstrumentField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int,
	bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspSubForQuoteRsp called",
		slog.Any("this", this),
	)

	(*ThostFtdcMdSpi)(
		(*C.CThostFtdcMdSpiExt)(this).spi,
	).callback.OnRspSubForQuoteRsp(
		(*mini.CThostFtdcSpecificInstrumentField)(unsafe.Pointer(pSpecificInstrument)),
		(*mini.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID),
		bool(bIsLast),
	)
}

//export CgoOnRspUnSubForQuoteRsp
func CgoOnRspUnSubForQuoteRsp(
	this unsafe.Pointer,
	pSpecificInstrument *C.struct_CThostFtdcSpecificInstrumentField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int,
	bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspUnSubForQuoteRsp called",
		slog.Any("this", this),
	)

	(*ThostFtdcMdSpi)(
		(*C.CThostFtdcMdSpiExt)(this).spi,
	).callback.OnRspUnSubForQuoteRsp(
		(*mini.CThostFtdcSpecificInstrumentField)(unsafe.Pointer(pSpecificInstrument)),
		(*mini.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID),
		bool(bIsLast),
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

	(*ThostFtdcMdSpi)(
		(*C.CThostFtdcMdSpiExt)(this).spi,
	).callback.OnRtnDepthMarketData(
		(*mini.CThostFtdcDepthMarketDataField)(unsafe.Pointer(pDepthMarketData)),
	)
}

//export CgoOnRtnMBLMarketData
func CgoOnRtnMBLMarketData(
	this unsafe.Pointer,
	pMBLMarketData *C.struct_CThostFtdcMBLMarketDataField,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRtnMBLMarketData called",
		slog.Any("this", this),
	)

	(*ThostFtdcMdSpi)(
		(*C.CThostFtdcMdSpiExt)(this).spi,
	).callback.OnRtnMBLMarketData(
		(*mini.CThostFtdcMBLMarketDataField)(unsafe.Pointer(pMBLMarketData)),
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

	(*ThostFtdcMdSpi)(
		(*C.CThostFtdcMdSpiExt)(this).spi,
	).callback.OnRtnForQuoteRsp(
		(*mini.CThostFtdcForQuoteRspField)(unsafe.Pointer(pForQuoteRsp)),
	)
}
