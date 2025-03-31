// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// Bot is the golang structure for table bot.
type Bot struct {
	Id             int    `json:"id"             orm:"id"              description:""` //
	Account        string `json:"account"        orm:"account"         description:""` //
	Name           string `json:"name"           orm:"name"            description:""` //
	Greeting       string `json:"greeting"       orm:"greeting"        description:""` //
	GreetingStatus int    `json:"greetingStatus" orm:"greeting_status" description:""` //
	Status         int    `json:"status"         orm:"status"          description:""` //
	Photo          string `json:"photo"          orm:"photo"           description:""` //
}
