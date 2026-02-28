# edu-cli 使用文档（中文版）

edu-cli 是 LTEdu 平台的命令行客户端，可通过后端 API 管理考纲、试卷、用户、班级、学习计划等资源。

---

## 目录

- [安装与配置](#安装与配置)
- [config – 配置管理](#config--配置管理)
- [organisation – 机构管理](#organisation--机构管理)
- [qualification – 考试管理](#qualification--考试管理)
- [syllabus – 考纲管理](#syllabus--考纲管理)
- [chapter – 章节管理](#chapter--章节管理)
- [question – 题目管理](#question--题目管理)
- [paper – 试卷管理](#paper--试卷管理)
- [exam-node – 考试节点管理](#exam-node--考试节点管理)
- [class – 班级管理](#class--班级管理)
- [user – 用户管理](#user--用户管理)
- [teacher – 教师管理](#teacher--教师管理)
- [learning-plan – 学习计划管理](#learning-plan--学习计划管理)
- [phase-plan – 阶段性计划管理](#phase-plan--阶段性计划管理)
- [generate-docs – 生成文档](#generate-docs--生成文档)

---

## 安装与配置

### 环境变量

| 变量 | 说明 |
|---|---|
| `EDU_BASE_URL` | 后端地址，例如 `https://api.example.com` |
| `EDU_TOKEN` | 从 Web UI 获取的 JWT 认证 Token |

```bash
export EDU_BASE_URL=https://api.example.com
export EDU_TOKEN=your_jwt_token_here
```

### 快速配置命令

```bash
# 设置后端地址
edu-cli config set-url https://api.example.com

# 设置认证 Token
edu-cli config set-token YOUR_JWT_TOKEN

# 查看当前配置
edu-cli config show
```

---

## config – 配置管理

管理 edu-cli 的配置项。

### `config set-url <url>`

设置后端服务地址。

```bash
edu-cli config set-url https://api.example.com
```

### `config set-token <token>`

设置认证 Token（从 Web UI 获取的 JWT）。

```bash
edu-cli config set-token eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...
```

### `config show`

显示当前配置（Token 会部分遮蔽）。

```bash
edu-cli config show
```

---

## organisation – 机构管理

管理考试机构（如剑桥、Edexcel）。

### `organisation list`

列出所有机构。

| 参数 | 默认值 | 说明 |
|---|---|---|
| `--page` | `1` | 页码 |
| `--page-size` | `20` | 每页数量 |
| `--name` | | 按名称过滤 |

```bash
edu-cli organisation list
edu-cli organisation list --name Cambridge
```

### `organisation get <id>`

按 ID 获取机构详情。

```bash
edu-cli organisation get 1
```

### `organisation create`

创建新机构。

| 参数 | 是否必填 | 说明 |
|---|---|---|
| `--name` | ✓ | 机构名称 |

```bash
edu-cli organisation create --name "Cambridge Assessment"
```

### `organisation edit`

编辑机构信息。

| 参数 | 是否必填 | 说明 |
|---|---|---|
| `--id` | ✓ | 机构 ID |
| `--name` | | 新名称 |

```bash
edu-cli organisation edit --id 1 --name "Cambridge"
```

### `organisation delete <id>`

按 ID 删除机构。

```bash
edu-cli organisation delete 1
```

---

## qualification – 考试管理

管理考试类型（如 A-Level、IGCSE）。

### `qualification list`

列出考试类型。

| 参数 | 默认值 | 说明 |
|---|---|---|
| `--page` | `1` | 页码 |
| `--page-size` | `20` | 每页数量 |
| `--organisation-id` | | 按机构 ID 过滤 |

```bash
edu-cli qualification list
edu-cli qualification list --organisation-id 1
```

### `qualification get <id>`

按 ID 获取考试类型详情。

```bash
edu-cli qualification get 2
```

### `qualification create`

创建新考试类型。

| 参数 | 是否必填 | 说明 |
|---|---|---|
| `--name` | ✓ | 考试名称 |
| `--organisation-id` | ✓ | 机构 ID |

```bash
edu-cli qualification create --name "A-Level" --organisation-id 1
```

### `qualification edit`

编辑考试类型信息。

| 参数 | 是否必填 | 说明 |
|---|---|---|
| `--id` | ✓ | 考试 ID |
| `--name` | | 新名称 |
| `--organisation-id` | | 新机构 ID |

```bash
edu-cli qualification edit --id 2 --name "A Level"
```

### `qualification delete <id>`

按 ID 删除考试类型。

```bash
edu-cli qualification delete 2
```

---

## syllabus – 考纲管理

管理考纲（如物理 A-Level 9702）。

### `syllabus list`

列出考纲。

| 参数 | 默认值 | 说明 |
|---|---|---|
| `--page` | `1` | 页码 |
| `--page-size` | `20` | 每页数量 |
| `--qualification-id` | | 按考试 ID 过滤 |

```bash
edu-cli syllabus list
edu-cli syllabus list --qualification-id 1
```

### `syllabus get <id>`

按 ID 获取考纲详情。

```bash
edu-cli syllabus get 3
```

---

## chapter – 章节管理

管理考纲内的章节。

### `chapter list`

列出章节（必须指定考纲 ID）。

| 参数 | 默认值 | 说明 |
|---|---|---|
| `--syllabus-id` | *(必填)* | 考纲 ID |
| `--page` | `1` | 页码 |
| `--page-size` | `20` | 每页数量 |
| `--parent-id` | | 按父章节 ID 过滤 |

```bash
edu-cli chapter list --syllabus-id 3
edu-cli chapter list --syllabus-id 3 --parent-id 10
```

### `chapter tree`

获取考纲的完整章节树形结构。

| 参数 | 默认值 | 说明 |
|---|---|---|
| `--syllabus-id` | *(必填)* | 考纲 ID |

```bash
edu-cli chapter tree --syllabus-id 3
```

### `chapter get <id>`

按 ID 获取章节详情。

```bash
edu-cli chapter get 10
```

### `chapter create`

创建新章节。

| 参数 | 是否必填 | 说明 |
|---|---|---|
| `--name` | ✓ | 章节名称 |
| `--syllabus-id` | ✓ | 考纲 ID |
| `--parent-id` | | 父章节 ID（0 表示根节点）|

```bash
edu-cli chapter create --name "力学" --syllabus-id 3
edu-cli chapter create --name "运动学" --syllabus-id 3 --parent-id 10
```

### `chapter edit`

编辑章节信息。

| 参数 | 是否必填 | 说明 |
|---|---|---|
| `--id` | ✓ | 章节 ID |
| `--name` | | 新名称 |

```bash
edu-cli chapter edit --id 10 --name "力学与受力分析"
```

### `chapter delete <id>`

按 ID 删除章节。

```bash
edu-cli chapter delete 10
```

---

## question – 题目管理

管理题库中的题目。

### `question list`

列出题目（必须指定考纲 ID）。

| 参数 | 默认值 | 说明 |
|---|---|---|
| `--syllabus-id` | *(必填)* | 考纲 ID |
| `--page` | `1` | 页码 |
| `--page-size` | `20` | 每页数量 |
| `--stem` | | 按题干文字过滤 |
| `--difficult` | | 难度 1–5 |
| `--status` | | 状态：1=正常，2=禁用，3=已删除 |
| `--past-paper-id` | | 按真题 ID 过滤 |

```bash
edu-cli question list --syllabus-id 3
edu-cli question list --syllabus-id 3 --difficult 3 --status 1
```

### `question get <id>`

按 ID 获取题目详情。

```bash
edu-cli question get 100
```

### `question delete <id>`

按 ID 删除题目。

```bash
edu-cli question delete 100
```

---

## paper – 试卷管理

管理真题、试卷代码和试卷系列。

### `paper past list`

列出真题。

| 参数 | 默认值 | 说明 |
|---|---|---|
| `--page` | `1` | 页码 |
| `--page-size` | `20` | 每页数量 |
| `--syllabus-id` | | 按考纲 ID 过滤 |
| `--year` | | 按年份过滤 |
| `--code-id` | | 按试卷代码 ID 过滤 |
| `--series-id` | | 按系列 ID 过滤 |

```bash
edu-cli paper past list
edu-cli paper past list --syllabus-id 3 --year 2023
```

### `paper past get <id>`

按 ID 获取真题详情。

```bash
edu-cli paper past get 50
```

### `paper code list`

列出试卷代码。

| 参数 | 默认值 | 说明 |
|---|---|---|
| `--page` | `1` | 页码 |
| `--page-size` | `20` | 每页数量 |
| `--syllabus-id` | | 按考纲 ID 过滤 |

```bash
edu-cli paper code list --syllabus-id 3
```

### `paper code get <id>`

按 ID 获取试卷代码详情。

```bash
edu-cli paper code get 5
```

### `paper series list`

列出试卷系列。

| 参数 | 默认值 | 说明 |
|---|---|---|
| `--page` | `1` | 页码 |
| `--page-size` | `20` | 每页数量 |
| `--syllabus-id` | | 按考纲 ID 过滤 |

```bash
edu-cli paper series list --syllabus-id 3
```

---

## exam-node – 考试节点管理

管理考纲内的考试节点，每个节点关联相应章节和试卷代码。

### `exam-node list`

列出考纲的所有考试节点。

| 参数 | 默认值 | 说明 |
|---|---|---|
| `--syllabus-id` | *(必填)* | 考纲 ID |

```bash
edu-cli exam-node list --syllabus-id 3
```

### `exam-node get <id>`

按 ID 获取考试节点详情。

```bash
edu-cli exam-node get 7
```

### `exam-node create`

为考纲创建考试节点。

| 参数 | 是否必填 | 说明 |
|---|---|---|
| `--syllabus-id` | ✓ | 考纲 ID |
| `--name` | ✓ | 节点名称 |
| `--description` | | 描述 |
| `--sort-order` | | 排序 |

```bash
edu-cli exam-node create --syllabus-id 3 --name "AS Paper 1" --sort-order 1
```

### `exam-node edit`

更新考试节点信息。

| 参数 | 是否必填 | 说明 |
|---|---|---|
| `--id` | ✓ | 节点 ID |
| `--name` | | 新名称 |
| `--description` | | 新描述 |
| `--sort-order` | | 新排序 |

```bash
edu-cli exam-node edit --id 7 --name "AS Paper 1（更新）"
```

### `exam-node delete <id>`

按 ID 删除考试节点。

```bash
edu-cli exam-node delete 7
```

### `exam-node add-chapter`

将章节及其所有子章节添加到考试节点。

| 参数 | 是否必填 | 说明 |
|---|---|---|
| `--exam-node-id` | ✓ | 节点 ID |
| `--chapter-id` | ✓ | 章节 ID（自动包含所有子章节）|

```bash
edu-cli exam-node add-chapter --exam-node-id 7 --chapter-id 10
```

### `exam-node remove-chapter`

从考试节点移除章节。

| 参数 | 是否必填 | 说明 |
|---|---|---|
| `--exam-node-id` | ✓ | 节点 ID |
| `--chapter-id` | ✓ | 章节 ID |

```bash
edu-cli exam-node remove-chapter --exam-node-id 7 --chapter-id 10
```

### `exam-node add-paper-code`

为考试节点添加试卷代码。

| 参数 | 是否必填 | 说明 |
|---|---|---|
| `--exam-node-id` | ✓ | 节点 ID |
| `--paper-code-id` | ✓ | 试卷代码 ID |

```bash
edu-cli exam-node add-paper-code --exam-node-id 7 --paper-code-id 5
```

### `exam-node remove-paper-code`

从考试节点移除试卷代码。

| 参数 | 是否必填 | 说明 |
|---|---|---|
| `--exam-node-id` | ✓ | 节点 ID |
| `--paper-code-id` | ✓ | 试卷代码 ID |

```bash
edu-cli exam-node remove-paper-code --exam-node-id 7 --paper-code-id 5
```

---

## class – 班级管理

管理班级、学生及班级成员关系。

> **班级类型说明**：  
> `1` = 教学班 — 学生可加入多个  
> `2` = 行政班 — 每位学生只能属于一个

### `class list`

列出所有班级。

| 参数 | 默认值 | 说明 |
|---|---|---|
| `--page` | `1` | 页码 |
| `--page-size` | `20` | 每页数量 |

```bash
edu-cli class list
```

### `class get <id>`

按 ID 获取班级详情。

```bash
edu-cli class get 1
```

### `class create`

创建新班级（自动生成邀请码）。

| 参数 | 是否必填 | 说明 |
|---|---|---|
| `--name` | ✓ | 班级名称 |
| `--type` | | 班级类型：1=教学班（默认），2=行政班 |

```bash
edu-cli class create --name "物理2026届" --type 1
edu-cli class create --name "高二甲班" --type 2
```

### `class edit`

编辑班级信息。

| 参数 | 是否必填 | 说明 |
|---|---|---|
| `--id` | ✓ | 班级 ID |
| `--name` | | 新名称 |

```bash
edu-cli class edit --id 1 --name "物理高级班2026届"
```

### `class delete <id>`

按 ID 删除班级。

```bash
edu-cli class delete 1
```

### `class bind-syllabus`

为教学班绑定考纲。

| 参数 | 是否必填 | 说明 |
|---|---|---|
| `--class-id` | ✓ | 班级 ID |
| `--syllabus-id` | ✓ | 考纲 ID |

```bash
edu-cli class bind-syllabus --class-id 1 --syllabus-id 3
```

### `class unbind-syllabus`

解除教学班的考纲绑定。

| 参数 | 是否必填 | 说明 |
|---|---|---|
| `--class-id` | ✓ | 班级 ID |

```bash
edu-cli class unbind-syllabus --class-id 1
```

### `class assign-teacher`

为班级指定教师。

| 参数 | 是否必填 | 说明 |
|---|---|---|
| `--class-id` | ✓ | 班级 ID |
| `--teacher-id` | ✓ | 教师用户 ID |

```bash
edu-cli class assign-teacher --class-id 1 --teacher-id 42
```

### `class apply`

学生使用邀请码申请加入班级。

| 参数 | 是否必填 | 说明 |
|---|---|---|
| `--invite-code` | ✓ | 邀请码 |
| `--message` | | 留言（可选）|

```bash
edu-cli class apply --invite-code ABC12345
edu-cli class apply --invite-code ABC12345 --message "请批准我加入"
```

### `class request list`

列出班级待审核的加入申请（管理员操作）。

| 参数 | 是否必填 | 说明 |
|---|---|---|
| `--class-id` | ✓ | 班级 ID |
| `--page` | | 页码 |
| `--page-size` | | 每页数量 |

```bash
edu-cli class request list --class-id 1
```

### `class request approve <request-id>`

批准加入申请。

```bash
edu-cli class request approve 15
```

### `class request reject <request-id>`

拒绝加入申请。

```bash
edu-cli class request reject 15
```

### `class student list`

列出班级中的学生。

| 参数 | 是否必填 | 说明 |
|---|---|---|
| `--class-id` | ✓ | 班级 ID |

```bash
edu-cli class student list --class-id 1
```

### `class student add`

直接将学生加入班级（超级管理员专用）。

| 参数 | 是否必填 | 说明 |
|---|---|---|
| `--class-id` | ✓ | 班级 ID |
| `--user-id` | ✓ | 学生用户 ID |

```bash
edu-cli class student add --class-id 1 --user-id 99
```

### `class student remove`

从班级移除学生。

| 参数 | 是否必填 | 说明 |
|---|---|---|
| `--class-id` | ✓ | 班级 ID |
| `--user-id` | ✓ | 学生用户 ID |

```bash
edu-cli class student remove --class-id 1 --user-id 99
```

---

## user – 用户管理

管理平台用户。

### `user list`

列出用户（支持多种过滤条件）。

| 参数 | 默认值 | 说明 |
|---|---|---|
| `--page` | `1` | 页码 |
| `--page-size` | `20` | 每页数量 |
| `--username` | | 按用户名过滤 |
| `--realname` | | 按真实姓名过滤 |
| `--status` | | 按状态过滤 |
| `--class-id` | | 按班级 ID 过滤 |
| `--student-id` | | 按 ID 查找特定用户 |
| `--basic` | | 仅显示基本信息 |
| `--admin-class` | | 显示行政班信息 |
| `--teaching-classes` | | 显示教学班信息 |
| `--teachers` | | 显示各班级的教师信息 |
| `--all` | | 显示全部信息（不指定任何展示参数时的默认行为）|

```bash
edu-cli user list
edu-cli user list --username alice --basic
edu-cli user list --class-id 1 --all
```

### `user get <id>`

按 ID 获取用户详情。

```bash
edu-cli user get 42
```

### `user create`

创建新用户（默认密码：123456）。

| 参数 | 是否必填 | 说明 |
|---|---|---|
| `--username` | ✓ | 用户名 |
| `--realname` | | 真实姓名 |
| `--email` | | 邮箱地址 |
| `--mobile` | | 手机号 |

```bash
edu-cli user create --username alice --realname "王小明" --email alice@example.com
```

### `user edit`

编辑用户信息。

| 参数 | 是否必填 | 说明 |
|---|---|---|
| `--id` | ✓ | 用户 ID |
| `--realname` | | 真实姓名 |
| `--nickname` | | 昵称 |
| `--engname` | | 英文名 |
| `--sex` | | 性别：1=男，2=女 |
| `--status` | | 状态：1=正常，2=停用，3=暂停，4=封禁 |

```bash
edu-cli user edit --id 42 --realname "王晓明" --status 1
```

### `user delete <id>`

按 ID 删除用户。

```bash
edu-cli user delete 42
```

---

## teacher – 教师管理

管理教师与班级的关系。

### `teacher apply`

教师申请加入班级。

| 参数 | 是否必填 | 说明 |
|---|---|---|
| `--class-id` | ✓ | 班级 ID |
| `--message` | | 留言（可选）|

```bash
edu-cli teacher apply --class-id 1 --message "我是该班级的物理老师"
```

### `teacher approve <application-id>`

审核通过教师加入申请。

```bash
edu-cli teacher approve 20
```

### `teacher reject <application-id>`

审核拒绝教师加入申请。

```bash
edu-cli teacher reject 20
```

### `teacher list-applications`

列出班级的教师加入申请。

| 参数 | 是否必填 | 说明 |
|---|---|---|
| `--class-id` | ✓ | 班级 ID |
| `--page` | | 页码 |
| `--page-size` | | 每页数量 |

```bash
edu-cli teacher list-applications --class-id 1
```

---

## learning-plan – 学习计划管理

管理学生学习计划，支持版本历史记录和批量模板生成。

> **计划类型**：`long`（长期计划）、`mid`（中期计划）、`short`（短期计划）

### `learning-plan list`

列出学习计划。

| 参数 | 默认值 | 说明 |
|---|---|---|
| `--page` | `1` | 页码 |
| `--page-size` | `20` | 每页数量 |
| `--class-id` | | 按班级 ID 过滤 |
| `--user-id` | | 按用户 ID 过滤 |
| `--plan-type` | | 按类型过滤：long/mid/short |

```bash
edu-cli learning-plan list --class-id 1
edu-cli learning-plan list --user-id 99 --plan-type long
```

### `learning-plan get <id>`

按 ID 获取学习计划详情。

```bash
edu-cli learning-plan get 10
```

### `learning-plan create`

为学生创建学习计划。

| 参数 | 是否必填 | 说明 |
|---|---|---|
| `--class-id` | ✓ | 班级 ID |
| `--user-id` | ✓ | 学生用户 ID |
| `--plan-type` | ✓ | 计划类型：long/mid/short |
| `--content` | | 计划内容 |
| `--comment` | | 初始版本备注 |

```bash
edu-cli learning-plan create --class-id 1 --user-id 99 --plan-type long --content "学习计划内容..."
```

### `learning-plan edit`

更新学习计划（自动记录新版本）。

| 参数 | 是否必填 | 说明 |
|---|---|---|
| `--id` | ✓ | 学习计划 ID |
| `--content` | | 新内容 |
| `--comment` | | 版本备注 |

```bash
edu-cli learning-plan edit --id 10 --content "更新后的计划内容..." --comment "考后修订"
```

### `learning-plan delete <id>`

按 ID 删除学习计划。

```bash
edu-cli learning-plan delete 10
```

### `learning-plan versions`

查看学习计划的历史版本列表。

| 参数 | 是否必填 | 说明 |
|---|---|---|
| `--plan-id` | ✓ | 学习计划 ID |
| `--page` | | 页码 |
| `--page-size` | | 每页数量 |

```bash
edu-cli learning-plan versions --plan-id 10
```

### `learning-plan rollback`

回滚学习计划到历史版本。

| 参数 | 是否必填 | 说明 |
|---|---|---|
| `--plan-id` | ✓ | 学习计划 ID |
| `--version` | ✓ | 目标版本号 |
| `--comment` | | 回滚备注 |

```bash
edu-cli learning-plan rollback --plan-id 10 --version 2 --comment "恢复到原始版本"
```

### `learning-plan generate-template`

为班级所有学生批量生成模板学习计划（长期、中期、短期三类）。

| 参数 | 是否必填 | 说明 |
|---|---|---|
| `--class-id` | ✓ | 目标教学班 ID |
| `--syllabus-id` | ✓ | 考纲 ID |
| `--phase-ratios` | | 4 个逗号分隔的阶段比例（总和 ≤100），默认 `30,20,20,10` |
| `--exam-node` | | 单独指定考试节点时间段：`id:YYYY-MM:YYYY-MM`（可重复）|
| `--start-month` | | 全局开始月份 YYYY-MM（不使用 `--exam-node` 时必填）|
| `--end-month` | | 全局结束月份 YYYY-MM（不使用 `--exam-node` 时必填）|
| `--comment` | | 版本备注 |

**示例：使用全局时间范围**
```bash
edu-cli learning-plan generate-template \
  --class-id 1 \
  --syllabus-id 3 \
  --start-month 2026-01 \
  --end-month 2026-12 \
  --phase-ratios 30,20,20,10
```

**示例：按考试节点分别设置时间**
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

## phase-plan – 阶段性计划管理

管理学习计划内的阶段性计划，每个阶段对应一个考试节点的学习周期。

### `phase-plan list`

列出学习计划的所有阶段性计划。

| 参数 | 是否必填 | 说明 |
|---|---|---|
| `--plan-id` | ✓ | 学习计划 ID |

```bash
edu-cli phase-plan list --plan-id 10
```

### `phase-plan get <id>`

按 ID 获取阶段性计划详情。

```bash
edu-cli phase-plan get 20
```

### `phase-plan create`

为学习计划创建阶段性计划。

| 参数 | 是否必填 | 说明 |
|---|---|---|
| `--plan-id` | ✓ | 学习计划 ID |
| `--exam-node-id` | ✓ | 考试节点 ID |
| `--title` | | 阶段标题 |
| `--start-date` | | 开始日期 YYYY-MM-DD |
| `--end-date` | | 结束日期 YYYY-MM-DD |
| `--sort-order` | | 排序 |

```bash
edu-cli phase-plan create \
  --plan-id 10 \
  --exam-node-id 7 \
  --title "AS Paper 1 备考阶段" \
  --start-date 2026-01-01 \
  --end-date 2026-05-31
```

### `phase-plan edit`

更新阶段性计划。

| 参数 | 是否必填 | 说明 |
|---|---|---|
| `--id` | ✓ | 阶段计划 ID |
| `--title` | | 新标题 |
| `--start-date` | | 新开始日期 YYYY-MM-DD |
| `--end-date` | | 新结束日期 YYYY-MM-DD |
| `--sort-order` | | 新排序 |

```bash
edu-cli phase-plan edit --id 20 --end-date 2026-06-15
```

### `phase-plan delete <id>`

按 ID 删除阶段性计划。

```bash
edu-cli phase-plan delete 20
```

### `phase-plan add-chapter`

为阶段性计划添加章节。

| 参数 | 是否必填 | 说明 |
|---|---|---|
| `--phase-plan-id` | ✓ | 阶段计划 ID |
| `--chapter-id` | ✓ | 章节 ID |

```bash
edu-cli phase-plan add-chapter --phase-plan-id 20 --chapter-id 10
```

### `phase-plan remove-chapter`

从阶段性计划移除章节。

| 参数 | 是否必填 | 说明 |
|---|---|---|
| `--phase-plan-id` | ✓ | 阶段计划 ID |
| `--chapter-id` | ✓ | 章节 ID |

```bash
edu-cli phase-plan remove-chapter --phase-plan-id 20 --chapter-id 10
```

---

## generate-docs – 生成文档

为所有 edu-cli 命令生成独立的 Markdown 文档文件（每个命令单独一个文件）。

```bash
# 生成文档到默认目录 ./docs/cli
edu-cli generate-docs

# 生成文档到自定义目录
edu-cli generate-docs --output /path/to/output
```

| 参数 | 默认值 | 说明 |
|---|---|---|
| `--output`, `-o` | `./docs/cli` | 输出目录 |

---

## 完整工作流示例

以下示例展示了从零开始的典型配置和使用流程：

```bash
# 第一步：配置后端连接
edu-cli config set-url https://api.example.com
edu-cli config set-token YOUR_TOKEN

# 第二步：创建机构和考试类型
edu-cli organisation create --name "剑桥大学考评部"
edu-cli qualification create --name "A-Level" --organisation-id 1

# 第三步：查看可用考纲
edu-cli syllabus list --qualification-id 1

# 第四步：创建班级并绑定考纲
edu-cli class create --name "物理2026届" --type 1
edu-cli class bind-syllabus --class-id 1 --syllabus-id 3

# 第五步：创建用户并加入班级
edu-cli user create --username alice --realname "王小明"
edu-cli class student add --class-id 1 --user-id 42

# 第六步：创建考试节点并关联章节
edu-cli exam-node create --syllabus-id 3 --name "AS Paper 1" --sort-order 1
edu-cli exam-node add-chapter --exam-node-id 7 --chapter-id 10

# 第七步：为班级所有学生批量生成学习计划
edu-cli learning-plan generate-template \
  --class-id 1 \
  --syllabus-id 3 \
  --exam-node 7:2026-01:2026-05 \
  --phase-ratios 30,20,20,10
```

---

## 自动化文档更新

在 CI/CD 流水线中定期重新生成文档：

```bash
edu-cli generate-docs --output ./docs/cli
```

`docs/cli/` 目录下的文件按命令独立生成，每次运行都会自动同步最新的命令定义。
