package model

type CourseUserRecord struct {
	Model
	UserID    uint `json:"userId"`
	CourseID  uint `json:"courseId"`
	IsWatched int  `json:"isWatched"`
}
