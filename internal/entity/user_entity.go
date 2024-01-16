package entity

type User struct {
	ID        string
	Username  string
	Password  string
	Email     string
	Phone     string
	Avatar    string
	CreatedAt int64     `gorm:"autoCreateTime:milli"`
	UpdatedAt int64     `gorm:"autoCreateTime:milli;autoUpdateTime:milli"`
	Articles  []Article `gorm:"foreignKey:user_id;references:id"`
	Comments  []Comment `gorm:"foreignKey:user_id;references:id"`
}
