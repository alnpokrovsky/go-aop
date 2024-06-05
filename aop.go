package aop

import (
	"fmt"
	"reflect"
	"runtime"
)

type WrappedFunc = func(args []reflect.Value) []reflect.Value

// WrapFunc returns new function with same signature as funcIn func
// funcIn will be argument of every wrapper.
// so you should manually call Proceed(fptr, args)
// to get result of inner function inside wrapper
func WrapFunc[FuncType any](
	funcIn FuncType,
	wrappers ...func(any) WrappedFunc,
) (funcOut FuncType) {
	funcOut = funcIn
	for _, wrapper := range wrappers {
		funcOut = wrapFunc(funcOut, wrapper)
	}
	return
}

func wrapFunc[FuncType any](
	funcIn FuncType,
	wrapper func(any) WrappedFunc,
) (funcOut FuncType) {
	fn := reflect.ValueOf(&funcOut).Elem()
	rf2 := reflect.MakeFunc(fn.Type(), wrapper(funcIn))
	fn.Set(rf2)
	return
}

// Proceed is usabale for wrapper, when you just want
// to call wrapped func with provided arguments
func Proceed(fptr any, args []reflect.Value) []reflect.Value {
	return reflect.ValueOf(fptr).Call(args)
}

// IsImplements checks if reflectValue rv implements interface T
func IsImplements[T any](rv reflect.Value) bool {
	return rv.Type().Implements(reflect.TypeOf((*T)(nil)).Elem())
}

// As casts reflectValue rv to your interface T
func As[T any](rv reflect.Value) T {
	return rv.Interface().(T)
}

// FuncName returns function name with it's argument and return types
func FuncName(fn any) string {
	fnName := runtime.FuncForPC(reflect.ValueOf(fn).Pointer()).Name()
	fnType := reflect.TypeOf(fn).String()
	return fmt.Sprintf("%s %s", fnName, fnType)
}
