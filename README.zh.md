<!-- TEMPLATE (ZH) BEGIN: BADGES -->

[![GitHub Workflow Status (branch)](https://img.shields.io/github/actions/workflow/status/yylego/gormzhcn/release.yml?branch=main&label=BUILD)](https://github.com/yylego/gormzhcn/actions/workflows/release.yml?query=branch%3Amain)
[![GoDoc](https://pkg.go.dev/badge/github.com/yylego/gormzhcn)](https://pkg.go.dev/github.com/yylego/gormzhcn)
[![Coverage Status](https://img.shields.io/coveralls/github/yylego/gormzhcn/main.svg)](https://coveralls.io/github/yylego/gormzhcn?branch=main)
[![Supported Go Versions](https://img.shields.io/badge/Go-1.27%2B-lightgrey.svg)](https://go.dev/)
[![GitHub Release](https://img.shields.io/github/release/yylego/gormzhcn.svg)](https://github.com/yylego/gormzhcn/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/yylego/gormzhcn)](https://goreportcard.com/report/github.com/yylego/gormzhcn)

<!-- TEMPLATE (ZH) CLOSE: BADGES -->

# GORMZHCN - GORM 中文命名把关器

**gormzhcn** 是 gormmom 命名校验的**中文本地化封装**。用母语中文字段编写 GORM 模型，再**把关**模型产出的每一个数据库名字——表名、每个列名、每个索引名——都是数据库安全的，全程通过直观的中文类型名和方法名完成。它作为一个普通单元测试运行，**完全不连接数据库**。

> 🎯 **用中文编程，交付数据库安全的表结构。** gormzhcn 在中文命名进入数据库之前，就替你抓出数据库会拒绝的那些名字。

## 生态系统

![GORM Type-Safe Ecosystem](https://github.com/yylego/gormcnm/raw/main/assets/gormcnm-ecosystem.svg)

---

<!-- TEMPLATE (ZH) BEGIN: LANGUAGE NAVIGATION -->

## 英文文档

[ENGLISH README](README.md)

<!-- TEMPLATE (ZH) CLOSE: LANGUAGE NAVIGATION -->

---

## 免责声明

使用中文编写 Go 代码在技术上可行，但在实际工程中极不推荐。这种做法不适用于任何正式或商业场景，采用它的团队或公司可能面临同行的轻视和行业的负面评价，在商业公司中尤其可能成为被业内议论的对象。该项目仅供学习和研究，请勿在正式工程中采用。

---

## 🤖 设计意图

gormzhcn 过去会替你用拼音拼出列名和索引名。现在你定义字段的那一刻，AI 助手就能写出又好又数据库安全的名字，所以生成这一层已经去掉了。

仍然会出问题的是漏网的名字：字段没写列名、索引没写名字，GORM 就会产出非 ASCII 的名字，数据库要到迁移或运行时才拒绝。gormzhcn 只做这一项检查，在不连数据库的单元测试里完成，底层用的是 [gormmom](https://github.com/yylego/gormmom)，gormzhcn 给它配上中文的类型名和方法名。

---

## 🚀 安装

```bash
go get github.com/yylego/gormzhcn
```

---

## 🌟 问题与解决方案

### ⚡ 中文母语模型的陷阱

中文字段读起来很自然，但 GORM 从中推导出的名字，可能正是数据库拒绝的名字：

```go
type T账户信息 struct {
    ID   uint   `gorm:"primaryKey"`
    Z账号 string `gorm:"uniqueIndex"` // ❌ 没写索引名 → GORM 用字段名拼索引 → 非 ASCII → 被拒绝
    N昵称 string `gorm:""`            // ❌ 没写列名 → 列名落到非 ASCII → 被拒绝
}
```

### ✅ gormzhcn 的做法

保留中文字段名，但显式给出数据库安全的列名 / 索引名 / 表名——然后让 gormzhcn 在单元测试里替你把关：

```go
type T账户信息 struct {
    ID   uint   `gorm:"primaryKey"`
    Z账号 string `gorm:"column:account;uniqueIndex:udx_account"`
    N昵称 string `gorm:"column:nickname;index:idx_nickname"`
    A年龄 int    `gorm:"column:age"`
    S状态 string `gorm:"column:status;index:idx_status"`
}

func (*T账户信息) TableName() string {
    return "accounts" // 数据库安全的 ASCII 表名
}
```

---

## 🛠️ 使用方法 —— 在单元测试里校验

gormzhcn 从不连接数据库，它作为一个普通单元测试运行：

```go
func TestValidateNaming(t *testing.T) {
    objects := []any{
        &T账户信息{},
    }

    // 批量把关：表名 / 列名 / 索引名都必须数据库安全
    require.NoError(t, gormzhcn.NewT批量校验器(objects, gormzhcn.NewT配置项()).Validate())

    // 也可以单独校验一个模型
    require.NoError(t, gormzhcn.NewT校验器(&T账户信息{}, gormzhcn.NewT配置项()).Validate())
}
```

校验失败会返回一个清晰的错误，指出是哪个结构体、哪个名字出了问题（涉及具体字段时也会一并指出）——远在任何迁移发生之前就被抓住。名字长度上限和 GORM 命名策略都可以调：

```go
v配置项 := gormzhcn.NewT配置项().
    With最大名字长度(63).                  // 名字长度上限（Postgres 为 63）
    With命名策略(schema.NamingStrategy{}) // 传入项目里 GORM 用的命名策略
```

---

## 🤝 配合生态使用

gormzhcn 负责把关命名，兄弟包负责类型安全列和 CRUD。一个典型流程，全程中文：

**1. 定义中文模型**，给出显式的数据库安全名字（见上）：一个 `T学生`，字段 `V名字`、`V性别`、`V班级编码`……

**2. 用 gormzhcn 把关命名**——不连数据库的单元测试：

```go
require.NoError(t, gormzhcn.NewT批量校验器([]any{&T学生{}}, gormzhcn.NewT配置项()).Validate())
```

**3. 用 [gormcngen](https://github.com/yylego/gormcngen) 生成类型安全列**：

```go
gormcngen.NewConfigs([]any{&T学生{}}, gormcngen.NewOptions().WithUseTagName(true), srcPath).Gen()
// → 生成带中文访问器的 T学生Columns：V名字、V性别、V班级编码……
```

**4. 用 [gormrepo](https://github.com/yylego/gormrepo) 通过中文列查询**：

```go
repo := gormrepo.NewRepo(gormclass.Use(&T学生{}))
student, err := repo.With(ctx, db).First(func(db *gorm.DB, cls *T学生Columns) *gorm.DB {
    return db.Where(cls.V名字.Eq("杨亦乐"))
})
```

生成测试会先跑校验，通过了才生成列，见 [example1](internal/examples/example1/internal/models/models.gen_test.go)。

---

## 📝 完整示例

| 示例                                   | 讲什么                                                        |
| -------------------------------------- | ------------------------------------------------------------- |
| [example1](internal/examples/example1) | 中文字段加 cnm 标签别名：通过校验，再用生成的列配合 GORM 查询 |
| [example2](internal/examples/example2) | 中文字段不写 cnm 标签：生成的列沿用字段名                     |

---

## 🔧 API 参考

| 中文名                            | 封装                 | 说明               |
| --------------------------------- | -------------------- | ------------------ |
| `NewT批量校验器(objects, 配置项)` | `gormmom.NewConfigs` | 构造批量校验器     |
| `T批量校验器.Validate()`          | `Configs.Validate`   | 校验整批模型       |
| `NewT校验器(object, 配置项)`      | `gormmom.NewConfig`  | 构造单模型校验器   |
| `T校验器.Validate()`              | `Config.Validate`    | 校验单个模型       |
| `NewT配置项()`                    | `gormmom.NewOptions` | 构造配置           |
| `With最大名字长度(n)`             | `WithMaxIdentLength` | 设置名字长度上限   |
| `With命名策略(namer)`             | `WithNamingStrategy` | 设置 GORM 命名策略 |

---

## 关联项目

探索完整的 GORM 生态系统集成包：

### 核心生态

- **[gormcnm](https://github.com/yylego/gormcnm)** - GORM 基础包，提供类型安全列和语句构建
- **[gormcngen](https://github.com/yylego/gormcngen)** - 基于 AST 的代码生成，支持类型安全 GORM 操作
- **[gormrepo](https://github.com/yylego/gormrepo)** - 仓储模式实现，遵循 GORM 最佳实践
- **[gormmom](https://github.com/yylego/gormmom)** - 原生语言 GORM 命名把关器（gormzhcn 封装的底层引擎）
- **[gormzhcn](https://github.com/yylego/gormzhcn)** - GORM 中文命名把关器（本包）

这些包涵盖 GORM 开发的各个方面：本地化、类型保护和命名校验。

---

<!-- TEMPLATE (ZH) BEGIN: STANDARD PROJECT FOOTER -->
<!-- VERSION 2025-11-25 03:52:28.131064 +0000 UTC -->

## 📄 许可证类型

MIT 许可证 - 详见 [LICENSE](LICENSE)。

---

## 💬 联系与反馈

非常欢迎贡献代码！报告 BUG、建议功能、贡献代码：

- 🐛 **问题报告？** 在 GitHub 上提交问题并附上重现步骤
- 💡 **新颖思路？** 创建 issue 讨论
- 📖 **文档疑惑？** 报告问题，帮助我们完善文档
- 🚀 **需要功能？** 分享使用场景，帮助理解需求
- ⚡ **性能瓶颈？** 报告慢操作，协助解决性能问题
- 🔧 **配置困扰？** 询问复杂设置的相关问题
- 📢 **关注进展？** 关注仓库以获取新版本和功能
- 🌟 **成功案例？** 分享这个包如何改善工作流程
- 💬 **反馈意见？** 欢迎提出建议和意见

---

## 🔧 代码贡献

新代码贡献，请遵循此流程：

1. **Fork**：在 GitHub 上 Fork 仓库（使用网页界面）
2. **克隆**：克隆 Fork 的项目（`git clone https://github.com/yourname/repo-name.git`）
3. **导航**：进入克隆的项目（`cd repo-name`）
4. **分支**：创建功能分支（`git checkout -b feature/xxx`）
5. **编码**：实现您的更改并编写全面的测试
6. **测试**：（Golang 项目）确保测试通过（`go test ./...`）并遵循 Go 代码风格约定
7. **文档**：面向用户的更改需要更新文档
8. **暂存**：暂存更改（`git add .`）
9. **提交**：提交更改（`git commit -m "Add feature xxx"`）确保向后兼容的代码
10. **推送**：推送到分支（`git push origin feature/xxx`）
11. **PR**：在 GitHub 上打开 Merge Request（在 GitHub 网页上）并提供详细描述

请确保测试通过并包含相关的文档更新。

---

## 🌟 项目支持

非常欢迎通过提交 Merge Request 和报告问题来贡献此项目。

**项目支持：**

- ⭐ **给予星标**如果项目对您有帮助
- 🤝 **分享项目**给团队成员和（golang）编程朋友
- 📝 **撰写博客**关于开发工具和工作流程 - 我们提供写作支持
- 🌟 **加入生态** - 致力于支持开源和（golang）开发场景

**祝你用这个包编程愉快！** 🎉🎉🎉

<!-- TEMPLATE (ZH) CLOSE: STANDARD PROJECT FOOTER -->
