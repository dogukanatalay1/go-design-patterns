# Builder Pattern

## What's the Point?

Ever tried to create an object that has like 10 different fields, half of which are optional? You end up with constructors that look like this:

```go
user := NewUser("John", "Doe", "john@example.com", "", "", "", true, false, nil, time.Now())
```

What the heck do all those parameters mean? You have no idea without looking at the function signature. And what if you only want to set 3 of them? You're stuck passing empty strings or nil values everywhere.

That's where the Builder pattern comes in. It lets you construct complex objects step by step, in a readable way. Instead of that mess above, you get:

```go
user := NewUserBuilder().
    FirstName("John").
    LastName("Doe").
    Email("john@example.com").
    IsActive(true).
    Build()
```

Much better, right? You can see exactly what each value is, and you only set what you need.

## Real-World Scenarios

Think about these situations:

- **Database queries**: Building SQL queries with optional WHERE clauses, JOINs, ORDER BY, etc.
- **HTTP requests**: Setting headers, body, timeout, retries - lots of optional stuff
- **Configuration objects**: Server configs with 20+ optional settings
- **UI components**: Building complex widgets with many styling and behavior options
- **API clients**: Creating request objects with many optional parameters

Basically, any time you have an object with many fields (especially optional ones), Builder makes your life easier.

## The Problem Without Builder

Let's say you're building a `ServerConfig` struct. Without builder, you might do:

```go
// Option 1: One massive constructor
config := NewServerConfig("localhost", 8080, true, false, 30, "", "", nil, ...)
// What does each parameter mean? Who knows!

// Option 2: Set fields after creation
config := &ServerConfig{}
config.Host = "localhost"
config.Port = 8080
config.SSL = true
// ... 15 more lines
```

Problems:
- Massive constructors are unreadable and error-prone
- Setting fields after creation means the object might be in an invalid state
- No way to enforce required vs optional fields
- Hard to make the construction process immutable

## How Builder Solves It

The Builder pattern gives you a fluent interface to construct objects step by step. You create a builder, call methods to set what you need, then call `Build()` to get your final object.

```go
config := NewServerConfigBuilder().
    Host("localhost").
    Port(8080).
    EnableSSL(true).
    Timeout(30 * time.Second).
    Build()
```

Benefits:
- **Readable**: You can see what each value represents
- **Flexible**: Set only what you need
- **Safe**: The builder can validate before creating the object
- **Fluent**: Method chaining makes it feel natural

## When Should You Use It?

**Good times to use Builder:**
- Objects with many fields (5+)
- Many optional parameters
- You want readable object construction
- You need to validate before creating the object
- You want to make construction immutable

**Maybe think twice if:**
- The object is simple (just 2-3 required fields)
- All fields are required (a normal constructor is fine)
- You don't need validation or complex construction logic

## The Basic Idea

1. **The Product**: The complex object you want to build

2. **The Builder**: A struct that holds temporary values and provides setter methods

3. **Fluent Methods**: Each setter returns `*Builder` so you can chain calls

4. **Build Method**: Validates and creates the final object

In Go, builders are often just structs with methods that return `*Builder` for chaining. Simple and effective!

## Bottom Line

Builder pattern is about making complex object construction readable and flexible. Instead of massive constructors or scattered field assignments, you get a clean, step-by-step process that's easy to read and understand. It's especially great when you have lots of optional parameters.

Let's build one together step by step!
