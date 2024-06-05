# go-aop
This package provide you ability to use AOP in your Golang programs

To use package in your project:
```bash 
go get github.com/alnpokrovsky/go-aop 
```

Simple wrapper example:
```go
var callsCounter = 0

func CleanWrapperCountAfter(fptr any) aop.WrappedFunc {
	return func(args []reflect.Value) []reflect.Value {
        log.Println(aop.FuncName(fptr)) // print func name before use
        results := aop.Proceed(fptr, args) // proceed wrapped func
		callsCounter++ // inc counter after func preceeded
		return results
	}
}

func main() {
    wrappedFunc := aop.WrapFunc(func(a int, b int) int {
		return a + b
	},
		CleanWrapperCountAfter,
	) // generate new func
    c := wrappedFunc(1,2) // call wrapped func
    log.Println(c)
    log.Println(callsCounter) // callsCounter == 1
}
```

