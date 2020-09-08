package handler

import (
	"net/http"
	"strconv"

	"github.com/alifudin-a/go-echo-api/database"
	"github.com/alifudin-a/go-echo-api/models"
	"github.com/labstack/echo/v4"

	_ "github.com/lib/pq"
)

// List all data from Character
func List(c echo.Context) (err error) {
	db := database.DBConn()
	defer db.Close()

	limit, _ := strconv.Atoi(c.QueryParam("limit"))
	if limit == 0 {
		limit = 10
	}

	var list []*models.Character

	row, err := db.Queryx("SELECT * FROM character ORDER BY id ASC LIMIT $1", limit)
	if err != nil {
		return err
	}

	for row.Next() {
		var character models.Character

		err = row.StructScan(&character)
		if err != nil {
			return err
		}

		list = append(list, &character)
	}

	return c.JSON(http.StatusOK, list)
}
