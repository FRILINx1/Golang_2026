package main

import (
	"database/sql"
	"log"
	"strconv"
	"os"
	"lab8/models" 

	"github.com/gofiber/fiber/v3"
	_ "github.com/lib/pq" 
	"github.com/mailru/easyjson"
)

var db *sql.DB

func main() {

	connStr := os.Getenv("DATABASE_URL")
	if connStr == "" {
		connStr = "postgres://mukutko:1234@localhost:5432/go_bd_lab_8?sslmode=disable"
	}
	
	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Не вдалося ініціалізувати драйвер БД: ", err)
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		log.Fatal("Не вдалося підключитися до бази даних: ", err)
	}
	log.Println("Успішне підключення до PostgreSQL!")

	app := fiber.New()

	app.Get("/contacts", getContacts)
	app.Get("/contacts/:id", getContactByID)
	app.Post("/contacts", createContact)
	app.Put("/contacts/:id", updateContact)
	app.Delete("/contacts/:id", deleteContact)

	log.Println("Сервер запущено на http://localhost:8080")
	log.Fatal(app.Listen(":8080"))
}

func getContacts(c fiber.Ctx) error {
	rows, err := db.Query("SELECT id, name, phone FROM contacts")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Помилка бази даних"})
	}
	defer rows.Close()

	var contacts models.ContactList
	for rows.Next() {
		var contact models.Contact
		if err := rows.Scan(&contact.ID, &contact.Name, &contact.Phone); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Помилка читання даних"})
		}
		contacts = append(contacts, contact)
	}

	bytes, _ := easyjson.Marshal(contacts)
	c.Set("Content-Type", "application/json")
	return c.Send(bytes)
}

func getContactByID(c fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Невалідний ID"})
	}

	var contact models.Contact
	err = db.QueryRow("SELECT id, name, phone FROM contacts WHERE id=$1", id).Scan(&contact.ID, &contact.Name, &contact.Phone)
	
	if err == sql.ErrNoRows {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Контакт не знайдено"})
	} else if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Помилка БД"})
	}

	bytes, _ := easyjson.Marshal(contact)
	c.Set("Content-Type", "application/json")
	return c.Send(bytes)
}

func createContact(c fiber.Ctx) error {
	var newContact models.Contact 

	if err := easyjson.Unmarshal(c.Body(), &newContact); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Невалідний JSON"})
	}

	err := db.QueryRow(
		"INSERT INTO contacts (name, phone) VALUES ($1, $2) RETURNING id",
		newContact.Name, newContact.Phone,
	).Scan(&newContact.ID)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Помилка при створенні контакту"})
	}

	bytes, _ := easyjson.Marshal(newContact)
	c.Set("Content-Type", "application/json")
	return c.Status(fiber.StatusCreated).Send(bytes)
}

func updateContact(c fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Невалідний ID"})
	}

	var updatedData models.Contact
	if err := easyjson.Unmarshal(c.Body(), &updatedData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Невалідний JSON"})
	}

	res, err := db.Exec("UPDATE contacts SET name=$1, phone=$2 WHERE id=$3", updatedData.Name, updatedData.Phone, id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Помилка оновлення"})
	}

	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Контакт не знайдено"})
	}

	updatedData.ID = id
	bytes, _ := easyjson.Marshal(updatedData)
	c.Set("Content-Type", "application/json")
	return c.Send(bytes)
}

func deleteContact(c fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Невалідний ID"})
	}

	res, err := db.Exec("DELETE FROM contacts WHERE id=$1", id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Помилка видалення"})
	}

	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Контакт не знайдено"})
	}

	return c.SendStatus(fiber.StatusNoContent) 
}