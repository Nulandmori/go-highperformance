package concurrentsum3

import (
	"os"

	"github.com/Nulandmori/go-highperformance/sum4"
)

func ConcurrentSum3(fileName string, workers int) (ret int64, _ error) {
	b, err := os.ReadFile(fileName)
	if err != nil {
		return 0, err
	}

	var (
		bytesPerWorker = len(b) / workers
		resultCh       = make(chan int64)
	)

	for i := 0; i < workers; i++ {
		go func(i int) {
			// Coordination-free algorithm, which shards
			// buffered file deterministically.
			begin, end := shardedRange(i, bytesPerWorker, b)

			var sum int64
			for last := begin; begin < end; begin++ {
				if b[begin] != '\n' {
					continue
				}
				num, err := sum4.ParseInt(b[last:begin])
				if err != nil {
					continue
				}
				sum += num
				last = begin + 1
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

func shardedRange(i, bytesPerWorker int, b []byte) (int, int) {
	panic("unimplemented")
}
