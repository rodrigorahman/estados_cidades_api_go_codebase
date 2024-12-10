package dto

type BrasilApiStateResponse struct {
	Sigla string `json:"sigla"`
	Nome  string `json:"nome"`
}

type BrasilApiCityResponse struct {
	Nome string `json:"nome"`
	Ibge string `json:"codigo_ibge"`
}
