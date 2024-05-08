package api

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

var srv http.Handler

func init() {
	e := echo.New()
	e.GET("/books", Test)
	e.GET("/aprs_passcode", Test)
	srv = e
}

type APRSCallSign struct {
	CallSign string `json:"call_sign" form:"call_sign" query:"call_sign"`
}

// aprsPass 函数根据提供的呼号（callsign）计算一个哈希值
func AprsPass(callsign string) (uint16, string) {
	// 查找 '-' 字符的位置，如果找到则截断呼号
	fmt.Println("callSign:", callsign)
	stopHere := strings.IndexByte(callsign, '-')
	if stopHere >= 0 {
		callsign = callsign[:stopHere]
	}
	// 将呼号转换为大写
	realCall := strings.ToUpper(callsign)

	fmt.Println("realCall:", realCall)
	hash := uint32(0x73e2)
	for i, str := range realCall {
		leftBit := 0
		if i%2 == 0 {
			leftBit = 8
		}
		hash ^= uint32(str) << leftBit
		//fmt.Println(str)
	}

	return uint16(hash & 0x7fff), realCall
}
func Test(c echo.Context) error {
	defer func() {
		if r := recover(); r != nil {
			c.Logger().Info("client resp error", r)
			c.String(200, "data error, please contact BH4FWA use Wechat.")
			return
		}
	}()
	var data = APRSCallSign{}
	err := c.Bind(&data)
	if err != nil {
		c.String(200, "data error, please contact BH4FWA use Wechat.")
	}
	passcode, realCall := AprsPass(data.CallSign)
	return c.String(200, "InputCallSign:"+data.CallSign+"Calc CallSign:"+realCall+" APRS Pass Code:"+fmt.Sprintf("%d", passcode))
}

func MainFunc(w http.ResponseWriter, r *http.Request) {
	srv.ServeHTTP(w, r)
}
