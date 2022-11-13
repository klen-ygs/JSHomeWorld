package DaoModle

// Image 图片模型
type Image struct {
	Id        uint64
	ShopId    uint64
	Flag      bool
	ImageData []byte
}
