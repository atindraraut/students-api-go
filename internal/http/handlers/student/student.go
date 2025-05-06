package student

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
	"io"
	"github.com/atindraraut/crudgo/internal/types"
	"github.com/atindraraut/crudgo/internal/utils/response"
)

func New() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var student types.Student
		err:=json.NewDecoder(r.Body).Decode(&student)
		if errors.Is(err, io.EOF) {
			response.WriteJSON(w, http.StatusBadRequest, response.GenerelError(errors.New("request body is empty")))
			return
		}
		slog.Info("Creating a student.",slog.String("method", r.Method), "Body:",slog.AnyValue(student))
		responseData := map[string]interface{}{
			"Message": "Student created successfully",
			"Student": student,
		}
		response.WriteJSON(w, http.StatusCreated, responseData)
	}
}