package mongo

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID             primitive.ObjectID `bson:"_id,omitempty"`
	Username       string             `json:"username"`
	Contacts       Contacts           `json:"contacts"`
	BaseInfo       BaseInfo           `json:"base_info"`
	Special        Special            `json:"special"`
	WorkExperience []WorkExperience   `json:"work_experience"`
	About          string             `json:"about"`
	Skills         []string           `json:"skills"`
	Study          Study              `json:"study"`
	Languages      Languages          `json:"languages"`
	Employment     string             `json:"employment"`
}

type Contacts struct {
	First_Name   string `json:"first_name" validate:"required,first_name"`
	Last_Name    string `json:"last_name" validate:"required,last_name"`
	Phone_Number string `json:"phone_number" validate:"required,phone_number"`
	City         string `json:"city" validate:"required,city"`
}

type BaseInfo struct {
	BirthdayDate       string `json:"birthdaydate"`
	Gender             bool   `json:"gender"`
	HaveWorkExperience bool   `json:"haveworkexperience"`
}

type Special struct {
	CareerObjective string `json:"careerobjective"`
	Payment         int32  `json:"payment"`
}

type WorkExperience struct {
	WorkStart        string `json:"workstart"`
	WorkNow          bool   `json:"worknow"`
	WorkEnd          string `json:"workend"`
	Organization     string `json:"organization"`
	Position         string `json:"position"`
	Responsibilities string `json:"responsibilities"`
}

type Study struct {
	Level        string        `json:"level"`
	Institutions []Institution `json:"institutions"`
}

type Institution struct {
	Institution    string `json:"institution"`
	Faculty        string `json:"faculty"`
	Specialization string `json:"specialization"`
	EndYear        int32  `json:"endyear"`
}

type Languages struct {
	NativeLang       string            `json:"nativelang"`
	ForeignLanguages []ForeignLanguage `json:"foreignlanguages"`
}

type ForeignLanguage struct {
	ForeignLanguage      string `json:"foreignlanguage"`
	ForeignLanguageLevel string `json:"foreignlanguagelevel"`
}

type Cv struct {
	User User `json:"user"`
}
