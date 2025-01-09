package serializer

import (
	"path"
	"test_mysql/conf"
	"test_mysql/model"
)

//vo view object 数据流

type User struct {
	ID       uint   `json:"id"`
	UserName string `json:"user_name"`
	NickName string `json:"nick_name"`
	Type     int    `json:"type"`
	Email    string `json:"email"`
	Status   string `json:"status"`
	Avatar   string `json:"avatar"`
	CreateAt int64  `json:"create_at"`
}

func BuildUser(user *model.User) *User {
	return &User{
		ID:       user.ID,
		UserName: user.Username,
		NickName: user.Nickname,
		Email:    user.Email,
		Status:   user.Status,
		Avatar:   path.Join(conf.Host, conf.HttpPort, conf.AvatarPath, user.Avatar),
		CreateAt: user.CreatedAt.Unix(),
	}
}
