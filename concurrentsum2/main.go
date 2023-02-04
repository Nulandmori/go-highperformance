package concurrentsum2

import (
	"os"
	"sync"
	"sync/atomic"

	"github.com/Nulandmori/go-highperformance/sum4"
)

func ConcurrentSum2(fileName string, workers int) (ret int64, _ error) {
	b, err := os.ReadFile(fileName)
	if err != nil {
		return 0, err
	}

	var (
		wg     = sync.WaitGroup{}
		workCh = make(chan []byte, 10)
	)

	wg.Add(workers + 1)
	go func() {
		var last int
		for i := 0; i < len(b); i++ {
			if b[i] != '\n' {
				continue
			}
			workCh <- b[last:i]
			last = i + 1
		}
		close(workCh)
		wg.Done()
	}()

	for i := 0; i < workers; i++ {
		go func() {
			var sum int64
			for line := range workCh {
				num, err := sum4.ParseInt(line)
				if err != nil {
					continue
				}
				sum += num
			}
			atomic.AddInt64(&ret, sum)
			wg.Done()
		}()
	}
	wg.Wait()
	return ret, nil
}
