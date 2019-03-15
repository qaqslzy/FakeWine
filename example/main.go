package main

/**
*
* @author Liu Weiyi
* @date 2019-03-14 23:20
 */
import "FakeWine"

func main() {
	fake := FakeWine.New()
	fake.GET("/", func(ctx *FakeWine.Context) {
		ctx.JSON(200, FakeWine.H{
			"hello": "world",
		})
	})
	_ = fake.Run(":8000")
}
