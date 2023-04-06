package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"

	"github.com/gofiber/fiber/v2"
)

type User struct {
	ID         int    `json:"id"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Email      string `json:"email"`
	Age        int    `json:"age"`
	AvatarPath string `json:"avatar_path"`
}

func main() {
	// Логирование
	f, err := os.OpenFile("file.log", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()
	log.SetOutput(f)
	//----------------------------- 



	app := fiber.New()

	app.Get("/users", func(c *fiber.Ctx) error {
		users, err := loadUsers()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		user := users[rand.Intn(len(users))]
		user.AvatarPath = fmt.Sprintf("/images/%d.jpg", user.ID)

		
		log.Println(fmt.Sprintf("Был отправлен пользовалет %d", user.ID))

		return c.JSON(user)
	})

	app.Static("/images", "./images")

	fmt.Println("Server started on port 8080")
	app.Listen(":8080")
}

func loadUsers() ([]User, error) {
	users := make([]User, 0)

	data, err := readFile("users.json")
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &users)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func readFile(filename string) ([]byte, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return data, nil
}