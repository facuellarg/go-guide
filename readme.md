# GOLANG GUIDE 
<p align="center">
<img src="./gopher.png" width=10%></img>
</p>

This guide is an attempt to create a basic path of learning for go programming language. This guide is based in [A Tour Of Go][tour_guide_golang], an interactive guide in golang official web page.

In every explanation ,we will do some example or a link will be left that redirects to some another page to see better that topic.

In some cases the explanation in the official go tour is enough and for that reason we only will be left the link of it.
<!-- TODO Context -->
- [GOLANG GUIDE](#golang-guide)
	- [Basic](#basic)
		- [Variables](#variables)
		- [Constants](#constants)
		- [Type Conversions](#type-conversions)
		- [Type Assertions](#type-assertions)
		- [Pointers](#pointers)
		- [Type](#type)
			- [Alias](#alias)
			- [Struct](#struct)
				- [Literals](#literals)
			- [Interfaces](#interfaces)
		- [Functions](#functions)
			- [Error](#error)
		- [GOTO](#goto)
		- [Panic](#panic)
		- [Flow Control](#flow-control)
			- [Defer](#defer)
		- [Packages](#packages)
		- [Modules](#modules)
		- [Imports](#imports)
		- [Exports](#exports)
	- [Advance](#advance)
		- [JSON](#json)
		- [Concurrency](#concurrency)
			- [Go Routine](#go-routine)
			- [Channels](#channels)
				- [Buffered Channel](#buffered-channel)
				- [Range and Close](#range-and-close)
				- [Select](#select)
			- [Sync](#sync)
				- [Mutex](#mutex)
				- [Wait Group](#wait-group)
		- [Web](#web)
			- [Native](#native)
				- [Server](#server)
				- [Handler](#handler)
				- [Route](#route)
				- [Listen and Serve](#listen-and-serve)
				- [Example](#example)
			- [Echo](#echo)
				- [Server](#server-1)
				- [Handler](#handler-1)
				- [Route](#route-1)
				- [Listen and Serve](#listen-and-serve-1)
				- [Example](#example-1)
			- [Gin](#gin)
				- [Server](#server-2)
				- [Handler](#handler-2)
				- [Route](#route-2)
				- [Listen and Serve](#listen-and-serve-2)
				- [Example](#example-2)
		- [Database Connection](#database-connection)
			- [Query](#query)
				- [Native](#native-1)
				- [Query Builder](#query-builder)
		- [Testing](#testing)
			- [Mock](#mock)
		- [Time](#time)
		- [Swagger](#swagger)
			- [Swaggo](#swaggo)
				- [Install](#install)
				- [Comments](#comments)
				- [Usage](#usage)
					- [Echo](#echo-1)
					- [Example](#example-3)
## Basic
### Variables
Go is a typed language and have many ways to declare a variable.
``` Go
var variable int //the value of variable take the zero-value the type, in this case 0;
var variable = 0 // In this case the type is inferred, this way can be grouped
var (
    variable1 = 0
    variable2 = 'a'
    variable3 = "string"
)
variable4 := 0 // Is exactly the same than before but this way only can be used inside a function
```
The `:=` operator make the initialization and declaration.<br>
In `go` we must use the variables that we declared but in the case that we need to declare a variable but don't use it we have the operator "_" it  says to compiler that we don't use it. It is useful in the case that you have a function that return two values but you only need one.<br>
*Example*
```Go
a,_ := functionThatReturnTwoValues()
```

### Constants

For constants we use the keyword *const*  and the usage is exactly the same than the *var*. A const must be initialized.

``` Go
const variable int //error
const variable = 0 // const inferring the type
const ( 
    variable1 = 0
    variable2 = 'a'
    variable3 = "string"
)
```
The operator `:=` only works with variables.

[Example][tour_constants]

### Type Conversions
[Type Conversions][tour_type_conversions]


### Type Assertions
[Type Assertion](https://tour.golang.org/methods/15)


### Pointers
In Go a pointer is a memory address of value. the type `*T` holds the memory address of a `T` data type. In a variable declaration de `*` define a pointer, in a pointer declared it give's the value in that memory address. To get the memory address of a variable we use the operator `&` that return the pointer of that value. <br> 

The zero value for pointers is ``nil``

```Go
package main
func main(){
	var i * int //is a pointer, this means that `i` will holds the memory address of a int
	var p = 43
	i = &p //memory address that have `i` is the memory address of p
	*i= 12 // puts the 12 value in the memory address that i points
	println(p) //print 12
}
```
[Example](https://tour.golang.org/moretypes/1).
### Type
`type` keyword allow us to define new data type like structs or alias for another existing types.
#### Alias
You can put an alias to another data types
```Go
type IntAlias int //this is a alias for int types
type IntArrAlias []int // this is a alias for int slice
type FuncAlias func(string)(bool,error) // this is a alias for functions with this sing
```
This is useful if you want that a datatype implement some interface, for example the usage of sort library
```Go
package main

import (
	"fmt"
	"sort"
)

type ArrInt []int

func (arr ArrInt) Len() int           { return len(arr) }
func (arr ArrInt) Swap(i, j int)      { arr[i], arr[j] = arr[j], arr[i] }
func (arr ArrInt) Less(i, j int) bool { return arr[i] > arr[j] }
func main() {
	var arr ArrInt = ArrInt{1, 5, 3, 7, 3, 8, 3}
	sort.Sort(arr)
	fmt.Printf("%+v\n", arr)
}
```
#### Struct
A struct is a collection of fields, but this fields can be anything, types, functions, interfaces, etc.
```Go
type Validator func()bool,error
type StructA struct{
	Field1 int
	Field2 int
	Field3 *int
}
type StructB struct{
	Field1 StructA //this field is  a struct
	Field2 Validator //this field is a function
	Field3 string
}
```
Is not needed to create an alias for a function but is recommended.


##### Literals
Literal is the way that we can declare and initialize a struct, [here](https://tour.golang.org/moretypes/5) is an explanation.


#### Interfaces
The interfaces in Go define a set of methods, and any struct can implement it by implementing the defined methods by the interface.
```Go
type ReadWritable interface {
	Read() string, error
	Write() error
}
type User strcut{
	Name string
}
func (u User) Write()error{
	println("this user was written")
	return nil
}
func (u User) Read()string, error{
	return "User read", nil
}
```
In the above example the struct `User` implements the interface `ReadWritable` because implements its two defined methods.<br>
 For convention  one-method interfaces are named his method-er

 ```Go
 type Reader interface{
	 Read()string, error
 }
 type Writer interface{
	 Write()error
 } 
 ```

### Functions
A function in Go can return many values, and receive any parameters defined in it. For the definition of the function we use the keyword `func`, after this we put the name of the function.
``` Go
 func nameOfFunction(param1 typeOfParam1, param2 typeOfParam2) returnedType1...{}
```

We can group the params and put its type to the end instead of the end of each one.
``` Go
func nameOfFunction(param1, param2 typeOfParam1&2) 
```


To return multiple values we put the return types in the sign separated by commas and in parentheses after params. To return the values in the function body we put the values separated by commas without parentheses after the keyword *return*.<br>
To get the values from a function that return multiple values we declare a variables separated by commas.
To see all this we going to do an example

``` Go
func AddAndSubs(a, b int)(int, int){
    return a+b, a-b
}

a,b := AddAndSubs(4,1) // in this case a = 5 and b = 3
```
[Example][tour_function_example]


In a function you can receive undetermined number of variables using the **...** operator, also you can send an array or slice using the same operator.


```Go
package main

func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}
func main() {
	arr := []int{1, 2, 3, 4, 5, 6}
	total1 := sum(1, 2, 3, 4, 5, 6)
	total2 := sum(arr...)
	println(total1)
	println(total2)
}
```

Because a function can be data type functions can return functions


```Go
package main

type DB struct {
	DbName string
}

type Closer func() error

func (db DB) Close() error {
	println("I am closing, I am: ",db.DbName)
	return nil
}

func connection() (DB, Closer) {
	var db DB
	return db, db.Close
}

func main() {
	db, closer := connection()
	defer closer() //this don't print the DbName because the function pass is passed when this field was not declared. 
	db.DbName = "testing"
}
```

#### Error
In Go exists an `error` data type, and it is used to handle errors in functions, usually this is the last value returned by the function if the function return an error. <br>
The way to handle this error data type is make a comparison to `nil` or to another error.
```Go
package main

import (
	"errors"
	"log"
)

func sum(nums ...int) (int, error) {
	if len(nums) < 1 {
		return 0, errors.New("the len cant be less than 1")
	}
	total := 0
	for _, num := range nums {
		total += num
	}
	return total, nil
}
func main() {
	arr := []int{1, 2, 3, 4, 5, 6}
	_, err := sum()
	if err != nil {
		log.Print(err)
	}
	total, err := sum(arr...)
	if err != nil {
		log.Fatal(err)
	}
	println(total)
}
```


### GOTO
The `goto` keyword make that current flow goes to a specific tag, this is useful when we want to break some flow and go to a specific place in our code


```Go
package main

func main() {
	count := 0
	for i := 0; i < 100; i++ {
		count++
		for j := 0; j < 1000; j++ {
			count++
			if j == 3 {
				goto FINISHED // jump to FINISHED tag
			}
		}
	}
	count *= 10// the program don't reach here
FINISHED: //finished TAG
	println(count) //print 5
}
```

### Panic
In Go a `panic` happens when something went unexpectedly wrong, unlike errors that are sended in an expected situation. we can throw a new panic `panic(string)`  but it's not recommend because it breaks the app, and for the definition of panic.<br>

Panic's breaks the normal flow of app and finish it, if it is not handled.
We can handle `panic` with the in-built function `recover()`  this function return an error with the panic message.

```Go
package main

import (
	"log"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
		}
	}()
	throwPanic()
}

func throwPanic() {
	panic("new panic")
}
```


### Flow Control

All this section is perfectly explained in the tour and we don't have nothing more to add, for this reason we only left the link and we ask the reader to view all flow control section in the tour.
[Flow Control section][tour_flow_control]

#### Defer

`defer` allow us to call a function or expression when a function is finished for any reason. Multiples `defer` in a function will be stacked, this means that the last `defer` declared is the first `defer` called. We can create anonymous functions or call a created function in `defer`.
```Go
package main

func printMessage(msg string){
	println(msg)
}

func main(){
	defer func(msg string){
		println(msg)
	}("defer1") //last called defer
	defer printMessage("defer2") //second called defer
	defer println("defer3") //first called defer

	println("first print")
}
```

### Packages
All programs in go start running in package *main*. The  definition of package is made in the top of the script. By convention, the package name is the same as the last element of the import path.

[Example][tour_package] <br>
[Style Guide](https://rakyll.org/style-packages/) for go package

### Modules 
Previously we must create our projects in the path that is declared in `GOPATH` or `GOROOT` environment variables, now we use go modules to create it in any place 
```sh
go mod init example/user/hello
```
The above command create two files, `go.mod` and `go.sum`
* go.mod: In this file are the dependencies and the current version of go used in the project. When we use ``go get ...`` that dependency is automatically added to `go.mod`. After running any package building command like `go build` for first time it will install all packages with the specific version.
* go.sum: This file maintains the checksum of the dependencies, so when you run the project again it will no install all packages again, and if its dependencies change the checksum will return in an error.

In order to keep our `go.mod` file updated we have the command
```sh
go mod tidy
```
It bind our project with the `go.mod` file, remove the unused imports and add the missing.
Another useful command is 
```sh
go mod vendor
```
It create a vendor folder with all third-party dependencies and uses this local dependencies instead of download it.

And another commonly used command is 
```sh
go mod download
```
It download all dependencies declared in our `go.mod` file.

### Imports
The imports is the way to use a package from some another place in our current code. we can group the imports
```  Go
import (
    "fmt"
    "math/rand"
)
```
or we can put them separated


``` Go
import "fmt"
import "math/rand"
import mutils "meli/utils"
import outils "own/utils"
import _ "indirectNeededLibrary"
```


As with variables we can use the symbol `_` to say that it library won't be used directly.


### Exports
In Go, the way to say that we want export some variable or function is made using the name of it, if the name begins with capital letter the resource will be exported, in other case doesn't


``` Go
var Value int // is exported
var value int // is private
func DoSomething(){...} // is Exported
func doSomething(){...} // is private
```

## Advance 
In this section we're going to view some advance and specific topics for go.


### JSON
To work with JSON we can use the json library that come by default with Go, and it is enough for almost all cases.

First we need to create a struct with public fields and put a tag to each field, the tag tells to `json` library what is the name of the field when is decoded or encoded. If we don't put tag, the name in encoding and decoding will be the same that the field.

```Go
type Person struct {
	Name     string `json:"name"`
	LastName string `json:"last_name"`
}
```

To decode a struct into a []byte we use the `marshal` function, this function return two values, first a []byte representing the struct and second an error if something wrong happened in the process.

```Go
	personString, err := json.Marshal(person)
```

To pass from `[]byte` to a struct we use the Unmarshal function that receives two parameters, first the []byte with the information and second a pointer to the variable where we want to store the struct.This function return an error.

```Go
var person1 Person
personString := []byte(`{"name":"freddy","last_name":"cuellar"}`)
json.Unmarshal(personString, &person1)
```
putting all together we have the next example


```Go
package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	Name     string `json:"name"`
	LastName string `json:"last_name"`
}

func main() {
	person := Person{
		"freddy",
		"cuellar",
	}
	personString, err := json.Marshal(person)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(personString))
	var person1 Person
	if err := json.Unmarshal(personString, &person1); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v", person1)
}
```


### Concurrency
#### Go Routine
To create a thread in Go we use the keyword *go* and after it call a function. we can create anonymous functions or use previous defined functions.

``` Go
go func(a int){
    return a + 2
}(4)
```

In the above example we create an anonymous function and call it with four as its param, this function is going to run in a new goroutine because we use the keyword *go*.

#### Channels
The way to communicate the goroutines in go is using channels. The chanel must be declared before its use and usually the declaration is done using the keyword *make*. This type of variable have two functions, receive and send values using the **<-** operator
```Go
ch := make(chan int)
ch <- v    // Send v to channel ch.
r := <-ch  // Receive from ch, and
           // assign value to v.
```
By default send and receive blocks the app until the other side is ready. That means sending messages will block the app until the channel can receive messages.  


```Go
package main

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // receive from c

	fmt.Println(x, y, x+y)
}

```
You can run the above example in [Tour Go][tour_concurrency_channel]



##### Buffered Channel
A buffered channel is a channel that only will be block when is full and you can give it a length in its initialization.
``` go
ch := make(chan int, 100)
```

##### Range and Close
You can *close* a channel to notify that no more values will be sent it in it channel, and receiver can  test whether a channel is closed using a second variable.

```Go
v, ok := <-ch

```
ok is false if there are no more values to receive and the channel is closed.

you can loop a channel with *range* and receive its value with a variable in the right side.

```Go
c := make(chan int, 10)
 for i := range c // this loop will continue until c be closed
```
[Here][tour_concurrency_close_range] you can run an example.



##### Select
With *select* we can wait for multiple communication and take the first one. In  *select* statement you can put a default value if the another options are block.

```Go
select {
case i := <-c:
    // use i
default:
    // receiving from c would block
}
```
[Here][tour_concurrency_select] you can run an example




#### Sync
##### Mutex
Mutex is the way that you can ensure that a resource is access only for one goroutine at time. you can found a fully explanation [Here](https://tour.golang.org/concurrency/9)



##### Wait Group
Some times  we want wait that all threads launched finish its work before continue, for this cases exists WaitGroup from sync library. We will use three functionalities `Add`, `Done` and `Wait`
Supposes that `Add` add a number to a count, `Done` subtracts one to that count and `Wait`  blocks the program until the count becomes 0.
```Go
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func worker(iteration, timeToSleep int, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Second * time.Duration(timeToSleep))
	fmt.Printf("iteration %d finished\n", iteration)

}
func main() {
	iterations := 10
	wg.Add(iterations)
	rand.Seed(time.Now().Unix())
	for i := 0; i < iterations; i++ {
		go worker(i, rand.Intn(10), &wg)
	}
	wg.Wait()
	println("all workers finished his work")
}
```

Note that we pass a pointer to wg, if we don't do like that the function `worker` create a copy of wg and don't use the same that was declared in the main function.<br>
If we remove the `wg.Wait`   the program will be finished before the workers do his work.

Now suppose that we want finish when the first worker do its job and stop others, we can use `select` and `close` for it 
```Go
func worker(iteration, timeToSleep int, wg *sync.WaitGroup, end chan int) {
	defer wg.Done()
	select {
	case <-end:
		fmt.Println("acabe antes", iteration)
	case <-time.After(time.Second * time.Duration(timeToSleep)):
		fmt.Printf("iteration %d finished\n", iteration)
		close(end)
	}
}
```
The above example can throw a panic because two works can do its job in the same time an both of them are going to try to close the same channel. We are going to see how solve this using the `context` library.

### Web
First we are going to create an end point with the default go library, after it we will use another third-party library to do this more easy and quickly.
#### Native
In this section we will use net/http default library to build a end point 
##### Server
The first that we need is create a server where we will use our handlers
```Go
server := http.NewServeMux()
```
##### Handler
A handler function in go is defined as a function that receives two variables first is a ResponseWriter that contains the information where we can do the response and a pointer to request, where comes the request information.
```Go
func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}
```
In this case we write the "Hello Word" but if we want write a json  we must create a struct and transform it to json, its can be do it in two ways, first like we see in the [JSON](#json) section, the other ways is:

``` Go
func hello(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(struct{
        Name string `json:"name"`
        LastName string `json:"last_name"`
    }{
        "freddy",
        "cuellar",
    })
}
```


##### Route

Using the above server we can route our handler to a route

```Go
server.HandleFunc("/hello", hello)
```

In this case the function `hello` will be called when a petition to /hello be do it. <br>
`Note`:the filter for the method must be do it using the request parameter in the handler function.

##### Listen and Serve

Using http library we can serve our server to start to listen petitions.

```Go
http.ListenAndServe(":8000", server)
```

In the above line we serve our server in localhost:8000

##### Example


putting all together we have the next example
```Go
package main

import (
	"encoding/json"
	"net/http"
)

type Person struct {
	Name     string `json:"name"`
	LastName string `json:"last_name"`
}

func main() {
	server := http.NewServeMux()
	server.HandleFunc("/hello", hello)
	http.ListenAndServe(":8000", server)
}

func hello(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(struct {
			Name     string `json:"name"`
			LastName string `json:"last_name"`
		}{
			"freddy",
			"cuellar",
		})
	}
}


```
#### Echo
[Echo](https://echo.labstack.com/) is a minimalists go web framework,
first we need to install it.

``` bash
go get -u github.com/labstack/echo/v4
```

and for use it in our project we must import it.


```Go
import 	"github.com/labstack/echo/v4"
```

##### Server
we create our server using the `New` function
```Go
server := echo.New()
```


##### Handler
In `Echo` a handler function receive an echo context and return an error, in this variable we have all that we need and  some others useful functions.

```Go
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
```

In this case we wrote the "Hello Word" as string, but if we want write a json we can do the next:

``` Go
func hello(c echo.Context) error {
	return c.JSON(http.StatusOK, struct{
        Name string `json:"name"`
        LastName string `json:"last_name"`
    }{
        "freddy",
        "cuellar",
    })
}
```


##### Route
Using the above server we can route our handler to a route

```Go
server.GET("/hello", hello)
```

this definition allow us filter the petition by method directly there.

##### Listen and Serve
To start the server call the `start` function 
```Go
server.Start(":8000")
```

in this case the function return an error, and our app is listening the port 8000


##### Example
putting all together we have the next example
``` Go
package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	server := echo.New()
	server.GET("/hello", hello)
	server.Logger.Fatal(server.Start(":8000"))
}
func hello(c echo.Context) error {
	return c.JSON(http.StatusOK, struct {
		Name     string `json:"name"`
		LastName string `json:"last_name"`
	}{
		"freddy",
		"cuellar",
	})
}

```
#### Gin
[Gin](https://gin-gonic.com/docs/) is another framework to do apis in Go, its works is very similar to `Echo`. To install we run the command
```sh
go get -u  github.com/gin-gonic/gin
``` 

and import it in our project

```Go
import "github.com/gin-gonic/gin"
```


##### Server
To start the server 

```Go
server := gin.Default()
```

##### Handler 
The handler in gin is defined as follows
```Go
// func(c *gin.Context) ->template
//example
func hello(c *gin.Context){
		c.String(200,"hello)
}
```
##### Route
```Go
server.Get("/hello", hello)
```

##### Listen and Serve
```Go
server.Run() //run in localhos
```

##### Example
```Go
package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/hello", hello)
	r.Run() // listen and serve on 0.0.0.0:8080
}

func hello(c *gin.Context) {
	c.JSON(200, struct {
		Name     string `json:"name"`
		LastName string `json:"last_name"`
	}{
		"freddy",
		"cuellar",
	})
}


```
### Database Connection
For database connection we need to use the driver for connection, for example, if we want connect to mysql database we must install the driver like that 
``` bash
go get -u github.com/go-sql-driver/mysql
```
And after this we stablish the connection using a URL connection string and database/sql library
```Go
import "database/sql"
db, err := sql.Open("mysql", "root:1234@tcp(project-db:3306)/project")
```
This function receives two parameters, the driver for the database and the URL for connection and return the database connection and an error in case that something wrong happen.
#### Query 
We can make queries using the struct returned by `sql.Open` or using a query builder
##### Native
We can prepare a `statement` to execute it later with the function Prepare, with the symbol **?** we say that it is a parameter that will be passed later. the above function return a Statement struct that will be execute.

```Go
stmt, _ := db.Prepare("INSERT INTO projects values (?,?)")//id, name
	res, err := stmt.Exec(
		0,
		"testing",
		)
	if err != nil {
		log.Print(err.Error()) // proper error handling instead of 
	}

```
##### Query Builder 
[Goqu][goqu] is a query builder that help us to make easier the query management, first we must install it
``` sh
go get -u github.com/doug-martin/goqu/v9
```
now we need to said to `goqu` the dialect that we want and connect this query builder with our connection.
```Go
import (

  "github.com/doug-martin/goqu/v9"
  // import the dialect
  _ "github.com/doug-martin/goqu/v9/dialect/mysql"
)

// look up the dialect
dialect := goqu.Dialect("mysql")
goquDb = dialect.DB(db)// pass the db connection that we did before
```
Now, for do the query we can use our query builder
``` Go
_, err := goquDb.Insert("projects").
	Rows(&projectStruct).
	Executor().
	Exec()
```
this return us an error for validations.<br>
To read the documentation go to [Goqu][goqu]

### Testing
We can test our functions in golang using the `testing` library, it allow us make unit test. By convention the test functions are named with the word Test followed by the name of the function that we want test, and the file is named as `file_name_test.go`. The test function always have the parameter `*testing.T`.<br>
for example suppose that we have the next function in our main package
```GO
package main
import "errors"
var (
	ErrZeroInput = errors.New("invalid input")
)

func Sum(a, b int) (int, error) {
	if a == 0 || b == 0 {
		return 0, ErrZeroInput
	}
	return a + b, nil
}
```

Now we want test the error returned and look if the sum was did right.

```Go
package main

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	//Test invalid input
	t.Run("Test invalid input",
		func(t *testing.T) {
			var a, b int
			want := ErrZeroInput
			val, err := Sum(a, b)
			if !errors.Is(err, want) || val != 0 {
				t.Fatalf(
					"Sum(%d,%d) = %d,%v, want match with %d, %v",
					a, b, val, err, 0, want,
				)
			}
		})
}

//Test invalid input
func TestSumInvalidInput(t *testing.T) {
	var a, b int
	want := ErrZeroInput
	val, err := Sum(a, b)
	if !errors.Is(err, want) || val != 0 {
		t.Fatalf(
			"Sum(%d,%d) = %d,%v, want match with %d, %v",
			a, b, val, err, 0, want,
		)
	}
}
// For use the assert library we need install it running the command 
// `go get -u github.com/stretchr/testify/assert`

//Testing With assert a invalid input
func TestSumInvalidInputAssert(t *testing.T) {
	assert := assert.New(t)
	var a, b int
	want := ErrZeroInput
	val, err := Sum(a, b)
	assert.Equal(val, 0)
	assert.ErrorIs(err, want)
}

//Answer that we want for normal input
func TestSumAssert(t *testing.T) {
	assert := assert.New(t)
	var a, b int
	a, b = 4, 3
	val, err := Sum(a, b)
	assert.Equal(val, 7)
	assert.Nil(err)
}

```
#### Mock
The above library also have functionalities that allow us make mocks and used it in ours tests.
[Here](https://github.com/stretchr/testify#mock-package) is a full documentation.
### Time
In go we have the `time` library to handle time variables.
```Go
import (
	"time"
)
now := time.Now()
```
In the above example we got the current data time.
For all operations that we want to do with time the best way is using this library, use `time.Time` for instance of time and use  `time.Duration` for periods of time.

This library help us with comparisons, aggregations and parsing.
```Go
now := time.Now()
now = now.Add(2 * *time.Hour) //add two hours to current time
now = now.AddDate(0, 0, 1) //add one day to current data time
now.Before(time.Now()) // compare if the data time is before that the time passed as param
now.After(time.Now()) // compare if the data time is after that the time passed as param
now.Sub(time.Now())// return the difference between the two dates
layout := "2006-01-02 15:04:05 -0700" // define the format for parsing
timeS := now.Format(layout) //parse time to string following the format
fmt.Println(timeS)
t, err := time.Parse(layout, timeS) //parse string to time following the format
if err != nil {
	fmt.Println("Error while parsing date :", err)
}
fmt.Println(t.Format(layout))

//some examples about layouts in golang
monthLayout := "Month = January,Jan,01"
dayLayout := "Day = Monday, Mon,02"
yearLayout := "Year = 2006, 006, 06"
hourLayout := "Hour = 03:04:05 PM"
hour24Layout := "Hour = 15:04:05"
```
[Here](https://programming.guide/go/format-parse-string-time-date-example.html) we have a more complete documentation about format in go, and [Here](https://github.com/uber-go/guide/blob/master/style.md#use-time-to-handle-time) we can found some useful tips for manage the date times.

### Swagger 
Swagger is a set tools that help us to document ours projects.
#### Swaggo
##### Install
Swaggo is a library that converts Go annotations to Swagger Documentation 2.0. 
To install it follow [this instructions](https://github.com/swaggo/swag#getting-started).
##### Comments
To use this library we must comment or code with annotations, here is a [fully explanation](https://github.com/swaggo/swag#declarative-comments-format) about its comments. 
##### Usage
This depends the way that your are doing your api, if you are using `echo`, `gin` or the `http` library.
###### Echo
We will do this example with echo.<br>
First install  the echo swagger
```sh
go get -u github.com/swaggo/echo-swagger
```
later import the middleware installed in our code
```Go
import "github.com/swaggo/echo-swagger"
```
now, we create our function with comments
```Go
// CreateUser godoc
// @Summary      save user
// @Description   save sended user
// @Accept       json
// @Param	     user body User true "user that will be saved"
// @Produce      json
// @Success      200  {object}  User
// @Router       /user/save [post]
func CreateUser(c echo.Context) error {
	var user User
	if err := c.Bind(&user); err != nil {
		return err
	}
	fmt.Printf("guardando usuario user: %v\n", user) //logica de guardar
	return c.JSON(http.StatusCreated, user)
}
```
And run the command
```sh
swag init
```
Add this line to our server
```Go
e.GET("/swagger/*", echoSwagger.WrapHandler)
```
and run our application and in `localhost:8000/swagger/` we must see the swagger document deployed
###### Example
```Go
package main

import (
	"fmt"
	"guide/docs"
	"net/http"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

type User struct {
	Name     string `json:"name,omitempty"`
	LastName string `json:"last_name,omitempty"`
}

func main() {

	//---- swagger-----
	// programmatically set swagger info
	docs.SwaggerInfo.Title = "Swagger Example API"
	docs.SwaggerInfo.Description = "This is a sample server."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8000"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	// use echo Swagger middleware to serve the API docs
	server := echo.New()
	server.GET("/swagger/*any", echoSwagger.WrapHandler)

	user := server.Group("/user")
	user.POST("/save", CreateUser)

	server.Logger.Fatal(server.Start(":8000"))
}
// CreateUser godoc
// @Summary      save user
// @Description   save sended user
// @Accept       json
// @Param	     user body User true "user that will be saved"
// @Produce      json
// @Success      200  {object}  User
// @Router       /user/save [post]
func CreateUser(c echo.Context) error {
	var user User
	if err := c.Bind(&user); err != nil {
		return err
	}
	fmt.Printf("guardando usuario user: %v\n", user) //logica de guardar
	return c.JSON(http.StatusCreated, user)
}
```
<!-- links -->
[tour_flow_control]:https://tour.golang.org/flowcontrol/1
[tour_constants]:https://tour.golang.org/basics/16
[tour_function_example]:https://tour.golang.org/basics/6
[tour_guide_golang]: https://tour.golang.org/welcome/1
[tour_package]: https://tour.golang.org/basics/1
[tour_type_conversions]:https://tour.golang.org/basics/13
[tour_concurrency_channel]: https://tour.golang.org/concurrency/2
[tour_concurrency_close_range]: https://tour.golang.org/concurrency/4
[tour_concurrency_select]: https://tour.golang.org/concurrency/6
[goqu]: https://doug-martin.github.io/goqu
