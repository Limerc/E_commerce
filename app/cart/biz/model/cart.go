package model

import (
	"context"
	"errors"

	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model
	UserId    uint32 `gorm:"type:int(11);not null;index:idx_user_id"`
	ProductId uint32 `gorm:"type:int(11);not null;"`
	Qty       uint32 `gorm:"type:int(11);not null;"`
}

func (Cart) TableName() string {
	return "cart"
}

// 添加方法将购物车记录写入数据库
func AddItem(ctx context.Context, db *gorm.DB, item *Cart) error {
	var row Cart
	// 如果有记录就更新记录
	err := db.WithContext(ctx).
		Model((&Cart{})).
		Where(&Cart{UserId: item.UserId, ProductId: item.ProductId}).
		First(&row).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	if row.ID > 0 {
		return db.WithContext(ctx).
			Model((&Cart{})).
			Where(&Cart{UserId: item.UserId, ProductId: item.ProductId}).
			UpdateColumn("qty", gorm.Expr("qty + ?", item.Qty)).Error
	}
	// 如果没有记录就创建新的记录
	return db.WithContext(ctx).Create(item).Error
}

func EmptyCart(ctx context.Context, db *gorm.DB, userId uint32) error {
	if userId == 0 {
		return errors.New("user id is required")
	}
	return db.WithContext(ctx).Delete(&Cart{}, "user_id = ?", userId).Error
}

func GetCartByUserId(ctx context.Context, db *gorm.DB, userId uint32) ([]*Cart, error) {
	var rows []*Cart
	err := db.WithContext(ctx).
		Model(&Cart{}).
		Where(&Cart{UserId: userId}).
		Find(&rows).Error
	return rows, err
}