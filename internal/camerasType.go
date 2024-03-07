package internal

type Camera1 struct{
	TransportChars   string `json:"transport_chars"`
	TransportNumbers string `json:"transport_numbers"`
	TransportRegion  string `json:"transport_region"`
	CameraID         string `json:"camera_id"`
	ViolationID      string `json:"violation_id"`
	ViolationValue   string `json:"violation_value"`
	SkillValue       int    `json:"skill_value"`
	Datetime         string `json:"datetime"`
}

type Camera2 struct{
	Transport struct {
		Chars   string `json:"chars"`
		Numbers string `json:"numbers"`
		Region  string `json:"region"`
	} `json:"transport"`
	Camera struct {
		ID string `json:"id"`
	} `json:"camera"`
	Violation struct {
		ID    string `json:"id"`
		Value string `json:"value"`
	} `json:"violation"`
	Skill struct {
		Value int `json:"value"`
	} `json:"skill"`
	Datetime struct {
		Year      int    `json:"year"`
		Month     int    `json:"month"`
		Day       int    `json:"day"`
		Minute    int    `json:"minute"`
		Seconds   int    `json:"seconds"`
		UtcOffset string `json:"utc_offset"`
	} `json:"datetime"`
}

type Camera3 struct{
	Transport string `json:"transport"`
	Camera    struct {
		ID string `json:"id"`
	} `json:"camera"`
	Violation struct {
		ID    string `json:"id"`
		Value string `json:"value"`
	} `json:"violation"`
	Skill    int `json:"skill"`
	Datetime int `json:"datetime"`
}