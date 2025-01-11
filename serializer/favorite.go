package serializer

import (
	"context"
	"test_mysql/conf"
	"test_mysql/dao"
	"test_mysql/model"
)

type Favorite struct {
	UserId        uint   `json:"id"`
	ProductID     uint   `json:"product_id"`
	CreatedAt     int64  `json:"created_at"`
	Name          string `json:"name"`
	Category      uint   `json:"category"`
	Title         string `json:"title"`
	Info          string `json:"info"`
	ImgPath       string `json:"img_path"`
	Price         string `json:"price"`
	DiscountPrice string `json:"discount_price"`
	BossId        uint   `json:"boss_id"`
	Num           int    `json:"num"`
	OnSale        bool   `json:"on_sale"`
}

func BuildFavorite(favorite *model.Favorite, product *model.Product, boss *model.User) Favorite {
	return Favorite{
		UserId:        favorite.UserID,
		ProductID:     favorite.ProductId,
		CreatedAt:     favorite.CreatedAt.Unix(),
		Name:          product.Name,
		Category:      product.CategoryId,
		Title:         product.Title,
		Info:          product.Info,
		ImgPath:       conf.Host + conf.HttpPort + conf.ProductPath + product.ImgPath,
		Price:         product.Price,
		DiscountPrice: product.DiscountPrice,
		BossId:        boss.ID,
		Num:           product.Num,
		OnSale:        product.OnSale,
	}
}

func BuildFavorites(ctx context.Context, items []*model.Favorite) (favorites []Favorite) {
	productDao := dao.NewProductDao(ctx)
	bossDao := dao.NewUserDao(ctx)
	for _, item := range items {
		product, err := productDao.GetProductById(item.ProductId)
		if err != nil {
			continue
		}
		boss, err := bossDao.GetUserById(item.BossID)
		if err != nil {
			continue
		}
		favorite := BuildFavorite(item, product, boss)
		favorites = append(favorites, favorite)
	}
	return favorites
}
