// Package reporting binds the reporting domain set of routes into the specified app.
package reporting

import (
	"github.com/0XFF-96/Service-A/app/services/sales-api/v1/handlers/checkgrp"
	"github.com/0XFF-96/Service-A/app/services/sales-api/v1/handlers/usersummarygrp"
	v1 "github.com/0XFF-96/Service-A/business/web/v1"
	"github.com/0XFF-96/Service-A/foundation/web"
)

// Routes constructs the add value which provides the implementation of
// of RouteAdder for specifying what routes to bind to this instance.
func Routes() add {
	return add{}
}

type add struct{}

// Add implements the RouterAdder interface.
func (add) Add(app *web.App, cfg v1.APIMuxConfig) {
	checkgrp.Routes(app, checkgrp.Config{
		UsingWeaver: cfg.UsingWeaver,
		Build:       cfg.Build,
		DB:          cfg.DB,
	})

	usersummarygrp.Routes(app, usersummarygrp.Config{
		Log:  cfg.Log,
		Auth: cfg.Auth,
		DB:   cfg.DB,
	})
}
