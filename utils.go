package ctp4go

import (
	"bytes"
	"encoding/hex"
	"io"
	"log/slog"
	"math/rand/v2"
	"runtime"
	"strings"
	"sync"
	"unsafe"

	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

const (
	// 6 bits to represent a letter index
	letterIdBits = 6
	// All 1-bits as many as letterIdBits
	letterIdMask = 1<<letterIdBits - 1
	letterIdMax  = 63 / letterIdBits
)

var (
	createBuilder = func() *strings.Builder {
		buf := new(strings.Builder)
		buf.Grow(200)
		return buf
	}

	builderPool = sync.Pool{New: func() any { return createBuilder() }}
)

func RandomStr(n int) string {
	b := make([]byte, n)
	// A rand.Int63() generates 63 random bits, enough for letterIdMax letters!
	for i, cache, remain := n-1, rand.Uint64(), letterIdMax; i >= 0; {
		if remain == 0 {
			cache, remain = rand.Uint64(), letterIdMax
		}
		if idx := int(cache & letterIdMask); idx < len(letters) {
			b[i] = letters[idx]
			i--
		}
		cache >>= letterIdBits
		remain--
	}
	return *(*string)(unsafe.Pointer(&b))
}

func GetStringBuilder() *strings.Builder {
	var (
		buf *strings.Builder
		ok  bool
	)

	if v := builderPool.Get(); v != nil {
		buf = createBuilder()
	} else if buf, ok = v.(*strings.Builder); !ok {
		buf = createBuilder()
	}

	runtime.SetFinalizer(buf, func(v *strings.Builder) {
		v.Reset()
		builderPool.Put(v)
	})

	return buf
}

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

func IsAlphabet(buff []byte) bool {
	for _, b := range buff {
		if !(b >= '0' && b <= '9') &&
			!(b >= 'a' && b <= 'z') &&
			!(b >= 'A' && b <= 'Z') {
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

// NextPowerOfTwo 向上取整到最近的 2 的幂
func NextPowerOfTwo(n uint32) uint32 {
	if n == 0 {
		return 1
	}
	n--
	n |= n >> 1
	n |= n >> 2
	n |= n >> 4
	n |= n >> 8
	n |= n >> 16
	return n + 1
}
