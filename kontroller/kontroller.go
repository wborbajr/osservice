package kontroller

import (
	"database/sql"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/wborbajr/osservice/database"
	"github.com/wborbajr/osservice/model"
)

//
// https://goplay.space/#mWw59cjYPh7
// https://play.golang.org/p/mWw59cjYPh7
// https://stackoverflow.com/questions/27795036/create-chan-for-func-with-two-return-args#27795117

// GetAllOS - retrieve customer order service
func GetAllOS(c *fiber.Ctx) error {

	// Parsing parameters
	paramDoc := c.Params("doc")
	paramOs  := c.Params("os")

	serviceOrder := model.ServiceOrder{}

	if err := database.KonnektAra(); err != nil {
		log.Println(err)
	}

	row, err := database.DB.Query("SELECT * FROM TB_OS WHERE ID_OS=$1 AND ID_CLIENTE=$2", paramOs, paramDoc)

	if err != nil {
		return c.Status(500).JSON(&fiber.Map{
			"success": false,
			"message": err,
		})

	}

	defer row.Close()

	for row.Next() {
		switch err := row.Scan(&serviceOrder.IdOs, &serviceOrder.IdCliente, &serviceOrder.IdStatus); err {
		case sql.ErrNoRows:
			log.Println("No rows were returned!")
			return c.Status(500).JSON(&fiber.Map{
				"success": false,
				"message": err,
			})
		case nil:
			log.Println(serviceOrder.IdOs, serviceOrder.IdCliente, serviceOrder.IdStatus)
		default:
			return c.Status(500).JSON(&fiber.Map{
				"success": false,
				"message": err,
			})
		}
	}

	return nil
}
