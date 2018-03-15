package querybuilder

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"strings"
	"strconv"
)

type MyFloats []float64
func (floats MyFloats) QueryString() string {
	return "works QueryString"
}

func (floats MyFloats) String() string {
	strs := []string{}
	for _,v := range floats {
		strs = append(strs, strconv.FormatFloat(v, 'f', -1 , 64))
	}
	return strings.Join(strs, ",")
}

func TestBuildFromMap3(t *testing.T) {

	// int
	func () {
		testdata := map[string]interface{}{
			"x" : int(10),
		}
		expected := "x=10"
		actual := Build(testdata, New()).Encode()
		assert.Equal(t, expected, actual)
	}()

	// int pointer
	func () {
		i := int(10)
		testdata := map[string]interface{}{
			"x" : &i,
		}
		expected := "x=10"
		actual := Build(testdata, New()).Encode()
		assert.Equal(t, expected, actual)
	}()

	// int nil pointer
	func () {
		var nilInt *int
		testdata := map[string]interface{}{
			"x" : nilInt,
		}
		expected := ""
		actual := Build(testdata, New()).Encode()
		assert.Equal(t, expected, actual)
	}()

	// int8
	func () {
		testdata := map[string]interface{}{
			"x" : int8(10),
		}
		expected := "x=10"
		actual := Build(testdata, New()).Encode()
		assert.Equal(t, expected, actual)
	}()

	// int8 pointer
	func () {
		i := int8(10)
		testdata := map[string]interface{}{
			"x" : &i,
		}
		expected := "x=10"
		actual := Build(testdata, New()).Encode()
		assert.Equal(t, expected, actual)
	}()

	// int8 pointer pointer
	func () {
		var ptrPtr **int
		i := int(10)
		ptr := &i
		ptrPtr = &ptr

		testdata := map[string]interface{}{
			"x" : &ptrPtr,
		}
		expected := "x=10"
		actual := Build(testdata, New()).Encode()
		assert.Equal(t, expected, actual)
	}()

	// int8 nil pointer
	func () {
		var nilInt8 *int8
		testdata := map[string]interface{}{
			"x" : nilInt8,
		}
		expected := ""
		actual := Build(testdata, New()).Encode()
		assert.Equal(t, expected, actual)
	}()

	func () {
		f1 := float32(10)
		f2 := float32(20)
		float32s := []*float32 {
			&f1, &f2,
		}
		testdata := map[string]interface{}{
			"x" : &float32s,
		}
		expected := "x=10%2C20"
		actual := Build(testdata, New()).Encode()
		assert.Equal(t, expected, actual)
	}()


	// *[]*float64
	func () {
		f1 := float64(10)
		f2 := float64(20)
		float64s := []*float64 {
			&f1, &f2,
		}
		testdata := map[string]interface{}{
			"x" : &float64s,
		}
		expected := "x=10%2C20"
		actual := Build(testdata, New()).Encode()
		assert.Equal(t, expected, actual)
	}()

	// type (interface)
	func () {
		type MyInt int
		typedInt := MyInt(10)
		testdata := map[string]interface{}{
			"x" : typedInt,
		}
		expected := "x=10"
		actual := Build(testdata, New()).Encode()
		assert.Equal(t, expected, actual)
	}()

	// Type []float64
	func () {
		myFloats := MyFloats{
				float64(10), float64(20),
		}
		testdata := map[string]interface{}{
			"x" : myFloats,
		}
		expected := "x=works+QueryString"
		actual := Build(testdata, New()).Encode()
		assert.Equal(t, expected, actual)
	}()

	// type pointer
	func () {

		type MyInt int
		i := int(10)
		typedInt := MyInt(i)
		testdata := map[string]interface{}{
			"x" : &typedInt,
		}
		expected := ""
		actual := Build(testdata, New()).Encode()
		assert.Equal(t, expected, actual)
	}()

	// type nil pointer
	func () {
		type MyInt int
		var myInt *MyInt
		testdata := map[string]interface{}{
			"x" : myInt,
		}
		expected := ""
		actual := Build(testdata, New()).Encode()
		assert.Equal(t, expected, actual)
	}()

	// QueryBuilder.QueryString()
	func () {
		myFloats := &MyFloats{
			float64(10), float64(20),
		}
		testdata := map[string]interface{}{
			"x" : myFloats,
		}
		expected := "x=works+QueryString"
		actual := Build(testdata, New()).Encode()
		assert.Equal(t, expected, actual)
	}()

}
