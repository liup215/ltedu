package model

type AdminDashboard struct {
	TodayRegisterUserCount int `json:"todayRegisterUserCount"`
	UserCount              int `json:"userCount"`
	YestodayPaidSum        int `json:"yestodayPaidSum"`
	TodayPaidSum           int `json:"todayPaidSum"`
	TodayPaidUserNum       int `json:"todayPaidUserNum"`
	YestodayPaidUserNum    int `json:"yestodayPaidUserNum"`
	ThisMonthPaidSum       int `json:"thisMonthPaidSum"`
	LastMonthPaidSum       int `json:"lastMonthPaidSum"`
}
