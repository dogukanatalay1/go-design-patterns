# Singleton Pattern

## What's the Point?

Ever found yourself in a situation where you need one instance of something, and only one? Maybe a database connection, a logger, or a configuration manager. You don't want multiple instances floating around because that would be wasteful, confusing, or just plain wrong.

That's where the Singleton pattern comes in. It guarantees you'll always get the same instance, no matter how many times you ask for it.

## Real-World Scenarios

Think about these situations:

- **Database connections**: You probably don't want to open 10 connections to your database when one will do. A singleton ensures you reuse the same connection.
- **Logging**: Your app should have one logger that everything writes to, not a bunch of different loggers creating chaos.
- **Configuration**: Settings should be loaded once and shared everywhere, not re-read from a file every time.
- **Caches**: One cache manager that everyone uses, keeping things consistent.

Basically, if you need "one of these, and only one," singleton is your friend.

## The Problem Without Singleton

Let me paint you a picture. Imagine you're building an app and you need a database connection. Without singleton, you might do this:

```go
// In one file
db := NewDatabaseConnection()

// In another file  
db2 := NewDatabaseConnection()

// Oops, now you have two connections!
```

This wastes resources, and worse, if these connections have state, they can get out of sync. Plus, if multiple goroutines try to create connections at the same time, you might hit race conditions.

Singleton solves this by saying "nope, you get the same one every time."

## How to Do It in Go

Go doesn't have classes like Java or C++, but we can still do singleton. Here's the trick: use a struct, keep the instance private, and use `sync.Once` to make sure it only gets created once (even with goroutines running around).

Here's the basic setup:

```go
type Singleton struct {
    // Put whatever fields you need here
}

var (
    instance *Singleton
    once     sync.Once
)

func GetInstance() *Singleton {
    once.Do(func() {
        instance = &Singleton{
            // Initialize your fields
        }
    })
    return instance
}
```

The magic is `sync.Once`. It's a Go thing that guarantees the function you pass to `Do()` runs exactly once, no matter how many goroutines are calling `GetInstance()` at the same time. Pretty neat, right?

## Breaking It Down

1. **The struct**: This is your singleton object. Put whatever data and methods it needs.

2. **The private instance**: Notice `instance` starts with a lowercase letter? That means it's unexported (private). Only code in this package can see it.

3. **sync.Once**: This is Go's way of saying "do this thing once, and only once." Even if 100 goroutines call `GetInstance()` at the exact same moment, only one will actually create the instance.

4. **GetInstance()**: This is your public function. Call it whenever you need the singleton. First call creates it, every call after that just returns the same one.

## When Should You Use It?

**Good times to use singleton:**
- You genuinely need exactly one instance (like a database connection)
- You want lazy initialization (don't create it until someone actually needs it)
- You need it accessible from many places in your code
- The thing is expensive to create (so you don't want to make it multiple times)

**Maybe think twice if:**
- You might need multiple instances later (singleton makes that hard)
- You're big on testability (singletons can make testing trickier because of global state)
- You're using dependency injection (singletons can fight against DI)
- The thing is stateless (maybe just use a function instead?)

## The Good and The Bad

**What's good about it:**
- You control exactly how many instances exist (one!)
- Lazy initialization means you don't waste resources creating things you might not use
- `sync.Once` makes it thread-safe without you having to think about locks
- Easy to access from anywhere (just call `GetInstance()`)

**What's not so great:**
- Global state can make testing harder (harder to mock, harder to reset)
- It kind of breaks the Single Responsibility Principle (the object manages its own lifecycle)
- Dependencies can be hidden (if your code uses a singleton, it's not obvious from the function signature)
- Can lead to tight coupling (everyone depends on that one instance)

## Bottom Line

Singleton is a simple pattern that solves a specific problem: "I need one of these, and only one." In Go, `sync.Once` makes it easy and thread-safe. Just be careful not to overuse it - sometimes you think you need a singleton, but you actually just need better design.

Check out `example.go` to see a working example, and `main.go` to see it in action!
