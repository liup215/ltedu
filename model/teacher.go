package model

import "time"

type Teacher struct {
	Model
	Username            string     `json:"username"`
	Nickname            string     `json:"nickname"`
	Realname            string     `json:"realname"`
	Sex                 int        `json:"sex"`
	Engname             string     `json:"engname"`
	Password            string     `json:"-"`
	Salt                string     `json:"-"`
	Mobile              string     `json:"mobile" gorm:"unique"`
	Avatar              string     `json:"avatar"`
	Status              int        `json:"Status"`
	SubjectId           uint       `json:"subjectId"`
	SubjectName         string     `json:"subjectName" gorm:"-"`
	IsActive            int        `json:"IsActive"`
	InviteUserId        uint       `json:"inviteUserId"`
	InviteBalance       int        `json:"inviteBalance"`
	InviteUserExpiredAt *time.Time `json:"inviteUserExpiredAt"`
	IsPasswordSet       int        `json:"isPasswordSet"`
	IsUsernameSet       int        `json:"isUsernameSet"`
	IsNicknameSet       int        `json:"isNicknameSet"`
	RegisterIp          string     `json:"registerIp"`
	RegisterArea        string     `json:"registerArea"`
	LastLoginId         uint       `json:"lastLoginId"`
}

type TeacherQuery struct {
	Teacher
	Page
}

type TeacherRegisterRequest struct {
	Username        string `json:"username"`
	Mobile          string `json:"mobile"`
	Password        string `json:"password"`
	PasswordConfirm string `json:"passwordConfirm"`
	Realname        string `json:"realname"`
}

type SetTeacherUsernamePasswordRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
