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

/*
В рамках задачи будем работать с картотекой известного врача. Нужно будет написать модуль с несколькими версиями: v1.0.0, v1.1.0, v2.0.0, v2.1.0. Модуль должен прочитать файл со следующим содержимым:

{"name":"Ёжик","age":10,"email":"ezh@mail.ru"}
{"name":"Зайчик","age":2,"email":"zayac@mail.ru"}
{"name":"Лисичка","age":3,"email":"alice@mail.ru"}

v1.0.0 должна создавать файл с содержимым:

[{«name»:»Ёжик","age":10,"email":"ezh@mail.ru"},
{"name":"Зайчик","age":2,"email":"zayac@mail.ru"},
{«name":"Лисичка","age":3,"email":"alice@mail.ru"}]

v1.1.0 должна сортировать данные по полю age по
возрастанию:
[{«name":"Зайчик","age":2,"email":"zayac@mail.ru"},
{«name":"Лисичка","age":3,"email":"alice@mail.ru"}{«name»:»
Ёжик","age":10,"email":"ezh@mail.ru"}]

v2.0.0 должна создавать файл с содержимым:
<?xml version="1.0" encoding="UTF-8"?>
<patients>
<Patient>
<Name>Ёжик</Name>
<Age>10</Age>
<Email>ezh@mail.ru</Email>
</Patient>
<Patient>
<Name>Зайчик</Name>
<Age>2</Age>
<Email>zayac@mail.ru</Email>
</Patient>
<Patient>
<Name>Лисичка</Name>
<Age>3</Age>
<Email>alice@mail.ru</Email>
</Patient>
</patients>
v2.1.0 должна сортировать данные по полю age по возрастанию:
<?xml version="1.0" encoding="UTF-8"?>
<patients>
<Patient>
<Name>Зайчик</Name>
<Age>2</Age>
<Email>zayac@mail.ru</Email>
</Patient>
<Patient>
<Name>Лисичка</Name>
<Age>3</Age>
<Email>alice@mail.ru</Email>
</Patient>
<Patient>
<Name>Ёжик</Name>
<Age>10</Age>
<Email>ezh@mail.ru</Email>
</Patient>
</patients>
Модуль должен содержать функцию Do, которая принимает два
строковых параметра: путь файла откуда прочитать данные, путь
файла, в который записать данные в требуемом формате; и
возвращать ошибку. Пример использования модуля:
package main
import (
format ...
)
func main() {
err := format.Do("patients", «result")
if err != nil {
…
}
}
Должна быть возможность подключить любую из версий: v1.0.0, v1.1.0,
v2.0.0, v2.1.0.
*/

func main() {
	jsonParse()
	jsonCreateFromStruct()
	parseXml()
}
