package main

import (
	"fmt"
	"github.com/cinar/indicator/v2/strategy"
	"github.com/cinar/indicator/v2/strategy/trend"
	"github.com/cinar/indicatoralpaca"
	"github.com/joho/godotenv"
	"os"
)

func main() {
	_, errEnv := os.Stat(".env")
	if !os.IsNotExist(errEnv) {
		godotenv.Load(".env")
	}

	// Initialize a new Alpaca Markets repository
	repository := indicatoralpaca.NewAlpacaBarsRepository(os.Getenv("ALPACA_API_KEY"), os.Getenv("ALPACA_API_SECRET"))

	// Make any necessary changes in GetBarsRequest
	repository.GetBarsRequestTemplate.Adjustment = "raw"
	repository.GetBarsRequestTemplate.Currency = "usd"

	// Use the Alpaca Markets repository in backtesting
	backtest := strategy.NewBacktest(repository, "output")
	backtest.Names = append(backtest.Names, "AAPL")
	// backtest.Strategies = append(backtest.Strategies, trend.NewAroonStrategy())
	backtest.Strategies = append(backtest.Strategies, trend.NewMacdStrategy())
	backtest.LastDays = 180

	err := backtest.Run()
	if err != nil {
		fmt.Println(err)
	}
}
