package sum5

import (
	"bufio"
	"os"

	"github.com/Nulandmori/go-highperformance/sum4"
	"github.com/efficientgo/core/errcapture"
)

func Sum5(fileName string) (ret int64, err error) {
	f, err := os.Open(fileName)
	if err != nil {
		return 0, err
	}
	defer errcapture.Do(&err, f.Close, "close file")

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		num, err := sum4.ParseInt(scanner.Bytes())
		if err != nil {
			return 0, err
		}

		ret += num
	}
	return ret, scanner.Err()
}
