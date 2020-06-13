package hello

import (
	"github.com/mgusiew-guide/go-module-hello/commons"
	"rsc.io/quote/v3"
)

func Hello() string {
	//return quote.HelloV3()
	return "Hello Maciek " + commons.Hello()
}

func Proverb() string {
	return quote.Concurrency()
}
