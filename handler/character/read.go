package handler

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"

	"github.com/alifudin-a/go-echo-api/database"
	"github.com/alifudin-a/go-echo-api/models"
	"github.com/labstack/echo/v4"

	_ "github.com/lib/pq"
)

// ReadCharacter by id
func ReadCharacter(c echo.Context) (err error) {
	db := database.DBConn()
	defer db.Close()

	id, err := strconv.Atoi(c.Param("id"))

	var character models.Character

	row, err := db.Queryx("SELECT * FROM character WHERE id = $1", id)
	if err != nil {
		return sql.ErrNoRows
	}

	for row.Next() {
		err = row.StructScan(&character)
		if err != nil {
			log.Println(err)
		}
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"charcter": character})
}
