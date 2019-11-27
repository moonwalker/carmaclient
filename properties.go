package carmaclient

import (
	"fmt"
	"net/http"
	"encoding/json"

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