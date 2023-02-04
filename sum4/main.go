package sum4

import (
	"os"

	"github.com/efficientgo/core/errors"
)

func ParseInt(input []byte) (n int64, _ error) {
	factor := int64(1)
	k := 0

	if input[0] == '-' {
		factor *= -1
		k++
	}

	for i := len(input) - 1; i >= k; i-- {
		if input[i] < '0' || input[i] > '9' {
			return 0, errors.Newf("not a valid integer: %v", input)
		}

		n += factor * int64(input[i]-'0')
		factor *= 10
	}
	return n, nil
}

func Sum4(fileName string) (ret int64, err error) {
	b, err := os.ReadFile(fileName)
	if err != nil {
		return 0, err
	}

	var last int
	for i := 0; i < len(b); i++ {
		if b[i] != '\n' {
			continue
		}
		num, err := ParseInt(b[last:i])
		if err != nil {
			return 0, err
		}

		ret += num
		last = i + 1
	}
	return ret, nil
}
