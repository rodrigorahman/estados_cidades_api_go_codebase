package dto

type StateResponse struct {
	Acronym string `json:"sigla"`
	Name    string `json:"nome"`
}

type CityResponse struct {
	Name string `json:"nome"`
	Ibge string `json:"codigo_ibge"`
}
