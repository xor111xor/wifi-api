package collect

import (
	"os"
	"strconv"
	"strings"
)

func GetTraffic() (uint, error) {
	var result string

	out, err := os.ReadFile("/proc/net/dev")
	if err != nil {
		return 0, err
	}

	str := strings.Fields(string(out))

	for i, x := range str {
		if strings.HasPrefix(x, "wlan") {
			result = str[i+1]
		}
	}

	getByte, err := strconv.Atoi(result)
	if err != nil {
		return 0, err
	}
	mb := getByte / (1024 * 1024)

	return uint(mb), nil
}
