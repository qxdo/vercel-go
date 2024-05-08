package aprs_passcode

import "strings"

// aprsPass 函数根据提供的呼号（callsign）计算一个哈希值
func AprsPass(callsign string) (uint16, string) {
	// 查找 '-' 字符的位置，如果找到则截断呼号
	//fmt.Println("callSign:", callsign)
	stopHere := strings.IndexByte(callsign, '-')
	if stopHere >= 0 {
		callsign = callsign[:stopHere]
	}
	// 将呼号转换为大写
	realCall := strings.ToUpper(callsign)

	//fmt.Println("realCall:", realCall)
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
