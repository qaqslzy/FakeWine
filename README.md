# FakeWine
This project just to learn how to design a go web framework.

And project's code 100% from [gin](https://github.com/gin-gonic/gin), 0 % from me.

I just copy that code for learning.

## How To Use
```go
package main
import "FakeWine/fakewine"

func main() {
	fake := fakewine.New()
	fake.GET("/", func(ctx *fakewine.Context) {
		ctx.JSON(200, fakewine.H{
			"hello": "world",
		})
	})
	_ = fake.Run(":8000")
}
```

The same to [gin](https://github.com/gin-gonic/gin), but lose almost all of them functions.

## Read Gin's code
[Gin源码浅读（一）]()

[Gin源码浅读（二）]()

[Gin源码浅读（三）]()

## The end
Continuous coding , but not always coding.

(Maybe I write this because of drinking the fake wine...)