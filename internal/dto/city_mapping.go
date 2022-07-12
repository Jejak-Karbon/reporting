package dto

import "github.com/born2ngopi/alterra/basic-echo-mvc/internal/model"

type CityMapping struct {
	CityName   			  string `json:"city_name"`
	AmountAbsorption      float32 `json:"amount_absorption"`
}

type CityMappingResponse struct {
	UserProductCarbonAbsorption model.UserProductCarbonAbsorption
	User UserProfileResponse
}
