package main

import "net/http"

type UserInfo struct {
	ID                   string     `json:"id"`
	Title                string     `json:"title"`
	Type                 string     `json:"type"`
	Properties           Properties `json:"properties"`
	AdditionalProperties bool       `json:"additionalProperties"`
}
type Key struct {
	Type string `json:"type"`
}
type Name struct {
	Type string `json:"type"`
}
type Password struct {
	Type string `json:"type"`
}
type EmailAddress struct {
	Type string `json:"type"`
}
type DisplayName struct {
	Type string `json:"type"`
}
type Notification struct {
	Type string `json:"type"`
}
type Items struct {
	Type string `json:"type"`
}
type ApplicationKeys struct {
	Type  string `json:"type"`
	Items Items  `json:"items"`
}
type Properties struct {
	Key             Key             `json:"key"`
	Name            Name            `json:"name"`
	Password        Password        `json:"password"`
	EmailAddress    EmailAddress    `json:"emailAddress"`
	DisplayName     DisplayName     `json:"displayName"`
	Notification    Notification    `json:"notification"`
	ApplicationKeys ApplicationKeys `json:"applicationKeys"`
}

type Req struct {
	ID                   string      `json:"id"`
	Title                string      `json:"title"`
	Type                 string      `json:"type"`
	Properties           Properties  `json:"properties"`
	Definitions          Definitions `json:"definitions"`
	AdditionalProperties bool        `json:"additionalProperties"`
	Required             []string    `json:"required"`
}
type Self struct {
	Type   string `json:"type"`
	Format string `json:"format"`
}
type Key struct {
	Type string `json:"type"`
}
type Name struct {
	Type string `json:"type"`
}
type EmailAddress struct {
	Type string `json:"type"`
}
type NAMING_FAILED struct {
	Type   string `json:"type"`
	Format string `json:"format"`
}
type PatternProperties struct {
	NAMING_FAILED NAMING_FAILED `json:".+"`
}
type AvatarUrls struct {
	Type                 string            `json:"type"`
	PatternProperties    PatternProperties `json:"patternProperties"`
	AdditionalProperties bool              `json:"additionalProperties"`
}
type DisplayName struct {
	Type string `json:"type"`
}
type Active struct {
	Type string `json:"type"`
}
type TimeZone struct {
	Type string `json:"type"`
}
type Locale struct {
	Type string `json:"type"`
}
type Groups struct {
	Ref string `json:"$ref"`
}
type ApplicationRoles struct {
	Ref string `json:"$ref"`
}
type Expand struct {
	Type string `json:"type"`
}
type Properties struct {
	Self             Self             `json:"self"`
	Key              Key              `json:"key"`
	Name             Name             `json:"name"`
	EmailAddress     EmailAddress     `json:"emailAddress"`
	AvatarUrls       AvatarUrls       `json:"avatarUrls"`
	DisplayName      DisplayName      `json:"displayName"`
	Active           Active           `json:"active"`
	TimeZone         TimeZone         `json:"timeZone"`
	Locale           Locale           `json:"locale"`
	Groups           Groups           `json:"groups"`
	ApplicationRoles ApplicationRoles `json:"applicationRoles"`
	Expand           Expand           `json:"expand"`
}
type Size struct {
	Type string `json:"type"`
}
type MaxResults struct {
	Type string `json:"type"`
}
type Properties struct {
	Name Name `json:"name"`
	Self Self `json:"self"`
}
type Items struct {
	Title                string     `json:"title"`
	Type                 string     `json:"type"`
	Properties           Properties `json:"properties"`
	AdditionalProperties bool       `json:"additionalProperties"`
}
type Items struct {
	Type  string `json:"type"`
	Items Items  `json:"items"`
}
type Properties struct {
	Size       Size       `json:"size"`
	MaxResults MaxResults `json:"max-results"`
	Items      Items      `json:"items"`
}
type SimpleListWrapper struct {
	Title                string     `json:"title"`
	Type                 string     `json:"type"`
	Properties           Properties `json:"properties"`
	AdditionalProperties bool       `json:"additionalProperties"`
	Required             []string   `json:"required"`
}
type Definitions struct {
	SimpleListWrapper SimpleListWrapper `json:"simple-list-wrapper"`
}

func (u *UserInfo) UserInfo() {

}
func (u *Req) UserInfo() {

}

func GetUserInfo( info *UserInfo) *Req  { 
}

func GetUser()  {
	
}

func main() {
	http.HandleFunc("/jira/user", )
		
}
