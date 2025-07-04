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

# Why Go? An Introduction to GoLang

Peter Preeper
2025-06-18
ppreeper@gmail.com

---

## What is Go?

- Created at Google in 2009
- Designed for _simplicity_, _performance_, _concurrency_
- Open-source, community-driven

<!--
- Developed by Robert Griesemer, Rob Pike, and Ken Thompson
- First released in November 2009
- Combines efficiency of C with readability of Python
- Created to solve Google's large-scale development challenges
- Rapidly growing ecosystem and community support
- Currently at version 1.24 (as of 2025), soon to be 1.25 in August
-->

---

## The Creators' Legacy

- Robert Griesemer: V8 JavaScript engine, Java HotSpot VM
- Rob Pike: Unix, Plan 9, UTF-8 co-creator
- Ken Thompson: Unix, B language, C language, UTF-8, Regular expressions

<!--
These guys had seriously big brains
Much of what people use to build systems was invented or co-invented by these ones
Their basis setup the conservative limits of Go
-->

---

## Key Features

- Simple syntax, minimal keywords
- Strong static typing
- Built-in concurrency with goroutines and channels
- Garbage collection for memory management
- Long term compatiblity

<style>
  code {
    font-size: 25px;
  }
</style>

```go
break, default, func, interface, select, case, defer, go,
map, struct, chan, else, goto, package, switch, const,
fallthrough, if, range, type, continue, for, import, return, var
```

<!--
- Go has only 25 keywords, making it easy to learn and read

- Strong typing catches errors at compile time, improving reliability
- Static typing also enables better tooling and code completion

- Goroutines are lightweight threads (< 4KB overhead)
- Channels enable safe communication between goroutines
- Can run thousands of goroutines concurrently

- Garbage collection removes memory management burden from developers
- Memory management is automatic but still efficient

- Go 1.0 compatibility promise ensures code won't break
- Programs written in Go 1.0 still compile and run today
- Stable language evolution without breaking changes
-->

---

## Real-World Use Cases

- Docker: Containerization platform
- Kubernetes: Orchestration system
- Cloud services: AWS CLI, Terraform

<!--
Cloud Infrastructure: Go is heavily used in cloud-native applications, offering tools and APIs for major cloud providers like AWS, Azure, and Google Cloud.

Web Development: Go's performance and concurrency capabilities make it suitable for building scalable web servers and APIs, often used as a backend for web applications.

Microservices: Go is well-suited for creating small, focused microservices due to its lightweight nature and efficient concurrency model.

Network Services: Go's built-in concurrency primitives make it ideal for building network services, including those for distributed systems and high-performance applications.

Command-Line Interfaces (CLIs): Go's ability to create fast and efficient executables makes it a good choice for building powerful and user-friendly command-line tools.

DevOps and SRE: Go is used to develop tools for infrastructure management, automation, and monitoring, assisting DevOps and Site Reliability Engineers.

Data Processing and Analysis: Go can be used for data conversion, normalization, and ingestion, especially when dealing with large datasets.

High-Performance Systems: Go's performance and concurrency features are beneficial in building high-performance systems like reputation systems or background simulations.

Kubernetes: Kubernetes, a system for container orchestration, is written in Go, demonstrating its power in managing complex systems.

Specific Company Examples:
Uber: Uses Go for its ridesharing platform, particularly for handling geofence lookups.
Netflix: Leverages Go for application data caching.
Monzo: Uses Go for building microservices within its banking application.
Twitch: Utilizes Go for various backend services.

In essence, Go is a versatile language that excels in building scalable, performant, and reliable systems, making it a strong choice for a wide range of production environments.
-->

---

## Why Choose Go?

- Fast compilation and execution
- Easy to learn for beginners
- Scales for large systems

<!--
- Fast compilation means quick development cycles and efficient CI/CD pipelines
- Simple, clean syntax reduces cognitive load and maintenance costs
- Built-in testing framework promotes reliable code
- Cross-platform compilation is straightforward
- Strong standard library reduces dependency management
- Corporate backing from Google ensures long-term support
- Growing job market with competitive salaries
- Active community provides excellent resources and tools
- Memory safety features prevent common bugs
- Excellent tooling ecosystem (go fmt, go vet, go mod)
-->

---

## Q&A

- Questions? Let’s discuss!

---

## About Me

- Name: Peter Preeper
- Contact: ppreeper@gmail.com
- I Work For: Thinksoft Inc.
- Job Title: Senior Implementation Specialist
