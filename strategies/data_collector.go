// Copyright © 2017 Alessandro Sanino <saninoale@gmail.com>
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
	"github.com/influxdata/influxdb/client/v2"
	"github.com/thomasxnguy/golang-crypto-bot/util/influxdb"
)

var c client.Client

const (
	DB          = "crypto_bot"
	username    = "golang"
	password    = "golang"
	precision   = "s"
	measurement = "binance_candle_15m"
)

var DataCollector Strategy = IntervalStrategy{
	Model: StrategyModel{
		Name: "DataCollector",
		Setup: func(wrappers []exchanges.ExchangeWrapper, markets []*environment.Market) error {
			fmt.Println("DataCollector starting")
			var err error
			c, err = client.NewHTTPClient(client.HTTPConfig{
				Addr:     "http://localhost:8086",
				Username: username,
				Password: password,
			})
			if err != nil {
				fmt.Println(err)
			}
			return nil
		},
		OnUpdate: func(wrappers []exchanges.ExchangeWrapper, markets []*environment.Market) error {

			lastTs, err := influxdb.GetLastTimestamp(c,DB,precision,measurement,"open")

			if err != nil || lastTs == 0 {
				fmt.Println("first data")
				lastTs = 1509579154
			}
			fmt.Println(lastTs*1000)
			results, err := wrappers[0].GetKlines(lastTs*1000, "BTCUSDT")

			bp, err := client.NewBatchPoints(client.BatchPointsConfig{
				Database:  DB,
				Precision: precision,
			})

			if err != nil {
				fmt.Println(err)
				return err
			}

			for _, val := range results.CandleSticks {

				fields := map[string]interface{}{
					"open":     val.Open,
					"close":    val.Close,
					"low":      val.Low,
					"high":     val.High,
					"trade_nb": val.TradeNb,
					"volume":   val.Volume,
				}

				pt, err := client.NewPoint(
					measurement,
					nil,
					fields,
					time.Unix(val.OpenTime/1000, 0),
				)
				if err != nil {
					fmt.Println(err)
					return err
				}
				bp.AddPoint(pt)
			}

			influxdb.WritePoints(c, bp)
			fmt.Println("write points")
			return nil
		},
		OnError: func(err error) {
			fmt.Println(err)
		},
		TearDown: func(wrappers []exchanges.ExchangeWrapper, markets []*environment.Market) error {
			c.Close()
			fmt.Println("DataCollector exited")
			return nil
		},
	},
	Interval: 5 * time.Second,
}
