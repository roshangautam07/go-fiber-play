package main

import (
	"api/go/dto"
	"time"
	"unicode/utf8"
	"unsafe"

	// "api/go/helper"
	"api/go/helper/utility"
	"api/go/middleware/validation"
	"fmt"
	"reflect"

	"github.com/gofiber/fiber/v2"
)

// type Person struct {
// 	ID         int
// 	firstName  string
// 	middleName string
// 	lastName   string
// 	isMarried  bool
// }
// func fun(c *fiber.Ctx) error {

// 	payload := make(map[string]interface{})
// 	if err := c.BodyParser(&payload); err != nil {
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 			"error": "Invalid request payload",
// 		})
// 	}

// 	// Perform validation
// 	errors := helper.NewValidationBuilder(payload).
// 		ValidateRequiredKeys([]string{"id", "firstName", "middleName", "lastName"}).
// 		// CheckLength([]string{"randomToken"}).
// 		IsEmptyOrNull().
// 		IsString([]string{"firstName", "middleName", "lastName"}).
// 		// IsInt([]string{"id"}).
// 		Build()

// 	// Check for validation errors
// 	if len(errors) > 0 {
// 		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"error": errors})

// 	} else {
// 		fmt.Println("Validation passed!")
// 	}

//		return c.Next()
//	}
func printNumbers() []int {
	var number []int
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
		number = append(number, i)

	}
	return number
}

// func scatterPlot() statsviz.TimeSeriesPlot {
// 	// Describe the 'sine' time series.
// 	sine := statsviz.TimeSeries{
// 		Name:     "short sin",
// 		Unitfmt:  "%{y:.4s}B",
// 		GetValue: updateSine,
// 	}

// 	// Build a new plot, showing our sine time series
// 	plot, err := statsviz.TimeSeriesPlotConfig{
// 		Name:  "sine",
// 		Title: "Sine",
// 		Type:  statsviz.Scatter,
// 		InfoText: `This is an example of a 'scatter' type plot, showing a single time series.<br>
// InfoText field (this) accepts any HTML tags like <b>bold</b>, <i>italic</i>, etc.`,
// 		YAxisTitle: "y unit",
// 		Series:     []statsviz.TimeSeries{sine},
// 	}.Build()
// 	if err != nil {
// 		log.Fatalf("failed to build timeseries plot: %v", err)
// 	}

// 	return plot
// }

// func barPlot() statsviz.TimeSeriesPlot {
// 	// Describe the 'user logins' time series.
// 	logins := statsviz.TimeSeries{
// 		Name:     "user logins",
// 		Unitfmt:  "%{y:.4s}",
// 		GetValue: logins,
// 	}

// 	// Describe the 'user signins' time series.
// 	signins := statsviz.TimeSeries{
// 		Name:     "user signins",
// 		Unitfmt:  "%{y:.4s}",
// 		GetValue: signins,
// 	}

// 	// Build a new plot, showing both time series at once.
// 	plot, err := statsviz.TimeSeriesPlotConfig{
// 		Name:  "users",
// 		Title: "Users",
// 		Type:  statsviz.Bar,
// 		InfoText: `This is an example of a 'bar' type plot, showing 2 time series.<br>
// InfoText field (this) accepts any HTML tags like <b>bold</b>, <i>italic</i>, etc.`,
// 		YAxisTitle: "users",
// 		Series:     []statsviz.TimeSeries{logins, signins},
// 	}.Build()
// 	if err != nil {
// 		log.Fatalf("failed to build timeseries plot: %v", err)
// 	}

// 	return plot
// }

// func stackedPlot() statsviz.TimeSeriesPlot {
// 	// Describe the 'user logins' time series.
// 	logins := statsviz.TimeSeries{
// 		Name:     "user logins",
// 		Unitfmt:  "%{y:.4s}",
// 		Type:     statsviz.Bar,
// 		GetValue: logins,
// 	}

// 	// Describe the 'user signins' time series.
// 	signins := statsviz.TimeSeries{
// 		Name:     "user signins",
// 		Unitfmt:  "%{y:.4s}",
// 		Type:     statsviz.Bar,
// 		GetValue: signins,
// 	}

// 	// Build a new plot, showing both time series at once.
// 	plot, err := statsviz.TimeSeriesPlotConfig{
// 		Name:    "users-stack",
// 		Title:   "Stacked Users",
// 		Type:    statsviz.Bar,
// 		BarMode: statsviz.Stack,
// 		InfoText: `This is an example of a 'bar' plot showing 2 time series stacked on top of each other with <b>BarMode:Stack</b>.<br>
// InfoText field (this) accepts any HTML tags like <b>bold</b>, <i>italic</i>, etc.`,
// 		YAxisTitle: "users",
// 		Series:     []statsviz.TimeSeries{logins, signins},
// 	}.Build()
// 	if err != nil {
// 		log.Fatalf("failed to build timeseries plot: %v", err)
// 	}

// 	return plot
// }

// var val = 0.

// func updateSine() float64 {
// 	val += 0.5
// 	return math.Sin(val)
// }

// func logins() float64 {
// 	return (rand.Float64() + 2) * 1000
// }

