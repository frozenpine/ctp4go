package types

import "strconv"

type Double float64

func (d Double) String() string {
	return strconv.FormatFloat(float64(d), 'f', 6, 64)
}

type Byte2 [2]byte

func (s *Byte2) SetString(v string) int {
	return SetCString(([]byte)((*s)[:]), v)
}
func (s Byte2) String() string {
	return DecodeGBK(s[:])
}

type Byte3 [3]byte

func (s *Byte3) SetString(v string) int {
	return SetCString(([]byte)((*s)[:]), v)
}
func (s Byte3) String() string {
	return DecodeGBK(s[:])
}

type Byte4 [4]byte

func (s *Byte4) SetString(v string) int {
	return SetCString(([]byte)((*s)[:]), v)
}
func (s Byte4) String() string {
	return DecodeGBK(s[:])
}

type Byte5 [5]byte

func (s *Byte5) SetString(v string) int {
	return SetCString(([]byte)((*s)[:]), v)
}
func (s Byte5) String() string {
	return DecodeGBK(s[:])
}

type Byte6 [6]byte

func (s *Byte6) SetString(v string) int {
	return SetCString(([]byte)((*s)[:]), v)
}
func (s Byte6) String() string {
	return DecodeGBK(s[:])
}

type Byte7 [7]byte

func (s *Byte7) SetString(v string) int {
	return SetCString(([]byte)((*s)[:]), v)
}
func (s Byte7) String() string {
	return DecodeGBK(s[:])
}

type Byte9 [9]byte

func (s *Byte9) SetString(v string) int {
	return SetCString(([]byte)((*s)[:]), v)
}
func (s Byte9) String() string {
	return DecodeGBK(s[:])
}

type Byte10 [10]byte

func (s *Byte10) SetString(v string) int {
	return SetCString(([]byte)((*s)[:]), v)
}
func (s Byte10) String() string {
	return DecodeGBK(s[:])
}

type Byte11 [11]byte

func (s *Byte11) SetString(v string) int {
	return SetCString(([]byte)((*s)[:]), v)
}
func (s Byte11) String() string {
	return DecodeGBK(s[:])
}

type Byte12 [12]byte

func (s *Byte12) SetString(v string) int {
	return SetCString(([]byte)((*s)[:]), v)
}
func (s Byte12) String() string {
	return DecodeGBK(s[:])
}

type Byte13 [13]byte

func (s *Byte13) SetString(v string) int {
	return SetCString(([]byte)((*s)[:]), v)
}
func (s Byte13) String() string {
	return DecodeGBK(s[:])
}

type Byte15 [15]byte

func (s *Byte15) SetString(v string) int {
	return SetCString(([]byte)((*s)[:]), v)
}
func (s Byte15) String() string {
	return DecodeGBK(s[:])
}

type Byte16 [16]byte

func (s *Byte16) SetString(v string) int {
	return SetCString(([]byte)((*s)[:]), v)
}
func (s Byte16) String() string {
	return DecodeGBK(s[:])
}

type Byte17 [17]byte

func (s *Byte17) SetString(v string) int {
	return SetCString(([]byte)((*s)[:]), v)
}
func (s Byte17) String() string {
	return DecodeGBK(s[:])
}

type Byte20 [20]byte

func (s *Byte20) SetString(v string) int {
	return SetCString(([]byte)((*s)[:]), v)
}
func (s Byte20) String() string {
	return DecodeGBK(s[:])
}

type Byte21 [21]byte

func (s *Byte21) SetString(v string) int {
	return SetCString(([]byte)((*s)[:]), v)
}
func (s Byte21) String() string {
	return DecodeGBK(s[:])
}

type Byte22 [22]byte

func (s *Byte22) SetString(v string) int {
	return SetCString(([]byte)((*s)[:]), v)
}
func (s Byte22) String() string {
	return DecodeGBK(s[:])
}

type Byte23 [23]byte

func (s *Byte23) SetString(v string) int {
	return SetCString(([]byte)((*s)[:]), v)
}
func (s Byte23) String() string {
	return DecodeGBK(s[:])
}

type Byte24 [24]byte

func (s *Byte24) SetString(v string) int {
	return SetCString(([]byte)((*s)[:]), v)
}
func (s Byte24) String() string {
	return DecodeGBK(s[:])
}

type Byte25 [25]byte

func (s *Byte25) SetString(v string) int {
	return SetCString(([]byte)((*s)[:]), v)
}
func (s Byte25) String() string {
	return DecodeGBK(s[:])
}

type Byte31 [31]byte

func (s *Byte31) SetString(v string) int {
	return SetCString(([]byte)((*s)[:]), v)
}
func (s Byte31) String() string {
	return DecodeGBK(s[:])
}

type Byte33 [33]byte

func (s *Byte33) SetString(v string) int {
	return SetCString(([]byte)((*s)[:]), v)
}
func (s Byte33) String() string {
	return DecodeGBK(s[:])
}

type Byte36 [36]byte

func (s *Byte36) SetString(v string) int {
	return SetCString(([]byte)((*s)[:]), v)
}
func (s Byte36) String() string {
	return DecodeGBK(s[:])
}

type Byte41 [41]byte

func (s *Byte41) SetString(v string) int {
	return SetCString(([]byte)((*s)[:]), v)
}
func (s Byte41) String() string {
	return DecodeGBK(s[:])
}

type Byte51 [51]byte

