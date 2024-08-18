package main

import (
	"api/go/dto"
	"encoding/json"
	"os"
	"sync"
	"time"
	"unicode/utf8"
	"unsafe"

	// "api/go/helper"
	"api/go/helper/utility"
	"api/go/middleware/validation"
	"fmt"
	"reflect"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/sirupsen/logrus"
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

// Function to add log entry to the global slice and write to file
// func addLogEntry(entry LogEntry) {
// 	mutex.Lock()
// 	defer mutex.Unlock()

// 	// Append the new log entry to the slice
// 	logEntries = append(logEntries, entry)

// 	// Write the entire slice as a JSON array to the file
// 	writeLogsToFile()
// }

// // Function to write all log entries to the file as a JSON array
// func writeLogsToFile() {
// 	file, err := os.OpenFile("logs.json", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
// 	if err != nil {
// 		logrus.Fatalf("Failed to open log file: %v", err)
// 	}
// 	defer file.Close()

// 	encoder := json.NewEncoder(file)
// 	encoder.SetIndent("", "  ") // Pretty-print with indentation
// 	if err := encoder.Encode(logEntries); err != nil {
// 		logrus.Fatalf("Failed to write log entries: %v", err)
// 	}
// }

// // Function to write log entry to a file
func writeLogToFile(entry LogEntry) {
	mutex.Lock()
	defer mutex.Unlock()

	// Open the file in append mode
	file, err := os.OpenFile("logs.json", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		logrus.Fatalf("Failed to open log file: %v", err)
	}
	defer file.Close()

	// Encode log entry to JSON
	encoder := json.NewEncoder(file)
	if err := encoder.Encode(entry); err != nil {
		logrus.Fatalf("Failed to write log entry: %v", err)
	}
}
func listRoutes(app *fiber.App) []fiber.Route {
	var routers []fiber.Route

	for _, routes := range app.Stack() {
		for _, route := range routes {
			routers = append(routers, fiber.Route{
				Method:   route.Method,
				Path:     route.Path,
				Handlers: route.Handlers,
				Params:   route.Params,
				Name:     route.Name,
			})
		}

	}
	return routers
}

type LogEntry struct {
	Time         string                 `json:"time"`
	Method       string                 `json:"method"`
	URL          string                 `json:"url"`
	Request      map[string]interface{} `json:"request"`
	Response     map[string]interface{} `json:"response"`
	ResponseTime int64                  `json:"response_time"`
}

var logEntries []LogEntry
var mutex sync.Mutex

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
	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		},
	})

	// Initialize logrus
	// log := logrus.New()

	// // Set log output to a file
	// file, err := os.OpenFile("logs.json", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	// if err != nil {
	// 	panic(err)
	// }
	// log.SetOutput(file)
	// log.SetFormatter(&logrus.JSONFormatter{})

	// Middleware to log request and response
	app.Use(func(c *fiber.Ctx) error {
		// Record the start time
		startTime := time.Now()

		// Capture request data
		reqData := make(map[string]interface{})
		reqData["body"] = string(c.Body())

		// Continue to the next handler
		err := c.Next()

		// Capture response data
		resData := make(map[string]interface{})
		resData["status"] = c.Response().StatusCode()
		resData["body"] = string(c.Response().Body())

		// Calculate response time in milliseconds
		elapsedTime := time.Since(startTime).Milliseconds()
		// Create log entry
		entry := LogEntry{
			Time:         time.Now().Format(time.RFC3339),
			Method:       c.Method(),
			URL:          c.OriginalURL(),
			Request:      reqData,
			Response:     resData,
			ResponseTime: elapsedTime,
		}

		// Add log entry to the slice in a thread-safe manner
		mutex.Lock()
		logEntries = append(logEntries, entry)
		mutex.Unlock()
		// Write log entry to file
		writeLogToFile(entry)
		return err
	})
	// Initialize default config
	app.Use(logger.New())
	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))
	// Initialize default config
	app.Use(compress.New())

	// Or extend your config for customization
	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed, // 1
	}))

	// Initialize default config
	// app.Use(csrf.New())

	// // Or extend your config for customization
	// app.Use(csrf.New(csrf.Config{
	// 	KeyLookup:      "header:X-Csrf-Token",
	// 	CookieName:     "csrf_",
	// 	CookieSameSite: "Lax",
	// 	Expiration:     1 * time.Hour,
	// 	KeyGenerator:   utils.UUIDv4,
	// }))
	// Or extend your config for customization
	// Logging remote IP and Port

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
			// time.Sleep(5 * time.Second)
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
	routes := listRoutes(app)
	for _, route := range routes {
		fmt.Println(route.Path, " ", route.Method)
	}
	// fmt.Println("ROUTES", routes)
	app.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "Route Not founds"})
	})
	// app.Use(func(c *fiber.Ctx) error {
	// 	return c.Status(fiber.StatusMethodNotAllowed).JSON(fiber.Map{"message": "Method Not Allowed"})
	// })

	app.Listen(":4000")
	// log.Fatal(http.ListenAndServe(":8093", mux))

}
