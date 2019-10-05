package entity

// Customer
type Customer struct {
	FirstName      string    `json:"first-name"`
	MiddleName   string `json:"middle-name"`
	LastName   string `json:"last-name"`
	Dob    string `json:"date-of-birth"`
	MobileNumber string `json:"mobile-number"`
	Gender string `json:"gender"`
	CustomerNumber string `json:"customer-number"`
	CountryOfBirth string `json:"country-of-birth"`
	CountryOfResidence string `json:"country-of-residence"`
	CustomerSegment string `json:"customer-segment"`

	Addresses []Address `json:"addresses"`
}

