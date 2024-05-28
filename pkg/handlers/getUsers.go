package handlers

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/anes011/gop/pkg/controllers"
	"github.com/anes011/gop/pkg/json"
	"github.com/gorilla/mux"
)

func HandleGetUsers(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)

		offsetInt, err := strconv.Atoi(params["offset"])
		limitInt, err := strconv.Atoi(params["limit"])

		pagination := &controllers.Pagination{
			Offset: offsetInt,
			Limit:  limitInt,
		}

		users, err := controllers.GetUsers(db, pagination)

		if err != nil {
			json.WriteJSON(w, 500, "couldn't get users!")
		}

		json.WriteJSON(w, 200, users)
	}
}
