package main

import (
	"api/go/dto"
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
func main() {
	app := fiber.New()
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
}
