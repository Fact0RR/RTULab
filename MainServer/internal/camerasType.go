package internal

import (
	"strconv"
	"strings"
	"time"
)

type Camera1 struct {
	Data struct {
		TransportChars   string `json:"transport_chars"`
		TransportNumbers string `json:"transport_numbers"`
		TransportRegion  string `json:"transport_region"`
		CameraID         string `json:"camera_id"`
		ViolationID      string `json:"violation_id"`
		ViolationValue   string `json:"violation_value"`
		SkillValue       int    `json:"skill_value"`
		Datetime         string `json:"datetime"`
	} `json:"data"`
	Photo string `json:"photo"`
}

func (c *Camera1) ToUnionStruct() *UnionCamera {
	c.Data.Datetime = strings.Replace(c.Data.Datetime, "T", " ", 1)
	uc := UnionCamera{
		Transport:      c.Data.TransportChars + c.Data.TransportNumbers + c.Data.TransportRegion,
		CameraID:       c.Data.CameraID,
		ViolationID:    c.Data.ViolationID,
		ViolationValue: c.Data.ViolationValue,
		Skill:          c.Data.SkillValue,
		DateTime:       c.Data.Datetime[:len(c.Data.Datetime)-6],
		Photo:          c.Photo,
	}
	return &uc
}

type Camera2 struct {
	Data struct {
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
			Hour      int    `json:"hour"`
			Minute    int    `json:"minute"`
			Seconds   int    `json:"seconds"`
			UtcOffset string `json:"utc_offset"`
		} `json:"datetime"`
	} `json:"data"`
	Photo string `json:"photo"`
}

func digitToNumber(i int) string{
	if i>=0 && i<10{
		return "0"+strconv.Itoa(i)
	}else {
		return strconv.Itoa(i)
	}
}

func (c *Camera2) ToUnionStruct() *UnionCamera {
	
	timestr := strconv.Itoa(c.Data.Datetime.Year) + "-" + digitToNumber(c.Data.Datetime.Month) + "-" + digitToNumber(c.Data.Datetime.Day) + " " +
		digitToNumber(c.Data.Datetime.Hour) + ":" + digitToNumber(c.Data.Datetime.Minute) + ":" + digitToNumber(c.Data.Datetime.Seconds)

	uc := UnionCamera{
		Transport:      c.Data.Transport.Chars + c.Data.Transport.Numbers + c.Data.Transport.Region,
		CameraID:       c.Data.Camera.ID,
		ViolationID:    c.Data.Violation.ID,
		ViolationValue: c.Data.Violation.Value,
		Skill:          c.Data.Skill.Value,
		DateTime:       timestr,
		Photo:          c.Photo,
	}
	return &uc
}

type Camera3 struct {
	Data struct {
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
	} `json:"data"`
	Photo string `json:"photo"`
}

func (c *Camera3) ToUnionStruct() *UnionCamera {

	unixTimeUTC := time.Unix(int64(c.Data.Datetime), 0).UTC()

	uc := UnionCamera{
		Transport:      c.Data.Transport,
		CameraID:       c.Data.Camera.ID,
		ViolationID:    c.Data.Violation.ID,
		ViolationValue: c.Data.Violation.Value,
		Skill:          c.Data.Skill,
		DateTime:       unixTimeUTC.String()[:len(unixTimeUTC.String())-10],
		Photo:          c.Photo,
	}
	return &uc
}

type UnionCamera struct {
	Transport      string
	CameraID       string
	ViolationID    string
	ViolationValue string
	Skill          int
	DateTime       string
	Photo          string
}
