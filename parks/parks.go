package parks

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"npsparkbrowser/auth"
	"npsparkbrowser/httputils"
	"strconv"
	"strings"
)

type ParksConfig struct {
	relativePath string
	method       string
	body         io.Reader
	headers      httputils.Headers
}

var config ParksConfig

func Configure() {
	config = ParksConfig{
		relativePath: "/v1/parks",
		method:       "GET",
		body:         strings.NewReader(""),
		headers: httputils.Headers{
			"Accept": "application/json",
		},
	}
}

func GetAll() {
	data := []ParkModel{}
	start := 0

	for {
		resp, err := get(start)
		if err != nil {
			fmt.Println(err)
		}
		data = append(data, resp.Data...)
		start += len(resp.Data)
		total, err := strconv.Atoi(resp.Total)
		if err != nil {
			fmt.Println(err)
		}
		if start == total {
			fmt.Printf(`Retrieved %v records from /v1/parks`, len(data))
			break
		}
	}

}

func get(start int) (ParksResponse, error) {
	client := &http.Client{}
	req, err := buildRequest(start)
	if err != nil {
		fmt.Println(err)
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}

	resp, err := parseResponse(body)

	return resp, err
}

func buildRequest(start int) (*http.Request, error) {
	url := fmt.Sprintf(`%v%v?api_key=%v&start=%v`, httputils.BaseUrl, config.relativePath, auth.ApiKey(), start)
	req, err := http.NewRequest(config.method, url, config.body)

	return req, err
}

func parseResponse(body []byte) (ParksResponse, error) {
	var response ParksResponse
	err := json.Unmarshal(body, &response)

	return response, err
}

type ParksResponse struct {
	Total string      `json:"total"`
	Start string      `json:"start"`
	Limit string      `json:"limit"`
	Data  []ParkModel `json:"data"`
}

type ParkModel struct {
	Id             string                    `json:"id"`
	Url            string                    `json:"url"`
	FullName       string                    `json:"fullName"`
	ParkCode       string                    `json:"parkCode"`
	Description    string                    `json:"description"`
	Latitude       string                    `json:"latitude"`
	Longitude      string                    `json:"longitude"`
	LatLong        string                    `json:"latLong"`
	Activities     []TopicActivityModel      `json:"activities"`
	Topics         []TopicActivityModel      `json:"topic"`
	States         string                    `json:"states"`
	Contacts       ContactModel              `json:"contacts"`
	EntranceFees   []EntrancePassesFeesModel `json:"entranceFees"`
	EntrancePasses []EntrancePassesFeesModel `json:"entrancePasses"`
	Fees           []string                  `json:"fees"`
	DirectionsInfo string                    `json:"directionsInfo"`
	DirectionsUrl  string                    `json:"directionsUrl"`
	OperatingHours []OperatingHoursModel     `json:"operatingHours"`
	Addresses      []AddressModel            `json:"addresses"`
	Images         []ImageModel              `json:"images"`
	WeatherInfo    string                    `json:"weatherInfo"`
	Name           string                    `json:"name"`
	Designation    string                    `json:"designation"`
	RelevanceScore float32                   `json:"relevanceScore"`
}

type ImageModel struct {
	Credit  string `json:"credit"`
	Title   string `json:"title"`
	AltText string `json:"altText"`
	Caption string `json:"caption"`
	Url     string `json:"url"`
}

type AddressModel struct {
	PostalCode            string `json:"postalCode"`
	City                  string `json:"city"`
	StateCode             string `json:"stateCode"`
	CountryCode           string `json:"countryCode"`
	ProvinceTerritoryCode string `json:"provinceTerritoryCode"`
	Line1                 string `json:"line1"`
	Type                  string `json:"type"`
	Line2                 string `json:"line2"`
	Line3                 string `json:"line3"`
}

type OperatingHoursModel struct {
	Exceptions    []OperatingHoursExceptionModel `json:"exceptions"`
	Description   string                         `json:"description"`
	StandardHours HoursMapModel                  `json:"standardHours"`
	Name          string                         `json:"name"`
}

type HoursMapModel map[string]string

type OperatingHoursExceptionModel struct {
	ExceptionHours HoursMapModel `json:"exceptionHours"`
	StartDate      string        `json:"startDate"`
	Name           string        `json:"name"`
	EndDate        string        `json:"endDate"`
}

type ContactModel struct {
	PhoneNumbers   []PhoneNumberModel  `json:"phoneNumbers"`
	EmailAddresses []EmailAddressModel `json:"emailAddresses"`
}

type PhoneNumberModel struct {
	PhoneNumber string `json:"phoneNumber"`
	Description string `json:"description"`
	Extension   string `json:"extension"`
	Type        string `json:"type"`
}

type EmailAddressModel struct {
	Description  string `json:"description"`
	EmailAddress string `json:"emailAddress"`
}

type TopicActivityModel struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type EntrancePassesFeesModel struct {
	Cost        string `json:"cost"`
	Description string `json:"description"`
	Title       string `json:"title"`
}
