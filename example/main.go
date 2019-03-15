package main

/**
*
* @author Liu Weiyi
* @date 2019-03-14 23:20
 */
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
