package service

import (
"edu/model"
"encoding/json"
"time"
)

var StatisticSvr *StatisticService = &StatisticService{baseService: newBaseService()}

type StatisticService struct {
	baseService
}

type ChartData struct {
	Labels  []string `json:"labels"`
	Dataset []int    `json:"dataset"`
}

func (svr *StatisticService) UserRegister(startAt, endAt time.Time) ChartData {

users := []struct {
CreatedAt time.Time
}{} // TODO: 替换为实际用户查询逻辑

	r := map[string]int{}
	data := ChartData{}
	tempDate := startAt
	for tempDate.Before(endAt) {
		date := tempDate.Format("2006-01-02")
		data.Labels = append(data.Labels, date)
		r[date] = 0
		tempDate = tempDate.Add(24 * time.Hour)
	}

	for _, u := range users {
		date := u.CreatedAt.Format("2006-01-02")
		r[date] = r[date] + 1
	}

	for _, date := range data.Labels {
		data.Dataset = append(data.Dataset, r[date])
	}

	log := model.AdminLog{Module: model.ADMINLOG_MODULE_STATS, Opt: model.ADMINLOG_OPT_VIEW}
	if remarkByte, err := json.Marshal(map[string]time.Time{
		"startAt": startAt,
		"endAt":   endAt,
	}); err == nil {
		log.Remark = string(remarkByte)
	}
/* TODO: 替换为实际日志创建逻辑 */

	return data
}

func (svr *StatisticService) OrderCreated(startAt, endAt time.Time) ChartData {

orders := []struct {
CreatedAt time.Time
}{} // TODO: 替换为实际订单查询逻辑

	r := map[string]int{}
	data := ChartData{}
	tempDate := startAt
	for tempDate.Before(endAt) {
		date := tempDate.Format("2006-01-02")
		data.Labels = append(data.Labels, date)
		r[date] = 0
		tempDate = tempDate.Add(24 * time.Hour)
	}

	for _, o := range orders {
		date := o.CreatedAt.Format("2006-01-02")
		r[date] = r[date] + 1
	}

	for _, date := range data.Labels {
		data.Dataset = append(data.Dataset, r[date])
	}

	log := model.AdminLog{Module: model.ADMINLOG_MODULE_STATS, Opt: model.ADMINLOG_OPT_VIEW}
	if remarkByte, err := json.Marshal(map[string]time.Time{
		"startAt": startAt,
		"endAt":   endAt,
	}); err == nil {
		log.Remark = string(remarkByte)
	}
/* TODO: 替换为实际日志创建逻辑 */

	return data
}

func (svr *StatisticService) OrderPaidCount(startAt, endAt time.Time) ChartData {

orders := []struct {
CreatedAt time.Time
Charge    int
}{} // TODO: 替换为实际已支付订单查询逻辑

	r := map[string]int{}
	data := ChartData{}
	tempDate := startAt
	for tempDate.Before(endAt) {
		date := tempDate.Format("2006-01-02")
		data.Labels = append(data.Labels, date)
		r[date] = 0
		tempDate = tempDate.Add(24 * time.Hour)
	}

	for _, o := range orders {
		date := o.CreatedAt.Format("2006-01-02")
		r[date] = r[date] + 1
	}

	for _, date := range data.Labels {
		data.Dataset = append(data.Dataset, r[date])
	}

	log := model.AdminLog{Module: model.ADMINLOG_MODULE_STATS, Opt: model.ADMINLOG_OPT_VIEW}
	if remarkByte, err := json.Marshal(map[string]time.Time{
		"startAt": startAt,
		"endAt":   endAt,
	}); err == nil {
		log.Remark = string(remarkByte)
	}
/* TODO: 替换为实际日志创建逻辑 */

	return data
}

func (svr *StatisticService) OrderPaidSum(startAt, endAt time.Time) ChartData {
orders := []struct {
CreatedAt time.Time
Charge    int
}{} // TODO: 替换为实际已支付订单查询逻辑

	r := map[string]int{}
	data := ChartData{}
	tempDate := startAt
	for tempDate.Before(endAt) {
		date := tempDate.Format("2006-01-02")
		data.Labels = append(data.Labels, date)
		r[date] = 0
		tempDate = tempDate.Add(24 * time.Hour)
	}

	for _, o := range orders {
		date := o.CreatedAt.Format("2006-01-02")
		r[date] = r[date] + o.Charge
	}

	for _, date := range data.Labels {
		data.Dataset = append(data.Dataset, r[date])
	}

	log := model.AdminLog{Module: model.ADMINLOG_MODULE_STATS, Opt: model.ADMINLOG_OPT_VIEW}
	if remarkByte, err := json.Marshal(map[string]time.Time{
		"startAt": startAt,
		"endAt":   endAt,
	}); err == nil {
		log.Remark = string(remarkByte)
	}
/* TODO: 替换为实际日志创建逻辑 */

return data
}
