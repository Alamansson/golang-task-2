package golangtask

type Employee struct {
	firstName   string `json:"firstName"`
	middleName  string `json:"middleName"`
	lastName    string `json:"lastName"`
	inn         int `json:"inn"`
	position    string `json:"position"`
	phone       string `json:"phone"`
	description string `json:"description"`
	attributes  []map[string]string `json:"attributes"`
}

var Meta = `{
	"meta":{
		"href": "https://online.moysklad.ru/api/remap/1.2/entity/employee/metadata/attributes/c7c8933f-ca37-11ec-0a80-0cd300228fd0",
		"type": "attributemetadata",
		"mediaType": "application/json"
	},
	"value": null
}`
