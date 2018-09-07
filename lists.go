package carmaclient

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/moonwalker/carmaclient/dto"
	"net/http"
	"net/url"
)

type ListsService service

func (s *ListsService) GetContact(listID int64, originalID string) (*dto.ContactDTO, error) {
	response := s.client.carmaRequest(fmt.Sprintf("lists/%d/contacts/%s", listID, url.QueryEscape(originalID)), http.MethodGet, nil)
	if response.err != nil {
		return nil, response.err
	}

	if response.statusCode == http.StatusNotFound {
		return nil, errors.New(http.StatusText(response.statusCode))
	}

	responseDTO := &dto.ContactDTO{}

	err := json.Unmarshal(response.data, responseDTO)
	if err != nil {
		return nil, err
	}

	return responseDTO, nil
}

func (s *ListsService) PutContactUpdate(listID int64, originalID string, contact dto.ContactDTO) (*dto.ContactDTO, error) {
	response := s.client.carmaRequest(fmt.Sprintf("lists/%d/contacts/%s/update?force=true", listID, url.QueryEscape(originalID)), http.MethodPut, contact)
	if response.err != nil {
		return nil, response.err
	}

	responseDTO := &dto.ContactDTO{}

	err := json.Unmarshal(response.data, responseDTO)
	if err != nil {
		return nil, err
	}

	return responseDTO, nil
}
