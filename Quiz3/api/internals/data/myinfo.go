package data

type Me struct {
	Name       string   `json:"name"`
	Age        int      `json:"age"`
	Location   string   `json:"location"`
	Email      string   `json:"email"`
	Interest   []string `json:"interest"`
	Occupation string   `json:"occupation"`
}
