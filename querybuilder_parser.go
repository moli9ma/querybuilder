package querybuilder

import (
	"strconv"
	"strings"
)

type Parser struct {
	// build in types
	ParseString func(value string) string

	ParseBool func(value bool) string

	ParseInt func(value int) string

	ParseInt8 func(value int8) string

	ParseInt32 func(value int32) string

	ParseInt64 func(value int64) string

	ParseUint func(value uint) string

	ParseUint8 func(value uint8) string

	ParseUint32 func(value uint32) string

	ParseUint64 func(value uint64) string

	ParseFloat32 func(value float32) string

	ParseFloat64 func(value float64) string

	// array of build in types
	ParseStrings func(values []string) string

	ParseBools func(values []bool) string

	ParseInts func(values []int) string

	ParseInt8s func(values []int8) string

	ParseInt32s func(values []int32) string

	ParseInt64s func(values []int64) string

	ParseUints func(values []uint) string

	ParseUint8s func(values []uint8) string

	ParseUint32s func(values []uint32) string

	ParseUint64s func(values []uint64) string

	ParseFloat32s func(values []float32) string

	ParseFloat64s func(values []float64) string
}

func (parser *Parser) SetIntParsers(intParser func(value int64) string) {
	parser.ParseInt64 = intParser
	parser.ParseInt = func(value int) string {
		return parser.ParseInt64(int64(value))
	}

	parser.ParseInt8 = func(value int8) string {
		return parser.ParseInt64(int64(value))
	}

	parser.ParseInt32 = func(value int32) string {
		return parser.ParseInt64(int64(value))
	}
}

func (parser *Parser) SetUintParsers(uintParser func(value uint64) string) {
	parser.ParseUint64 = uintParser
	parser.ParseUint = func(value uint) string {
		return parser.ParseUint64(uint64(value))
	}

	parser.ParseUint8 = func(value uint8) string {
		return parser.ParseUint64(uint64(value))
	}

	parser.ParseUint32 = func(value uint32) string {
		return parser.ParseUint64(uint64(value))
	}
}

func (parser *Parser) SetFloatParsers(floatParser func(value float64) string) {
	parser.ParseFloat64 = floatParser
	parser.ParseFloat32 = func(value float32) string {
		return parser.ParseFloat64(float64(value))
	}
}

func (parser *Parser) SetSliceTypeParsers(parseStrings func(values []string) string) {
	var parseBools = func(values []bool) string {
		var stringValues []string
		for _, value := range values {
			stringValues = append(stringValues, parser.ParseBool(value))
		}
		return parseStrings(stringValues)
	}

	var parseInts = func(values []int) string {
		var stringValues []string
		for _, value := range values {
			stringValues = append(stringValues, parser.ParseInt(value))
		}
		return parseStrings(stringValues)
	}

	var parseInt8s = func(values []int8) string {
		var stringValues []string
		for _, value := range values {
			stringValues = append(stringValues, parser.ParseInt8(value))
		}
		return parseStrings(stringValues)
	}

	var parseInt32s = func(values []int32) string {
		var stringValues []string
		for _, value := range values {
			stringValues = append(stringValues, parser.ParseInt32(value))
		}
		return parseStrings(stringValues)
	}

	var parseInt64s = func(values []int64) string {
		var stringValues []string
		for _, value := range values {
			stringValues = append(stringValues, parser.ParseInt64(value))
		}
		return parseStrings(stringValues)
	}

	var parseUints = func(values []uint) string {
		var stringValues []string
		for _, value := range values {
			stringValues = append(stringValues, parser.ParseUint(value))
		}
		return parseStrings(stringValues)
	}

	var parseUint8s = func(values []uint8) string {
		var stringValues []string
		for _, value := range values {
			stringValues = append(stringValues, parser.ParseUint8(value))
		}
		return parseStrings(stringValues)
	}

	var parseUint32s = func(values []uint32) string {
		var stringValues []string
		for _, value := range values {
			stringValues = append(stringValues, parser.ParseUint32(value))
		}
		return parseStrings(stringValues)
	}

	var parseUint64s = func(values []uint64) string {
		var stringValues []string
		for _, value := range values {
			stringValues = append(stringValues, parser.ParseUint64(value))
		}
		return parseStrings(stringValues)
	}

	var parseFloat32s = func(values []float32) string {
		var stringValues []string
		for _, value := range values {
			stringValues = append(stringValues, parser.ParseFloat32(value))
		}
		return parseStrings(stringValues)
	}

	var parseFloat64s = func(values []float64) string {
		var stringValues []string
		for _, value := range values {
			stringValues = append(stringValues, parser.ParseFloat64(value))
		}
		return parseStrings(stringValues)
	}

	parser.ParseStrings = parseStrings
	parser.ParseBools = parseBools
	parser.ParseInts = parseInts
	parser.ParseInt8s = parseInt8s
	parser.ParseInt32s = parseInt32s
	parser.ParseInt64s = parseInt64s

	parser.ParseUints = parseUints
	parser.ParseUint8s = parseUint8s
	parser.ParseUint32s = parseUint32s
	parser.ParseUint64s = parseUint64s

	parser.ParseFloat32s = parseFloat32s
	parser.ParseFloat64s = parseFloat64s
}