//	func signins() float64 {
//		return (rand.Float64() + 1.5) * 100
//	}
func DoneAsync() chan int {
	r := make(chan int)
	fmt.Println("Warming up ...")
	go func() {
		time.Sleep(3 * time.Second)
		r <- 1
		fmt.Println("Done ...")
	}()
	return r
}
func main() {
	// mux := http.NewServeMux()

	// // Register statsviz handlers and 3 addition user plots.
	// if err := statsviz.Register(mux,
	// 	statsviz.TimeseriesPlot(utility.ScatterPlot()),
	// 	statsviz.TimeseriesPlot(utility.BarPlot()),
	// 	statsviz.TimeseriesPlot(utility.StackedPlot()),
	// ); err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println("Point your browser to http://localhost:8093/debug/statsviz/")
	app := fiber.New()
	a := dto.Mystruct{}
	fmt.Println(unsafe.Sizeof(a))

	user := dto.Person{
		ID:         1,
		FirstName:  "roshan",
		MiddleName: "Ma",
		LastName:   "Gautam",
		IsMarried:  false,
		// FullName: func(FirstName string, MiddleName string, LastName string) string {
		// 	return strings.Join([]string{FirstName, MiddleName, LastName}, " ")
		// },
		Contacts: []dto.Contact{
			{Type: "email", Detail: "roshan@example.com"},
			{Type: "phone", Detail: "123-456-7890"},
		},
		PAN: dto.PAN{
			Type:   "Individual",
			Number: "ABCDE1234F",
		},
		Hobbies: []string{"playing", "dancing"},
	}
	users := []dto.Person{
		{ID: 1, FirstName: "roshan", MiddleName: "M", LastName: "Gautam", IsMarried: false},
		{ID: 2, FirstName: "roshan", MiddleName: "M", LastName: "Gautam", IsMarried: false, Salary: "1000"},
	}

	subjectMarks := map[string]any{"Golang": 85, "Java": "80", "Python": 81}
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Hello, World!")
	})

	app.Route("/api", func(router fiber.Router) {
		router.Get("get", func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusOK).JSON(user)
		})
		router.Get("/all", func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusOK).JSON(users)
		})
		router.Get("/map", func(c *fiber.Ctx) error {
			return c.Status((fiber.StatusOK)).JSON(subjectMarks)
		})
		router.Get("/fiber", func(c *fiber.Ctx) error {
			build := make(map[string]interface{})
			build["id"] = 10
			build["name"] = "roshan"
			return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": fiber.Map{"person": build}})

		})
		router.Get("/long", func(c *fiber.Ctx) error {
			go printNumbers()
			time.Sleep(1 * time.Second)
			return c.JSON(fiber.Map{"number": ":herr"})
		})
		router.Get("/defer", func(c *fiber.Ctx) error {
			// fmt.Println("Let's start ...")
			// val := DoneAsync()
			// fmt.Println("Done is running ...")
			// fmt.Println(<-val)
			//We should know defer statements will run in LIFO (Last In, First Out) order:
			defer fmt.Println("Statement 1")
			defer fmt.Println("Statement 2")
			defer fmt.Println("Statement 3")
			//https://rezakhademix.medium.com/defer-functions-in-golang-common-mistakes-and-best-practices-96eacdb551f0
			for i := 0; i < 10; i++ {
				defer fmt.Println(i)
				//We expected the first printed value to be 0,
				//but using defer keyword will delay the result,
				//stack them and by LIFO behavior the result will be 9876543210
			}
			return c.JSON(fiber.Map{"number": "val"})

		})
		router.Get("/string", func(c *fiber.Ctx) error {
			//In Golang, strings are made up of bytes (slice of bytes)
			//and some characters need to store in multiple bytes e.g: "♥"
			str := "hss♥"
			// st := "é"
			return c.JSON(fiber.Map{"byte": len(str), "len": utf8.RuneCountInString(str)})
		})
		router.Post("/user", validation.ValidateUser, func(c *fiber.Ctx) error {
			p := new(dto.User) // Use new to create a pointer to the struct

			if err := c.BodyParser(p); err != nil {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
			}
			fmt.Println(p)

			return c.Status(fiber.StatusOK).JSON(p)
		})
		router.Get("filter/:attributes/:value", func(c *fiber.Ctx) error {
			attributes := c.Params("attributes")
			value := c.Params("value")
			// fmt.Println(attributes, value)
			filteredUser := utility.Filter(users, func(us dto.Person) bool {
				r := reflect.ValueOf(us)
				field := reflect.Indirect(r).FieldByName(attributes)
				fmt.Println(field.String())

				if !field.IsValid() {
					fmt.Println("Invalid field name:", attributes)
					return false
				}

				// Handle different types appropriately
				switch field.Kind() {
				case reflect.String:
					fmt.Println("Field is a string:", field.String())
					return field.String() == value
				case reflect.Int:
					fmt.Println("Field is an int:", field.Int())
					return fmt.Sprint(field.Int()) == value
				case reflect.Bool:
					fmt.Println("Field is a bool:", field.Bool())
					return fmt.Sprint(field.Bool()) == value
				case reflect.Interface:
					fmt.Println("Field is an interface:", field.Interface())
					return fmt.Sprint(field.Interface()) == value
				default:
					fmt.Println("Field type is not handled:", field.Kind())
					return false
				}
			})

			return c.Status(fiber.StatusOK).JSON(filteredUser)
		})
	})

	app.Listen(":4000")
	// log.Fatal(http.ListenAndServe(":8093", mux))

}
