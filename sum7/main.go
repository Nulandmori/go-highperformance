package sum7

import "github.com/Nulandmori/go-highperformance/sum1"

var sumByFile = map[string]int64{}

func Sum7(fileName string) (int64, error) {
	if s, ok := sumByFile[fileName]; ok {
		return s, nil
	}

	ret, err := sum1.Sum(fileName)
	if err != nil {
		return 0, err
	}

	sumByFile[fileName] = ret
	return ret, nil
}
