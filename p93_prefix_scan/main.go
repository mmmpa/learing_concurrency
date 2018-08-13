package p93

import (
	"github.com/mmmpa/parallel/common"
)

func compute(array []int) []int {
	l := len(array)
	sum := make([]int, l)

	sum[0] = array[0]
	for i := 1; i < l; i++ {
		sum[i] = sum[i-1] + array[i]
	}

	return sum
}

type SumResult struct {
	Index int
	Sum   []int
}
func PrefixScan(array []int, rawWorkers int) []int {
	return computeC(array, rawWorkers)
}

func computeC(array []int, rawWorkers int) []int {
	l := len(array)

	ch := make(chan *SumResult)
	works := common.SplitWorks(l, rawWorkers)
	workersNum := len(works)
	head := 0

	for i, n := range works {
		go func(ch chan *SumResult, index, length, head int) {
			sum := make([]int, length)

			sum[0] = array[head]
			for i := 1; i < length; i++ {
				sum[i] = sum[i-1] + array[head+i]
			}

			ch <- &SumResult{
				Index: index,
				Sum:   sum,
			}
		}(ch, i, n, head)

		head += n
	}

	partPrefixResults := make([][]int, workersNum)

	for i := 0; i < workersNum; i++ {
		s := <-ch
		partPrefixResults[s.Index] = s.Sum
	}

	// 部分プリフィックススキャン結果を排他プリフィックススキャンする

	partResultTailsPrefix := make([]int, workersNum)

	partResultTailsPrefix[0] = 0
	for i := 1; i < workersNum; i++ {
		partResult := partPrefixResults[i-1]
		partResultTailsPrefix[i] = partResultTailsPrefix[i-1] + partResult[len(partResult)-1]
	}

	// 排他プリフィックススキャン結果を部分プリフィックススキャンの結果全てに足す

	allPrefix := make([]int, l)
	offset := 0

	for resultIndex, partResult := range partPrefixResults {
		go func(partResult []int, resultIndex, offset int) {
			for i, _ := range partResult {
				allPrefix[offset+i] = partResult[i] + partResultTailsPrefix[resultIndex]
			}
			ch <- nil
		}(partResult, resultIndex, offset)

		offset += len(partResult)
	}

	for i := 0; i < workersNum; i++ {
		<-ch
	}

	return allPrefix
}
