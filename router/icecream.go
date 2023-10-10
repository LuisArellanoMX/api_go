package router

import (
	"github.com/LuisArellanoMX/api_go/common"
	"github.com/LuisArellanoMX/api_go/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddIcecreamGroup(app *fiber.App) {
	icecreamGroup := app.Group("/icecream")

	icecreamGroup.Get("/", getIcecreams)
	icecreamGroup.Get("/:id", getIcecream)
	icecreamGroup.Post("/", createIcecream)
	icecreamGroup.Put("/:id", updateIcecream)
	icecreamGroup.Delete("/:id", deleteIcecream)
}

func getIcecreams(c *fiber.Ctx) error {
	coll := common.GetDBCollection("icecreams")

	// Find all icecreams
	icecreams := make([]models.Icecream, 0)
	cursor, err := coll.Find(c.Context(), bson.M{})
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Iterate over the cursor
	for cursor.Next(c.Context()) {
		icecream := models.Icecream{}
		err := cursor.Decode(&icecream)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		icecreams = append(icecreams, icecream)
	}

	return c.Status(200).JSON(fiber.Map{"data": icecreams})
}

func getIcecream(c *fiber.Ctx) error {
	coll := common.GetDBCollection("icecreams")

	// Find the icecream  by id
	id := c.Params("id")
	if id == "" {
		return c.Status(400).JSON(fiber.Map{
			"error": "id is required :(",
		})
	}
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid id :(",
		})
	}

	icecream := models.Icecream{}

	err = coll.FindOne(c.Context(), bson.M{"_id": objectId}).Decode(&icecream)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(200).JSON(fiber.Map{"data": icecream})
}

// Create a new structure to help us create a new document
type createDTO struct {
	Flavor string `json:"flavor" bson:"flavor"`
	Stock  string `json:"stock" bson:"stock"`
}

func createIcecream(c *fiber.Ctx) error {
	// Validate the body
	b := new(createDTO)
	if err := c.BodyParser(b); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid body :(",
		})
	}

	// Create the new icecream
	coll := common.GetDBCollection("icecreams")
	result, err := coll.InsertOne(c.Context(), b)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error":   "Failed to create icecream :(",
			"message": err.Error(),
		})
	}

	// Return the icecream
	return c.Status(201).JSON(fiber.Map{
		"result": result,
	})
}

// Create a new structure to help us update a document
type updateDTO struct {
	Flavor string `json:"flavor,omitempty" bson:"flavor,omitempty"`
	Stock  string `json:"stock,omitempty" bson:"stock,omitempty"`
}

func updateIcecream(c *fiber.Ctx) error {
	// Validate the body
	b := new(updateDTO)
	if err := c.BodyParser(b); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid body :(",
		})
	}

	// Get the id of params
	id := c.Params("id")
	if id == "" {
		return c.Status(400).JSON(fiber.Map{
			"error": "id is required :(",
		})
	}
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid id :(",
		})
	}

	// Update the icecream
	coll := common.GetDBCollection("icecreams")
	result, err := coll.UpdateOne(c.Context(), bson.M{"_id": objectId}, bson.M{"$set": b})
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error":   "Failed to update icecream :(",
			"message": err.Error(),
		})
	}

	// Return the icecream
	return c.Status(200).JSON(fiber.Map{
		"result": result,
	})
}

func deleteIcecream(c *fiber.Ctx) error {
	// Get the id of params
	id := c.Params("id")
	if id == "" {
		return c.Status(400).JSON(fiber.Map{
			"error": "id is required :(",
		})
	}
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid id :(",
		})
	}

	// Delete the icecream
	coll := common.GetDBCollection("icecreams")
	result, err := coll.DeleteOne(c.Context(), bson.M{"_id": objectId})
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error":   "Failed to delete icecream :(",
			"message": err.Error(),
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"result": result,
	})
}
