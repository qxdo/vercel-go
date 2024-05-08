package api

import (
	"fmt"
	"github.com/qxdo/vercel-go/aprs_passcode"
	"testing"
)

func TestAPRSCode(t *testing.T) {

	fmt.Println(aprs_passcode.AprsPass("BH4FW"))

}
