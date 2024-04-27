package services

import (
	"gateway_go/dao"
	"gateway_go/global"
	"time"
)

type userFriendsService struct {
}

var UserFriendsService = new(userFriendsService)

func (userFriendsServer *userFriendsService) TableName() string {
	return "user_friends"
}

// Register
func (userFriendsServer *userFriendsService) AddFriend(userid string, friendId string) (err error, user dao.UserFriends) {
	user = dao.UserFriends{UserId: userid, FriendId: friendId, CreatedTime: time.Now().Unix(), UpdatedTime: time.Now().Unix(), IsDelete: 0}
	err = global.App.DB.Table(userFriendsServer.TableName()).Create(&user).Error
	return
}
