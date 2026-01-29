package model

type Menu struct {
	Id       string `json:"id"`
	Icon     string `json:"icon"`
	Title    string `json:"title"`
	Children []Menu `json:"children" gorm:"-"`
}

var AdminMenu = []Menu{
	{Id: "/dashboard", Icon: "layui-icon-home", Title: "主页"},
	{Id: "/decoration", Icon: "layui-icon-home", Title: "装修", Children: []Menu{
		{Id: "/decoration/pc", Icon: "layui-icon-home", Title: "PC"}},
	},
	{Id: "/user", Icon: "layui-icon-user", Title: "用户管理", Children: []Menu{
		{Id: "/user/user", Icon: "layui-icon-user", Title: "学生管理"},
		{Id: "/user/teacher", Icon: "layui-icon-user", Title: "教师管理"}},
	},
	{Id: "/school", Icon: "layui-icon-user", Title: "教务管理", Children: []Menu{
		{Id: "/school/grade", Icon: "layui-icon-user", Title: "年级管理"},
		{Id: "/school/classType", Icon: "layui-icon-user", Title: "班级类型管理"},
		{Id: "/school/class", Icon: "layui-icon-user", Title: "班级管理"}},
	},
	{Id: "/qualification", Icon: "layui-icon-home", Title: "学科管理", Children: []Menu{
		{Id: "/qualification/organisation", Icon: "layui-icon-home", Title: "考试局"},
		{Id: "/qualification/qualification", Icon: "layui-icon-home", Title: "考试"},
		{Id: "/qualification/syllabus", Icon: "layui-icon-home", Title: "考纲"},
		{Id: "/qualification/series", Icon: "layui-icon-home", Title: "考试季"},
		{Id: "/qualification/code", Icon: "layui-icon-home", Title: "试卷代码"}},
	},
	{Id: "/resource", Icon: "layui-icon-home", Title: "资源管理", Children: []Menu{
		{Id: "/resource/image", Icon: "layui-icon-home", Title: "图片列表"},
		{Id: "/resource/video", Icon: "layui-icon-home", Title: "视频列表"},
	}},
	{Id: "/course", Icon: "layui-icon-home", Title: "课程管理", Children: []Menu{
		{Id: "/course/list", Icon: "layui-icon-home", Title: "课程列表"},
	},
	},
	{Id: "/document", Icon: "layui-icon-home", Title: "文档管理", Children: []Menu{
		{Id: "/document/list", Icon: "layui-icon-home", Title: "文档列表"},
		{Id: "/document/category", Icon: "layui-icon-home", Title: "文档分类"},
	},
	},
	{Id: "/paper", Icon: "layui-icon-home", Title: "题卷管理", Children: []Menu{
		{Id: "/paper/past", Icon: "layui-icon-home", Title: "真题试卷"},
		{Id: "/paper/random", Icon: "layui-icon-home", Title: "随机试卷"},
		{Id: "/paper/exam", Icon: "layui-icon-home", Title: "选题组卷"},
		{Id: "/paper/question", Icon: "layui-icon-home", Title: "试题列表"},
	}},
	{Id: "/vocabularySet", Icon: "layui-icon-home", Title: "词汇管理", Children: []Menu{
		{Id: "/vocabularySet/list", Icon: "layui-icon-home", Title: "词汇集列表"},
	}},
	{Id: "/slide", Icon: "layui-icon-home", Title: "课件管理", Children: []Menu{
		{Id: "/slide/list", Icon: "layui-icon-home", Title: "课件列表"},
	}},
	{Id: "/system", Icon: "layui-icon-home", Title: "系统", Children: []Menu{
		{Id: "/system/sysset", Icon: "layui-icon-home", Title: "系统设置"},
		{Id: "/system/admin", Icon: "layui-icon-home", Title: "管理员"},
		{Id: "/system/role", Icon: "layui-icon-home", Title: "角色配置"}},
	}}

var TeacherMenu = []Menu{
	{Id: "/dashboard", Icon: "layui-icon-home", Title: "主页"},
	{Id: "/user", Icon: "layui-icon-user", Title: "用户管理", Children: []Menu{
		{Id: "/user/user", Icon: "layui-icon-user", Title: "学生管理"},
	},
	},
	{Id: "/school", Icon: "layui-icon-user", Title: "教务管理", Children: []Menu{
		{Id: "/school/class", Icon: "layui-icon-user", Title: "班级管理"}},
	},
	{Id: "/questionPaper", Icon: "layui-icon-home", Title: "试卷管理", Children: []Menu{
		{Id: "/questionPaper/pastPaper", Icon: "layui-icon-home", Title: "真题试卷"},
		{Id: "/questionPaper/question", Icon: "layui-icon-home", Title: "试题列表"},
	}},
	{Id: "/performMark", Icon: "layui-icon-home", Title: "操行管理", Children: []Menu{
		{Id: "/performMark/record", Icon: "layui-icon-home", Title: "分数记录"}},
	},
}
