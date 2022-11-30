package DaoModle

type HotSearchSlice []HotSearch

func (h *HotSearchSlice) Len() int {
	//TODO implement me
	return len(*h)
}

func (h *HotSearchSlice) Less(i, j int) bool {
	//TODO implement me
	return (*h)[i].SearchCount > (*h)[j].SearchCount
}

func (h *HotSearchSlice) Swap(i, j int) {
	//TODO implement me
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *HotSearchSlice) Push(x any) {
	//TODO implement me
	*h = append(*h, x.(HotSearch))
}

func (h *HotSearchSlice) Pop() any {
	//TODO implement me
	tar := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return tar
}
