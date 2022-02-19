# Go language 101
Learning to program with the Go language. Most projects are 
from the Udemy course [***Go: The Complete Developer's Guide***](https://www.udemy.com/share/101Xzy3@sLnIUzctBBQGBPVAGtwXSNuUSUeJvsDT7YT1srtaIuiuyQttgpAoZWFCMkoKc92I/) by ***Stephen Grider***.  
## Hello World
The tried and true starter program. Main takeaways from this classic:
- The overall structure of a Go file with ***package***, ***import*** and function locations
- Declaring **package main** signifies to the Go compiler that this will be an executable file
  instead of a reusable library
- If **package main** is declared then a **main()** function is required and will be the function
  that is executed
- The **import** statement allows other reusable library packages to be used within the file
## Cards
Go program that will be built during the first part of the Udemy course. Main takeaways:
- Creating custom types
- Import and use of standard packages from golang.org
- What "slices" are and how to interact with them (Similar to arrays)
- Syntax for declaring functions:
  - Receivers (Similar cocept to extension methods in C#)
  - Input paramters
  - Multiple return objects
- Syntax for creating a for loop using the ***range*** keyword
  - Using the index as well as the object in the slice
  - Ignoring the index by using the underscore
## Even & Odd
Short assignment during the course to write a small program that lists out the numbers 0 to 10 and print beside the number if it is even or odd. Simple use of a for loop and the modulus operator.
## Structs
Go program to demonstrate the creation and use of structs. This program was also used to demonstrate how and when to use pointers. Go is considered a "pass by value" language and whenever a value is passed into a function Go makes a copy if that value to work with inside the function. I am having serious flashbacks to writing programs in C back in my college days. Main takeaways:
- Go is a "pass by value" language by default
- Declaring structs with values depends on the order of the values if the field name is not specified
- When declaring structs using the ***var*** keyword the fields will be automatically assigned with "zero values"
- The command `fmt.Printf("%+v", [struct])` will print out all the field names and their values of a struct
- Pointers are needed when dealing with value types and not reference types

|Value Types|Reference Types|
|:---:|:---:|
|int|slices|
|float|maps|
|string|channels|
|bool|pointers|
|structs|functions|

## Maps
Go program to demostrate the creation and use of maps. Very similar concept to dictionaries in C#. Main takeaways:
- Maps are a collection of key/value pairs
- All the keys must be of the same type
- All the values must be of the same type
- Key/value pairs can be easily added and removed making maps dynamic
- Key/value pairs are indexed and can be iterated over
- Maps are a reference type

## Interfaces
Go program to demostrate interfaces. Similar concepts to how interfaces are used in C#. The main difference being C# interface are declared and not implicit. Main takeaways:
- Interfaces are a special ***type*** that cannot have values assigned to them
- Interfaces are not like ***generic types*** in other languages like C#
- Interfaces are ***implicit*** (no formal link or declaration needed)

## Shapes
The second short assignment to learn how to work with interfaces. Building structs to represent squares and triangles and a ***shape*** interface that requires a ***getArea()*** function. Then a print function that accepts a shape interface argument to print the area to the screen.

## Print File
A more difficult assignment to demonstrate the use of interfaces. The assignment was more about looking through the Golang documentation to determine if there are types that implement interfaces to help read the contents of a text file and write it to the screen. The assignment also demonstrated how to pass arguments into a program by using ***os.Args*** (in this case we're passing in the name of the text file). Once the name of the file is passed into the program it was a matter of researching the ***File*** type and its ***Open()*** and ***Read()*** functions. 

## Channels & Go Routines
Go program to demonstrate the use of ***Go routines*** and ***channels***. Go routines are a way of *"scheduling"* the execution of code concurrently using one or more CPU cores. This is similar to the async/await pattern used in C#. Channels are a way to pass messages between Go routines. Main takeaways:
- Go routines are *scheduled* and are constrained by the number of CPU core available
- Channels are strongly typed (only values matching the channel type can be passed through that channel)
- Receving a message from a channel is a blocking call
- There are multiple ways to define an infinite for loop to check a channel
  ```go
  // With just the 'for' keyword
  for {
    fmt.Println(<-c)
  }
  // With the 'range' keyword
  for m =: range c {
    fmt.Println(m)
  }
  ```