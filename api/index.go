package api

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/qxdo/vercel-go/aprs_passcode"
	"net/http"
	"time"
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

	passcode, realCall := aprs_passcode.AprsPass(data.CallSign)
	beijingLocation, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		panic(err) // 处理错误
	}
	beijingTime := time.Now().In(beijingLocation).Format(time.DateTime)
	returnStr := "Beijing Time: " + beijingTime + ", InputCallSign:" + data.CallSign + ",  Calc CallSign:" + realCall + ", APRS Pass Code:" + fmt.Sprintf("%d", passcode)
	fmt.Println(returnStr)
	return c.String(200, returnStr)
}

func MainFunc(w http.ResponseWriter, r *http.Request) {
	srv.ServeHTTP(w, r)
}
