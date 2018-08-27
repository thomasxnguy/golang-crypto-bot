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

package exchanges

import (
	"github.com/thomasxnguy/golang-crypto-bot/environment"
)

// BitmexWrapper represents the wrapper for the Bitmex exchange.
type BitmexWrapper struct {
	//Add client here
}

// Name returns the name of the wrapped exchange.
func (wrapper BitmexWrapper) Name() string {
	return "Bitmex"
}

func (wrapper BitmexWrapper) String() string {
	return wrapper.Name()
}

// GetMarkets Gets all the markets info.
func (wrapper BitmexWrapper) GetMarkets() ([]*environment.Market, error) {
	return nil, nil
}

// GetOrderBook gets the order(ASK + BID) book of a market.
func (wrapper BitmexWrapper) GetOrderBook(market *environment.Market) (*environment.OrderBook, error) {
	return nil, nil
}

// BuyLimit performs a limit buy action.
func (wrapper BitmexWrapper) BuyLimit(market *environment.Market, amount float64, limit float64) (string, error) {
	return "", nil
}

// SellLimit performs a limit sell action.
func (wrapper BitmexWrapper) SellLimit(market *environment.Market, amount float64, limit float64) (string, error) {
	return "", nil
}

// GetTicker gets the updated ticker for a market.
func (wrapper BitmexWrapper) GetTicker(market *environment.Market) (*environment.Ticker, error) {
	return nil, nil
}

// GetMarketSummary gets the current market summary.
func (wrapper BitmexWrapper) GetMarketSummary(market *environment.Market) (*environment.MarketSummary, error) {
	return nil, nil
}

// CalculateTradingFees calculates the trading fees for an order on a specified market.
//
//     NOTE: In Binance fees are currently hardcoded.
func (wrapper BitmexWrapper) CalculateTradingFees(market *environment.Market, amount float64, limit float64, orderType TradeType) float64 {
	return 0
}

// CalculateWithdrawFees calculates the withdrawal fees on a specified market.
func (wrapper BitmexWrapper) CalculateWithdrawFees(market *environment.Market, amount float64) float64 {
	panic("Not Implemented")
}

// FeedConnect connects to the feed of the exchange.
func (wrapper BitmexWrapper) FeedConnect() {
	//empty
}

// SubscribeMarketSummaryFeed subscribes to the Market Summary Feed service.
func (wrapper BitmexWrapper) SubscribeMarketSummaryFeed(market *environment.Market, onUpdate func(environment.MarketSummary)) {
}

// UnsubscribeMarketSummaryFeed unsubscribes from the Market Summary Feed service.
func (wrapper BitmexWrapper) UnsubscribeMarketSummaryFeed(market *environment.Market) {
}

// GetKlines Gets candlestick bar information
func (wrapper BitmexWrapper) GetKlines(start int64, symbol string, interval string) (*environment.CandleStickChart, error) {
	return nil, nil
}
