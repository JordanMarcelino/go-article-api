package entity

type Comment struct {
	ID        string
	Body      string
	CreatedAt int64 `gorm:"autoCreateTime:milli"`
	User      `gorm:"foreignKey:user_id;references:id"`
	Article   `gorm:"foreignKey:article_id;references:id"`
}
