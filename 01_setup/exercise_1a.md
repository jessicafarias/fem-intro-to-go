# Exercise 1A: Find Stuff

## Goals

- Familiarize yourself with the golang.org website and resources

## Setup

- Visit `golang.org`

## Directions

Answer the following questions

1. Read about `for loops` in the _Effective Go_ document

There is a lot of similarities on for loops 

```go
// Like a C for
for init; condition; post { }
// Like a C while
for condition { }
// Like a C for(;;)
for { }
```

 Short declarations make it easy to declare the index variable right in the loop. 

```go
sum := 0
for i := 0; i < 10; i++ {
    sum += i
}
```

If you're looping over an array, slice, string, or map, or reading from a channel, a range clause can manage the loop. 

```go
for key, value := range oldMap {
    newMap[key] = value
}
```

 If you only need the first item in the range (the key or index), drop the second: 

```go
for key := range m {
    if key.expired() {
        delete(m, key)
    }
}
```

If you only need the second item in the range (the value), use the blank identifier, an underscore, to discard the first: 

```go
sum := 0
for _, value := range array {
    sum += value
}
```

- What kind of loop doesn’t exist in Go?
answer: So the do while loop dosn't exist.


2. Read about the `fmt` _package_

Formatted printing in Go uses a style similar to C's printf family but is richer and more general. The functions live in the fmt package and have capitalized names: fmt.Printf, fmt.Fprintf, fmt.Sprintf and so on. The string functions (Sprintf etc.) return a string rather than filling in a provided buffer.

You don't need to provide a format string. For each of Printf, Fprintf and Sprintf there is another pair of functions, for instance Print and Println. These functions do not take a format string but instead generate a default format for each argument. The Println versions also insert a blank between arguments and append a newline to the output while the Print versions add blanks only if the operand on neither side is a string. In this example each line produces the same output. 

```go
fmt.Printf("Hello %d\n", 23)
fmt.Fprint(os.Stdout, "Hello ", 23, "\n")
fmt.Println("Hello", 23)
fmt.Println(fmt.Sprint("Hello ", 23))
```

- What does `fmt.Println()` return?

answer: It returns an integer and an error, the integer will be the number of bytes writen and any error encounter

3. Find a _blog post_ about the recent release of Go 1.16

- What are some of the new features?

The go command now builds packages in module-aware mode by default, even when no go.mod is present. This is a big step toward using modules in all projects.

It's still possible to build packages in GOPATH mode by setting the GO111MODULE environment variable to off. You can also set GO111MODULE to auto to enable module-aware mode only when a go.mod file is present in the current directory or any parent directory. This was previously the default. Note that you can set GO111MODULE and other variables permanently with go env -w:

go env -w GO111MODULE=auto

We plan to drop support for GOPATH mode in Go 1.17. In other words, Go 1.17 will ignore GO111MODULE. If you have projects that do not build in module-aware mode, now is the time to migrate. If there is a problem preventing you from migrating, please consider filing an issue or an experience report.
