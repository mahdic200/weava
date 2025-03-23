package Http

import (
	"strings"

	"github.com/gofiber/fiber/v2"
)

type QueryArg struct {
	Key   string
	Value string
}

func queryString(c *fiber.Ctx) *string {
	uri := c.Request().URI().String()
	sliced_uri := strings.Split(uri, "?")
	if len(sliced_uri) == 1 {
		return nil
	}
	return &sliced_uri[1]
}

func QueryParser(c *fiber.Ctx) []QueryArg {
	query := ""
	if q := queryString(c); q != nil {
		query = *q
	}
	raw_partial := strings.Split(query, "&&")
	result := []QueryArg{}
	for _, item := range raw_partial {
		if sliced_arg := strings.Split(item, "="); len(sliced_arg) == 1 {
			result = append(result, QueryArg{Key: sliced_arg[0], Value: ""})
		} else {
			result = append(result, QueryArg{Key: sliced_arg[0], Value: sliced_arg[1]})
		}
	}
	return result
}

func GetQueryArg(c *fiber.Ctx, key string) *string {
	for _, query_arg := range QueryParser(c) {
		if query_arg.Key == key {
			return &query_arg.Value
		}
	}
	return nil
}
