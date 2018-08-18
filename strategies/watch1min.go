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

	"github.com/thomasxnguy/golang-crypto-trading-bot/environment"
	"github.com/thomasxnguy/golang-crypto-trading-bot/exchanges"
)

// Watch1Min prints out the info of the market every 5 minutes.
var Watch1Min Strategy = IntervalStrategy{
	Model: StrategyModel{
		Name: "Watch1Min",
		Setup: func(wrappers []exchanges.ExchangeWrapper, markets []*environment.Market) error {
			fmt.Println("Watch1Min starting")
			return nil
		},
		OnUpdate: func(wrappers []exchanges.ExchangeWrapper, markets []*environment.Market) error {
			markets, err := wrappers[0].GetMarkets()
			if err != nil {
				return err
			}
			fmt.Println(markets)
			return nil
		},
		OnError: func(err error) {
			fmt.Println(err)
		},
		TearDown: func(wrappers []exchanges.ExchangeWrapper, markets []*environment.Market) error {
			fmt.Println("Watch1Min exited")
			return nil
		},
	},
	Interval: time.Minute,
}
