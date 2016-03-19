package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

/***

CREATE TABLE `blog`.`<table_name>` (
	`user_id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
	`user_name` varchar(32) DEFAULT NULL,
	`password` varchar(255) DEFAULT NULL,
	`nick_name` varchar(32) DEFAULT NULL,
	`age` int(11) DEFAULT 0,
	`sex` char(32) DEFAULT NULL,
	`state` int(11) DEFAULT 0,
	`address` varchar(255) DEFAULT NULL,
	`create_time` datetime DEFAULT NULL,
	`update_time` datetime DEFAULT NULL,
	`mobile` varchar(32) DEFAULT NULL,
	`email` varchar(32) DEFAULT NULL,
	PRIMARY KEY (`user_id`),
	UNIQUE `index_user_name` USING BTREE (`user_name`) comment '',
	UNIQUE `index_user_mobile` USING BTREE (`mobile`) comment '',
	UNIQUE `index_user_email` USING BTREE (`email`) comment ''
) ENGINE=`InnoDB` AUTO_INCREMENT=2 DEFAULT CHARACTER SET utf8 COLLATE utf8_general_ci ROW_FORMAT=COMPACT COMMENT='' CHECKSUM=0 DELAY_KEY_WRITE=0;

*/
type User struct {
	UserId     int `orm:"pk;auto"`
	UserName   string
	Password   string
	NickName   string
	Age        int
	Sex        string
	State      int
	Address    string
	Mobile     string
	Email      string
	CreateTime time.Time
	UpdateTime time.Time
}

func (this *User) TableName() string {
	return "t_user"
}

/***
用户登录
**/
func Login(userName, password string) (User, error) {

	var user = User{}
	err := orm.NewOrm().QueryTable("t_user").Filter("user_name",
		userName).Filter("password",
		password).Filter("state", 0).One(&user)
	return user, err

}
