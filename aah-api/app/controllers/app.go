package controllers

import (
  "aahframework.org/aah.v0"
  "github.com/fedir/gmm02/aah-api/app/models"
)

// AppController struct application controller
type AppController struct {
  *aah.Context
}

// Index method is application root API endpoint.
func (a *AppController) Index() {
  a.Reply().Ok().JSON(models.Greet{
    Message: "Welcome to aah framework - API application",
  })
}
