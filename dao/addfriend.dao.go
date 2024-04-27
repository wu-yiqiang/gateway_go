package dao

type UserFriends struct {
	Id          int    `json:"id" gorm:"primary_key" description:"自增主键"`
	UserId      string `json:"user_id" gorm:"user_id" description:"主用户id"`
	FriendId    string `json:"friend_id" gorm:"friend_id" description:"从用户id"`
	UpdatedTime int64  `json:"updated_Time" gorm:"column:updated_time" description:"更新时间"`
	CreatedTime int64  `json:"created_time" gorm:"column:created_time" description:"创建时间"`
	IsDelete    int    `json:"is_delete" gorm:"column:is_delete" description:"是否删除"`
}
