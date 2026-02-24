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

// @Summary      用户注册统计
// @Description  获取指定时间范围内的用户注册数量统计
// @Tags         系统设置
// @Accept       json
// @Produce      json
// @Param        startAt  query  string  false  "开始日期 (2006-01-02)"
// @Param        endAt    query  string  false  "结束日期 (2006-01-02)"
// @Success      200  {object}  map[string]interface{}  "成功"
// @Security     BearerAuth
// @Router       /v1/statistic/userRegister [get]
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

// @Summary      订单创建统计
// @Description  获取指定时间范围内的订单创建数量统计
// @Tags         系统设置
// @Accept       json
// @Produce      json
// @Param        startAt  query  string  false  "开始日期 (2006-01-02)"
// @Param        endAt    query  string  false  "结束日期 (2006-01-02)"
// @Success      200  {object}  map[string]interface{}  "成功"
// @Security     BearerAuth
// @Router       /v1/statistic/orderCreated [get]
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

// @Summary      订单支付数量统计
// @Description  获取指定时间范围内的订单支付数量统计
// @Tags         系统设置
// @Accept       json
// @Produce      json
// @Param        startAt  query  string  false  "开始日期 (2006-01-02)"
// @Param        endAt    query  string  false  "结束日期 (2006-01-02)"
// @Success      200  {object}  map[string]interface{}  "成功"
// @Security     BearerAuth
// @Router       /v1/statistic/orderPaidCount [get]
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

// @Summary      订单支付金额统计
// @Description  获取指定时间范围内的订单支付金额统计
// @Tags         系统设置
// @Accept       json
// @Produce      json
// @Param        startAt  query  string  false  "开始日期 (2006-01-02)"
// @Param        endAt    query  string  false  "结束日期 (2006-01-02)"
// @Success      200  {object}  map[string]interface{}  "成功"
// @Security     BearerAuth
// @Router       /v1/statistic/orderPaidSum [get]
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
