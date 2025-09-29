package interactive

var interactiveRoutes = map[string]interface{}{
	"hostlookup": "/hostlookup",
	"user.login":   "/user/session",
	"user.logout":  "/user/session",
	"user.balance": "/user/balance",
	"user.profile": "/interactive/user/profile",
	"user.bokerage": "/interactive/user/calculatebrokerage",

	"exchange.message": "/interactive/messages/exchange",
	"exchange.status":  "/interactive/status/exchange",

	"order":           "/interactive/orders",
	"order.tradebook": "/interactive/orders/trades",
	"order.cancelall": "/interactive/orders/cancelall",
	"order.bracket":   "/interactive/orders/bracket",
	"order.cover":     "/interactive/orders/cover",

	"order.spread": "/interactive/orders/spread",

	"portfolio.holdings":        "/interactive/portfolio/holdings",
	"portfolio.positions":       "/interactive/portfolio/positions",
	"portfolio.positionConvert": "/interactive/portfolio/positions/convert",
	"portfolio.squareoff":       "/interactive/portfolio/squareoff",
	"portfolio.squareoffall":    "/interactive/portfolio/squareoffall",
}
