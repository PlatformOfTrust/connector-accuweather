package main

import (
	"fmt"
	"net/http"

	"github.com/alecthomas/jsonschema"
	"github.com/go-chi/chi"
	"github.com/go-playground/validator/v10"
)

var schemas = make(map[string][]byte)

func init() {
	jsonschema.Version = "http://json-schema.org/draft-06/schema#"

	//schema, _ := jsonschema.Reflect(&models.RequestParameters{}).MarshalJSON()
	schema := []byte(requestSchema)
	schemas["/schemas/params"] = schema

	//schema, _ = jsonschema.Reflect(&models.ForecastResponse{}).MarshalJSON()
	schema = []byte(responseSchema)
	schemas["/schemas/response"] = schema
}

func handleGetSchema(w http.ResponseWriter, r *http.Request) {
	s := chi.URLParam(r, "schemaName")

	var schema []byte

	switch s {
	case "params":
		schema = schemas["/schemas/params"]
		break
	case "response":
		schema = schemas["/schemas/response"]
		break
	}

	w.Write(schema)
}

func handleGetSchemas(w http.ResponseWriter, r *http.Request) {
	prefix := "http://" + r.Host
	res := []string{}

	for k, _ := range schemas {
		res = append(res, prefix+k)
	}

	jdata, _ := json.Marshal(res)

	w.Write(jdata)
}

type ValidationError struct {
	ValidationErrors []validator.FieldError
}

func (e ValidationError) Error() string {
	var errs string
	for _, er := range e.ValidationErrors {
		errs = fmt.Sprintf(`%s%s`, errs, er.Error())
	}
	return errs
}

var validate = validator.New()

func validateParams(data interface{}) error {
	err := validate.Struct(data)
	if err != nil {
		verr := err.(validator.ValidationErrors)
		return ValidationError{
			ValidationErrors: verr,
		}
	}

	return err
}
