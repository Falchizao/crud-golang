package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/falchizao/crud-golang/models"
	"github.com/falchizao/crud-golang/storage"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func (r *Repository) UpdateUser(context *fiber.Ctx) error {

	id := context.Params("id")

	if id == "" {
		context.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "id nao pode ser vazio",
		})
		return nil
	}

	var idcast = StringToUint(id)

	user := models.Users{ID: idcast}

	err := context.BodyParser(&user)

	if err != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{"message": "request failed"})
		return err
	}

	err = r.DB.Updates(&user).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{"message": "erro ao fazer update do usuario"})
		return err
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{"message": "usuario atualizado com sucesso"})

	return nil
}

func (r *Repository) CreateUser(context *fiber.Ctx) error {
	user := User{}

	err := context.BodyParser(&user)

	if err != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{"message": "request failed"})
		return err
	}

	err = r.DB.Create(&user).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{"message": "erro ao criar usuario"})
		return err
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{"message": "usuario criado com sucesso"})
	return nil
}

func (r *Repository) GetUsers(context *fiber.Ctx) error {
	userModels := &[]models.Users{}

	err := r.DB.Find(userModels).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{"message": "erro ao buscar usuarios"})
		return err
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "sucesso ao buscar usuarios",
		"data":    userModels,
	})

	return nil
}

func StringToUint(s string) uint {
	i, _ := strconv.Atoi(s)
	return uint(i)
}

type User struct {
	Nome  string `json:"nome"`
	Email string `json:"email"`
	Senha string `json:"senha"`
	Idade int    `json:"idade"`
}

func (r *Repository) DeleteUser(context *fiber.Ctx) error {
	user := models.Users{}
	id := context.Params("id")

	if id == "" {
		context.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "id nao pode ser vazio",
		})
		return nil
	}

	err := r.DB.Delete(user, id)

	if err.Error != nil {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "nao foi possivel deletar o usuario",
		})

		return err.Error
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "usuario deletado com sucesso",
	})

	return nil

}
func (r *Repository) GetUserByID(context *fiber.Ctx) error {
	id := context.Params("id")
	userModel := &models.Users{}
	if id == "" {
		context.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "id nao pode ser vazio",
		})
		return nil
	}

	err := r.DB.Where("id = ?", id).First(userModel).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{
				"message": "nao foi possivel encontrar o usuario",
			})
		return err
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "usuario encontrado", "data": userModel,
	})
	return nil

}

func (r *Repository) SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Post("/createUser", r.CreateUser)
	api.Delete("/deleteUser/:id", r.DeleteUser)
	api.Get("/get_user/:id", r.GetUserByID)
	api.Get("/users", r.GetUsers)
	api.Put("/update_user/:id", r.UpdateUser)
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	config := &storage.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Password: os.Getenv("DB_PASS"),
		User:     os.Getenv("DB_USER"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
		DBName:   os.Getenv("DB_NAME"),
	}

	db, err := storage.NewConnection(config)

	if err != nil {
		log.Fatal("Erro ao conectar ao banco")
	}

	err = models.MigrateUsers(db)

	if err != nil {
		log.Fatal("erro ao fazer a migration")
	}

	r := Repository{
		DB: db,
	}
	app := fiber.New()

	r.SetupRoutes(app)
	fmt.Println(app)
	app.Listen(":8484")

}
