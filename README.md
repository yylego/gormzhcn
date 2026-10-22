<!-- TEMPLATE (EN) BEGIN: BADGES -->

[![GitHub Workflow Status (branch)](https://img.shields.io/github/actions/workflow/status/yylego/gormzhcn/release.yml?branch=main&label=BUILD)](https://github.com/yylego/gormzhcn/actions/workflows/release.yml?query=branch%3Amain)
[![GoDoc](https://pkg.go.dev/badge/github.com/yylego/gormzhcn)](https://pkg.go.dev/github.com/yylego/gormzhcn)
[![Coverage Status](https://img.shields.io/coveralls/github/yylego/gormzhcn/main.svg)](https://coveralls.io/github/yylego/gormzhcn?branch=main)
[![Supported Go Versions](https://img.shields.io/badge/Go-1.27%2B-lightgrey.svg)](https://go.dev/)
[![GitHub Release](https://img.shields.io/github/release/yylego/gormzhcn.svg)](https://github.com/yylego/gormzhcn/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/yylego/gormzhcn)](https://goreportcard.com/report/github.com/yylego/gormzhcn)

<!-- TEMPLATE (EN) CLOSE: BADGES -->

# 🇨🇳 GORMZHCN - Chinese GORM Naming Guardian

**gormzhcn** puts a **Chinese-localized interface** on gormmom naming validation. Write GORM models with native Chinese fields, then **guard** that each database name a struct produces — the table name, each column name, each index name — stays database-safe, through intuitive Chinese type and method names. It runs as a plain unit test and touches no database.

> 🎯 **Code in Chinese, ship database-safe schemas.** gormzhcn catches Chinese names that a database would reject — long before such names reach the database.

## Ecosystem

![GORM Type-Safe Ecosystem](https://github.com/yylego/gormcnm/raw/main/assets/gormcnm-ecosystem.svg)

---

<!-- TEMPLATE (EN) BEGIN: LANGUAGE NAVIGATION -->

## CHINESE README

[中文说明](README.zh.md)

<!-- TEMPLATE (EN) CLOSE: LANGUAGE NAVIGATION -->

---

## DISCLAIMER

Writing Go code in Chinese is a viable technique, but something to avoid in production engineering. This approach should not be used in serious and business settings. Teams and companies that embrace it could face contempt from peers and negative judgment across the profession. In business companies, this practice often becomes a target of public criticism. This project is dedicated to research and academic studies. Do not use this approach in production.

---

## 🤖 How gormzhcn became just a guardian (the AI-era simplification)

gormzhcn used to wrap a **generation engine**: you wrote Chinese fields, tags left blank, and it auto-synthesized the column names and index names on its own (pinyin and pattern-based encodings).

**The AI era made that generation pointless.** The moment you define a field, an AI assistant gives you a good, database-safe column name — in context, a cut above a fixed rule engine. A machine-synthesized `z_zhang_hao` was not what someone wanted. So the entire generation step lost its reason to exist.

What AI does **not** hand you on its own is the part that bites: **validation**. When you program in Chinese, GORM might turn a field into a non-ASCII column name, and it can build a composite index name out of Chinese field names — names the database rejects. That failure shows up late, at migration time, at runtime.

So gormzhcn dropped generation outright and now does just one thing: **catch those naming problems in a database-free unit test**, so none of them blow up in production.

- ❌ Gone: code generation, pinyin/pattern column synthesis, the old code-gen types.
- ✅ Kept: one job — validate that the table / column / index names are database-safe.
- ✅ Zero database connection — a pure static check that runs as a plain unit test.

---

## 🚀 Installation

```bash
go get github.com/yylego/gormzhcn
```

---

## 🌟 The Problem & Solution

### ⚡ The trap of Chinese-language models

Chinese fields read great, but the names GORM derives from them can be names the database refuses:

```go
type T账户信息 struct {
    ID   uint   `gorm:"primaryKey"`
    Z账号 string `gorm:"uniqueIndex"` // ❌ no index name → GORM builds one from the field name → non-ASCII → rejected
    N昵称 string `gorm:""`            // ❌ no column name → column falls back to a non-ASCII name → rejected
}
```

### ✅ The gormzhcn approach

Keep the Chinese field names, but give explicit database-safe column / index / table names — and let gormzhcn guard them in a unit test:

```go
type T账户信息 struct {
    ID   uint   `gorm:"primaryKey"`
    Z账号 string `gorm:"column:account;uniqueIndex:udx_account"`
    N昵称 string `gorm:"column:nickname;index:idx_nickname"`
    A年龄 int    `gorm:"column:age"`
    S状态 string `gorm:"column:status;index:idx_status"`
}

func (*T账户信息) TableName() string {
    return "accounts" // database-safe ASCII table name
}
```

---

## 🛠️ Usage — validate in a unit test

gormzhcn opens no database connection; it runs as a plain unit test:

```go
func TestValidateNaming(t *testing.T) {
    objects := []any{
        &T账户信息{},
    }

    // Guard a batch: table / column / index names must be database-safe.
    require.NoError(t, gormzhcn.NewT模型校验器(objects, gormzhcn.NewT配置项()).Validate())

    // Guard a single struct on its own.
    require.NoError(t, gormzhcn.NewT校验器(&T账户信息{}, gormzhcn.NewT配置项()).Validate())
}
```

A failing name comes back as a precise message — naming the struct, the offending name, and the field when one applies — caught long before a migration. The name length cap and the GORM naming strategy are tunable:

```go
v配置项 := gormzhcn.NewT配置项().
    With最大名字长度(63).        // name length cap (Postgres uses 63)
    With命名策略(schema.NamingStrategy{}) // pass a custom strategy when a project needs one
```

---

## 🤝 Use with the ecosystem

gormzhcn guards the naming; its sibling packages handle type-safe columns and repo-style CRUD. A common flow, start to finish in Chinese:

**1. Define a Chinese struct** with explicit database-safe names (see above): a `T学生` with fields `V名字`, `V性别`, `V班级编码`, ...

**2. Guard the naming with gormzhcn** — a database-free unit test:

```go
require.NoError(t, gormzhcn.NewT模型校验器([]any{&T学生{}}, gormzhcn.NewT配置项()).Validate())
```

**3. Generate type-safe columns with [gormcngen](https://github.com/yylego/gormcngen)**:

```go
gormcngen.NewConfigs([]any{&T学生{}}, gormcngen.NewOptions().WithUseTagName(true), srcPath).Gen()
// → produces T学生Columns with Chinese columns: V名字, V性别, V班级编码, ...
```

**4. Run queries with [gormrepo](https://github.com/yylego/gormrepo)** through the Chinese columns:

```go
repo := gormrepo.NewRepo(gormclass.Use(&T学生{}))
student, err := repo.Repo(db).First(func(db *gorm.DB, cls *T学生Columns) *gorm.DB {
    return db.Where(cls.V名字.Eq("杨亦乐"))
})
```

Complete runnable examples live in [examples](internal/examples/).

---

## 🔧 API Reference

| Chinese name                      | Wraps                | Description                             |
| --------------------------------- | -------------------- | --------------------------------------- |
| `NewT模型校验器(objects, 配置项)` | `gormmom.NewConfigs` | Build a batch check                     |
| `T模型校验器.Validate()`          | `Configs.Validate`   | Validate the whole batch                |
| `NewT校验器(object, 配置项)`      | `gormmom.NewConfig`  | Build a single-struct check             |
| `T校验器.Validate()`              | `Config.Validate`    | Validate one struct                     |
| `NewT配置项()`                    | `gormmom.NewOptions` | Build a config                          |
| `With最大名字长度(n)`             | `WithMaxNameLength`  | Set the name length cap                 |
| `With命名策略(namer)`             | `WithNamingStrategy` | Set the GORM naming strategy            |
| `GetOptions()`                    | —                    | Access the underlying `gormmom.Options` |

---

## Related Projects

Explore the complete GORM ecosystem with these integrated packages:

### Core Ecosystem

- **[gormcnm](https://github.com/yylego/gormcnm)** - GORM base package with type-safe columns and statement building
- **[gormcngen](https://github.com/yylego/gormcngen)** - AST-based code generation with type-safe GORM operations
- **[gormrepo](https://github.com/yylego/gormrepo)** - Repo pattern implementation with GORM best practices
- **[gormmom](https://github.com/yylego/gormmom)** - Native language GORM naming guardian (the engine gormzhcn wraps)
- **[gormzhcn](https://github.com/yylego/gormzhcn)** - Chinese GORM naming guardian (this package)

These packages handle GORM development aspects: localization, type protection, and naming validation.

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

<!-- TEMPLATE (EN) CLOSE: STANDARD PROJECT FOOTER -->
