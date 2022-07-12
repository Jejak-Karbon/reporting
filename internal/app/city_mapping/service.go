package city_mapping

import (
	"sort"
	"log"
	"context"
	"encoding/json"
	"net/http"

	"github.com/born2ngopi/alterra/basic-echo-mvc/internal/factory"
	"github.com/born2ngopi/alterra/basic-echo-mvc/internal/dto"
	_ "github.com/born2ngopi/alterra/basic-echo-mvc/internal/model"
	_ "github.com/born2ngopi/alterra/basic-echo-mvc/pkg/constant"
	_ "github.com/born2ngopi/alterra/basic-echo-mvc/pkg/util/response"
)

type service struct {
	
}

type Service interface {
	Find(ctx context.Context) ([]dto.CityMapping,error)
}

func NewService(f *factory.Factory) Service {
	return &service{
		
	}
}

func (s *service) Find(ctx context.Context) ([]dto.CityMapping,error) {

	var data []dto.CityMapping
	url := "http://localhost:3000/user_product_carbon_absorption"

	var client = &http.Client{}
	request, _ := http.NewRequest("GET", url, nil)

	response, err := client.Do(request)
	if err != nil {
		return data,err
	}
	
	defer response.Body.Close()

	var result map[string][]dto.CityMappingResponse

	json.NewDecoder(response.Body).Decode(&result)

	for _,value := range  result["data"]{

		// check if cityname exitst
		check,index := checkSlice(value.User.CityID,data)
		log.Println(index)

		if check.CityName == ""{
					
			data = append(data,dto.CityMapping{
					CityName : value.User.CityID,
					AmountAbsorption : value.UserProductCarbonAbsorption.ProductCarbonAbsorption.Amount,
				},
			)

		}else{
			amount_before := check.AmountAbsorption
			data = RemoveIndex(data,index)
			data = append(data,dto.CityMapping{
				CityName : check.CityName,
				AmountAbsorption : check.AmountAbsorption + amount_before,
			},
		)
		}

	}

	sort.Slice(data, func(i, j int) bool {
	return data[i].AmountAbsorption > data[j].AmountAbsorption
	})
	
	return data[:3],nil

}

func RemoveIndex(s []dto.CityMapping, index int) []dto.CityMapping {
	return append(s[:index], s[index+1:]...)
}

func checkSlice(CityName string, data[]dto.CityMapping) (dto.CityMapping,int){
	var x dto.CityMapping
	index := -1

	for i,val := range(data){
		if val.CityName == CityName{
			x.CityName = val.CityName
			x.AmountAbsorption = val.AmountAbsorption
			index = i
			break
		}
	}

	return x,index

}
