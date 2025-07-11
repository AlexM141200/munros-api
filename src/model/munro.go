package model

type Munro struct {
	RunningNo    int     `json:"running_no"`
	DoBIHNumber  int     `json:"dobih_number"`
	Name         string  `json:"name"`
	SMCSection   string  `json:"smc_section"`
	RHBSection   string  `json:"rhb_section"`
	HeightM      float64 `json:"height_m"`
	HeightFt     int     `json:"height_ft"`
	Map1to50k    string  `json:"map_1_50k"`
	Map1to25k    string  `json:"map_1_25k"`
	GridRef      string  `json:"grid_ref"`
	GridRefXY    string  `json:"grid_ref_xy"`
	XCoord       float64 `json:"x_coord"`
	YCoord       float64 `json:"y_coord"`
	Latitude     float64 `json:"latitude"`
	Longitude    float64 `json:"longitude"`
	Classification string `json:"classification"`
	Comments     string  `json:"comments"`
	StreetmapURL string  `json:"streetmap_url"`
	GeographURL  string  `json:"geograph_url"`
	HillBaggingURL string `json:"hill_bagging_url"`
}
