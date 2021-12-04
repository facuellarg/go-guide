package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func worker(iteration, timeToSleep int, wg *sync.WaitGroup, end chan int) {
	defer wg.Done()
	select {
	// case <-ctx.Done():
	case <-end:
		fmt.Println("acabe antes", iteration)
	case <-time.After(time.Second * time.Duration(timeToSleep)):
		fmt.Printf("iteration %d finished\n", iteration)
		if !IsClosed(end) {
			close(end)
		}
	}
}

func IsClosed(ch <-chan int) bool {
	select {
	case <-ch:
		return true
	default:
	}

	return false
}

func main() {
	end := make(chan int)
	test(end)
	println("fuera de la funcion que llamo las gorutines")
	wg.Wait()
	// time.Sleep(time.Second * 10)
}

func test(end chan int) {
	iterations := 10
	wg.Add(iterations)
	rand.Seed(time.Now().Unix())
	for i := 0; i < iterations; i++ {
		go worker(i, rand.Intn(10), &wg, end)
	}
	println("all workers finished his work")
}

// package main

// import (
// 	"fmt"
// 	"guide/docs"
// 	"net/http"

// 	"github.com/labstack/echo/v4"
// 	echoSwagger "github.com/swaggo/echo-swagger"
// )

// type ErrorEcho struct {
// 	Message string `json:"message"`
// }
// type User struct {
// 	Name     string `json:"name,omitempty"`
// 	LastName string `json:"last_name,omitempty"`
// }

// func main() {

// 	//---- swagger-----
// 	// programmatically set swagger info
// 	docs.SwaggerInfo.Title = "Swagger Example API"
// 	docs.SwaggerInfo.Description = "This is a sample server."
// 	docs.SwaggerInfo.Version = "1.0"
// 	docs.SwaggerInfo.Host = "localhost:8000"
// 	docs.SwaggerInfo.Schemes = []string{"http", "https"}

// 	// use ginSwagger middleware to serve the API docs
// 	server := echo.New()

// 	server.GET("/swagger/*any", echoSwagger.WrapHandler)
// 	server.Use(myMiddlewareFunc)
// 	server.GET("/bye", bye)

// 	user := server.Group("/user")
// 	user.GET("/hello", hello)
// 	user.POST("/save", CreateUser)

// 	server.Logger.Fatal(server.Start(":8000"))
// }
// func hello(c echo.Context) error {
// 	name := c.Request().Header.Get("name-header")
// 	if name == "" {
// 		return c.JSON(http.StatusNotFound, ErrorEcho{"no se encontro la cabecera"})
// 	}
// 	val := c.Get("userVal").(int)
// 	fmt.Printf("val: %v\n", val)
// 	return c.JSON(http.StatusOK, User{
// 		name,
// 		"cuellar",
// 	})
// }

// CreateUser godoc
// @Summary      save user
// @Description   save sended user
// @Accept       json
// @Param	     user body User true "user that will be saved"
// @Produce      json
// @Success      200  {object}  User
// @Router       /user/save [post]
// func CreateUser(c echo.Context) error {
// 	var user User
// 	if err := c.Bind(&user); err != nil {
// 		return err
// 	}
// 	fmt.Printf("guardando usuario user: %v\n", user) //logica de guardar
// 	return c.JSON(http.StatusCreated, user)
// }

// func bye(c echo.Context) error {
// 	return c.JSON(http.StatusOK, User{
// 		"bye",
// 		"bye",
// 	})
// }

// func myMiddlewareFunc(next echo.HandlerFunc) echo.HandlerFunc {
// 	println("hecho por nosostros")
// 	return func(c echo.Context) error {
// 		c.Set("userVal", 10)
// 		return next(c)
// 	}
// }
