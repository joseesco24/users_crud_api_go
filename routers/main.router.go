package routers

import (
	"fmt"

	fgg "github.com/cckwes/fiber-graphql-go"
	graphql "github.com/graph-gophers/graphql-go"
	resolvers "github.com/joseesco24/users_crud_api_go/resolvers"
	logging "github.com/sirupsen/logrus"
)

var MainRouter *graphQlRouter = newMainRouter()

type graphQlRouter struct {
	Handler fgg.Handler
}

func newMainRouter() *graphQlRouter {

	var schema *graphql.Schema
	var handler fgg.Handler
	var rawSchema string
	var err error

	rawSchema = resolvers.MainResolver.RawSchemaString
	if err != nil {
		logging.Warn(fmt.Sprintf("error reading schema file %+v", err.Error()))
		return &graphQlRouter{}
	}

	schema = graphql.MustParseSchema(rawSchema, resolvers.MainResolver)
	handler = fgg.Handler{Schema: schema}

	return &graphQlRouter{
		Handler: handler,
	}

}