// New default customTypeParser
func New() Parser {

	var parseString = func(value string) string {
		return value
	}

	var parseBool = func(value bool) string {
		return strconv.FormatBool(value)
	}

	var parseInt64 = func(value int64) string {
		return strconv.FormatInt(value, 10)
	}

	var parseInt = func(value int) string {
		return parseInt64(int64(value))
	}

	var parseInt8 = func(value int8) string {
		return parseInt64(int64(value))
	}

	var parseInt32 = func(value int32) string {
		return parseInt64(int64(value))
	}

	var parseUint64 = func(value uint64) string {
		return strconv.FormatUint(value, 10)
	}

	var parseUint = func(value uint) string {
		return parseUint64(uint64(value))
	}

	var parseUint8 = func(value uint8) string {
		return parseUint64(uint64(value))
	}

	var parseUint32 = func(value uint32) string {
		return parseUint64(uint64(value))
	}

	var parseFloat64 = func(value float64) string {
		return strconv.FormatFloat(value, 'f', -1, 64)
	}

	var parseFloat32 = func(value float32) string {
		return parseFloat64(float64(value))
	}

	var parseStrings = func(values []string) string {
		return strings.Join(values, ",")
	}

	var parseBools = func(values []bool) string {
		var stringValues []string
		for _, value := range values {
			stringValues = append(stringValues, strconv.FormatBool(value))
		}
		return parseStrings(stringValues)
	}

	var parseInts = func(values []int) string {
		var stringValues []string
		for _, value := range values {
			stringValues = append(stringValues, parseInt(value))
		}
		return parseStrings(stringValues)
	}

	var parseInt8s = func(values []int8) string {
		var stringValues []string
		for _, value := range values {
			stringValues = append(stringValues, parseInt8(value))
		}
		return parseStrings(stringValues)
	}

	var parseInt32s = func(values []int32) string {
		var stringValues []string
		for _, value := range values {
			stringValues = append(stringValues, parseInt32(value))
		}
		return parseStrings(stringValues)
	}

	var parseInt64s = func(values []int64) string {
		var stringValues []string
		for _, value := range values {
			stringValues = append(stringValues, parseInt64(value))
		}
		return parseStrings(stringValues)
	}

	var parseUints = func(values []uint) string {
		var stringValues []string
		for _, value := range values {
			stringValues = append(stringValues, parseUint(value))
		}
		return parseStrings(stringValues)
	}

	var parseUint8s = func(values []uint8) string {
		var stringValues []string
		for _, value := range values {
			stringValues = append(stringValues, parseUint8(value))
		}
		return parseStrings(stringValues)
	}

	var parseUint32s = func(values []uint32) string {
		var stringValues []string
		for _, value := range values {
			stringValues = append(stringValues, parseUint32(value))
		}
		return parseStrings(stringValues)
	}

	var parseUint64s = func(values []uint64) string {
		var stringValues []string
		for _, value := range values {
			stringValues = append(stringValues, parseUint64(value))
		}
		return parseStrings(stringValues)
	}

	var parseFloat32s = func(values []float32) string {
		var stringValues []string
		for _, value := range values {
			stringValues = append(stringValues, parseFloat32(value))
		}
		return parseStrings(stringValues)
	}

	var parseFloat64s = func(values []float64) string {
		var stringValues []string
		for _, value := range values {
			stringValues = append(stringValues, parseFloat64(value))
		}
		return parseStrings(stringValues)
	}

	parser := Parser{
		ParseString:  parseString,
		ParseBool:    parseBool,
		ParseInt:     parseInt,
		ParseInt8:    parseInt8,
		ParseInt32:   parseInt32,
		ParseInt64:   parseInt64,
		ParseUint:    parseUint,
		ParseUint8:   parseUint8,
		ParseUint32:  parseUint32,
		ParseUint64:  parseUint64,
		ParseFloat32: parseFloat32,
		ParseFloat64: parseFloat64,

		ParseStrings:  parseStrings,
		ParseBools:    parseBools,
		ParseInts:     parseInts,
		ParseInt8s:    parseInt8s,
		ParseInt32s:   parseInt32s,
		ParseInt64s:   parseInt64s,
		ParseUints:    parseUints,
		ParseUint8s:   parseUint8s,
		ParseUint32s:  parseUint32s,
		ParseUint64s:  parseUint64s,
		ParseFloat32s: parseFloat32s,
		ParseFloat64s: parseFloat64s,
	}

	return parser
}