[![GitHub Workflow Status (branch)](https://img.shields.io/github/actions/workflow/status/yylego/gormzhcn/release.yml?branch=main&label=BUILD)](https://github.com/yylego/gormzhcn/actions/workflows/release.yml?query=branch%3Amain)
[![GoDoc](https://pkg.go.dev/badge/github.com/yylego/gormzhcn)](https://pkg.go.dev/github.com/yylego/gormzhcn)
[![Coverage Status](https://img.shields.io/coveralls/github/yylego/gormzhcn/main.svg)](https://coveralls.io/github/yylego/gormzhcn?branch=main)
[![Supported Go Versions](https://img.shields.io/badge/Go-1.25+-lightgrey.svg)](https://go.dev/)
[![GitHub Release](https://img.shields.io/github/release/yylego/gormzhcn.svg)](https://github.com/yylego/gormzhcn/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/yylego/gormzhcn)](https://goreportcard.com/report/github.com/yylego/gormzhcn)

# 🇨🇳 GORMZHCN - GORM 完整中文编程接口

**gormzhcn** 是 **GORM 生态系统的完整中文编程接口**，为中文开发者提供纯中文 API 的类型安全数据库开发体验。基于强大的 gormmom 原生语言引擎构建。

> 🎯 **中文语言解放**: 使用中文原生字段名编程 - gormzhcn 连接中文业务术语和数据库需求。

## 生态系统

![GORM Type-Safe Ecosystem](https://github.com/yylego/gormcnm/raw/main/assets/gormcnm-ecosystem.svg)

---

<!-- TEMPLATE (ZH) BEGIN: LANGUAGE NAVIGATION -->

## 英文文档

[ENGLISH README](README.md)
<!-- TEMPLATE (ZH) END: LANGUAGE NAVIGATION -->

---

## 免责声明

使用中文编写 Go 代码在技术上可行，但在实际工程中极不推荐。这种做法不适用于任何正式或商业场景，采用它的团队或公司可能面临同行的轻视和行业的负面评价，在商业公司中尤其可能成为被业内议论的对象。该项目仅供学习和研究，请勿在正式工程中采用。

---

## 🚀 安装

```bash
go get github.com/yylego/gormzhcn
```

---

## 🔄 技术对比

| 生态系统            | Java MyBatis Plus  | Python SQLAlchemy | Go GORM 生态系统   |
|--------------------|--------------------|--------------------|-------------------|
| **类型安全列**      | `Example::getName` | `Example.name`     | `cls.Name.Eq()`   |
| **代码生成**        | ✅ 插件支持         | ✅ 反射机制         | ✅ AST 精准生成    |
| **仓储模式**        | ✅ BaseMapper       | ✅ Session API     | ✅ GormRepo        |
| **原生语言**        | 🟡 有限支持         | 🟡 有限支持        | ✅ 完整支持        |

---

## 🌟 问题与解决方案

### ⚡ 传统方式
```go
// ❌ 常见方式：开发者只能使用英文命名
type Account struct {
    ID       uint   `gorm:"primaryKey"`
    Username string `gorm:"column:username;uniqueIndex"`
    Nickname string `gorm:"column:nickname;index"`
    Age      int    `gorm:"column:age"`
    PhoneNum string `gorm:"column:phone_num"`
    Mailbox  string `gorm:"column:mailbox"`
    Address  string `gorm:"column:address"`
    Status   string `gorm:"column:status;index"`
}
```

### ✅ GORMZHCN 解决方案
```go
// ✅ GORMZHCN：使用中文原生业务字段名编程！
type T账户信息 struct {
    ID   uint   `gorm:"primaryKey"`
    Z账号 string `gorm:"uniqueIndex"`
    N昵称 string `gorm:"index"`
    A年龄 int    `gorm:""`
    D电话 string `gorm:""`
    E邮箱 string `gorm:""`
    J住址 string `gorm:""`
    S状态 string `gorm:"index"`
}

func (*T账户信息) TableName() string {
    return "accounts" // 数据库兼容表名
}
```

---

## 💡 核心特性

### 🌍 语言支持
- **中文 Unicode 兼容**：原生中文字段名和业务术语
- **智能映射**：中文名到数据库安全列名的转换
- **确定性输出**：跨运行的一致生成结果
- **Go 导出规则**：大写前缀（如 `V用户名`、`V姓名`）

### 🤖 智能生成
- **自动 GORM 标签**：`column`、`index`、`uniqueIndex` 等
- **命名模式**：`S63 编码`、`拼音`、自定义格式
- **索引管理**：带自动生成的描述性名称
- **标签保留**：添加新标签时保持现有标签

### ⚠️ 使用要求
- **结构体命名**：以大写开头 `type T账户信息`（Go 导出规则）
- **TableName() 方法**：必需，避免 `regexp does not match` 错误
- **两遍生成**：第一遍更新标签，第二遍生成方法

---

## 🎯 快速开始

### 1. 自动标签生成

gormzhcn 执行后，结构体获得数据库兼容的列标签：

```go
// 生成带数据库兼容列名的结构体
type T账户信息 struct {
    ID    uint   `gorm:"primaryKey"`
    Z账号  string `gorm:"column:z_zhang_hao;uniqueIndex"`
    N昵称  string `gorm:"column:n_ni_cheng;index"`
    A年龄  int    `gorm:"column:a_nian_ling"`
    D电话  string `gorm:"column:d_dian_hua"`
    E邮箱  string `gorm:"column:e_you_xiang"`
    J住址  string `gorm:"column:j_zhu_zhi"`
    S状态  string `gorm:"column:s_zhuang_tai;index"`
}
```

### 2. 生成命令

```bash
# 第一步：为原生语言字段生成 GORM 标签
go test -v -run TestGen/GenGormMom

# 第二步：生成类型安全列方法（使用 gormcngen）
go test -v -run TestGen/GenGormCnm
```

### 3. 使用类型安全查询
```go
// 创建 repo
repo := gormrepo.NewGormRepo(&T账户信息{}, (&T账户信息{}).Columns())

// 查询 - First（按账号）
account, err := repo.With(ctx, db).First(func(db *gorm.DB, cls *T账户信息Columns) *gorm.DB {
    return db.Where(cls.Z账号.Eq("wang-xiao-ming"))
})

// 查询 - First（按昵称）
account, err := repo.With(ctx, db).First(func(db *gorm.DB, cls *T账户信息Columns) *gorm.DB {
    return db.Where(cls.N昵称.Eq("王小明"))
})

// 查询 - Find
accounts, err := repo.With(ctx, db).Find(func(db *gorm.DB, cls *T账户信息Columns) *gorm.DB {
    return db.Where(cls.A年龄.Gte(18))
})

// 查询 - FindPage
accounts, err := repo.With(ctx, db).FindPage(
    func(db *gorm.DB, cls *T账户信息Columns) *gorm.DB {
        return db.Where(cls.A年龄.Gte(18))
    },
    func(cls *T账户信息Columns) gormcnm.OrderByBottle {
        return cls.ID.OrderByBottle("DESC")
    },
    &gormrepo.Pagination{Limit: 10, Offset: 0},
)

// 创建
err := repo.With(ctx, db).Create(&T账户信息{Z账号: "han-mei-mei", N昵称: "韩梅梅", A年龄: 25})

// 更新
err := repo.With(ctx, db).Updates(
    func(db *gorm.DB, cls *T账户信息Columns) *gorm.DB {
        return db.Where(cls.ID.Eq(1))
    },
    func(cls *T账户信息Columns) map[string]interface{} {
        return cls.Kw(cls.A年龄.Kv(26)).AsMap()
    },
)

// 删除
err := repo.With(ctx, db).DeleteW(func(db *gorm.DB, cls *T账户信息Columns) *gorm.DB {
    return db.Where(cls.ID.Eq(1))
})
```

### 4. 自动标签生成原理

创建测试文件来运行标签生成，将中文字段名转换为数据库兼容的列标签：

```go
func TestGen(t *testing.T) {
    // 第一步：为中文语言字段生成 GORM 标签
    t.Run("GenGormMom", func(t *testing.T) {
        srcPath := runpath.PARENT.Join("models.go")
        t.Log(srcPath)

        param := gormzhcn.NewT表结构V2[T账户信息](srcPath)
        cfg := gormzhcn.NewT编码器(param, gormzhcn.NewT配置项())
        cfg.Gen替换源码()
    })

    // 第二步：生成类型安全列方法
    t.Run("GenGormCnm", func(t *testing.T) {
        absPath := osmustexist.FILE(runtestpath.SrcPath(t))
        t.Log(absPath)

        options := gormcngen.NewOptions().
            WithColumnClassExportable(true).
            WithColumnsMethodRecvName("T").
            WithColumnsCheckFieldType(true)
        cfg := gormcngen.NewConfigs([]interface{}{&T账户信息{}}, options, absPath)
        cfg.Gen()
    })
}
```

查看 [examples](internal/examples/) 目录获取完整配置。

---

## 🔧 配置和使用

```go
// 同时处理多个结构体
srcPath := runpath.PARENT.Join("models.go")
p表结构1 := gormzhcn.NewT表结构V2[T学生](srcPath)
p表结构2 := gormzhcn.NewT表结构V2[T班级](srcPath)

cfg := gormzhcn.NewT批量编码器([]*gormzhcn.T表结构{
    p表结构1,
    p表结构2,
}, gormzhcn.NewT配置项())

result := cfg.Gen替换源码()
if result.Has有变化() {
    // 标签已更新
}
```

### 配置选项

```go
options := gormzhcn.NewT配置项().
    With命名总标签名("mom").           // 系统标签名
    With列名模式标签("mcp").           // 列命名子标签
    With列名命名样式(pattern).         // 列命名模式
    With重新命名索引(true)             // 自动重命名索引
```

查看 [examples](internal/examples/) 目录获取完整示例。

---

## 🔧 完整 API 参考

### T编码器（单结构体）

| 中文方法 | 封装方法 | 参数 | 返回值 | 描述 |
|---------|--------|------|--------|------|
| `NewT编码器` | `gormmom.NewConfig` | `v表结构, v配置项` | `*T编码器` | 创建编码器 |
| `Gen替换源码` | `config.Generate` | 无 | 无 | 生成并替换 |
| `Get预览代码` | `config.Preview` | 无 | `*R替换的结果` | 预览变化 |

### T批量编码器（批量处理）

| 中文方法 | 封装方法 | 参数 | 返回值 | 描述 |
|---------|--------|------|--------|------|
| `NewT批量编码器` | `NewT编码器` 循环 | `v表结构s, v配置项` | `*T批量编码器` | 创建批量编码器 |
| `Gen替换源码` | `configs.Generate` | 无 | `*R批量替换的结果` | 批量生成 |
| `Get预览代码` | `configs.Preview` | 无 | `*R批量替换的结果` | 批量预览 |
| `Validate验证标签` | `configs.ValidateGormTags` | 无 | `error` | 验证标签 |

### T表结构（表结构定义）

| 中文方法 | 封装方法 | 参数 | 返回值 | 描述 |
|---------|--------|------|--------|------|
| `NewT表结构` | `gormmom.NewGormStruct` | `sourcePath, structName, schema` | `*T表结构` | 基础创建 |
| `NewT表结构V2[T]` | `gormmom.ParseStruct[T]` | `sourcePath` | `*T表结构` | 从泛型创建 |
| `NewT表结构V3` | `gormmom.ParseObject` | `sourcePath, object` | `*T表结构` | 从对象创建 |
| `NewT表结构s` | `gormmom.ParseObjects` | `root, objects` | `[]*T表结构` | 批量创建 |

### T配置项（配置选项）

| 中文方法 | 封装方法 | 参数 | 返回值 | 描述 |
|---------|--------|------|--------|------|
| `NewT配置项` | `gormmom.NewOptions` | 无 | `*T配置项` | 创建配置 |
| `With命名总标签名` | `WithTagName` | `systemTagName` | `*T配置项` | 设置系统标签 |
| `With列名模式标签` | `WithSubTagName` | `subTagName` | `*T配置项` | 设置列标签 |
| `With列名命名样式` | `WithColumnPattern` | `pattern` | `*T配置项` | 设置列模式 |
| `With默认列名样式` | `WithDefaultColumnPattern` | `pattern` | `*T配置项` | 设置默认列 |
| `With跳过基本列名` | `WithSkipBasicColumnName` | `skip` | `*T配置项` | 跳过基本字段 |
| `With重新命名索引` | `WithAutoIndexName` | `rename` | `*T配置项` | 自动重命名索引 |
| `With索引命名样式` | `WithIndexPattern` | `pattern` | `*T配置项` | 设置索引模式 |
| `With默认索引样式` | `WithDefaultIndexPattern` | `pattern` | `*T配置项` | 设置默认索引 |

---

## 关联项目

探索完整的 GORM 生态系统集成包：

### 核心生态

- **[gormcnm](https://github.com/yylego/gormcnm)** - GORM 基础包，提供类型安全列和语句构建
- **[gormcngen](https://github.com/yylego/gormcngen)** - 基于 AST 的代码生成，支持类型安全 GORM 操作
- **[gormrepo](https://github.com/yylego/gormrepo)** - 仓储模式实现，遵循 GORM 最佳实践
- **[gormmom](https://github.com/yylego/gormmom)** - 原生语言 GORM 标签生成引擎，支持名称映射
- **[gormzhcn](https://github.com/yylego/gormzhcn)** - 完整的 GORM 中文编程接口（本包）

这些包涵盖 GORM 开发的各个方面：本地化、类型保护和代码生成。

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

<!-- TEMPLATE (ZH) END: STANDARD PROJECT FOOTER -->

---

## GitHub Stars

[![Stargazers](https://starchart.cc/yylego/gormzhcn.svg?variant=adaptive)](https://starchart.cc/yylego/gormzhcn)

