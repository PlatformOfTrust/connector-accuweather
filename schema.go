package main

import (
	"fmt"
	"net/http"

	"github.com/alecthomas/jsonschema"
	"github.com/go-chi/chi"
	"github.com/goreleaser/goreleaser/pkg/context"
	jsc "github.com/qri-io/jsonschema"
)

var schemas = make(map[string][]byte)

func init() {
	jsonschema.Version = "http://json-schema.org/draft-07/schema#"

	// schema, _ := jsonschema.Reflect(&models.RequestParameters{}).MarshalJSON()
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
	ValidationErrors []jsc.KeyError
}

func (e ValidationError) Error() string {
	var errs string
	for _, er := range e.ValidationErrors {
		errs = fmt.Sprintf(`%s%s with value: %v, %s; `, errs, er.PropertyPath, er.InvalidValue, er.Message)
	}
	return errs
}

func validateSchema(schemaName string, data []byte) error {
	schema := schemas[schemaName]

	ctx := context.Context{}

	rs := &jsc.Schema{}
	err := json.Unmarshal(schema, rs)
	if err != nil {
		return err
	}

	verr, err := rs.ValidateBytes(ctx, data)
	if err != nil {
		return ValidationError{
			ValidationErrors: verr,
		}
	}

	return nil
}
