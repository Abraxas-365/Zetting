package user_models

type Features struct {
	Gender string `json:"gender"`
	Age    int    `json:"age"`
	// desde aqui es actores
	Height     int      `json:"height"`
	Body       string   `json:"body"`
	Skin       string   `json:"skin"`
	HairType   string   `json:"hair_type"`
	HairZise   string   `json:"hair_zise"`
	Skills     []string `json:"skills"`
	FacialHair string   `json:"facial_hair"`
	HairColor  string   `json:"hair_color"`
	EyeColor   string   `json:"eye_color"`
	// Aqui acaba actores
}
