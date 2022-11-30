package DaoModle

// Shop 商店模型
type Shop struct {
	Id            uint64
	Name          string
	ShopTitleText string
	Text          string
	Having        uint64
	Price         int16
	Address       string
}

type ShopSlice []Shop

func (s *ShopSlice) Less(i, j int) bool {
	return (*s)[i].Id < (*s)[j].Id
}

func (s *ShopSlice) Len() int {
	return len(*s)
}

func (s *ShopSlice) Swap(i, j int) {
	(*s)[i], (*s)[j] = (*s)[j], (*s)[i]
}

func (s *ShopSlice) Push(x any) {
	*s = append(*s, x.(Shop))
}

func (s *ShopSlice) Pop() any {
	tar := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return tar
}
