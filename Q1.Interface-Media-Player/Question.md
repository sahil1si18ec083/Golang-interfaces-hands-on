# Golang Interfaces --- Hands-On Assignments

## 🧩 Assignment 1 --- Media Player (Warm-up)

### 🎯 Objective

Learn how Golang interfaces enable polymorphic behavior without
inheritance.

------------------------------------------------------------------------

### 📘 Problem Statement

Build a simple **Media Player system** where different media types can
be played using a common interface.

------------------------------------------------------------------------

### 📌 Requirements

1.  Create an interface named `Player` with the following method:

``` go
Play()
```

2.  Create two structs:

-   `Music`
-   `Video`

3.  Implement the `Play()` method:

-   `Music` should print:

```{=html}
<!-- -->
```
    Playing music 🎵

-   `Video` should print:

```{=html}
<!-- -->
```
    Playing video 🎬

4.  Create a function:

``` go
func Start(p Player)
```

This function should call the `Play()` method.

5.  In the `main()` function:

-   Create instances of `Music` and `Video`
-   Pass them to the `Start()` function

------------------------------------------------------------------------

### 🚫 Constraints

-   Do not use `if`, `switch`, or type assertions
-   Do not use inheritance
-   Use interfaces only for behavior abstraction

------------------------------------------------------------------------

### ✅ Expected Output

    Playing music 🎵
    Playing video 🎬

------------------------------------------------------------------------

### 🧠 Key Learnings

-   Interfaces define behavior, not data
-   Any type that implements interface methods satisfies the interface
    automatically
-   Interfaces help decouple logic from concrete implementations
