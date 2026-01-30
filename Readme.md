# Using Interfaces in Golang --- Practical Guide

This document explains **how to use interfaces in Golang**, why they
exist, and how they are applied in real-world Go programs.\
It is written for developers who already know basic Go syntax and want
to **truly understand interfaces beyond toy examples**.

------------------------------------------------------------------------

## 1. What is an Interface in Go?

In Golang, an interface defines **behavior**, not data.

``` go
type Speaker interface {
    Speak()
}
```

Any type that implements the `Speak()` method **automatically
satisfies** this interface.

There is: - ❌ No `implements` keyword - ❌ No inheritance - ✅ Implicit
implementation based on method signatures

------------------------------------------------------------------------

## 2. Why Interfaces Exist (Core Idea)

Interfaces are used to: - Decouple code from concrete implementations -
Allow multiple implementations of the same behavior - Make code
extensible and easier to maintain

> Interfaces answer the question:\
> **"What can this thing do?"**, not **"What is this thing?"**

------------------------------------------------------------------------

## 3. Basic Interface Example

### Interface

``` go
type Player interface {
    Play()
}
```

### Implementations

``` go
type Music struct{}

func (m Music) Play() {
    fmt.Println("Playing music")
}

type Video struct{}

func (v Video) Play() {
    fmt.Println("Playing video")
}
```

### Function Using Interface

``` go
func Start(p Player) {
    p.Play()
}
```

### Usage

``` go
func main() {
    m := Music{}
    v := Video{}

    Start(m)
    Start(v)
}
```

### Key Takeaway

The `Start` function does not know or care whether it receives `Music`
or `Video`.\
It only cares that the type can `Play()`.

------------------------------------------------------------------------

## 4. Interfaces at Function Boundaries (Important Pattern)

Interfaces are most useful when used as **function parameters**.

``` go
func Process(p Player) {
    p.Play()
}
```

This allows: - Replacing implementations easily - Adding new types
without changing existing code

------------------------------------------------------------------------

## 5. Multiple Implementations, One Interface

Example: Payment system

### Interface

``` go
type Payment interface {
    Pay(amount int)
}
```

### Implementations

``` go
type UPI struct{}

func (u UPI) Pay(amount int) {
    fmt.Println("Paid using UPI:", amount)
}

type Card struct{}

func (c Card) Pay(amount int) {
    fmt.Println("Paid using Card:", amount)
}
```

### Usage

``` go
func Checkout(p Payment) {
    p.Pay(500)
}
```

------------------------------------------------------------------------

## 6. Collections of Interfaces

Interfaces allow storing different concrete types in the same slice.

``` go
players := []Player{
    Music{},
    Video{},
}

for _, p := range players {
    p.Play()
}
```

This is not possible without interfaces.

------------------------------------------------------------------------

## 7. Interface Design Best Practices

### Keep Interfaces Small

❌ Bad:

``` go
type UserService interface {
    CreateUser()
    DeleteUser()
    SendEmail()
    LogActivity()
}
```

✅ Good:

``` go
type UserSaver interface {
    Save(user User) error
}
```

------------------------------------------------------------------------

## 8. Accept Interfaces, Return Concrete Types

✅ Good:

``` go
func Process(p Player)
```

❌ Avoid (unless required):

``` go
func GetPlayer() Player
```

------------------------------------------------------------------------

## 9. When NOT to Use Interfaces

Do NOT use interfaces when: - There is only one implementation -
Behavior is not expected to change - Code is small and local (e.g.,
inside `main`)

Interfaces are tools, not mandatory abstractions.

------------------------------------------------------------------------

## 10. Mental Model to Remember

-   Interface = ability
-   Struct = concrete thing
-   If a struct has required methods → it satisfies the interface

> "In Go, you don't design hierarchies.\
> You design behaviors."

------------------------------------------------------------------------

## 11. What to Learn Next

After understanding basic interfaces, move on to: - Pointer vs value
receivers with interfaces - Interface method sets - Using interfaces for
testing and mocking - Real-world patterns (repository, logger, payment
gateway)

------------------------------------------------------------------------

## Summary

Interfaces in Go: - Enable polymorphism without inheritance - Reduce
coupling - Make code extensible and maintainable - Are central to Go's
standard library design
