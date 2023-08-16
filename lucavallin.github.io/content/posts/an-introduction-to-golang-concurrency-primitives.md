---
showToc: true
TocOpen: false
hidemeta: false
comments: true
ShowCanonicalLink: false
disableShare: false
searchHidden: true
hideSummary: false
ShowReadingTime: true
ShowBreadCrumbs: true
ShowPostNavLinks: true
# ShowWordCount: true
ShowRssButtonInSectionTermList: true
UseHugoToc: true
editPost:
    URL: "https://github.com/lucavallin/lucavallin.github.io/tree/main/content"
    Text: "Suggest Changes"
    appendFilePath: true
# CanonicalLinkText
# disableAnchoredHeadings
# hideFooter
# ShowCodeCopyButtons
# ShareButtons
# robotsNoIndex

author: "Luca Cavallin"
date: "2022-12-06"
title: "An introduction to Golang concurrency primitives"
description: "Concurrency is a powerful feature of the Go programming language that allows developers to write efficient and scalable programs. Go makes it easy to write concurrent programs by providing several high-level abstractions, such as goroutines, channels, the select statement, mutexes, and wait groups."
canonicalURL: ""
tags: ["golang", "concurrency"]
draft: false
cover:
    image: "/luca.png"
    alt: "An introduction to Golang concurrency primitives"
    caption: "An introduction to Golang concurrency primitives" # display caption under cover
    relative: false # when using page bundles set this to true
    hidden: false
# weight: 1
# aliases: ["/first"]
---
Goroutines are lightweight threads of execution that can be easily created and managed by the Go runtime. Goroutines are similar to threads in other languages, but they are much cheaper to create and manage, which makes them ideal for concurrent programming. Here is an example of how to create a goroutine in Go:

## Goroutines

```go
package main

import "fmt"

func main() {
    go fmt.Println("Hello from goroutine 1")
    go fmt.Println("Hello from goroutine 2")
}
```

This code creates a new goroutine that prints a message to the console. The `go` keyword is used to indicate that the function should be run in a new goroutine. Because goroutines are concurrent, the main function will continue to execute while the goroutine is running.

## Channels

Channels are another important concurrency primitive in Go. Channels provide a way for goroutines to communicate and synchronize with each other. By using channels, goroutines can send and receive values from each other, allowing them to coordinate and share data. Here is an example of how to use channels in Go:

```go
package main

import "fmt"

func main() {
    // Create a channel that can be used to send and receive strings.
    c := make(chan string)

    // Create two goroutines that send messages to the channel.
    go func() { c <- "Hello from goroutine 1" }()
    go func() { c <- "Hello from goroutine 2" }()

    // Read the messages from the channel and print them to the console.
    fmt.Println(<-c)
    fmt.Println(<-c)
}
```

This code creates a new channel and a goroutine that sends a value on the channel. The main function receives the value from the channel and prints it to the console. Channels are a powerful tool for coordinating and synchronizing goroutines.

## Select

The `select` statement allows a goroutine to wait for multiple channels to become ready, then execute a block of code depending on which channel is ready. Here's an example of how to use `select` to implement a simple concurrent timer:

```go
package main

import "fmt"
import "time"

func main() {
    // Create two channels: one for sending a signal after 1 second,
    // and one for sending a signal after 2 seconds.
    c1 := time.After(1 * time.Second)
    c2 := time.After(2 * time.Second)

    // Use a select statement to wait for one of the channels to become ready.
    select {
    case <-c1:
        fmt.Println("1 second elapsed")
    case <-c2:
        fmt.Println("2 seconds elapsed")
    }
}
```

## Mutex

To prevent concurrent access to shared data from causing race conditions and other bugs, Go provides a `mutex` (short for mutual exclusion) type that can be used to protect critical sections of code. Here's an example of how to use a mutex to safely increment a shared counter variable:

```go
package main

import "fmt"
import "sync"

func main() {
    // Create a mutex to protect the counter variable.
    var mu sync.Mutex

    // Create a counter variable and initialize it to 0.
    var counter int

    // Create 10 goroutines that concurrently increment the counter.
    for i := 0; i < 10; i++ {
        go func() {
            mu.Lock()
            defer mu.Unlock()

            counter++
        }()
    }

    // Wait for all the goroutines to finish, then print the final value of the counter.
    time.Sleep(1 * time.Second)
    fmt.Println(counter)
}
```

## Wait Groups
The `sync` package in Go provides additional tools for working with concurrency, including `WaitGroups` for coordinating the termination of multiple goroutines, `Once` for ensuring that a piece of code is only executed once, and `Pool` for managing and reusing a pool of resources.

Here's an example of how to use the `WaitGroup` type from the `sync` package to wait for a group of goroutines to finish:

```go
package main

import "fmt"
import "sync"

func main() {
    // Create a WaitGroup to wait for the goroutines to finish.
    var wg sync.WaitGroup

    // Launch 10 goroutines that concurrently print messages.
    wg.Add(10)
    for i := 0; i < 10; i++ {
        go func() {
            fmt.Println("Hello from goroutine")
            wg.Done()
        }()
    }

    // Wait for the goroutines to finish.
    wg.Wait()
}
```

## Summary

In conclusion, concurrency is a powerful and important concept in Go, and the language provides a rich set of tools for working with concurrency, including goroutines, channels, select statements, the `mutex` and `sync` packages. By using these tools effectively, Go programmers can write concurrent programs that are efficient, scalable, and easy to understand.
