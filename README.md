# golang-crypto-trading-bot

[![GoDoc](https://godoc.org/github.com/thomasxnguy/golang-crypto-trading-bot?status.svg)](https://godoc.org/github.com/thomasxnguy/golang-crypto-trading-bot)
[![license](https://img.shields.io/github/license/thomasxnguy/golang-crypto-trading-bot.svg?maxAge=2592000)](https://github.com/thomasxnguy/golang-crypto-trading-bot/LICENSE)


A golang implementation of a console-based trading bot for cryptocurrency exchanges. 

## Supported Exchanges
Bittrex, Poloniex, Binance, Bitfinex and Kraken, other in progress.

## Usage

Download a release or directly build the code from this repository.
``` bash
$ go get github.com/thomasxnguy/golang-crypto-bot
```

If you need to, you can create a strategy and bind it to the bot:
``` go
import bot "github.com/thomasxnguy/golang-crypto-bot/cmd"

bot.AddCustomStrategy(myStrategy)
bot.Execute()
```

For strategy reference see the [Godoc documentation](https://godoc.org/github.com/thomasxnguy/golang-crypto-bot).

# Configuration file template
Create a configuration file from this example or run the `init` command of the compiled executable.
``` yaml
exchange_configs: 
  - exchange: bittrex
    public_key: your_bittrex_public_key
    secret_key: your_bittrex_secret_key
  - exchange: binance
    public_key: your_binance_public_key
    secret_key: your_binance_secret_key
  - exchange: bitfinex
    public_key: your_bitfinex_public_key
    secret_key: your_bitfinex_secret_key
strategies:
  - strategy: your_strategy_name
    markets:
      - market: market_logical_name
        bindings:
        - exchange: bittrex
          market_name: market_name_on_bittrex
        - exchange: binance
          market_name: market_name_on_binance
        - exchange: bitfinex
          market_name: market_name_on_bitfinex
      - market: another_market_logical_name
        bindings:
        - exchange: bittrex
          market_name: market_name_on_bittrex
        - exchange: binance
          market_name: market_name_on_binance
        - exchange: bitfinex
          market_name: market_name_on_bitfinex
```