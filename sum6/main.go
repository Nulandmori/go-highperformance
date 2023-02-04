package sum6

import (
	"io"
	"os"

	"github.com/Nulandmori/go-highperformance/sum4"
	"github.com/efficientgo/core/errcapture"
)

func Sum6(fileName string) (ret int64, err error) {
	f, err := os.Open(fileName)
	if err != nil {
		return 0, err
	}
	defer errcapture.Do(&err, f.Close, "close file")

	buf := make([]byte, 8*1024)
	return Sum6Reader(f, buf)
}

func Sum6Reader(r io.Reader, buf []byte) (ret int64, err error) {
	var offset, n int
	for err != io.EOF {
		n, err = r.Read(buf[offset:])
		if err != nil && err != io.EOF {
			return 0, err
		}
		n += offset

		var last int
		for i := range buf[:n] {
			if buf[i] != '\n' {
				continue
			}
			num, err := sum4.ParseInt(buf[last:i])
			if err != nil {
				return 0, err
			}

			ret += num
			last = i + 1
		}

		offset = n - last
		if offset > 0 {
			_ = copy(buf, buf[last:n])
		}
	}
	return ret, nil
}
