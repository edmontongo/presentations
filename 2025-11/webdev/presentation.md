---
marp: true
theme: default
class:
  - invert
paginate: true
---

<style>
section {
  padding-bottom: 20%;
}
</style>

# Go for Web Development

Peter Preeper
2025-11-19
ppreeper@gmail.com

---

## Slide 1: Introduction

- Go (Golang) is a modern, performant systems language designed at Google.
- Built-in concurrency, strong standard library, and simple syntax.
- Excellent for building scalable web backends and efficient web tooling.

---

## Slide 2: Why Use Go for Web Development?

- High performance (compiled language)
- Simple concurrency model with goroutines
- Powerful standard library for HTTP, JSON, templates
- Easy deployment (single static binary)
- Strong ecosystem: Gin, Echo, Fiber, Chi, Buffalo

---

## Slide 3: Basic Web Server in Go

```go
package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello, Go Web!")
    })

    http.ListenAndServe(":8080", nil)
}
```

---

## Slide 4: Web APIs With Go

- JSON serialization built-in (`encoding/json`)
- Clean routing with frameworks like Chi or Gin
- Easy middleware for logging, auth, rate-limiting

---

## Slide 5: Example REST API With Gin

```go
package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()

    r.GET("/users/:id", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "id":   c.Param("id"),
            "name": "John Doe",
        })
    })

    r.Run(":8080")
}
```

---

## Slide 6: Middleware Example

```go
func Logger() gin.HandlerFunc {
    return func(c *gin.Context) {
        log.Println("Request received:", c.FullPath())
        c.Next()
    }
}

func main() {
    r := gin.Default()
    r.Use(Logger())
    // routes...
}
```

---

## Slide 7: Web UI With Go

- Go can render templates using `html/template`.
- Suitable for server-rendered apps (SSR).
- Great for dashboards, admin panels, internal systems.

---

## Slide 8: Template Example

```go
// layout.html
'text/html'
<html>
<body>
    <h1>{{ .Title }}</h1>
    <p>{{ .Message }}</p>
</body>
</html>
```

---

```go
package main

type PageData struct {
    Title   string
    Message string
}

func main() {
    tmpl := template.Must(template.ParseFiles("layout.html"))
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        tmpl.Execute(w, PageData{Title: "Go Templates", Message: "Rendered from Go!"})
    })

    http.ListenAndServe(":8080", nil)
}
```

---

## Slide 9: WebSockets in Go

- Go's concurrency makes WebSockets efficient.
- Gorilla WebSocket is a popular library.

```go
conn, _ := upgrader.Upgrade(w, r, nil)
for {
    _, msg, _ := conn.ReadMessage()
    conn.WriteMessage(websocket.TextMessage, msg)
}
```

---

## Slide 10: Deploying Go Web Applications

- Build static binaries: `go build -o app`
- Deploy via Docker, systemd, or serverless platforms
- Minimal runtime dependencies

---

## Slide 11: Summary

- Go is fast, efficient, and great for web APIs and backend services.
- Built-in HTTP server and powerful ecosystem.
- Supports server-side HTML rendering, APIs, WebSockets.
- Easy deployment and scaling.

---

## Slide 12: Q&A
