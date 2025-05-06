package student

import (
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"net/http"
	"github.com/atindraraut/crudgo/internal/types"
	"github.com/atindraraut/crudgo/internal/utils/response"
	"github.com/go-playground/validator/v10"
)

func New() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var student types.Student
		err:=json.NewDecoder(r.Body).Decode(&student)
		if errors.Is(err, io.EOF) {
			response.WriteJSON(w, http.StatusBadRequest, response.GeneralError(errors.New("request body is empty")))
			return
		}
		if err != nil {
			response.WriteJSON(w, http.StatusBadRequest, response.GeneralError(err))
			return
		}
		//request validation
		if err := validator.New().Struct(&student); err != nil {
			response.WriteJSON(w, http.StatusBadRequest, response.GeneralError(err))
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