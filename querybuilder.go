package querybuilder

import (
	"net/url"
	"fmt"
	"github.com/moli9ma/querybuilder/internal/util"
	"reflect"
)

// QueryBuilder
type QueryBuilder interface {
	QueryString() string
}

// Build
func Build(data map[string]interface{}) url.Values {
	return build(data, NewParser())
}

// Build
func BuildWithCustomParser(data map[string]interface{}, parser Parser) url.Values {
	return build(data, parser)
}

// build
func build(data map[string]interface{}, parser Parser) url.Values {
	var urlValues = url.Values{}
	for key, value := range data {

		switch val := value.(type) {
		case string:
			urlValues.Add(key, parser.ParseString(val))
		case *string:
			if val != nil {
				urlValues.Add(key, parser.ParseString(*val))
			}

		case bool:
			urlValues.Add(key, parser.ParseBool(val))
		case *bool:
			if val != nil {
				urlValues.Add(key, parser.ParseBool(*val))
			}

		case int:
			urlValues.Add(key, parser.ParseInt(val))
		case *int:
			if val != nil {
				urlValues.Add(key, parser.ParseInt(*val))
			}
		case ***int:
			if val != nil {
				urlValues.Add(key, parser.ParseInt(***val))
			}

		case int8:
			urlValues.Add(key, parser.ParseInt8(val))
		case *int8:
			if val != nil {
				urlValues.Add(key, parser.ParseInt8(*val))
			}

		case int32:
			urlValues.Add(key, parser.ParseInt32(val))
		case *int32:
			if val != nil {
				urlValues.Add(key, parser.ParseInt32(*val))
			}

		case int64:
			urlValues.Add(key, parser.ParseInt64(val))
		case *int64:
			if val != nil {
				urlValues.Add(key, parser.ParseInt64(*val))
			}

		case uint:
			urlValues.Add(key, parser.ParseUint(val))
		case *uint:
			if val != nil {
				urlValues.Add(key, parser.ParseUint(*val))
			}

		case uint8:
			urlValues.Add(key, parser.ParseUint8(val))
		case *uint8:
			if val != nil {
				urlValues.Add(key, parser.ParseUint8(*val))
			}

		case uint32:
			urlValues.Add(key, parser.ParseUint32(val))
		case *uint32:
			if val != nil {
				urlValues.Add(key, parser.ParseUint32(*val))
			}

		case uint64:
			urlValues.Add(key, parser.ParseUint64(val))
		case *uint64:
			if val != nil {
				urlValues.Add(key, parser.ParseUint64(*val))
			}

		case float32:
			urlValues.Add(key, parser.ParseFloat32(val))
		case *float32:
			if val != nil {
				urlValues.Add(key, parser.ParseFloat32(*val))
			}

		case float64:
			urlValues.Add(key, parser.ParseFloat64(val))
		case *float64:
			if val != nil {
				urlValues.Add(key, parser.ParseFloat64(*val))
			}

		case []string:
			urlValues.Add(key, parser.ParseStrings(val))
		case *[]string:
			if val != nil {
				urlValues.Add(key, parser.ParseStrings(*val))
			}

		case []bool:
			urlValues.Add(key, parser.ParseBools(val))
		case *[]bool:
			if val != nil {
				urlValues.Add(key, parser.ParseBools(*val))
			}

		case []int:
			urlValues.Add(key, parser.ParseInts(val))
		case []*int:
			urlValues.Add(key, parser.ParseInts(util.CastInt(val)))
		case *[]int:
			if val != nil {
				urlValues.Add(key, parser.ParseInts(*val))
			}

		case []int8:
			urlValues.Add(key, parser.ParseInt8s(val))
		case []*int8:
			urlValues.Add(key, parser.ParseInt8s(util.CastInt8(val)))
		case *[]int8:
			if val != nil {
				urlValues.Add(key, parser.ParseInt8s(*val))
			}

		case []int32:
			urlValues.Add(key, parser.ParseInt32s(val))
		case []*int32:
			urlValues.Add(key, parser.ParseInt32s(util.CastInt32(val)))
		case *[]int32:
			if val != nil {
				urlValues.Add(key, parser.ParseInt32s(*val))
			}

		case []int64:
			urlValues.Add(key, parser.ParseInt64s(val))
		case []*int64:
			urlValues.Add(key, parser.ParseInt64s(util.CastInt64(val)))
		case *[]int64:
			if val != nil {
				urlValues.Add(key, parser.ParseInt64s(*val))
			}

		case []uint:
			urlValues.Add(key, parser.ParseUints(val))
		case []*uint:
			urlValues.Add(key, parser.ParseUints(util.CastUint(val)))
		case *[]uint:
			if val != nil {
				urlValues.Add(key, parser.ParseUints(*val))
			}

		case []uint8:
			urlValues.Add(key, parser.ParseUint8s(val))
		case []*uint8:
			urlValues.Add(key, parser.ParseUint8s(util.CastUint8(val)))
		case *[]uint8:
			if val != nil {
				urlValues.Add(key, parser.ParseUint8s(*val))
			}

		case []uint32:
			urlValues.Add(key, parser.ParseUint32s(val))
		case []*uint32:
			urlValues.Add(key, parser.ParseUint32s(util.CastUint32(val)))
		case *[]uint32:
			if val != nil {
				urlValues.Add(key, parser.ParseUint32s(*val))
			}

		case []uint64:
			urlValues.Add(key, parser.ParseUint64s(val))
		case []*uint64:
			urlValues.Add(key, parser.ParseUint64s(util.CastUint64(val)))
		case *[]uint64:
			if val != nil {
				urlValues.Add(key, parser.ParseUint64s(*val))
			}

		case []float32:
			urlValues.Add(key, parser.ParseFloat32s(val))
		case []*float32:
			urlValues.Add(key, parser.ParseFloat32s(util.CastFloat32(val)))
		case *[]float32:
			if val != nil {
				urlValues.Add(key, parser.ParseFloat32s(*val))
			}
		case *[]*float32:
			if val != nil {
				urlValues.Add(key, parser.ParseFloat32s(util.CastFloat32(*val)))
			}

		case []float64:
			urlValues.Add(key, parser.ParseFloat64s(val))
		case *[]float64:
			if val != nil {
				urlValues.Add(key, parser.ParseFloat64s(*val))
			}
		case []*float64:
			urlValues.Add(key, parser.ParseFloat64s(util.CastFloat64(val)))
		case *[]*float64:
			if val != nil {
				urlValues.Add(key, parser.ParseFloat64s(util.CastFloat64(*val)))
			}

		case QueryBuilder:
			urlValues.Add(key, val.QueryString())

		case fmt.Stringer:
			urlValues.Add(key, val.String())

		case interface{}:

			if reflect.ValueOf(val).Kind() != reflect.Ptr {
				urlValues.Add(key, fmt.Sprintf("%v", val))
			}

			// unsupported

		default:
			// unsupported

		}
	}
	return urlValues
}
