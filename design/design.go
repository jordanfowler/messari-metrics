package design

import . "goa.design/goa/v3/dsl"

var _ = API("metrics", func() {
	Title("Messari Metrics")
	Description("Service providing metrics on single and aggregate assets from Messari Data.")
	Server("metrics", func() {
		Host("localhost", func() {
			URI("http://localhost:8000/api/")
		})
	})
})

var _ = Service("metrics", func() {
	Method("asset", func() {
		Payload(func() {
			Attribute("slug", String)
		})
		Result(func() {
			Attribute("price", Float64, "Current spot price in USD")
			Attribute("vlm24hr", Float64, "Volume traded over last 24 hours")
			Attribute("chg24hr", Float64, "Change in price over last 24 hours")
			Attribute("mktcap", Float64, "Market cap of asset")
			// Can add additional metrics here...
		})
		HTTP(func() {
			GET("/asset/{slug}")
			Response(StatusOK)
		})
	})
	Method("aggregate", func() {
		Payload(func() {
			Attribute("tags", String)
			Attribute("sector", String)
		})
		Result(func() {
			Attribute("price", Float64, "Current spot price in USD")
			Attribute("vlm24hr", Float64, "Volume traded over last 24 hours")
			Attribute("chg24hr", Float64, "Change in price over last 24 hours")
			Attribute("mktcap", Float64, "Market cap of asset")
			// Can add additional metrics here...
		})
		HTTP(func() {
			GET("/aggregate")
			Param("tags")
			Param("sector")
			Response(StatusOK)
		})
	})
})
