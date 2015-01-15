type TestStruct struct {
	ID            string  `json:"_id"`
	About         string  `json:"about"`
	Address       string  `json:"address"`
	Age           float64 `json:"age"`
	Balance       string  `json:"balance"`
	Company       string  `json:"company"`
	Email         string  `json:"email"`
	EyeColor      string  `json:"eyeColor"`
	FavoriteFruit string  `json:"favoriteFruit"`
	Friends       []struct {
		ID   float64 `json:"id"`
		Name string  `json:"name"`
	} `json:"friends"`
	Gender     string      `json:"gender"`
	Greeting   string      `json:"greeting"`
	Guid       string      `json:"guid"`
	Index      float64     `json:"index"`
	IsActive   bool        `json:"isActive"`
	Latitude   float64     `json:"latitude"`
	Longitude  float64     `json:"longitude"`
	Name       string      `json:"name"`
	Phone      string      `json:"phone"`
	Picture    string      `json:"picture"`
	Registered string      `json:"registered"`
	Tags       []string    `json:"tags"`
	Test_O_Q   interface{} `json:"test.o-q"`
}