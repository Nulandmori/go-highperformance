package concurrentsum4

import (
	"os"
	"testing"

	"github.com/efficientgo/core/testutil"
	"github.com/felixge/fgprof"
)

func BenchmarkSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = ConcurrentSum4("../testdata/test.10M.txt", 2)
	}
}

// BenchmarkSum_fgprof recommended run options:
//
//	$ export ver=v1fg && go test -run '^$' -bench '^BenchmarkSum_fgprof' \
//	   -benchtime 60s  -cpu 1 | tee ${ver}.txt
func BenchmarkSum_fgprof(b *testing.B) {
	f, err := os.Create("fgprof.pprof")
	testutil.Ok(b, err)

	defer func() { testutil.Ok(b, f.Close()) }()

	closeFn := fgprof.Start(f, fgprof.FormatPprof)
	BenchmarkSum(b)
	testutil.Ok(b, closeFn())
}
