package v1

import (
	"edu/lib/net/http"
	"edu/service"
	"time"

	"github.com/gin-gonic/gin"
)

var StatisticCtrl = &StatisticController{
	statisticSvr: service.StatisticSvr,
}

type StatisticController struct {
	statisticSvr *service.StatisticService
}

func (sc *StatisticController) UserRegister(c *gin.Context) {
	startAtString := c.Query("startAt")
	endAtString := c.Query("endAt")
	now := time.Now()

	startAt, e1 := time.Parse("2006-01-02", startAtString)
	if e1 != nil {
		day := now.AddDate(0, -1, 0)
		startAt = time.Date(day.Year(), day.Month(), day.Day(), 0, 0, 0, 0, time.Local)
	}
	endAt, e2 := time.Parse("2006-01-02", endAtString)
	if e2 != nil {
		endAt = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local).Add(24 * time.Hour)
	}

	data := sc.statisticSvr.UserRegister(startAt, endAt)
	http.SuccessData(c, "数据获取成功!", data)
}

func (sc *StatisticController) OrderCreated(c *gin.Context) {
	startAtString := c.Query("startAt")
	endAtString := c.Query("endAt")
	now := time.Now()

	startAt, e1 := time.Parse("2006-01-02", startAtString)
	if e1 != nil {
		day := now.AddDate(0, -1, 0)
		startAt = time.Date(day.Year(), day.Month(), day.Day(), 0, 0, 0, 0, time.Local)
	}
	endAt, e2 := time.Parse("2006-01-02", endAtString)
	if e2 != nil {
		endAt = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local).Add(24 * time.Hour)
	}

	data := sc.statisticSvr.OrderCreated(startAt, endAt)
	http.SuccessData(c, "数据获取成功!", data)
}

func (sc *StatisticController) OrderPaidCount(c *gin.Context) {
	startAtString := c.Query("startAt")
	endAtString := c.Query("endAt")
	now := time.Now()

	startAt, e1 := time.Parse("2006-01-02", startAtString)
	if e1 != nil {
		day := now.AddDate(0, -1, 0)
		startAt = time.Date(day.Year(), day.Month(), day.Day(), 0, 0, 0, 0, time.Local)
	}
	endAt, e2 := time.Parse("2006-01-02", endAtString)
	if e2 != nil {
		endAt = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local).Add(24 * time.Hour)
	}

	data := sc.statisticSvr.OrderPaidCount(startAt, endAt)
	http.SuccessData(c, "数据获取成功!", data)
}

func (sc *StatisticController) OrderPaidSum(c *gin.Context) {
	startAtString := c.Query("startAt")
	endAtString := c.Query("endAt")
	now := time.Now()

	startAt, e1 := time.Parse("2006-01-02", startAtString)
	if e1 != nil {
		day := now.AddDate(0, -1, 0)
		startAt = time.Date(day.Year(), day.Month(), day.Day(), 0, 0, 0, 0, time.Local)
	}
	endAt, e2 := time.Parse("2006-01-02", endAtString)
	if e2 != nil {
		endAt = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local).Add(24 * time.Hour)
	}

	data := sc.statisticSvr.OrderPaidSum(startAt, endAt)
	http.SuccessData(c, "数据获取成功!", data)
}
