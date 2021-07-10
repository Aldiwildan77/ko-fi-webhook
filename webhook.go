package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type KoFiRequest struct {
	MessageID                  string    `json:"message_id"`
	KofiTransactionID          string    `json:"kofi_transaction_id"`
	Timestamp                  time.Time `json:"timestamp"`
	Type                       string    `json:"type"`
	FromName                   string    `json:"from_name"`
	Message                    string    `json:"message"`
	Amount                     string    `json:"amount"`
	Currency                   string    `json:"currency"`
	URL                        string    `json:"url"`
	IsSubscriptionPayment      bool      `json:"is_subscription_payment"`
	IsFirstSubscriptionPayment bool      `json:"is_first_subscription_payment"`
}

func webhook(c echo.Context) error {
	data := c.FormValue("data")

	var kofi KoFiRequest
	if err := json.Unmarshal([]byte(data), &kofi); err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}

	// pending, should be connected with rust plugin
	// so players can have a benefit after they are donated

	return c.JSON(http.StatusOK, nil)
}
