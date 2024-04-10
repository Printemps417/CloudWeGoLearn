package main

import (
	stock "example_shop/kitex_gen/example/shop/stock/stockservice"
	"log"
)

func main() {
	svr := stock.NewServer(new(StockServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
