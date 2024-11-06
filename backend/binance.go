package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"github.com/gorilla/websocket"
)

const binanceWS = "wss://stream.binance.com:9443/ws/btcusdt@depth"

type PriceLevel struct {
	Price    float64
	Quantity float64
}

type OrderBook struct {
	Bids []PriceLevel
	Asks []PriceLevel
}

type BinanceOrderBookUpdate struct {
	Bids [][]string `json:"b"`
	Asks [][]string `json:"a"`
}

func connectToBinanceAndProcess() {
	c, _, err := websocket.DefaultDialer.Dial(binanceWS, nil)
	if err != nil {
		log.Fatal("Error connecting to Binance:", err)
	}
	defer c.Close()

	for {
		_, message, err := c.ReadMessage()
		if err != nil {
			log.Println("Error reading from Binance:", err)
			return
		}

		var update BinanceOrderBookUpdate
		if err := json.Unmarshal(message, &update); err != nil {
			log.Println("Error unmarshalling Binance message:", err)
			continue
		}

		orderBook := processOrderBookUpdate(update)
		averagePrice := calculateAveragePrice(orderBook)
		broadcast <- averagePrice
	}
}

func processOrderBookUpdate(update BinanceOrderBookUpdate) OrderBook {
	var orderBook OrderBook

	for _, bid := range update.Bids {
		price, quantity := parsePriceLevel(bid)
		orderBook.Bids = append(orderBook.Bids, PriceLevel{Price: price, Quantity: quantity})
	}

	for _, ask := range update.Asks {
		price, quantity := parsePriceLevel(ask)
		orderBook.Asks = append(orderBook.Asks, PriceLevel{Price: price, Quantity: quantity})
	}

	return orderBook
}

func calculateAveragePrice(orderBook OrderBook) float64 {
	totalPrice, totalCount := 0.0, 0.0

	for _, bid := range orderBook.Bids {
		totalPrice += bid.Price * bid.Quantity
		totalCount += bid.Quantity
	}

	for _, ask := range orderBook.Asks {
		totalPrice += ask.Price * ask.Quantity
		totalCount += ask.Quantity
	}

	if totalCount == 0 {
		return 0
	}
	return totalPrice / totalCount
}

func parsePriceLevel(level []string) (float64, float64) {
	price, _ := strconv.ParseFloat(level[0], 64)
	quantity, _ := strconv.ParseFloat(level[1], 64)
	return price, quantity
}
