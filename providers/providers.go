package models

type MetaPurchase struct {
	ID                      int     `json:"id"`
	TransactionID           string  `json:"transactionID"`
	EndUserName             string  `json:"EndUserName"`
	EndUserEmail            string  `json:"EndUserEmail"`
	EndUserCountry          string  `json:"EndUserCountry"`
	EndUserPhoneCountryCode string  `json:"EndUserPhoneCountryCode"`
	EndUserPhoneNumber      string  `json:"EndUserPhoneNumber"`
	EndUserSecondNumber     string  `json:"EndUserSecondNumber"`
	ProductID               int     `json:"productID"`
	TotalValue              float64 `json:"totalValue"`
	Currency                string  `json:"currency"`
}
