# Awesome Golang

<p align="center">
  <img src="https://go.dev/blog/gopher/gopher.png" width="200" alt="Go Gopher" />
</p>

A curated, comprehensive, and constantly updated collection of the best Golang (Go) learning resources, tools, libraries, books, courses, and tips — with a special focus on Persian/Farsi-speaking developers.

This repository aims to be the ultimate one-stop shop for anyone learning or mastering Go in the Persian community.

## Why this repository?

- Gathers scattered Persian resources in one place
- Organized from beginner to advanced levels
- Actively maintained and kept up-to-date
- Includes real-world tools, libraries, and best practices used by Iranian Go developers

## Table of Contents

- [Getting Started with Go](#getting-started-with-go)
- [Persian Learning Resources](#persian-learning-resources)
- [Recommended Books](#recommended-books)
- [Video Courses & Tutorials](#video-courses--tutorials)
- [Awesome Go Libraries (Curated for Persian Developers)](#awesome-go-libraries)
- [Project Ideas for Practice & Portfolio](#project-ideas)
- [Tools & Development Environment](#tools--development-environment)
- [Tips, Tricks & Best Practices](#tips-tricks--best-practices)
- [Code Style & Standards](#code-style--standards)
- [Persian Go Community](#persian-go-community)
- [Contributing](#contributing)

## Getting Started with Go

| Level              | Recommended Path                                              | Estimated Time |
|--------------------|---------------------------------------------------------------|----------------|
| Complete Beginner  | Official Tour of Go + a Persian video course                  | 2–4 weeks      |
| Beginner → Intermediate | "Learning Go" book + small projects + Effective Go         | 2–4 months     |
| Intermediate → Advanced | Go by Example + deep dives + real-world projects + reading popular package source code | 6+ months |

## Persian Learning Resources (Highly Recommended)

- [Roocket.ir Golang Series](https://roocket.ir/series/golang)
- [SabzLearn Comprehensive Go Course](https://sabzlearn.ir/course/golang/)
- [Quera College Go Path](https://quera.org/college)
- YouTube channels & Persian blogs (feel free to suggest more!)

## Recommended Books

### Free / Persian
- *Learning Go* (Persian translation or local editions)
- Go Bootcamp (partial translations available)

### English (Must-Read)
- *The Go Programming Language* – Alan A. A. Donovan & Brian W. Kernighan
- *Concurrency in Go* – Katherine Cox-Buday
- *100 Go Mistakes and How to Avoid Them* – Teiva Harsanyi
- *Efficient Go* – Bartłomiej Płotka

## Awesome Go Libraries (Popular in the Iranian Ecosystem)

| Category               | Top Pick                         | Great Alternatives              | Notes (in Persian context)                  |
|------------------------|----------------------------------|----------------------------------|--------------------------------------------|
| Web Framework / API    | [gin-gonic/gin](https://github.com/gin-gonic/gin) | Fiber, Echo, Chi               | Most popular in Iranian startups           |
| ORM                    | [go-gorm/gorm](https://github.com/go-gorm/gorm)   | Ent, sqlc, Bun                 | Easiest to get started                     |
| Raw Database Driver    | [jackc/pgx](https://github.com/jackc/pgx)         | —                               | Best PostgreSQL driver                     |
| CLI Apps               | [spf13/cobra](https://github.com/spf13/cobra) + viper | urfave/cli                     | Industry standard                          |
| Logging                | rs/zerolog or go.uber.org/zap    | logrus                           | Fast & structured                          |
| Configuration          | [spf13/viper](https://github.com/spf13/viper)     | koanf, godotenv                  | Extremely versatile                        |
| Validation             | [go-playground/validator](https://github.com/go-playground/validator) | ozzo-validation            | Widely used in Persian projects            |
| JSON Handling          | [json-iterator/go](https://github.com/json-iterator/go) | std encoding/json            | Much faster than standard library          |

## Project Ideas

(Great for learning and building your portfolio)

1. REST API with Gin + GORM + PostgreSQL
2. CLI tool using Cobra (e.g., file organizer, currency converter)
3. Real-time chat app with WebSockets (Gorilla or Melody)
4. Microservices with gRPC + Protocol Buffers
5. URL shortener (like bit.ly clone)
6. Telegram bot using telebot or tgbotapi

## Tools & Development Environment

- Editor: VS Code + Go extension / GoLand
- Formatter: `gofumpt`, `golines`
- Linter: `golangci-lint`
- Testing: built-in `testing` + `testify`
- Mocking: `gomock`, `mockgen`
- Debugging: Delve (`dlv`)

## Tips, Tricks & Best Practices

- Always use `go mod`
- Prefer `context` package everywhere
- Never ignore errors
- Use structured logging in production
- Write table-driven tests
- Keep your `main.go` thin

## Code Style & Standards

- Follow [Effective Go](https://golang.org/doc/effective_go)
- Use [Uber Go Style Guide](https://github.com/uber-go/guide)
- Run `golangci-lint` in CI

## Persian Go Community

- Telegram groups: @golang_ir, @golangfa
- Discord / Slack communities (add yours!)
- Local meetups & conferences

### Sample Code in This Repository

In addition to curated lists and resources, **this project also includes real, runnable sample source code**.  
You can clone the repository and directly use, run, study, or build upon these code examples for your own projects and learning.

These samples will be continuously improved and expanded over time — more topics, more real-world examples, and better patterns will be added gradually.

## Contributing

Contributions are very welcome! This list will never be perfect without your help.

How to contribute:

1. Found a great resource/course/library? → Open a Pull Request
2. Broken link or outdated info? → Open an Issue
3. Have a better structure idea? → Let’s discuss!

Every contribution matters — even fixing a typo.

If you find this repository helpful, please give it a ⭐ and share it with your Go friends!

Happy coding in Go! 🚀
