package service

import (
"edu/model"
)

var DashboardSvr *DashboardService = &DashboardService{baseService: newBaseService()}

type DashboardService struct {
	baseService
}

func (svr *DashboardService) AdminDashboard() model.AdminDashboard {
	rdb := model.AdminDashboard{}
	db := model.AdminDashboard{}

todayRegisterUserCount := 0 // TODO: 替换为实际用户统计逻辑
db.TodayRegisterUserCount = int(todayRegisterUserCount)

// // 总用户数
userCount := 0 // TODO: 替换为实际用户统计逻辑
db.UserCount = int(userCount)

// // 昨日订单支付总额
db.YestodayPaidSum = 0 // TODO: 替换为实际统计逻辑
// // 今日订单支付总额
db.TodayPaidSum = 0 // TODO: 替换为实际统计逻辑

	// // 进入付费用户数量
	todayPaidUserNum := int64(0)
	// svr.db.Model(&model.Order{}).Select("user_id").Where("created_at > ?", todayStart).Where("status = ?", model.ORDER_STATUS_PAID).Group("user_id").Count(&todayPaidUserNum)

	db.TodayPaidUserNum = int(todayPaidUserNum)

	// // 昨日付费用户数量
	yestodayPaidUserNum := int64(0)
	// svr.db.Model(&model.Order{}).Select("user_id").Where("created_at BETWEEN ? AND ?", todayStart.Add(-24*time.Hour), todayStart).Where("status = ?", model.ORDER_STATUS_PAID).Group("user_id").Count(&yestodayPaidUserNum)
	db.YestodayPaidUserNum = int(yestodayPaidUserNum)

	// // 本月收益
	// svr.db.Model(&model.Order{}).Where("created_at > ?", todayStart.AddDate(0, 0, -todayStart.Day()+1)).Where("status = ?", model.ORDER_STATUS_PAID).Select("sum(charge) as this_month_paid_sum").Scan(&rdb)
	db.ThisMonthPaidSum = rdb.ThisMonthPaidSum

	// // 上月收益
	// svr.db.Model(&model.Order{}).Where("created_at BETWEEN ? AND ?", todayStart.AddDate(0, -1, -todayStart.Day()+1), todayStart.AddDate(0, 0, -todayStart.Day()+1)).Where("status = ?", model.ORDER_STATUS_PAID).Select("sum(charge) as last_month_paid_sum").Scan(&rdb)
	db.LastMonthPaidSum = rdb.LastMonthPaidSum

	return db
}
