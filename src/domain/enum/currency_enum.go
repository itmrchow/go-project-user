package enum

type CurrencyEnum struct {
	PHP  string
	USD  string
	BTC  string
	USDT string
}

var Currency CurrencyEnum = CurrencyEnum{
	PHP:  "PHP",
	USD:  "USD",
	BTC:  "BTC",
	USDT: "USDT",
}
