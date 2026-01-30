# 🧩 Assignment 2 --- Payment Simulation (Golang Interfaces)

## 🎯 Objective

Understand how interfaces help in handling multiple implementations of
the same behavior, commonly used in real-world systems like payments.

------------------------------------------------------------------------

## 📘 Problem Statement

Design a simple payment system where different payment methods can be
used interchangeably through a common interface.

------------------------------------------------------------------------

## 📌 Requirements

1.  Create an interface named `Payment` with the following method:

``` go
Pay(amount int)
```

2.  Create two structs:

-   `UPI`
-   `Card`

3.  Implement the `Pay()` method:

-   `UPI` should print:

```{=html}
<!-- -->
```
    Paid ₹<amount> using UPI

-   `Card` should print:

```{=html}
<!-- -->
```
    Paid ₹<amount> using Card

4.  Create a function:

``` go
func Checkout(p Payment)
```

This function should call the `Pay()` method.

5.  In the `main()` function:

-   Create instances of `UPI` and `Card`
-   Call `Checkout()` for both payment methods with any amount

------------------------------------------------------------------------

## 🚫 Constraints

-   Do not use `if`, `switch`, or type assertions
-   Do not hardcode payment logic inside `Checkout`
-   Use interfaces to abstract behavior

------------------------------------------------------------------------

## ✅ Expected Output

    Paid ₹500 using UPI
    Paid ₹500 using Card

------------------------------------------------------------------------

## 🧠 Key Learnings

-   Interfaces allow plugging in different implementations
-   Business logic does not depend on concrete payment types
-   Adding a new payment method requires no changes to existing logic
