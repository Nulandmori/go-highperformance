package concurrentsum1

import (
	"os"
	"sync"
	"sync/atomic"

	"github.com/Nulandmori/go-highperformance/sum4"
)

func ConcurrentSum1(fileName string) (ret int64, _ error) {
	b, err := os.ReadFile(fileName)
	if err != nil {
		return 0, err
	}

	var wg sync.WaitGroup
	var last int
	for i := 0; i < len(b); i++ {
		if b[i] != '\n' {
			continue
		}

		wg.Add(1)
		go func(line []byte) {
			defer wg.Done()
			num, err := sum4.ParseInt(line)
			if err != nil {
				return
			}
			atomic.AddInt64(&ret, num)
		}(b[last:i])
		last = i + 1
	}
	wg.Wait()
	return ret, nil
}
