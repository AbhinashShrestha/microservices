package handlers

import  ( 
	"github.com/AbhinashShrestha/rest_api/models"
	"github.com/AbhinashShrestha/rest_api/database"
	"github.com/gofiber/fiber/v2"
)
func Home(c *fiber.Ctx) error {
	return c.SendString("Jumping in the wild <:3 !")
}

func CreateFact(c *fiber.Ctx) error{
	fact := new(models.Fact)

	if err := c.BodyParser(fact); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message" : err.Error(),
		})
	}

	database.DB.Db.Create(&fact)

	return c.Status(200).JSON(fact)
}