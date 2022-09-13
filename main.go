package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/getDocketData", getDocketData)
	http.HandleFunc("/postDocketData", postData)
	log.Fatal(http.ListenAndServe(":10000", nil))

}

func getDocketData(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(parseJson())

}

func postData(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "I can have a post request here, maybe.")
}

func main() {
	handleRequests()
}

type Response struct {
	CreatedAt         string      `json:"createdAt"`
	Date              string      `json:"date"`
	DocketNumber      string      `json:"docketNumber"`
	Id                string      `json:"id"`
	Organisation      string      `json:"organisation"`
	OrganisationGroup string      `json:"organisationGroup"`
	Resources         []Resources `json:"resources"`
	Status            string      `json:"status"`
	UpdatedAt         string      `json:"updatedAt"`
}

type Resources struct {
	Id           string         `json:"id"`
	Interactions []Interactions `json:"interactions"`
	Resource     Resource       `json:"resource"`
}

type Interactions struct {
	// createdAt string
	// customerFields map[string]string
	Id       string   `json:"id"`
	Quantity int      `json:"quantity"`
	Resource Resource `json:"resource"`
}

type Resource struct {
	Active   bool     `json:"active"`
	Category Category `json:"category"`

	CategoryId string `json:"categoryId"`
	Id         string `json:"id"`
	Identifier string `json:"identifier"`
	Name       string `json:"name"`
	Uom        UOM    `json:"uom"`
	UomId      string `json:"uomId"`
}

type Category struct {
	Active       bool   `json:"active"`
	Description  string `json:"description"`
	Id           string `json:"id"`
	Name         string `json:"name"`
	Organization string `json:"organization"`
}

type UOM struct {
	Abbreviation string `json:"abbreviation"`
	Active       bool   `json:"active"`
	CreatedAt    string `json:"createdAt"`
	Description  string `json:"description"`
	Id           string `json:"id"`
	IsTime       bool   `json:"isTime"`
	Organization string `json:"organisation"`
}

