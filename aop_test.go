package aop

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

var callsCounter = 0

func cleanWrapperCountAfter(fptr any) WrappedFunc {
	return func(args []reflect.Value) []reflect.Value {
		results := Proceed(fptr, args) // вызываем исходную функцию
		callsCounter++                 // увеличиваем счетчик вызовов функции
		return results
	}
}

// just to make sure it is working
func TestSimpleWrapper(t *testing.T) {
	wrappedFunc := WrapFunc(func(a int, b int) int {
		return a + b
	},
		cleanWrapperCountAfter,
	) // получили новую функцию
	result := wrappedFunc(1, 2) // применили

	assert.Equal(t, 3, result)
	assert.Equal(t, callsCounter, 1)

	result = wrappedFunc(10, 20)

	assert.Equal(t, 30, result)
	assert.Equal(t, callsCounter, 2)
}

var simpleFunc = func(a int, b int) int {
	return a + b
}

func BenchmarkNoReflection(b *testing.B) {
	for i := 0; i < b.N; i++ {
		simpleFunc(1, 2)
	}
}

var wrappedFunc = WrapFunc(simpleFunc, cleanWrapperCountAfter)

func BenchmarkWrappedFunc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		wrappedFunc(1, 2)
	}
}

var twiceWrappedFunc = WrapFunc(simpleFunc, cleanWrapperCountAfter, cleanWrapperCountAfter)

func BenchmarkTwiceWrappedFunc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		twiceWrappedFunc(1, 2)
	}
}
