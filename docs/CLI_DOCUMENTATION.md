# edu-cli 完整文档 / Complete Documentation

edu-cli is the command-line client for the LTEdu platform. It lets you manage
syllabuses, papers, users, classes, learning plans, and more through the backend API.

edu-cli 是 LTEdu 平台的命令行客户端，可通过后端 API 管理考纲、试卷、用户、班级、学习计划等。

---

## 目录 / Table of Contents

- [安装与配置 / Setup](#安装与配置--setup)
- [config – 配置管理](#config--配置管理--configuration)
- [organisation – 机构管理](#organisation--机构管理--organisations)
- [qualification – 考试管理](#qualification--考试管理--qualifications)
- [syllabus – 考纲管理](#syllabus--考纲管理--syllabuses)
- [chapter – 章节管理](#chapter--章节管理--chapters)
- [question – 题目管理](#question--题目管理--questions)
- [paper – 试卷管理](#paper--试卷管理--papers)
- [exam-node – 考试节点管理](#exam-node--考试节点管理--exam-nodes)
- [class – 班级管理](#class--班级管理--classes)
- [user – 用户管理](#user--用户管理--users)
- [teacher – 教师管理](#teacher--教师管理--teachers)
- [learning-plan – 学习计划管理](#learning-plan--学习计划管理--learning-plans)
- [phase-plan – 阶段性计划管理](#phase-plan--阶段性计划管理--phase-plans)
- [generate-docs – 生成文档](#generate-docs--生成文档--documentation-generation)

---

## 安装与配置 / Setup

### 环境变量 / Environment Variables

| 变量 / Variable | 说明 / Description |
|---|---|
| `EDU_BASE_URL` | Backend base URL, e.g. `https://api.example.com` |
| `EDU_TOKEN` | MCP token obtained from the CLI Tokens page in the web UI |

```bash
export EDU_BASE_URL=https://api.example.com
export EDU_TOKEN=your_mcp_token_here
```

### 快速配置 / Quick Configuration

```bash
# Set backend URL / 设置后端地址
edu-cli config set-url https://api.example.com

# Set authentication token / 设置认证 Token
edu-cli config set-token YOUR_MCP_TOKEN

# Show current configuration / 查看当前配置
edu-cli config show
```

---

## config – 配置管理 / Configuration

Manage edu-cli configuration settings.

管理 edu-cli 配置项。

### `config set-url <url>`

Set the backend base URL.  
设置后端服务地址。

```bash
edu-cli config set-url https://api.example.com
```

### `config set-token <token>`

Set the authentication token (MCP token obtained from the CLI Tokens page in the web UI).  
设置认证 Token（从 Web UI 的 CLI Tokens 页面获取的 MCP Token）。

```bash
edu-cli config set-token YOUR_MCP_TOKEN
```

### `config show`

Show the current configuration (token is partially masked).  
显示当前配置（Token 会部分遮蔽）。

```bash
edu-cli config show
```

---

## organisation – 机构管理 / Organisations

Manage exam organisations (e.g. Cambridge, Edexcel).

管理考试机构（如剑桥、Edexcel）。

### `organisation list`

List organisations.  
列出所有机构。

| Flag | Default | Description |
|---|---|---|
| `--page` | `1` | Page number / 页码 |
| `--page-size` | `20` | Results per page / 每页数量 |
| `--name` | | Filter by name / 按名称过滤 |

```bash
edu-cli organisation list
edu-cli organisation list --name Cambridge
```

### `organisation get <id>`

Get an organisation by ID.  
按 ID 获取机构详情。

```bash
edu-cli organisation get 1
```

### `organisation create`

Create a new organisation.  
创建新机构。

| Flag | Required | Description |
|---|---|---|
| `--name` | ✓ | Organisation name / 机构名称 |

```bash
edu-cli organisation create --name "Cambridge Assessment"
```

### `organisation edit`

Edit an organisation.  
编辑机构信息。

| Flag | Required | Description |
|---|---|---|
| `--id` | ✓ | Organisation ID / 机构 ID |
| `--name` | | New name / 新名称 |

```bash
edu-cli organisation edit --id 1 --name "Cambridge"
```

### `organisation delete <id>`

Delete an organisation by ID.  
按 ID 删除机构。

```bash
edu-cli organisation delete 1
```

---

## qualification – 考试管理 / Qualifications

Manage qualifications (e.g. A-Level, IGCSE).

管理考试类型（如 A-Level、IGCSE）。

### `qualification list`

List qualifications.  
列出考试列表。

| Flag | Default | Description |
|---|---|---|
| `--page` | `1` | Page number / 页码 |
| `--page-size` | `20` | Results per page / 每页数量 |
| `--organisation-id` | | Filter by organisation ID / 按机构 ID 过滤 |

```bash
edu-cli qualification list
edu-cli qualification list --organisation-id 1
```

### `qualification get <id>`

Get a qualification by ID.  
按 ID 获取考试详情。

```bash
edu-cli qualification get 2
```

### `qualification create`

Create a new qualification.  
创建新考试类型。

| Flag | Required | Description |
|---|---|---|
| `--name` | ✓ | Qualification name / 考试名称 |
| `--organisation-id` | ✓ | Organisation ID / 机构 ID |

```bash
edu-cli qualification create --name "A-Level" --organisation-id 1
```

### `qualification edit`

Edit a qualification.  
编辑考试信息。

| Flag | Required | Description |
|---|---|---|
| `--id` | ✓ | Qualification ID / 考试 ID |
| `--name` | | New name / 新名称 |
| `--organisation-id` | | New organisation ID / 新机构 ID |

```bash
edu-cli qualification edit --id 2 --name "A Level"
```

### `qualification delete <id>`

Delete a qualification by ID.  
按 ID 删除考试类型。

```bash
edu-cli qualification delete 2
```

---

## syllabus – 考纲管理 / Syllabuses

Manage syllabuses (e.g. Physics A-Level 9702).

管理考纲（如物理 A-Level 9702）。

### `syllabus list`

List syllabuses.  
列出考纲。

| Flag | Default | Description |
|---|---|---|
| `--page` | `1` | Page number / 页码 |
| `--page-size` | `20` | Results per page / 每页数量 |
| `--qualification-id` | | Filter by qualification ID / 按考试 ID 过滤 |

```bash
edu-cli syllabus list
edu-cli syllabus list --qualification-id 1
```

### `syllabus get <id>`

Get a syllabus by ID.  
按 ID 获取考纲详情。

```bash
edu-cli syllabus get 3
```

---

## chapter – 章节管理 / Chapters

Manage chapters within a syllabus.

管理考纲内的章节。

### `chapter list`

List chapters (requires `--syllabus-id`).  
列出章节（需要指定考纲 ID）。

| Flag | Default | Description |
|---|---|---|
| `--syllabus-id` | *(required)* | Syllabus ID / 考纲 ID |
| `--page` | `1` | Page number / 页码 |
| `--page-size` | `20` | Results per page / 每页数量 |
| `--parent-id` | | Filter by parent chapter ID / 按父章节 ID 过滤 |

```bash
edu-cli chapter list --syllabus-id 3
edu-cli chapter list --syllabus-id 3 --parent-id 10
```

### `chapter tree`

Get the full chapter tree for a syllabus.  
获取考纲的完整章节树。

| Flag | Default | Description |
|---|---|---|
| `--syllabus-id` | *(required)* | Syllabus ID / 考纲 ID |

```bash
edu-cli chapter tree --syllabus-id 3
```

### `chapter get <id>`

Get a chapter by ID.  
按 ID 获取章节详情。

```bash
edu-cli chapter get 10
```

### `chapter create`

Create a new chapter.  
创建新章节。

| Flag | Required | Description |
|---|---|---|
| `--name` | ✓ | Chapter name / 章节名称 |
| `--syllabus-id` | ✓ | Syllabus ID / 考纲 ID |
| `--parent-id` | | Parent chapter ID (0 = root) / 父章节 ID |

```bash
edu-cli chapter create --name "Mechanics" --syllabus-id 3
edu-cli chapter create --name "Kinematics" --syllabus-id 3 --parent-id 10
```

### `chapter edit`

Edit a chapter.  
编辑章节信息。

| Flag | Required | Description |
|---|---|---|
| `--id` | ✓ | Chapter ID / 章节 ID |
| `--name` | | New name / 新名称 |

```bash
edu-cli chapter edit --id 10 --name "Mechanics & Forces"
```

### `chapter delete <id>`

Delete a chapter by ID.  
按 ID 删除章节。

```bash
edu-cli chapter delete 10
```

---

## question – 题目管理 / Questions

Manage questions in the question bank.

管理题库中的题目。

### `question list`

List questions (requires `--syllabus-id`).  
列出题目（需要指定考纲 ID）。

| Flag | Default | Description |
|---|---|---|
| `--syllabus-id` | *(required)* | Syllabus ID / 考纲 ID |
| `--page` | `1` | Page number / 页码 |
| `--page-size` | `20` | Results per page / 每页数量 |
| `--stem` | | Filter by stem text / 按题干文字过滤 |
| `--difficult` | | Difficulty 1–5 / 难度 1–5 |
| `--status` | | Status: 1=Normal, 2=Forbidden, 3=Deleted |
| `--past-paper-id` | | Filter by past paper ID / 按真题 ID 过滤 |

```bash
edu-cli question list --syllabus-id 3
edu-cli question list --syllabus-id 3 --difficult 3 --status 1
```

### `question get <id>`

Get a question by ID.  
按 ID 获取题目详情。

```bash
edu-cli question get 100
```

### `question delete <id>`

Delete a question by ID.  
按 ID 删除题目。

```bash
edu-cli question delete 100
```

---

## paper – 试卷管理 / Papers

Manage past papers, paper codes, and paper series.

管理真题、试卷代码和试卷系列。

### `paper past list`

List past papers.  
列出真题。

| Flag | Default | Description |
|---|---|---|
| `--page` | `1` | Page number / 页码 |
| `--page-size` | `20` | Results per page / 每页数量 |
| `--syllabus-id` | | Filter by syllabus ID / 按考纲 ID 过滤 |
| `--year` | | Filter by year / 按年份过滤 |
| `--code-id` | | Filter by paper code ID / 按试卷代码 ID 过滤 |
| `--series-id` | | Filter by paper series ID / 按系列 ID 过滤 |

```bash
edu-cli paper past list
edu-cli paper past list --syllabus-id 3 --year 2023
```

### `paper past get <id>`

Get a past paper by ID.  
按 ID 获取真题详情。

```bash
edu-cli paper past get 50
```

### `paper past create`

Create a new past paper.  
创建新真题。

| Flag | Default | Description |
|---|---|---|
| `--name` | | Past paper name (required) / 真题名称（必填） |
| `--syllabus-id` | | Syllabus ID (required) / 考纲 ID（必填） |
| `--year` | | Year of the past paper (required) / 年份（必填） |
| `--code-id` | | Paper code ID (required) / 试卷代码 ID（必填） |
| `--series-id` | | Paper series ID (required) / 试卷系列 ID（必填） |

```bash
edu-cli paper past create --name "2023 真题" --syllabus-id 3 --year 2023 --code-id 7 --series-id 2
```

### `paper past edit`

Edit a past paper.  
修改真题信息。

| Flag | Default | Description |
|---|---|---|
| `--id` | | Past paper ID (required) / 真题 ID（必填） |
| `--name` | | New name / 新名称 |
| `--year` | | New year / 新年份 |
| `--code-id` | | New paper code ID / 新试卷代码 ID |
| `--series-id` | | New paper series ID / 新试卷系列 ID |

```bash
edu-cli paper past edit --id 50 --name "Updated 2023 真题"
```

### `paper past delete <id>`

Delete a past paper by ID.  
按 ID 删除真题。

```bash
edu-cli paper past delete 50
```

### `paper code list`

List paper codes.  
列出试卷代码。

| Flag | Default | Description |
|---|---|---|
| `--page` | `1` | Page number / 页码 |
| `--page-size` | `20` | Results per page / 每页数量 |
| `--syllabus-id` | | Filter by syllabus ID / 按考纲 ID 过滤 |

```bash
edu-cli paper code list --syllabus-id 3
```

### `paper code get <id>`

Get a paper code by ID.  
按 ID 获取试卷代码详情。

```bash
edu-cli paper code get 5
```

### `paper series list`

List paper series.  
列出试卷系列。

| Flag | Default | Description |
|---|---|---|
| `--page` | `1` | Page number / 页码 |
| `--page-size` | `20` | Results per page / 每页数量 |
| `--syllabus-id` | | Filter by syllabus ID / 按考纲 ID 过滤 |

```bash
edu-cli paper series list --syllabus-id 3
```

---

## exam-node – 考试节点管理 / Exam Nodes

Manage exam nodes within a syllabus. Each exam node groups related chapters and paper codes.

管理考纲内的考试节点，每个节点关联相应章节和试卷代码。

### `exam-node list`

List exam nodes for a syllabus.  
列出考纲的所有考试节点。

| Flag | Default | Description |
|---|---|---|
| `--syllabus-id` | *(required)* | Syllabus ID / 考纲 ID |

```bash
edu-cli exam-node list --syllabus-id 3
```

### `exam-node get <id>`

Get an exam node by ID.  
按 ID 获取考试节点详情。

```bash
edu-cli exam-node get 7
```

### `exam-node create`

Create an exam node for a syllabus.  
为考纲创建考试节点。

| Flag | Required | Description |
|---|---|---|
| `--syllabus-id` | ✓ | Syllabus ID / 考纲 ID |
| `--name` | ✓ | Exam node name / 节点名称 |
| `--description` | | Description / 描述 |
| `--sort-order` | | Sort order / 排序 |

```bash
edu-cli exam-node create --syllabus-id 3 --name "AS Paper 1" --sort-order 1
```

### `exam-node edit`

Update an exam node.  
更新考试节点信息。

| Flag | Required | Description |
|---|---|---|
| `--id` | ✓ | Exam node ID / 节点 ID |
| `--name` | | New name / 新名称 |
| `--description` | | New description / 新描述 |
| `--sort-order` | | New sort order / 新排序 |

```bash
edu-cli exam-node edit --id 7 --name "AS Paper 1 (Updated)"
```

### `exam-node delete <id>`

Delete an exam node by ID.  
按 ID 删除考试节点。

```bash
edu-cli exam-node delete 7
```

### `exam-node add-chapter`

Add a chapter (and all its sub-chapters) to an exam node.  
将章节及其所有子章节添加到考试节点。

| Flag | Required | Description |
|---|---|---|
| `--exam-node-id` | ✓ | Exam node ID / 节点 ID |
| `--chapter-id` | ✓ | Chapter ID / 章节 ID |

```bash
edu-cli exam-node add-chapter --exam-node-id 7 --chapter-id 10
```

### `exam-node remove-chapter`

Remove a chapter from an exam node.  
从考试节点移除章节。

| Flag | Required | Description |
|---|---|---|
| `--exam-node-id` | ✓ | Exam node ID / 节点 ID |
| `--chapter-id` | ✓ | Chapter ID / 章节 ID |

```bash
edu-cli exam-node remove-chapter --exam-node-id 7 --chapter-id 10
```

### `exam-node add-paper-code`

Add a paper code to an exam node.  
为考试节点添加试卷代码。

| Flag | Required | Description |
|---|---|---|
| `--exam-node-id` | ✓ | Exam node ID / 节点 ID |
| `--paper-code-id` | ✓ | Paper code ID / 试卷代码 ID |

```bash
edu-cli exam-node add-paper-code --exam-node-id 7 --paper-code-id 5
```

### `exam-node remove-paper-code`

Remove a paper code from an exam node.  
从考试节点移除试卷代码。

| Flag | Required | Description |
|---|---|---|
| `--exam-node-id` | ✓ | Exam node ID / 节点 ID |
| `--paper-code-id` | ✓ | Paper code ID / 试卷代码 ID |

```bash
edu-cli exam-node remove-paper-code --exam-node-id 7 --paper-code-id 5
```

---

## class – 班级管理 / Classes

Manage classes, students, and class membership.

管理班级、学生及班级成员关系。

> **班级类型 / Class Types**:  
> `1` = 教学班 (Teaching class) — a student may join multiple  
> `2` = 行政班 (Administrative class) — a student may only belong to one

### `class list`

List all classes.  
列出所有班级。

| Flag | Default | Description |
|---|---|---|
| `--page` | `1` | Page number / 页码 |
| `--page-size` | `20` | Results per page / 每页数量 |

```bash
edu-cli class list
```

### `class get <id>`

Get a class by ID.  
按 ID 获取班级详情。

```bash
edu-cli class get 1
```

### `class create`

Create a new class.  
创建新班级。

| Flag | Required | Description |
|---|---|---|
| `--name` | ✓ | Class name / 班级名称 |
| `--type` | | Class type: 1=教学班 (default), 2=行政班 |

```bash
edu-cli class create --name "Physics 2026" --type 1
edu-cli class create --name "Grade 11A" --type 2
```

### `class edit`

Edit a class.  
编辑班级信息。

| Flag | Required | Description |
|---|---|---|
| `--id` | ✓ | Class ID / 班级 ID |
| `--name` | | New name / 新名称 |

```bash
edu-cli class edit --id 1 --name "Physics Advanced 2026"
```

### `class delete <id>`

Delete a class by ID.  
按 ID 删除班级。

```bash
edu-cli class delete 1
```

### `class bind-syllabus`

Bind a syllabus to a teaching class.  
为教学班绑定考纲。

| Flag | Required | Description |
|---|---|---|
| `--class-id` | ✓ | Class ID / 班级 ID |
| `--syllabus-id` | ✓ | Syllabus ID / 考纲 ID |

```bash
edu-cli class bind-syllabus --class-id 1 --syllabus-id 3
```

### `class unbind-syllabus`

Unbind the syllabus from a teaching class.  
解除教学班的考纲绑定。

| Flag | Required | Description |
|---|---|---|
| `--class-id` | ✓ | Class ID / 班级 ID |

```bash
edu-cli class unbind-syllabus --class-id 1
```

### `class assign-teacher`

Assign a teacher to a class.  
为班级指定教师。

| Flag | Required | Description |
|---|---|---|
| `--class-id` | ✓ | Class ID / 班级 ID |
| `--teacher-id` | ✓ | Teacher user ID / 教师用户 ID |

```bash
edu-cli class assign-teacher --class-id 1 --teacher-id 42
```

### `class apply`

Apply to join a class using an invite code (student action).  
学生使用邀请码申请加入班级。

| Flag | Required | Description |
|---|---|---|
| `--invite-code` | ✓ | Invite code / 邀请码 |
| `--message` | | Optional message / 可选留言 |

```bash
edu-cli class apply --invite-code ABC12345
edu-cli class apply --invite-code ABC12345 --message "请批准我加入"
```

### `class request list`

List pending join requests for a class (admin action).  
列出班级待审核的加入申请（管理员操作）。

| Flag | Required | Description |
|---|---|---|
| `--class-id` | ✓ | Class ID / 班级 ID |
| `--page` | | Page number / 页码 |
| `--page-size` | | Results per page / 每页数量 |

```bash
edu-cli class request list --class-id 1
```

### `class request approve <request-id>`

Approve a join request.  
批准加入申请。

```bash
edu-cli class request approve 15
```

### `class request reject <request-id>`

Reject a join request.  
拒绝加入申请。

```bash
edu-cli class request reject 15
```

### `class student list`

List students in a class.  
列出班级学生。

| Flag | Required | Description |
|---|---|---|
| `--class-id` | ✓ | Class ID / 班级 ID |

```bash
edu-cli class student list --class-id 1
```

### `class student add`

Directly add a student to a class (admin only).  
直接将学生添加到班级（超级管理员专用）。

| Flag | Required | Description |
|---|---|---|
| `--class-id` | ✓ | Class ID / 班级 ID |
| `--user-id` | ✓ | User (student) ID / 学生用户 ID |

```bash
edu-cli class student add --class-id 1 --user-id 99
```

### `class student remove`

Remove a student from a class.  
从班级移除学生。

| Flag | Required | Description |
|---|---|---|
| `--class-id` | ✓ | Class ID / 班级 ID |
| `--user-id` | ✓ | User (student) ID / 学生用户 ID |

```bash
edu-cli class student remove --class-id 1 --user-id 99
```

---

## user – 用户管理 / Users

Manage platform users.

管理平台用户。

### `user list`

List users with optional filters.  
列出用户（支持多种过滤条件）。

| Flag | Default | Description |
|---|---|---|
| `--page` | `1` | Page number / 页码 |
| `--page-size` | `20` | Results per page / 每页数量 |
| `--username` | | Filter by username / 按用户名过滤 |
| `--realname` | | Filter by real name / 按真实姓名过滤 |
| `--status` | | Filter by status / 按状态过滤 |
| `--class-id` | | Filter by class ID / 按班级 ID 过滤 |
| `--student-id` | | Show specific user by ID / 按 ID 查找用户 |
| `--basic` | | Show basic info only / 仅显示基本信息 |
| `--admin-class` | | Show administrative class info / 显示行政班信息 |
| `--teaching-classes` | | Show teaching class info / 显示教学班信息 |
| `--teachers` | | Show teacher info for each class / 显示班级教师 |
| `--all` | | Show all info (default) / 显示全部信息（默认）|

```bash
edu-cli user list
edu-cli user list --username alice --basic
edu-cli user list --class-id 1 --all
```

### `user get <id>`

Get a user by ID.  
按 ID 获取用户详情。

```bash
edu-cli user get 42
```

### `user create`

Create a new user (default password: 123456).  
创建新用户（默认密码：123456）。

| Flag | Required | Description |
|---|---|---|
| `--username` | ✓ | Username / 用户名 |
| `--realname` | | Real name / 真实姓名 |
| `--email` | | Email address / 邮箱 |
| `--mobile` | | Mobile number / 手机号 |

```bash
edu-cli user create --username alice --realname "Alice Wang" --email alice@example.com
```

### `user edit`

Edit a user.  
编辑用户信息。

| Flag | Required | Description |
|---|---|---|
| `--id` | ✓ | User ID / 用户 ID |
| `--realname` | | Real name / 真实姓名 |
| `--nickname` | | Nickname / 昵称 |
| `--engname` | | English name / 英文名 |
| `--sex` | | Sex: 1=Male, 2=Female |
| `--status` | | Status: 1=Active, 2=Inactive, 3=Suspended, 4=Banned |

```bash
edu-cli user edit --id 42 --realname "Alice Li" --status 1
```

### `user delete <id>`

Delete a user by ID.  
按 ID 删除用户。

```bash
edu-cli user delete 42
```

---

## teacher – 教师管理 / Teachers

Manage teacher-class relationships.

管理教师与班级的关系。

### `teacher apply`

Apply to join a class as a teacher.  
教师申请加入班级。

| Flag | Required | Description |
|---|---|---|
| `--class-id` | ✓ | Class ID / 班级 ID |
| `--message` | | Optional message / 可选留言 |

```bash
edu-cli teacher apply --class-id 1 --message "我是该班级的物理老师"
```

### `teacher approve <application-id>`

Approve a teacher join application.  
审核通过教师加入申请。

```bash
edu-cli teacher approve 20
```

### `teacher reject <application-id>`

Reject a teacher join application.  
审核拒绝教师加入申请。

```bash
edu-cli teacher reject 20
```

### `teacher list-applications`

List teacher join applications for a class.  
列出班级的教师加入申请。

| Flag | Required | Description |
|---|---|---|
| `--class-id` | ✓ | Class ID / 班级 ID |
| `--page` | | Page number / 页码 |
| `--page-size` | | Results per page / 每页数量 |

```bash
edu-cli teacher list-applications --class-id 1
```

---

## learning-plan – 学习计划管理 / Learning Plans

Manage student learning plans with version history and template generation.

管理学生学习计划，支持版本历史和批量模板生成。

> **计划类型 / Plan Types**: `long`（长期）, `mid`（中期）, `short`（短期）

### `learning-plan list`

List learning plans.  
列出学习计划。

| Flag | Default | Description |
|---|---|---|
| `--page` | `1` | Page number / 页码 |
| `--page-size` | `20` | Results per page / 每页数量 |
| `--class-id` | | Filter by class ID / 按班级 ID 过滤 |
| `--user-id` | | Filter by user ID / 按用户 ID 过滤 |
| `--plan-type` | | Filter by type: long/mid/short / 按类型过滤 |

```bash
edu-cli learning-plan list --class-id 1
edu-cli learning-plan list --user-id 99 --plan-type long
```

### `learning-plan get <id>`

Get a learning plan by ID.  
按 ID 获取学习计划详情。

```bash
edu-cli learning-plan get 10
```

### `learning-plan create`

Create a learning plan for a student.  
为学生创建学习计划。

| Flag | Required | Description |
|---|---|---|
| `--class-id` | ✓ | Class ID / 班级 ID |
| `--user-id` | ✓ | Student user ID / 学生 ID |
| `--plan-type` | ✓ | Plan type: long/mid/short / 计划类型 |
| `--content` | | Plan content / 计划内容 |
| `--comment` | | Initial version comment / 初始版本备注 |

```bash
edu-cli learning-plan create --class-id 1 --user-id 99 --plan-type long --content "Study plan..."
```

### `learning-plan edit`

Update a learning plan (automatically creates a new version).  
更新学习计划（自动记录新版本）。

| Flag | Required | Description |
|---|---|---|
| `--id` | ✓ | Learning plan ID / 计划 ID |
| `--content` | | New content / 新内容 |
| `--comment` | | Version comment / 版本备注 |

```bash
edu-cli learning-plan edit --id 10 --content "Updated plan..." --comment "Revised after exam"
```

### `learning-plan delete <id>`

Delete a learning plan by ID.  
按 ID 删除学习计划。

```bash
edu-cli learning-plan delete 10
```

### `learning-plan versions`

List version history of a learning plan.  
查看学习计划的历史版本。

| Flag | Required | Description |
|---|---|---|
| `--plan-id` | ✓ | Learning plan ID / 计划 ID |
| `--page` | | Page number / 页码 |
| `--page-size` | | Results per page / 每页数量 |

```bash
edu-cli learning-plan versions --plan-id 10
```

### `learning-plan rollback`

Rollback a learning plan to a previous version.  
回滚学习计划到历史版本。

| Flag | Required | Description |
|---|---|---|
| `--plan-id` | ✓ | Learning plan ID / 计划 ID |
| `--version` | ✓ | Target version number / 目标版本号 |
| `--comment` | | Rollback comment / 回滚备注 |

```bash
edu-cli learning-plan rollback --plan-id 10 --version 2 --comment "Reverting to original"
```

### `learning-plan generate-template`

Batch-generate template learning plans for all students in a class.  
为班级所有学生批量生成模板学习计划。

| Flag | Required | Description |
|---|---|---|
| `--class-id` | ✓ | Target teaching class ID / 目标教学班 ID |
| `--syllabus-id` | ✓ | Syllabus ID / 考纲 ID |
| `--phase-ratios` | | Phase ratios (4 integers ≤100), e.g. `30,20,20,10` |
| `--exam-node` | | Per-exam-node schedule: `id:YYYY-MM:YYYY-MM` (repeatable) |
| `--start-month` | | Global start month YYYY-MM (when `--exam-node` not used) |
| `--end-month` | | Global end month YYYY-MM (when `--exam-node` not used) |
| `--comment` | | Version comment / 版本备注 |

**Example: Global date range / 使用全局时间范围**
```bash
edu-cli learning-plan generate-template \
  --class-id 1 \
  --syllabus-id 3 \
  --start-month 2026-01 \
  --end-month 2026-12 \
  --phase-ratios 30,20,20,10
```

**Example: Per-exam-node schedules / 按考试节点设置时间**
```bash
edu-cli learning-plan generate-template \
  --class-id 1 \
  --syllabus-id 3 \
  --exam-node 7:2026-01:2026-05 \
  --exam-node 8:2026-06:2026-11 \
  --phase-ratios 30,20,20,10 \
  --comment "2026学年学习计划"
```

---

## phase-plan – 阶段性计划管理 / Phase Plans

Manage phase plans within a learning plan. A phase plan is a time-bound study block for one exam node.

管理学习计划内的阶段性计划（每个阶段对应一个考试节点的学习周期）。

### `phase-plan list`

List phase plans for a learning plan.  
列出学习计划的所有阶段性计划。

| Flag | Required | Description |
|---|---|---|
| `--plan-id` | ✓ | Learning plan ID / 学习计划 ID |

```bash
edu-cli phase-plan list --plan-id 10
```

### `phase-plan get <id>`

Get a phase plan by ID.  
按 ID 获取阶段性计划详情。

```bash
edu-cli phase-plan get 20
```

### `phase-plan create`

Create a phase plan for a learning plan.  
为学习计划创建阶段性计划。

| Flag | Required | Description |
|---|---|---|
| `--plan-id` | ✓ | Learning plan ID / 学习计划 ID |
| `--exam-node-id` | ✓ | Exam node ID / 考试节点 ID |
| `--title` | | Phase plan title / 阶段标题 |
| `--start-date` | | Start date YYYY-MM-DD / 开始日期 |
| `--end-date` | | End date YYYY-MM-DD / 结束日期 |
| `--sort-order` | | Sort order / 排序 |

```bash
edu-cli phase-plan create \
  --plan-id 10 \
  --exam-node-id 7 \
  --title "AS Paper 1 Preparation" \
  --start-date 2026-01-01 \
  --end-date 2026-05-31
```

### `phase-plan edit`

Update a phase plan.  
更新阶段性计划。

| Flag | Required | Description |
|---|---|---|
| `--id` | ✓ | Phase plan ID / 阶段计划 ID |
| `--title` | | New title / 新标题 |
| `--start-date` | | New start date YYYY-MM-DD |
| `--end-date` | | New end date YYYY-MM-DD |
| `--sort-order` | | New sort order / 新排序 |

```bash
edu-cli phase-plan edit --id 20 --end-date 2026-06-15
```

### `phase-plan delete <id>`

Delete a phase plan by ID.  
按 ID 删除阶段性计划。

```bash
edu-cli phase-plan delete 20
```

### `phase-plan add-chapter`

Add a chapter to a phase plan.  
为阶段性计划添加章节。

| Flag | Required | Description |
|---|---|---|
| `--phase-plan-id` | ✓ | Phase plan ID / 阶段计划 ID |
| `--chapter-id` | ✓ | Chapter ID / 章节 ID |

```bash
edu-cli phase-plan add-chapter --phase-plan-id 20 --chapter-id 10
```

### `phase-plan remove-chapter`

Remove a chapter from a phase plan.  
从阶段性计划移除章节。

| Flag | Required | Description |
|---|---|---|
| `--phase-plan-id` | ✓ | Phase plan ID / 阶段计划 ID |
| `--chapter-id` | ✓ | Chapter ID / 章节 ID |

```bash
edu-cli phase-plan remove-chapter --phase-plan-id 20 --chapter-id 10
```

---

## generate-docs – 生成文档 / Documentation Generation

Generate per-command Markdown documentation for all edu-cli commands.

为所有 edu-cli 命令生成独立的 Markdown 文档文件。

```bash
# Generate docs to ./docs/cli (default)
# 生成文档到默认目录 ./docs/cli
edu-cli generate-docs

# Generate docs to a custom directory
# 生成文档到自定义目录
edu-cli generate-docs --output /path/to/output
```

| Flag | Default | Description |
|---|---|---|
| `--output`, `-o` | `./docs/cli` | Output directory / 输出目录 |

---

## 完整工作流示例 / Complete Workflow Example

The following example walks through a typical setup from scratch.  
以下示例展示从零开始的典型配置流程。

```bash
# 1. Configure backend / 配置后端
edu-cli config set-url https://api.example.com
edu-cli config set-token YOUR_TOKEN

# 2. Create an organisation and qualification / 创建机构和考试类型
edu-cli organisation create --name "Cambridge Assessment"
edu-cli qualification create --name "A-Level" --organisation-id 1

# 3. Create a syllabus / 创建考纲
# (syllabuses are typically managed via the web UI and imported)
edu-cli syllabus list --qualification-id 1

# 4. Create a class and bind a syllabus / 创建班级并绑定考纲
edu-cli class create --name "Physics 2026" --type 1
edu-cli class bind-syllabus --class-id 1 --syllabus-id 3

# 5. Create users and add them to the class / 创建用户并加入班级
edu-cli user create --username alice --realname "Alice Wang"
edu-cli class student add --class-id 1 --user-id 42

# 6. Create exam nodes / 创建考试节点
edu-cli exam-node create --syllabus-id 3 --name "AS Paper 1" --sort-order 1
edu-cli exam-node add-chapter --exam-node-id 7 --chapter-id 10

# 7. Generate learning plans for all students / 批量生成学习计划
edu-cli learning-plan generate-template \
  --class-id 1 \
  --syllabus-id 3 \
  --exam-node 7:2026-01:2026-05 \
  --phase-ratios 30,20,20,10
```

---

## 自动化文档更新 / Automated Documentation Update

To regenerate the per-command documentation (e.g. in a CI/CD pipeline):

在 CI/CD 流水线中定期更新文档：

```bash
# Regenerate all per-command docs
edu-cli generate-docs --output ./docs/cli
```

The generated files in `docs/cli/` are structured Markdown files, one per command,
and are automatically kept up to date with the cobra command definitions.

`docs/cli/` 目录下的文件按命令生成，与代码中的 cobra 命令定义保持同步。
