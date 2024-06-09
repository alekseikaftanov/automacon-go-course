package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
)

/*
Задача 13.1 Необходимо распарсить json
{"number":1,"landlord":"ОстапБендер","tenat":"ШураБалаганов»}
*/

type Contract struct {
	Number   int    `json:"number"`
	Landlord string `json:"landlord"`
	Tenant   string `json:"tenat"`
}

type Contracts struct {
	List []Contractxml `xml:"contract"`
}

type Contractxml struct {
	Number   int    `xml:"number"`
	SignDate string `xml:"sign_date"`
	Landlord string `xml:"landlord"`
	Tenant   string `xml:"tenat"`
}

func jsonParse() {

	jsonData := `{ "number": 1, "landlord": "Остап Бендер", "tenat": "Шура Балаганов"}`

	var contract Contract

	err := json.Unmarshal([]byte(jsonData), &contract)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("Number: %d\nLandlord: %s\nTenant: %s\n", contract.Number, contract.Landlord, contract.Tenant)

}

/*
Задача 13.2 Необходимо представить в виде json структуру contract
contract{
	Number:   2,
	Landlord: "Остап",
	tenat:    "Шура",}
*/

func jsonCreateFromStruct() {
	var contract Contract
	jsonData, err := json.Marshal(contract)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(string(jsonData))
}

/*
Задача 13.3 Необходимо распарсить xml
*/

func parseXml() {
	xmlData := `<?xml version="1.0" encoding="UTF-8"?>
<contracts>
  <contract>
     <number>1</number>
	 <sign_date>2023-09-02</sign_date>
	 <landlord>Остап</landlord>
	 <tenat>Шура</tenat>
	 </contract>
	 <contract>
	 <number>2</number>
	 <sign_date>2023-09-03</sign_date>
	 <landlord>Бендер</landlord>
	  <tenat>Балаганов</tenat>
	   </contract>
</contracts>`

	var contracts Contracts

	err := xml.Unmarshal([]byte(xmlData), &contracts)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for _, contract := range contracts.List {
		fmt.Printf("Number: %d\nSignDate: %s\nLandlord: %s\nTenant: %s\n", contract.Number, contract.SignDate, contract.Landlord, contract.Tenant)
	}
}

/*
Задача 13.4 Необходимо представить в виде xml структуру
 type contracts struct
 {
	List []contract `xml:"contract"`
}
 contract{
	Number:   1,
	Landlord: "ОстапБендер",
	tenat:    "ШураБалаганов"
}
*/

func xmlCreateFromStruct() {
	var contract Contract
	xmlData, err := xml.Marshal(contract)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(string(xmlData))
}

func main() {
	jsonParse() // 13.2
	jsonCreateFromStruct() // 13.3
	parseXml() // 13.4
}
