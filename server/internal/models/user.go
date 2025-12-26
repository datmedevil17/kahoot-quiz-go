package models

type UserRole string

const (
	RoleHost   UserRole = "host"
	RolePlayer UserRole = "player"
)

type User struct {
	BaseModel

	Name         string   `gorm:"not null"`
	Email        string   `gorm:"uniqueIndex;not null"`
	PasswordHash string   `gorm:"not null"`
	Role         UserRole `gorm:"type:varchar(20);not null"`
	Followers    []*User  `gorm:"many2many:follows;joinForeignKey:following_id;joinReferences:follower_id"`
	Following    []*User  `gorm:"many2many:follows;joinForeignKey:follower_id;joinReferences:following_id"`
}

type Follows struct {
	FollowerID  uint  `gorm:"primaryKey"`
	FollowingID uint  `gorm:"primaryKey"`
	CreatedAt   int64 `gorm:"autoCreateTime"` // Using simple int64 for timestamp if compatible, or just standard fields check base model
}
