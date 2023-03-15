package main

import "go.mongodb.org/mongo-driver/bson/primitive"

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
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	PhoneNumber string `json:"phone_number"`
	City        string `json:"city"`
}

type BaseInfo struct {
	BirthdayDate       string `json:"birthday_date"`
	Gender             bool   `json:"gender"`
	HaveWorkExperience bool   `json:"have_work_experience"`
}

type Special struct {
	CareerObjective string `json:"career_objective"`
	Payment         int32  `json:"payment"`
}

type WorkExperience struct {
	WorkStart        string `json:"work_start"`
	WorkNow          bool   `json:"work_now"`
	WorkEnd          string `json:"work_end"`
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
	EndYear        int32  `json:"end_year"`
}

type Languages struct {
	NativeLang       string            `json:"native_lang"`
	ForeignLanguages []ForeignLanguage `json:"foreign_languages"`
}

type ForeignLanguage struct {
	ForeignLanguage      string `json:"foreign_language"`
	ForeignLanguageLevel string `json:"foreign_language_level"`
}
