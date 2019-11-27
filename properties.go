package carmaclient

import (
	"encoding/json"
	"net/http"

	"github.com/moonwalker/carmaclient/dto"
)

type PropertiesService service

func (s *PropertiesService) GetProperties() (*[]dto.PropertyDTO, error) {
	response := s.client.carmaRequest("properties", http.MethodGet, nil)
	if response.err != nil {
		return nil, response.err
	}

	responseDTO := &[]dto.PropertyDTO{}

	err := json.Unmarshal(response.data, responseDTO)
	if err != nil {
		return nil, err
	}

	return responseDTO, nil
}
