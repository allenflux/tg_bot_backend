// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// CentralControl is the golang structure for table central_control.
type CentralControl struct {
	Id                int    `json:"id"                orm:"id"                  description:""` //
	Name              string `json:"name"              orm:"name"                description:""` //
	Domain            string `json:"domain"            orm:"domain"              description:""` //
	NumberOfCustomers int    `json:"numberOfCustomers" orm:"number_of_customers" description:""` //
	NumberOfBusiness  int    `json:"numberOfBusiness"  orm:"number_of_business"  description:""` //
	Note              string `json:"note"              orm:"note"                description:""` //
	Status            int    `json:"status"            orm:"status"              description:""` //
	SecretKey         string `json:"secretKey"         orm:"secret_key"          description:""` //
}
