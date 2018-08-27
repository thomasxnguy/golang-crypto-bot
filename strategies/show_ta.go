// Copyright © 2017 Thomas Nguy
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

package strategies

import (
	"fmt"
	"time"

	"github.com/thomasxnguy/golang-crypto-bot/environment"
	"github.com/thomasxnguy/golang-crypto-bot/exchanges"
	techan "github.com/sdcoffey/techan"
	"github.com/sdcoffey/big"
)

var ShowTa Strategy = IntervalStrategy{
	Model: StrategyModel{
		Name: "ShowTa",
		Setup: func(wrappers []exchanges.ExchangeWrapper, markets []*environment.Market) error {
			fmt.Println("ShowTa starting")
			return nil
		},
		OnUpdate: func(wrappers []exchanges.ExchangeWrapper, markets []*environment.Market) error {
			series := techan.NewTimeSeries()
			i := 0
			for i < 14 {
				results, err := wrappers[0].GetKlines(time.Now().UnixNano()/1000000-int64((14-i)*6060000), "BTCUSDT", "1m")
				if err != nil {
					return err
				}

				for _, c := range results.CandleSticks {
					period := techan.NewTimePeriod(time.Unix(c.OpenTime, 0), time.Minute)
					candle := techan.NewCandle(period)
					candle.OpenPrice = big.NewFromString(c.Open.String())
					candle.ClosePrice = big.NewFromString(c.Close.String())
					candle.MaxPrice = big.NewFromString(c.High.String())
					candle.MinPrice = big.NewFromString(c.Low.String())
					series.AddCandle(candle)
				}
				i++
			}
			closePrices := techan.NewClosePriceIndicator(series)
			movingAverage := techan.NewEMAIndicator(closePrices, 100)

			fmt.Printf("closePrices : %v \n", closePrices.Calculate(0).FormattedString(2))
			fmt.Printf("Moving average : %v \n",movingAverage.Calculate(0).FormattedString(2))
			return nil
		},
		OnError: func(err error) {
			fmt.Println(err)
		},
		TearDown: func(wrappers []exchanges.ExchangeWrapper, markets []*environment.Market) error {
			c.Close()
			fmt.Println("ShowTa exited")
			return nil
		},
	},
	Interval: 60 *time.Second,
}
