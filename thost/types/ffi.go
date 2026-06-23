package types

import (
	"bytes"
	"encoding/hex"
	"io"
	"log/slog"
	"strings"
	"unsafe"

	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

var decoder = simplifiedchinese.GB18030.NewDecoder()

func RawBytes(ptr unsafe.Pointer, len int) []byte {
	return unsafe.Slice((*byte)(ptr), len)
}

func SetCString(buff []byte, v string) int {
	size := copy(buff, ([]byte)(v))
	buff[len(buff)-1] = 0
	return size
}

func IsASCII(buff []byte) bool {
	for _, b := range buff {
		if b > 127 {
			return false
		}
	}

	return true
}

func ShadowString(buff []byte) string {
	idx := bytes.IndexByte(buff, 0)

	if idx <= 0 {
		return ""
	}

	return strings.Repeat("*", idx+1)
}

func DecodeGBK(buff []byte) string {
	idx := bytes.IndexByte(buff, 0)

	if idx == 0 {
		return ""
	}

	if idx < 0 {
		return hex.EncodeToString(buff)
	}

	if IsASCII(buff[:idx]) {
		return string(buff[:idx])
	}

	reader := transform.NewReader(bytes.NewReader(buff[:idx]), decoder)
	if decoded, err := io.ReadAll(reader); err != nil {
		slog.Error(
			"decode GB18030 failed",
			slog.Any("error", err),
			slog.Any("buff", buff),
		)
		return ""
	} else {
		return string(decoded)
	}
}
