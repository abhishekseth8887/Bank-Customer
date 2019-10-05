package entity

// Address
type Address struct {
	Type      string    `json:"type"`
	AddrLine1   string `json:"addr-line1"`
	AddrLine2   string `json:"addr-line2"`
	AddrLine3    string `json:"addr-line3"`
	AddrLine4 string `json:"addr-line4"`
	ZipCode string `json:"zipcode"`
	City string `json:"city"`
	State string `json:"state"`
	CountryCode string `json:"country-code"`
}

