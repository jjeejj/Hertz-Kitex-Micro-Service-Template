package model

import (
	"github.com/bwmarrin/snowflake"
	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/gorm"
)

type User struct {
	ID           int64 `gorm:"primarykey"`
	PhoneNumber  int64
	AvatarBlobId int64
	Username     string `gorm:"type:varchar(40)"`
	OpenID       string `gorm:"column:openid;type:varchar(100);not null"`
}

// BeforeCreate uses snowflake to generate an ID.
func (u *User) BeforeCreate(_ *gorm.DB) (err error) {
	sf, err := snowflake.NewNode(1)
	if err != nil {
		klog.Fatalf("generate id failed: %s", err.Error())
	}
	u.ID = sf.Generate().Int64()
	return nil
}
