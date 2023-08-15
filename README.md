# Go

[Introduction](#intro)
[what is Go?](#what)
[Installation](#install)
[Visual Studio Code Extensions for Go](#vs)
[Create first Application](#firstapp)
[Variable](#variable)
[Constants](#constants)
[In built Packages](#inbuilt_pkg)
[Print()](#print)
[Println()](#println)
[Formatted Data Printing](#printf)
[Input Scan()](#scan)
[Data Type](#dataType)
[Pointer](#pointer)
[Array](#array)
[Slice](#slice)
[Loops](#loops)
[break](#break)
[continue](#continue)
[Package strings](#strings)
[If..Else](#ifelse)
[Else...If](#elseif)
[Operators](#operators)
[Switch Statements](#switch)
[Function in Go](#functions)
[Packages Level Variables](#pkg_level_variable)
[Managing and Running Multiple Packages](#multiplePackages)
[Importing packages](#import)
[Export packages](#export)
[Level of Scopes](#scopes)

## Introduction

## What is Go?

## Installation

## Visual Studio Code Extensions for Go

## Create first Application

create module command

```go
go mod init booking-app
```

```go
package main
import (
    "fmt"
)

func main()
{
    fmt.Print("Welcome to ticket booking app")
    fmt.Println("We love to assist you. Please press Enter to continue")
}
```

In above example
`package main` Defines name of the package. As everything in Go organized in package.
`import ("fmt")` Defines you are trying to import in built package of Go i.e. `fmt`
`func main` Defines entry point of execution. Each package must contain `main()`
`fmt.Print()` Defines you want to use `fmt` and its `Print()` to print output on screen
`fmt.Println()` Defines you want to use `fmt` and its `Print()` to print output on screen

## Variable

Variables are declared with var keyword or := operator

```go
 var remainingTicket = 100

```

Variables with local scope can also be declared with := operator

```go
    remainingTicket := 100
```

When type is not defined. It implicity consider type as per assigned value.
Variable declaration with type

```go
var remainingTicket int = 100 // integer
var userName string = "Manisha" // string
var bookings [50] string{"Pune","Mumbai"} // array
```

## Constants

Const are declared with `const` keyword.
Constanst are used to store constant value.

```go
const totalTickets = 100
```

## In built Packages

## Print()

Print is used to print output on screen without next line.

```go
package main

import (
    "fmt"
)

func main()
{
    fmt.Print("Welcome to Ticket Booking app")
}
```

In above example
`package main` Defines name of the package. As everything in Go organized in package.
`import ("fmt")` Defines you are trying to import in built package of Go i.e. `fmt`
`func main` Defines entry point of execution. Each package must contain `main()`
`fmt.Print()` Defines you want to use `fmt` and its `Print()` to print output on screen

## Println()

`Println()` is used to print output on screen on new line.

```go
package main

import (
    "fmt"
)

func main()
{
    fmt.Println("Welcome to Ticket Booking app")
}
```

In above example
`func main` Defines entry point of execution. Each package must contain `main()`
`fmt.Print()` Defines you want to use `fmt` and its `Print()` to print output on screen with new line

## Printf()

`Printf()` is used to print formatted output on screen without new line.
For new line we need to use escape sequence `\n`

```go
package main

import (
    "fmt"
)

func main()
{
    var remainingTicket int = 100
    fmt.Printf("Welcome to Ticket Booking app. We have %v remaining ticket for the show",remainingTicket)
}
```

In above example
`func main` Defines entry point of execution. Each package must contain `main()`
`fmt.Printf()` Defines you want to use `fmt` and its `Printf()` to print formatted output on screen. `%v` holds place for variable.

## Input Scan()

## Data Type

Following is the list of data types in Go-lang

|     | Type               |
| --- | ------------------ |
| 1   | Interger `int`     |
| 2   | String `string`    |
| 3   | Array `[]dataType` |
| 4   | map                |
| 5   | Booleans `bool`    |

## Pointer

Pointer points to the address of another variable.

## Array

## Slice

## Loops

## break

## continue

## Package "strings"

## If...Else

`If...Else` is used to define control flow in Go.

```go
age := 20

if age > 18{
    fmt.Println("Welcome to Show !!! You have Successfully Booked ticket for the show.")
}else {
    fmt.Println("This show is probhited for audience less than age of 18")
}
```

## Else...If

`Else...If` is used to define control flow in Go.

```go
age := 20

if age > 18{
    fmt.Println("Welcome to Show !!! You have Successfully Booked ticket for the show.")
}else if age > 15 && age < 18{
      fmt.Println("Welcome to Show !!! You have Successfully Booked ticket for the show in Pune!!!")
}
else {
    fmt.Println("This show is probhited for audience less than age of 18")
}
```

## Operators

## Switch Statements

## Function in Go

Funtions are smallest unit. functions are delcared with keyword `func` in Go-lang.

`main.go`

```go
package main

import ("fmt")

func main(){
    userName := "Manisha"
    greeting(userName)
}

func greeting(userName string){
    fmt.Print("Welcome to show %v",userName)
}

```

Here `main()` function is the entry point of the go. Inside main we have called function `greeting()`.
`greeting(userName string)` userName is argument name that has been passed for the function. string defines type of argument.

## Packages Level Variable

## Managing and Running Multiple Packages

## Import packages

Package in Go import with keywork `import` with enclosing in round parenthesis `()`

```go
import (
    "fmt"
    "booking-app/greeting"
)
```

Here, `"fmt"` is in-built package of Go, where `"booking-app/greeting"` is custom package.
`"booking-app/greeting"`
`booking-app` : Denotes application or module name from `go.mod` file.
`greeting` : Denotes package name used while writing package.
Each new package is mentioned at new line

## Export packages

Exporting is done using captilizing first letter of enity you want to export

`greeting.go`

```go
package greeting
import (
    "fmt"
)
func WelcomeGreet(userName)
{
    fmt.Print("Welcome to Pune!! %v",userName)
}
```

`main.go`

```go
package main

import (
    "fmt"
    "booking-app/greeting"
)

func main()
{
    var userName = "Mayur"
    WelcomeGreet(userName)
}

```

In above example, there are two packages `greeting` and `main`. We have exported function `WelcomGreet()` from `greeting` package, by capitalizing first letter of function. `WelcomeGreet()` : `W`.

## Level of Scopes

Following is the list of scopes in Go-lang

|     | Type    |
| --- | ------- |
| 1   | Local   |
| 2   | Global  |
| 3   | Package |

`Local` - Scope local to function
`Global` - Global to all function
`Package` - Package level Scope
