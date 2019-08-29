package main

import (
	routers "my_blog/router"
)

func main() {
	r := routers.InitRouter()

	r.Run(":8088")
}
