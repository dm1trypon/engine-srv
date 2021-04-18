package jsonvalidator

import (
	"bytes"
	"fmt"

	logger "github.com/dm1trypon/easy-logger"
	"github.com/santhosh-tekuri/jsonschema"
)

// LC - logging category
const LC = "JSON_VALIDATOR"

func Check(jsonData []byte, schemaPath string) bool {
	schema, err := jsonschema.Compile(schemaPath)
	if err != nil {
		logger.Error(LC, fmt.Sprint("JSON schema loader error: ", err.Error()))
		return false
	}

	jsonReader := bytes.NewReader(jsonData)

	if err = schema.Validate(jsonReader); err != nil {
		logger.Error(LC, fmt.Sprint("JSON source is not valid: ", err.Error()))
		return false
	}

	return true
}
