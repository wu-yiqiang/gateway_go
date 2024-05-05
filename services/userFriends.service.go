package services

import (
	"gateway_go/dao"
	"gateway_go/global"
	"gateway_go/utils"
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

type UuidLists struct {
	Uuid string
}

func (userFriendsServer *userFriendsService) QueryFriends(userid string) (err error, friends []string) {
	// var uuidLists []UuidLists
	global.App.DB.Table("user_friends").Select("friend_id").Where("user_id = ?", userid).Find(&friends)
	var friends2 []string
	global.App.DB.Table("user_friends").Select("user_id").Where("friend_id = ?", userid).Find(&friends2)
	friends = append(friends, friends2...)
	friends = utils.UniqueArray(friends)
	return
}
