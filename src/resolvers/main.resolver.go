package resolvers

import (
	"fmt"
	"io/ioutil"
	"os"

	logging "github.com/sirupsen/logrus"
)

var MainResolver *resolver = newMainResolver()

type resolver struct {
	RawSchemaString string
}

func newMainResolver() *resolver {

	var rawSchemaString string
	var wokingDirectory string
	var rawSchemaByte []byte
	var err error

	const schemaName string = "main.schema.graphql"

	wokingDirectory, err = os.Getwd()
	if err != nil {
		logging.Warn(fmt.Sprintf("error reading current working directory %+v", err.Error()))
		return &resolver{}
	}

	rawSchemaByte, err = ioutil.ReadFile(fmt.Sprintf("%s/schemas/%s", wokingDirectory, schemaName))
	if err != nil {
		logging.Warn(fmt.Sprintf("error reading schema file %+v", err.Error()))
		return &resolver{}
	}

	rawSchemaString = string(rawSchemaByte)
	return &resolver{
		RawSchemaString: rawSchemaString,
	}

}

func (r *resolver) Name() *string {

	var nameLiteral string = "Pedro"
	var namePointer *string = &nameLiteral

	return namePointer

}

func (r *resolver) SecondName() *string {

	var secondNameLiteral string = "Perez"
	var secondNamePointer *string = &secondNameLiteral

	return secondNamePointer

}
