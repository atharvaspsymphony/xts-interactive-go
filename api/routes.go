package interactive

var interactiveRoutes = map[string]interface{}{
	"user.login":   "/interactive/user/session",
	"user.logout":  "/interactive/user/session",
	"user.balance": "/interactive/user/balance",
	"user.profile": "/interactive/user/profile",

	"exchange.message": "/interactive/messages/exchange",
	"exchange.status":  "/interactive/status/exchange",

	"order":           "/interactive/orders",
	"order.tradebook": "/interactive/orders/trades",
	"order.cancelall": "/interactive/orders/cancelall",
	"order.bracket":   "/interactive/orders/bracket",
	"order.cover":     "/interactive/orders/cover",

	"portfolio.holdings":        "/interactive/portfolio/holdings",
	"portfolio.positions":       "/interactive/portfolio/positions",
	"portfolio.positionConvert": "/interactive/portfolio/positions/convert",
	"portfolio.squareoff":       "/interactive/portfolio/squareoff",
	"portfolio.squareoffall":    "/interactive/portfolio/squareoffall",
}
