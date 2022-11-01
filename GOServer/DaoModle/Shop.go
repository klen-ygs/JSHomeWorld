package DaoModle

// Shop 商店模型
type Shop struct {
	Id            uint64
	ShopTitleText string
	TitleImageId  uint64
	ShowImageIds  string //json
	ShowTextId    uint64
	Having        uint64
	price         int16
}
