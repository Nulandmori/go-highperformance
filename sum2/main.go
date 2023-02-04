package sum2

import (
	"os"
	"strconv"
)

func Sum2(fileName string) (ret int64, _ error) {
	b, err := os.ReadFile(fileName)
	if err != nil {
		return 0, err
	}

	var last int
	for i := 0; i < len(b); i++ {
		if b[i] != '\n' {
			continue
		}
		num, err := strconv.ParseInt(string(b[last:i]), 10, 64)
		if err != nil {
			return 0, err
		}

		ret += num
		last = i + 1
	}
	return ret, nil
}
