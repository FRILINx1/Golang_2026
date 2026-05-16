package main

import (
	"log"
	"strconv"
	"myapi/models" 

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"github.com/mailru/easyjson"
)

var (
	notes    = models.NoteList{}
	nextID   = 1
	validate = validator.New()
)

func main() {
	
	app := fiber.New()

	app.Get("/notes", getNotes)
	app.Get("/notes/:id", getNoteByID)
	app.Post("/notes", createNote)
	app.Put("/notes/:id", updateNote)
	app.Delete("/notes/:id", deleteNote)

	log.Println("Сервер запущено на http://localhost:8080")
	log.Fatal(app.Listen(":8080"))
	
}

func getNotes(c fiber.Ctx) error {

	bytes, _ := easyjson.Marshal(notes)
	c.Set("Content-Type", "application/json")
	return c.Send(bytes)
	
}

func getNoteByID(c fiber.Ctx) error {
	
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Incorrect id number"})
	}

	for _, note := range notes {
		if note.ID == id {
			bytes, _ := easyjson.Marshal(note)
			c.Set("Content-Type", "application/json")
			return c.Send(bytes)
		}
	}

	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Note not found"})
	
}

func createNote(c fiber.Ctx) error {
	
	var newNote models.Note 

	if err := easyjson.Unmarshal(c.Body(), &newNote); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid JSON format"})
	}

	if err := validate.Struct(&newNote); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Validation failed"})
	}

	newNote.ID = nextID
	nextID++
	notes = append(notes, newNote)

	bytes, _ := easyjson.Marshal(newNote)
	c.Set("Content-Type", "application/json")
	return c.Status(fiber.StatusCreated).Send(bytes)
}

func updateNote(c fiber.Ctx) error {
	
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Incorrect id number"})
	}

	var updatedData models.Note
	if err := easyjson.Unmarshal(c.Body(), &updatedData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid JSON format"})
	}

	if err := validate.Struct(&updatedData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Validation failed"})
	}

	for i, note := range notes {
		if note.ID == id {
			notes[i].Title = updatedData.Title
			notes[i].Content = updatedData.Content

			bytes, _ := easyjson.Marshal(notes[i])
			c.Set("Content-Type", "application/json")
			return c.Send(bytes)
		}
	}

	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Note not found"})
}

func deleteNote(c fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Incorrect id number"})
	}

	for i, note := range notes {
		if note.ID == id {
			notes = append(notes[:i], notes[i+1:]...)
			return c.SendStatus(fiber.StatusNoContent)
		}
	}

	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Note not found"})
}