package get_query_operator

import "strings"

func GetQueryOperator(op string) (string, error) {
	op = strings.Replace(op, "$", "", -1)
	op = strings.Replace(op, " ", "", -1)

	switch op {
	case "eq":
		return "=", nil
	case "ne":
		return "!=", nil
	case "gt":
		return ">", nil
	case "gte":
		return ">=", nil
	case "lt":
		return "<", nil
	case "lte":
		return "<=", nil
	case "in":
		return "IN", nil
	case "nin":
		return "NOT IN", nil
	case "any":
		return "ANY", nil
	case "some":
		return "SOME", nil
	case "all":
		return "ALL", nil
	case "notnull":
		return "IS NOT NULL", nil
	case "null":
		return "IS NULL", nil
	case "true":
		return "IS TRUE", nil
	case "nottrue":
		return "IS NOT TRUE", nil
	case "false":
		return "IS FALSE", nil
	case "notfalse":
		return "IS NOT FALSE", nil
	case "like":
		return "LIKE", nil
	case "ilike":
		return "ILIKE", nil
	}

	err := errors.New("invalid operator")
	return "", err

}
