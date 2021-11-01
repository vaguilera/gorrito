package gorrito

type matchesByPuuidOptions struct {
	startTime *int64
	endTime   *int64
	queue     *int
	ttype     *string
	start     *int
	count     *int
}

func (c *Client) NewMatchesByPuuidOptions() *matchesByPuuidOptions {
	return &matchesByPuuidOptions{}
}

func (mo *matchesByPuuidOptions) StartTime(startTime int64) *matchesByPuuidOptions {
	mo.startTime = &startTime
	return mo
}

func (mo *matchesByPuuidOptions) EndTime(endTime int64) *matchesByPuuidOptions {
	mo.endTime = &endTime
	return mo
}

func (mo *matchesByPuuidOptions) Queue(queue int) *matchesByPuuidOptions {
	mo.queue = &queue
	return mo
}

func (mo *matchesByPuuidOptions) Type(ttype string) *matchesByPuuidOptions {
	mo.ttype = &ttype
	return mo
}

func (mo *matchesByPuuidOptions) Start(start int) *matchesByPuuidOptions {
	mo.start = &start
	return mo
}

func (mo *matchesByPuuidOptions) Count(count int) *matchesByPuuidOptions {
	mo.count = &count
	return mo
}
