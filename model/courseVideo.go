package model

import "time"

type CourseVideo struct {
	Model
	CourseId          uint       `json:"courseId"`
	Course            Course     `json:"course"`
	MediaVideoId      uint       `json:"mediaVideoId"`
	MediaVideo        MediaVideo `json:"mediaVideo"`
	ChapterId         uint       `json:"chapterId"`
	Chapter           Chapter    `json:"chapter"`
	Name              string     `json:"name"`
	PublishedAtString string     `json:"publishedAt" gorm:"-"`
	PublishedAt       *time.Time `json:"-"`
	Duration          int        `json:"duration"`
	BanDrag           int        `json:"banDrag"`
	IsShow            int        `json:"isShow"`
}

type CourseVideoCreateEditRequest struct {
	ID           uint   `json:"id"`
	CourseId     uint   `json:"courseId"`
	MediaVideoId uint   `json:"mediaVideoId"`
	Name         string `json:"name"`
	PublishedAt  string `json:"publishedAt"`
	Duration     int    `json:"duration"`
	BanDrag      int    `json:"banDrag"`
	IsShow       int    `json:"isShow"`
	ChapterId    uint   `json:"chapterId"`
}

type CourseVideoQueryRequest struct {
	ID       uint `json:"id"`
	CourseID uint `json:"course_id"`
	VideoID  uint `json:"video_id"`
	Page
}