func (s *Byte51) SetString(v string) int {
	return SetCString(([]byte)((*s)[:]), v)
}
func (s Byte51) String() string {
	return DecodeGBK(s[:])
}

type Byte61 [61]byte

func (s *Byte61) SetString(v string) int {
	return SetCString(([]byte)((*s)[:]), v)
}
func (s Byte61) String() string {
	return DecodeGBK(s[:])
}

type Byte65 [65]byte

func (s *Byte65) SetString(v string) int {
	return SetCString(([]byte)((*s)[:]), v)
}
func (s Byte65) String() string {
	return DecodeGBK(s[:])
}

type Byte71 [71]byte

func (s *Byte71) SetString(v string) int {
	return SetCString(([]byte)((*s)[:]), v)
}
func (s Byte71) String() string {
	return DecodeGBK(s[:])
}

type Byte81 [81]byte

func (s *Byte81) SetString(v string) int {
	return SetCString(([]byte)((*s)[:]), v)
}
func (s Byte81) String() string {
	return DecodeGBK(s[:])
}

type Byte100 [100]byte

func (s *Byte100) SetString(v string) int {
	return SetCString(([]byte)((*s)[:]), v)
}
func (s Byte100) String() string {
	return DecodeGBK(s[:])
}

type Byte101 [101]byte

func (s *Byte101) SetString(v string) int {
	return SetCString(([]byte)((*s)[:]), v)
}
func (s Byte101) String() string {
	return DecodeGBK(s[:])
}

type Byte129 [129]byte

func (s *Byte129) SetString(v string) int {
	return SetCString(([]byte)((*s)[:]), v)
}
func (s Byte129) String() string {
	return DecodeGBK(s[:])
}

type Byte151 [151]byte

func (s *Byte151) SetString(v string) int {
	return SetCString(([]byte)((*s)[:]), v)
}
func (s Byte151) String() string {
	return DecodeGBK(s[:])
}

type Byte161 [161]byte

func (s *Byte161) SetString(v string) int {
	return SetCString(([]byte)((*s)[:]), v)
}
func (s Byte161) String() string {
	return DecodeGBK(s[:])
}

type Byte201 [201]byte

func (s *Byte201) SetString(v string) int {
	return SetCString(([]byte)((*s)[:]), v)
}
func (s Byte201) String() string {
	return DecodeGBK(s[:])
}

type Byte256 [256]byte

func (s *Byte256) SetString(v string) int {
	return SetCString(([]byte)((*s)[:]), v)
}
func (s Byte256) String() string {
	return DecodeGBK(s[:])
}

type Byte257 [257]byte

func (s *Byte257) SetString(v string) int {
	return SetCString(([]byte)((*s)[:]), v)
}
func (s Byte257) String() string {
	return DecodeGBK(s[:])
}

type Byte261 [261]byte

func (s *Byte261) SetString(v string) int {
	return SetCString(([]byte)((*s)[:]), v)
}
func (s Byte261) String() string {
	return DecodeGBK(s[:])
}

type Byte273 [273]byte

func (s *Byte273) SetString(v string) int {
	return SetCString(([]byte)((*s)[:]), v)
}
func (s Byte273) String() string {
	return DecodeGBK(s[:])
}

type Byte301 [301]byte

func (s *Byte301) SetString(v string) int {
	return SetCString(([]byte)((*s)[:]), v)
}
func (s Byte301) String() string {
	return DecodeGBK(s[:])
}

type Byte349 [349]byte

func (s *Byte349) SetString(v string) int {
	return SetCString(([]byte)((*s)[:]), v)
}
func (s Byte349) String() string {
	return DecodeGBK(s[:])
}

type Byte365 [365]byte

func (s *Byte365) SetString(v string) int {
	return SetCString(([]byte)((*s)[:]), v)
}
func (s Byte365) String() string {
	return DecodeGBK(s[:])
}

type Byte401 [401]byte

func (s *Byte401) SetString(v string) int {
	return SetCString(([]byte)((*s)[:]), v)
}
func (s Byte401) String() string {
	return DecodeGBK(s[:])
}

type Byte501 [501]byte

func (s *Byte501) SetString(v string) int {
	return SetCString(([]byte)((*s)[:]), v)
}
func (s Byte501) String() string {
	return DecodeGBK(s[:])
}

type Byte513 [513]byte

func (s *Byte513) SetString(v string) int {
	return SetCString(([]byte)((*s)[:]), v)
}
func (s Byte513) String() string {
	return DecodeGBK(s[:])
}

type Byte1001 [1001]byte

func (s *Byte1001) SetString(v string) int {
	return SetCString(([]byte)((*s)[:]), v)
}
func (s Byte1001) String() string {
	return DecodeGBK(s[:])
}

type Byte1025 [1025]byte

func (s *Byte1025) SetString(v string) int {
	return SetCString(([]byte)((*s)[:]), v)
}
func (s Byte1025) String() string {
	return DecodeGBK(s[:])
}

type Byte2049 [2049]byte

func (s *Byte2049) SetString(v string) int {
	return SetCString(([]byte)((*s)[:]), v)
}
func (s Byte2049) String() string {
	return DecodeGBK(s[:])
}

type Byte2561 [2561]byte

func (s *Byte2561) SetString(v string) int {
	return SetCString(([]byte)((*s)[:]), v)
}
func (s Byte2561) String() string {
	return DecodeGBK(s[:])
}
