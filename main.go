package proper

import (
	"reflect"
	"runtime"
	"strings"
	"testing"

	"github.com/brianvoe/gofakeit"
)

func TestProperties(t *testing.T, propertyFunctions ...any) {
	for _, property := range propertyFunctions {
		if !isValidProperty(property) {
			continue
		}

		t.Log("Testing", getFunctionName(property))
		for i := 0; i < 100; i++ {
			arguments := generateArguments(property)
			success := reflect.ValueOf(property).Call(arguments)[0]
			if success.IsZero() {
				t.Log(getFunctionName(property), "failed with parameters:")
				for _, argument := range arguments {
					t.Log(argument.Interface())
				}
				t.Fail()
			}
		}
	}
}

func isValidProperty(property any) bool {
	propType := reflect.TypeOf(property)
	return propType.Kind() == reflect.Func &&
		propType.NumOut() == 1 &&
		propType.Out(0).Kind() == reflect.Bool
}

// NOTE: Assumes a valid property
func generateArguments(property any) (arguments []reflect.Value) {
	propType := reflect.TypeOf(property)
	for i := 0; i < propType.NumIn(); i++ {
		value := reflect.New(propType.In(i))
		gofakeit.Struct(value.Interface())

		arguments = append(arguments, value.Elem())
	}
	return arguments
}

func getFunctionName(temp any) string {
    strs := strings.Split((runtime.FuncForPC(reflect.ValueOf(temp).Pointer()).Name()), ".")
    return strs[len(strs)-1]
}
