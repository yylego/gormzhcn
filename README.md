[![GitHub Workflow Status (branch)](https://img.shields.io/github/actions/workflow/status/yylego/gormzhcn/release.yml?branch=main&label=BUILD)](https://github.com/yylego/gormzhcn/actions/workflows/release.yml?query=branch%3Amain)
[![GoDoc](https://pkg.go.dev/badge/github.com/yylego/gormzhcn)](https://pkg.go.dev/github.com/yylego/gormzhcn)
[![Coverage Status](https://img.shields.io/coveralls/github/yylego/gormzhcn/main.svg)](https://coveralls.io/github/yylego/gormzhcn?branch=main)
[![Supported Go Versions](https://img.shields.io/badge/Go-1.25+-lightgrey.svg)](https://go.dev/)
[![GitHub Release](https://img.shields.io/github/release/yylego/gormzhcn.svg)](https://github.com/yylego/gormzhcn/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/yylego/gormzhcn)](https://goreportcard.com/report/github.com/yylego/gormzhcn)

# 🇨🇳 GORMZHCN - Complete Chinese Programming Interface with GORM

**gormzhcn** is the **complete Chinese programming interface** that provides native Chinese localization with the GORM ecosystem, enabling Chinese development with complete type-safe GORM operations. Built on top of the robust gormmom native language engine.

> 🎯 **Chinese Language Liberation**: Code in Chinese using native field names - gormzhcn bridges Chinese business terms and database requirements.

## Ecosystem

![GORM Type-Safe Ecosystem](https://github.com/yylego/gormcnm/raw/main/assets/gormcnm-ecosystem.svg)

---

<!-- TEMPLATE (EN) BEGIN: LANGUAGE NAVIGATION -->

## CHINESE README

[中文说明](README.zh.md)
<!-- TEMPLATE (EN) END: LANGUAGE NAVIGATION -->

---

## DISCLAIMER

Writing Go code in Chinese is a viable technique, but something to avoid in production engineering. This approach should not be used in serious and business settings. Teams and companies that embrace it could face contempt from peers and negative judgment across the profession. In business companies, this practice is even more prone to becoming a target of public criticism. This project is dedicated to research and academic studies. Do not use this approach in production.

---

## 🚀 Installation

```bash
go get github.com/yylego/gormzhcn
```

---

## 🔄 Tech Comparison

| Ecosystem             | Java MyBatis Plus  | Python SQLAlchemy | Go GORM Ecosystem  |
|-----------------------|--------------------|-------------------|--------------------|
| **Type-Safe Columns** | `Example::getName` | `Example.name`    | `cls.Name.Eq()`    |
| **Code Generation**   | ✅ Plugin support   | ✅ Reflection      | ✅ AST precision    |
| **Repo Pattern**      | ✅ BaseMapper       | ✅ Session API     | ✅ GormRepo         |
| **Native Language**   | 🟡 Limited         | 🟡 Limited        | ✅ Complete support |

---

## 🌟 The Problem & Solution

### ⚡ Standard Approach
```go
// ❌ Common approach: Developers constrained to English naming
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

### ✅ GORMZHCN Solution
```go
// ✅ GORMZHCN: Program in Chinese with native business field names!
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
    return "accounts" // Database-compatible table name
}
```

---

## 💡 Core Features

### 🌍 Language Support
- **Chinese Unicode compatible**: Native Chinese field names and business terms
- **Smart mapping**: Chinese names to database-safe column names
- **Deterministic output**: Consistent generation across runs
- **Go export rules**: Uppercase prefixes (e.g., `V用户名`, `V姓名`)

### 🤖 Smart Generation
- **Auto GORM tags**: `column`, `index`, `uniqueIndex` and more
- **Naming patterns**: `S63 encoding`, `Pinyin`, custom formats
- **Index management**: Descriptive names with auto-generation
- **Tag preservation**: Maintains existing tags while adding new ones

### ⚠️ Requirements
- **Struct naming**: Start with uppercase `type T账户信息` (Go export rules)
- **TableName() method**: Required to avoid `regexp does not match` errors
- **Two-pass generation**: First run updates tags, second run generates methods

---

## 🎯 Getting Started

### 1. Auto Tag Generation

Once gormzhcn executes, the struct gets database-compatible column tags:

```go
// Generated with database-compatible column names
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

### 2. Generate Commands

```bash
# Step 1: Generate GORM tags with native language fields
go test -v -run TestGen/GenGormMom

# Step 2: Generate type-safe column methods (with gormcngen)
go test -v -run TestGen/GenGormCnm
```

### 3. Use Type-Safe Queries
```go
// Create repo
repo := gormrepo.NewGormRepo(&T账户信息{}, (&T账户信息{}).Columns())

// Select - First (by username)
account, err := repo.With(ctx, db).First(func(db *gorm.DB, cls *T账户信息Columns) *gorm.DB {
    return db.Where(cls.Z账号.Eq("wang-xiao-ming"))
})

// Select - First (by nickname)
account, err := repo.With(ctx, db).First(func(db *gorm.DB, cls *T账户信息Columns) *gorm.DB {
    return db.Where(cls.N昵称.Eq("王小明"))
})

// Select - Find
accounts, err := repo.With(ctx, db).Find(func(db *gorm.DB, cls *T账户信息Columns) *gorm.DB {
    return db.Where(cls.A年龄.Gte(18))
})

// Select - FindPage
accounts, err := repo.With(ctx, db).FindPage(
    func(db *gorm.DB, cls *T账户信息Columns) *gorm.DB {
        return db.Where(cls.A年龄.Gte(18))
    },
    func(cls *T账户信息Columns) gormcnm.OrderByBottle {
        return cls.ID.OrderByBottle("DESC")
    },
    &gormrepo.Pagination{Limit: 10, Offset: 0},
)

// Create
err := repo.With(ctx, db).Create(&T账户信息{Z账号: "han-mei-mei", N昵称: "韩梅梅", A年龄: 25})

// Update
err := repo.With(ctx, db).Updates(
    func(db *gorm.DB, cls *T账户信息Columns) *gorm.DB {
        return db.Where(cls.ID.Eq(1))
    },
    func(cls *T账户信息Columns) map[string]interface{} {
        return cls.Kw(cls.A年龄.Kv(26)).AsMap()
    },
)

// Delete
err := repo.With(ctx, db).DeleteW(func(db *gorm.DB, cls *T账户信息Columns) *gorm.DB {
    return db.Where(cls.ID.Eq(1))
})
```

### 4. How Auto Tag Generation Works

Create a test file to run the tag generation that transforms Chinese field names into database-compatible column tags:

```go
func TestGen(t *testing.T) {
    // Step 1: Generate GORM tags with Chinese language fields
    t.Run("GenGormMom", func(t *testing.T) {
        srcPath := runpath.PARENT.Join("models.go")
        t.Log(srcPath)

        param := gormzhcn.NewT表结构V2[T账户信息](srcPath)
        cfg := gormzhcn.NewT编码器(param, gormzhcn.NewT配置项())
        cfg.Gen替换源码()
    })

    // Step 2: Generate type-safe column methods
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

See [examples](internal/examples/) directories with complete setup.

---

## 🔧 Configuration and Usage

```go
// Process multiple structs at the same time
srcPath := runpath.PARENT.Join("models.go")
p表结构1 := gormzhcn.NewT表结构V2[T学生](srcPath)
p表结构2 := gormzhcn.NewT表结构V2[T班级](srcPath)

cfg := gormzhcn.NewT批量编码器([]*gormzhcn.T表结构{
    p表结构1,
    p表结构2,
}, gormzhcn.NewT配置项())

result := cfg.Gen替换源码()
if result.Has有变化() {
    // Tags were updated
}
```

### Configuration Options

```go
options := gormzhcn.NewT配置项().
    With命名总标签名("mom").           // System tag name
    With列名模式标签("mcp").           // Column naming sub-tag
    With列名命名样式(pattern).         // Column naming pattern
    With重新命名索引(true)             // Auto rename indexes
```

See the [examples](internal/examples/) directories with complete working examples.

---

## 🔧 Complete API Reference

### T编码器 (Single Struct)

| Chinese Method | Wraps | Parameters | Returns | Description |
|---------------|-------|-----------|---------|-------------|
| `NewT编码器` | `gormmom.NewConfig` | `v表结构, v配置项` | `*T编码器` | Create engine |
| `Gen替换源码` | `config.Generate` | none | none | Generate and replace |
| `Get预览代码` | `config.Preview` | none | `*R替换的结果` | Preview changes |

### T批量编码器 (Batch Processing)

| Chinese Method | Wraps | Parameters | Returns | Description |
|---------------|-------|-----------|---------|-------------|
| `NewT批量编码器` | `NewT编码器` loop | `v表结构s, v配置项` | `*T批量编码器` | Create batch engine |
| `Gen替换源码` | `configs.Generate` | none | `*R批量替换的结果` | Batch generate |
| `Get预览代码` | `configs.Preview` | none | `*R批量替换的结果` | Batch preview |
| `Validate验证标签` | `configs.ValidateGormTags` | none | `error` | Validate tags |

### T表结构 (Table Structure)

| Chinese Method | Wraps | Parameters | Returns | Description |
|---------------|-------|-----------|---------|-------------|
| `NewT表结构` | `gormmom.NewGormStruct` | `sourcePath, structName, schema` | `*T表结构` | Create with params |
| `NewT表结构V2[T]` | `gormmom.ParseStruct[T]` | `sourcePath` | `*T表结构` | Create from generic |
| `NewT表结构V3` | `gormmom.ParseObject` | `sourcePath, object` | `*T表结构` | Create from object |
| `NewT表结构s` | `gormmom.ParseObjects` | `root, objects` | `[]*T表结构` | Batch create |

### T配置项 (Configuration)

| Chinese Method | Wraps | Parameters | Returns | Description |
|---------------|-------|-----------|---------|-------------|
| `NewT配置项` | `gormmom.NewOptions` | none | `*T配置项` | Create config |
| `With命名总标签名` | `WithTagName` | `systemTagName` | `*T配置项` | Set system tag |
| `With列名模式标签` | `WithSubTagName` | `subTagName` | `*T配置项` | Set column tag |
| `With列名命名样式` | `WithColumnPattern` | `pattern` | `*T配置项` | Set column pattern |
| `With默认列名样式` | `WithDefaultColumnPattern` | `pattern` | `*T配置项` | Set default column |
| `With跳过基本列名` | `WithSkipBasicColumnName` | `skip` | `*T配置项` | Skip basic fields |
| `With重新命名索引` | `WithAutoIndexName` | `rename` | `*T配置项` | Auto rename index |
| `With索引命名样式` | `WithIndexPattern` | `pattern` | `*T配置项` | Set index pattern |
| `With默认索引样式` | `WithDefaultIndexPattern` | `pattern` | `*T配置项` | Set default index |

---

## Related Projects

Explore the complete GORM ecosystem with these integrated packages:

### Core Ecosystem

- **[gormcnm](https://github.com/yylego/gormcnm)** - GORM base package with type-safe columns and statement building
- **[gormcngen](https://github.com/yylego/gormcngen)** - AST-based code generation with type-safe GORM operations
- **[gormrepo](https://github.com/yylego/gormrepo)** - Repo pattern implementation with GORM best practices
- **[gormmom](https://github.com/yylego/gormmom)** - Native language GORM tag generation engine with name mapping
- **[gormzhcn](https://github.com/yylego/gormzhcn)** - Complete Chinese programming interface with GORM (this package)

These packages handle GORM development aspects: localization, type protection, and code generation.

---

<!-- TEMPLATE (EN) BEGIN: STANDARD PROJECT FOOTER -->
<!-- VERSION 2025-11-25 03:52:28.131064 +0000 UTC -->

## 📄 License

MIT License - see [LICENSE](LICENSE).

---

## 💬 Contact & Feedback

Contributions are welcome! Report bugs, suggest features, and contribute code:

- 🐛 **Mistake reports?** Open an issue on GitHub with reproduction steps
- 💡 **Fresh ideas?** Create an issue to discuss
- 📖 **Documentation confusing?** Report it so we can improve
- 🚀 **Need new features?** Share the use cases to help us understand requirements
- ⚡ **Performance issue?** Help us optimize through reporting slow operations
- 🔧 **Configuration problem?** Ask questions about complex setups
- 📢 **Follow project progress?** Watch the repo to get new releases and features
- 🌟 **Success stories?** Share how this package improved the workflow
- 💬 **Feedback?** We welcome suggestions and comments

---

## 🔧 Development

New code contributions, follow this process:

1. **Fork**: Fork the repo on GitHub (using the webpage UI).
2. **Clone**: Clone the forked project (`git clone https://github.com/yourname/repo-name.git`).
3. **Navigate**: Navigate to the cloned project (`cd repo-name`)
4. **Branch**: Create a feature branch (`git checkout -b feature/xxx`).
5. **Code**: Implement the changes with comprehensive tests
6. **Testing**: (Golang project) Ensure tests pass (`go test ./...`) and follow Go code style conventions
7. **Documentation**: Update documentation to support client-facing changes
8. **Stage**: Stage changes (`git add .`)
9. **Commit**: Commit changes (`git commit -m "Add feature xxx"`) ensuring backward compatible code
10. **Push**: Push to the branch (`git push origin feature/xxx`).
11. **PR**: Open a merge request on GitHub (on the GitHub webpage) with detailed description.

Please ensure tests pass and include relevant documentation updates.

---

## 🌟 Support

Welcome to contribute to this project via submitting merge requests and reporting issues.

**Project Support:**

- ⭐ **Give GitHub stars** if this project helps you
- 🤝 **Share with teammates** and (golang) programming friends
- 📝 **Write tech blogs** about development tools and workflows - we provide content writing support
- 🌟 **Join the ecosystem** - committed to supporting open source and the (golang) development scene

**Have Fun Coding with this package!** 🎉🎉🎉

<!-- TEMPLATE (EN) END: STANDARD PROJECT FOOTER -->

---

## GitHub Stars

[![Stargazers](https://starchart.cc/yylego/gormzhcn.svg?variant=adaptive)](https://starchart.cc/yylego/gormzhcn)