func parseJson() (resp string) {
	jsonobj := `[{"accessControl":{"docketbook.interview.process@6220972.on.docketbook.io":["creator","editor"],"stefan+interview@docketbook.com.au":["creator","editor"]},"additionalContacts":[],"attachments":[],"author":"39d12620-a6d6-4db1-81bb-fc74ea8aca52","authorName":"Stefan Interview Test","createdAt":"2022-09-07T04:17:53.828Z","date":"2022-09-07","docketNumber":"DOC-D-0","id":"32603b07-d46f-4935-a070-993be281e970","organisation":"da07c78c-8881-4627-b15f-3dd67fbdb1bc","organisationGroup":"053a217d-3e7d-4034-99c4-d731a456df53","resources":[{"id":"8b8d5afa-fd33-4ab3-b768-6a5f8659eee7","interactions":[{"createdAt":"2022-09-07T04:26:02.624Z","customFields":{},"id":"c0c1b843-8a7f-45bc-b9f6-19b450cf60f4","interactionName":"Hours Worked","interactionTemplate":{"billable":true,"fields":[],"name":"Hours Worked","refersParent":true,"xeroQuantity":false},"lat":-27.4522159,"long":153.0338269,"quantity":5,"resource":{"active":true,"category":{"active":true,"description":"Labour Items","id":"63c7e2e6-e276-416c-ad66-ea6890c6411f","name":"Labour","organisation":"da07c78c-8881-4627-b15f-3dd67fbdb1bc"},"categoryId":"63c7e2e6-e276-416c-ad66-ea6890c6411f","id":"1af35a5e-1774-47ea-a150-ba9d8f7a4dbb","identifier":"CW3","name":"CWP","organisation":"da07c78c-8881-4627-b15f-3dd67fbdb1bc","uom":{"abbreviation":"Hrs","active":true,"createdAt":"2022-09-07T04:16:05Z","description":"Hours","id":"3e2b63a2-9fab-4598-9d9f-e98d8151daae","isTime":false,"organisation":"da07c78c-8881-4627-b15f-3dd67fbdb1bc"},"uomId":"3e2b63a2-9fab-4598-9d9f-e98d8151daae"},"resourceIsParent":true}],"resource":{"active":true,"category":{"active":true,"description":"Labour Items","id":"63c7e2e6-e276-416c-ad66-ea6890c6411f","name":"Labour","organisation":"da07c78c-8881-4627-b15f-3dd67fbdb1bc"},"categoryId":"63c7e2e6-e276-416c-ad66-ea6890c6411f","id":"1af35a5e-1774-47ea-a150-ba9d8f7a4dbb","identifier":"CW3","name":"CWP","organisation":"da07c78c-8881-4627-b15f-3dd67fbdb1bc","uom":{"abbreviation":"Hrs","active":true,"createdAt":"2022-09-07T04:16:05Z","description":"Hours","id":"3e2b63a2-9fab-4598-9d9f-e98d8151daae","isTime":false,"organisation":"da07c78c-8881-4627-b15f-3dd67fbdb1bc"},"uomId":"3e2b63a2-9fab-4598-9d9f-e98d8151daae"}}],"startTime":"2022-09-06T23:25:00.000Z","status":"draft","subDockets":[],"supplierContact":{"addresses":[{"country":"Australia","designation":"Address","num":"","postCode":"4006","state":"QLD","streetName":"King Street","streetNum":"25","suburb":"Bowen Hills"}],"contactAddresses":{"docketbook.interview.process@6220972.on.docketbook.io":[]},"displayName":"Docketbook Interview Process","reassignments":[]},"template":{"active":true,"additionalContactNames":[],"attachments":[],"blocks":[{"fields":[{"active":true,"displayName":"Fit for work","id":"364c1506-cf30-4a5f-81cd-8187a801800f","mandatory":false,"organisation":"da07c78c-8881-4627-b15f-3dd67fbdb1bc","selectionFields":[],"text":"","type":"boolean"},{"active":true,"displayName":"Shift","id":"aa1a9eff-fc4f-40ab-a87a-85a27aecb216","mandatory":false,"organisation":"da07c78c-8881-4627-b15f-3dd67fbdb1bc","selectionFields":["Day","Night"],"type":"select"},{"active":true,"displayName":"Description of work","id":"f2686670-f898-4b27-a991-6979b5b2804e","mandatory":false,"organisation":"da07c78c-8881-4627-b15f-3dd67fbdb1bc","selectionFields":[],"type":"multiline"}],"id":"1d3f4026-f929-4a9b-8eff-16bb488d342a","name":"Details","type":"custom"},{"fields":[],"id":"2f3a81a8-f1cc-40fb-b916-e8b064af466c","type":"time"},{"fields":[],"id":"86323fac-9f9f-4e82-8534-963ac7ed62f9","type":"resources"},{"fields":[],"id":"2df26e4b-5ff2-43e9-afc4-922392a4e0ed","type":"signatures"}],"checklistContainers":[],"id":"419e828d-5c44-4e2c-bf87-bf60d2d4a8c6","interactionTemplates":[{"actorCategories":[{"active":true,"description":"Labour Items","id":"63c7e2e6-e276-416c-ad66-ea6890c6411f","name":"Labour","organisation":"da07c78c-8881-4627-b15f-3dd67fbdb1bc"}],"actorFromOrder":false,"billable":true,"categories":[],"childFromOrder":false,"description":"Hours Worked","fields":[],"name":"Hours Worked","refersParent":true,"xeroQuantity":false}],"minimumInteractionRequired":true,"name":"Labour Hire Docket","populateResourcesFromJob":false,"published":true,"signature":false},"tracked":false,"updateHash":"eb576830ed2819dcd53219927c5d9c9c5702b590a865e791e39f544fdaf3fc43","updatedAt":"2022-09-07T04:26:24.802Z","visibleInDraft":false}]`

	var response []Response
	err := json.Unmarshal([]byte(jsonobj), &response)
	if err != nil {
		fmt.Println((err))
	}
	x := &response
	data, _ := json.Marshal(x)
	return string(data)
}
