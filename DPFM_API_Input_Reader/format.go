package dpfm_api_input_reader

import (
	"data-platform-api-division-creates-rmq-kube/DPFM_API_Caller/requests"
)

func (sdc *SDC) ConvertToDivision() *requests.Division {
	data := sdc.Division
	return &requests.Division{
		Division: data.Division,
	}
}

func (sdc *SDC) ConvertToDivisionText() *requests.DivisionText {
	dataDivision := sdc.Division
	data := sdc.Division.DivisionText
	return &requests.DivisionText{
		Division:     dataDivision.Division,
		Language:     data.Language,
		DivisionName: data.DivisionName,
	}
}
