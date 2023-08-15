# Go

1. [Introduction](#intro)
2. [what is Go?](#what)
3. [Installation](#install)
4. [Visual Studio Code Extensions for Go](#vs)
5. [Create first Application](#firstapp)
6. [Variable](#variable)
7. [Constants](#constants)
8. [In built Packages](#inbuilt_pkg)
9. [Print()](#print)
10. [Println()](#println)
11. [Formatted Data Printing](#printf)
12. [Input Scan()](#scan)
13. [Data Type](#dataType)
14. [Pointer](#pointer)
15. [Array](#array)
16. [Slice](#slice)
17. [Loops](#loops)
18. [break](#break)
19. [continue](#continue)
20. [Package strings](#strings)
21. [If..Else](#ifelse)
22. [Else...If](#elseif)
23. [Operators](#operators)
24. [Switch Statements](#switch)
25. [Function in Go](#functions)
26. [Packages Level Variables](#pkg_level_variable)
27. [Managing and Running Multiple Packages](#multiplePackages)
28. [Importing packages](#import)
29. [Export packages](#export)
30. [Level of Scopes](#scopes)

## Introduction

- Open source Programming Language
- Developed By Google
- Designed with Focus on `Simplicity`, `efficiency` and `Concurrency`.
- Provides Robust toolset to build reliable and efficient software

### Features of Go

1. `Simplicity` - Emphasizes simplicity and readability
2. `Concurrency` - Includes built in support for cocurrent programming
3. `Efficiency` - Compiles to machine code, resulting efficient and performant executables
4. `Garbage Collection` - manages memory allocation and deallocation
5. `Static Typing` - Variable types are checked at compile time
6. `Standard Liabrary` - Provides essential packages for tasks like handling I/O, networking, data manipulation and more
7. `Open Source` - fostering community collaboration and contributions
8. `Cross-Platform` - Support various platforms, making suitable for developing application that can run on different Operating Systems
9. `Compiled` - Compiles to native machine code
10. `Modern Language Support` - Incorporates modern programming language features

## what is Go?

Open Source Programming Language developed by Google. Designed with focus on simplicity, efficiency and concurrency

## Installation

1. Dowload go binary

```url
https://go.dev/dl/
```

2. Extract Binary

```bash
tar -xzf binarytar.tar.gz
```

3. Set Environment Variable

```bash
export PATH=$PATH:/usr/local/go/bin
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin
```

4. Verify Installation

```bash
go version
```

## Visual Studio Code Extensions for Go

Go extension maintained by Google Team

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

```go
package main

import (
    "fmt"
)

func main(){
    var userTicket int
    fmt.Println("Please enter number of tickets you want to book")
    fmt.Scan(&userTicket)
    fmt.Println("You have Successfully booked %v no of tickets",userTicket)
}
```

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

```go
package main

import (
    "fmt"
)

func main(){
    fmt.Println("Welcome to Ticket Booking application")
}
```

## Slice

```go
package main

import (
    "fmt"
)

func main(){
    fmt.Println("Welcome to Ticket Booking application")
}
```

## Loops

Go has only one looping statement that is `for`.
Looping statements are concerned to run set of instruction repeteadly.

`main.go`

```go
package main

import (
    "fmt"
)

func main(){
    counter:=1
    fmt.Println("Welcome to Ticket Booking application")
    for {
        fmt.Println("Counter Value %v",counter)
        counter = counter + 1
        if(counter == 5)
        {
            break;
        }
    }
}
```

In above example for implements `infinite loop`

`main.go`

```go
package main

import (
    "fmt"
)

func main(){
    counter:=1
    fmt.Println("Welcome to Ticket Booking application")
    for counter > 6{
        fmt.Println("Counter Value %v",counter)
        counter = counter + 1
    }
}
```

In above example for have some condition, after that for condition goes to false, for exexution stops.

## break

It is used to terminate the execution of a loop prematurely.
`break` in loop is used to exit the loop.
`break` is applicable in for loop and switch statements only.

`main.go`

```go
package main

import (
    "fmt"
)

func main(){
    counter:=1
    fmt.Println("Welcome to Ticket Booking application")
    for {
        fmt.Println("Counter Value %v",counter)
        counter = counter + 1
        if(counter == 5)
        {
            break;
        }
    }
}
```

Here, we have defined for loop without condition which determines, that loop is infinite loop.

```go
for {
// code to be executed repeteadly
}

for true{
// code to be executed repeteadly
}
// both of the above statement mean same thing
```

## continue

`continue` used to skip the current iteration of the loop and move to next iteration.
Commonly used to skip the certain steps of the loop execution.

`main.go`

```go
package main

import "fmt"

func main() {
    for i := 1; i <= 5; i++ {
        if i == 3 {
            continue // Skip the rest of the loop for i equals 3
        }
        fmt.Println(i)
    }
}
```

Above example will skip the priniting i value when value of will be 3

## Package "strings"

`strings` package provides collection of functions for manipulating and working with strings.
`strings` package is part of standard liabrary

#### Following is the list of commonly used packages

- len(str) : String Length - returns length of the string. `lengthOfString := strings.len(userName)`

```go
username := "Manisha"
lengthOfString := string.len(userName)
```

- String Checking : strings.Contains() - checks if strings conatins a specified substring or nor

```go
email := "manishajadhav2323@gmail.com"
isValidEmail := strings.Contains(email, "@")
fmt.Print(isValidEmail)
```

```go
package main

import (
"fmt"
)

func main(){
    fmt.Println("Welcome to Ticket Booking application")
    email := "manishajadhav2323@gmail.com"
    isValidEmail := strings.Contains(email, "@")
    if !isValidEmail{
        fmt.Println("Incorrect Email ID")
    }else
    {
        fmt.Println("Ticket has been booked Successfully")
    }
}

```

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

Supports variety of operators for performing different types of operations

| Name                        | Symbols                           |
| --------------------------- | --------------------------------- |
| Arithmetic Operator         | `+, -, *, /, %`                   |
| Assignement Operator        | `=,  +=, -=, *=, /=, %=`          |
| Comparison Operator         | `==, !=, <, >, <=, >=`            |
| Logical Operator            | `&&, logical or , !`              |
| Unary Operators             | `+, -, ++, --`                    |
| Bitwise Operator            | `&,   or opertor, ^, <<, >>`      |
| Bitwise Assignment Operator | `&=,  or operator=, ^=, <<=, >>=` |
| Other Operator              | `&, *, ., ->, [], ()`             |

## Switch Statements

Used for making decisions based on different set possible values

```go
package main

import "fmt"

func main() {
    day := "Wednesday"

    switch day {
    case "Monday":
        fmt.Println("It's Monday!")
    case "Tuesday":
        fmt.Println("It's Tuesday!")
    case "Wednesday":
        fmt.Println("It's Wednesday!")
    default:
        fmt.Println("It's another day.")
    }
}

```

In above example it checks day match case. As case matches with `Wednesday`, it will print `It's Wednesday`

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
