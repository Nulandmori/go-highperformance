package concurrentsum4

import (
	"errors"
	"io"
	"os"

	"github.com/Nulandmori/go-highperformance/sum6"
	"github.com/efficientgo/core/errcapture"
)

func ConcurrentSum4(fileName string, workers int) (ret int64, _ error) {
	f, err := os.Open(fileName)
	if err != nil {
		return 0, err
	}
	defer errcapture.Do(&err, f.Close, "close file")

	s, err := f.Stat()
	if err != nil {
		return 0, err
	}

	var (
		size           = int(s.Size())
		bytesPerWorker = size / workers
		resultCh       = make(chan int64)
	)

	if bytesPerWorker < 10 {
		return 0, errors.New("can't have less bytes per goroutine than 10")
	}

	for i := 0; i < workers; i++ {
		go func(i int) {
			begin, end := shardedRangeFromReaderAt(i, bytesPerWorker, size, f)
			r := io.NewSectionReader(f, int64(begin), int64(end-begin))

			b := make([]byte, 8*1024)
			sum, err := sum6.Sum6Reader(r, b)
			if err != nil {
				// TODO
			}
			resultCh <- sum
		}(i)
	}

	for i := 0; i < workers; i++ {
		ret += <-resultCh
	}
	close(resultCh)
	return ret, nil
}

func shardedRangeFromReaderAt(i, bytesPerWorker, size int, f *os.File) (int, int) {
	panic("unimplemented")
}
