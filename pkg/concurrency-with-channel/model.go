package concurrency_with_channel

const channelLimit = 100

type Results [][]byte

func (r *Results) CheckChannelLimit(channelLimit int, channelCount int, resultChan chan *ChanResult) (isChannelLimitOver bool, err error) {
	if channelCount < channelLimit {
		return
	}

	var resultsFromChannel Results

	if resultsFromChannel, err = receiveResultFromChannel(resultChan, channelCount); err != nil {
		return
	}
	*r = append(*r, resultsFromChannel...)
	isChannelLimitOver = true
	return
}

type ChanResult struct {
	Result []byte
	Err    error
}

func newChanResult() *ChanResult {
	return &ChanResult{}
}

func (r *ChanResult) setError(err error) *ChanResult {
	r.Err = err

	return r
}
func (r *ChanResult) setResult(result []byte) *ChanResult {
	r.Result = result

	return r
}
