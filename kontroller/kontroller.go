package kontroller

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/wborbajr/osservice/database"
	"github.com/wborbajr/osservice/model"
)

//
// https://goplay.space/#mWw59cjYPh7
// https://play.golang.org/p/mWw59cjYPh7
// https://stackoverflow.com/questions/27795036/create-chan-for-func-with-two-return-args#27795117

var (
	errAra error
	errCwb error
	errLon error
	errNat error
	errRec error
)

// GetAllOS - retrieve customer order service
func GetAllOS(c *fiber.Ctx) error {
	count := 0

	// Parsing parameters
	paramDoc := c.Params("doc")
	paramOs := c.Params("os")

	serviceOrder := model.ServiceOrder{}

	if err := database.KonnektAra(); err == nil {
		errAra = database.DB.QueryRow("SELECT * FROM TB_OS WHERE ID_OS=? AND ID_CLIENTE=?", paramOs, paramDoc).Scan(&serviceOrder.IdOs, &serviceOrder.IdCliente, &serviceOrder.IdStatus)
	}
	if err := database.KonnektCwb(); err == nil {
		errCwb = database.DB.QueryRow("SELECT * FROM TB_OS WHERE ID_OS=? AND ID_CLIENTE=?", paramOs, paramDoc).Scan(&serviceOrder.IdOs, &serviceOrder.IdCliente, &serviceOrder.IdStatus)
	}
	if err := database.KonnektLon(); err == nil {
		errLon = database.DB.QueryRow("SELECT * FROM TB_OS WHERE ID_OS=? AND ID_CLIENTE=?", paramOs, paramDoc).Scan(&serviceOrder.IdOs, &serviceOrder.IdCliente, &serviceOrder.IdStatus)
	}
	if err := database.KonnektNat(); err == nil {
		errNat = database.DB.QueryRow("SELECT * FROM TB_OS WHERE ID_OS=? AND ID_CLIENTE=?", paramOs, paramDoc).Scan(&serviceOrder.IdOs, &serviceOrder.IdCliente, &serviceOrder.IdStatus)
	}
	if err := database.KonnektRec(); err == nil {
		errRec = database.DB.QueryRow("SELECT * FROM TB_OS WHERE ID_OS=? AND ID_CLIENTE=?", paramOs, paramDoc).Scan(&serviceOrder.IdOs, &serviceOrder.IdCliente, &serviceOrder.IdStatus)
	}

	if errAra == nil {
		count++
		log.Println("ARA Found")
		log.Println(serviceOrder.IdOs, serviceOrder.IdCliente, serviceOrder.IdStatus)
	} else {
		log.Println("ARA NOT Found")
	}
	if errCwb == nil {
		count++
		log.Println("CWB Found")
		log.Println(serviceOrder.IdOs, serviceOrder.IdCliente, serviceOrder.IdStatus)
	} else {
		log.Println("CWB NOT Found")
	}
	if errLon == nil {
		count++
		log.Println("LON Found")
		log.Println(serviceOrder.IdOs, serviceOrder.IdCliente, serviceOrder.IdStatus)
	} else {
		log.Println("LON NOT Found")
	}
	if errNat == nil {
		count++
		log.Println("NAT Found")
		log.Println(serviceOrder.IdOs, serviceOrder.IdCliente, serviceOrder.IdStatus)
	} else {
		log.Println("NAT NOT Found")
	}
	if errRec == nil {
		count++
		log.Println("REC Found")
		log.Println(serviceOrder.IdOs, serviceOrder.IdCliente, serviceOrder.IdStatus)
	} else {
		log.Println("REC NOT Found")
	}

	if count != 0 {
		if err := c.JSON(&fiber.Map{
			"success":   false,
			"message":   "Successfully fetched OSService",
			"osservice": serviceOrder,
		}); err != nil {
			c.Status(500).JSON(&fiber.Map{
				"success": false,
				"message": err,
			})
			return nil
		}
	} else {
		c.Status(500).JSON(&fiber.Map{
			"success": false,
			"message": "Service Order not found!",
		})
	}

	// if errAra != sql.ErrNoRows {

	// }

	// row, err := database.DB.Query("SELECT * FROM TB_OS WHERE ID_OS=$1 AND ID_CLIENTE=$2", paramOs, paramDoc)

	// if err != nil {
	// 	return c.Status(500).JSON(&fiber.Map{
	// 		"success": false,
	// 		"message": err,
	// 	})

	// }

	// defer row.Close()

	// for row.Next() {
	// 	switch err := row.Scan(&serviceOrder.IdOs, &serviceOrder.IdCliente, &serviceOrder.IdStatus); err {
	// 	case sql.ErrNoRows:
	// 		log.Println("No rows were returned!")
	// 		return c.Status(500).JSON(&fiber.Map{
	// 			"success": false,
	// 			"message": err,
	// 		})
	// 	case nil:
	// 		log.Println(serviceOrder.IdOs, serviceOrder.IdCliente, serviceOrder.IdStatus)
	// 	default:
	// 		return c.Status(500).JSON(&fiber.Map{
	// 			"success": false,
	// 			"message": err,
	// 		})
	// 	}
	// }

	return nil
}
