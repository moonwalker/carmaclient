package carmaclient

import (
	"encoding/json"
	"net/http"
	"fmt"
	"errors"

	"github.com/moonwalker/carmaclient/dto"
)

type TriggersService service

func (s *TriggersService) GetTrigger(triggerID int64) (*dto.TriggerDTO, error) {
	response := s.client.carmaRequest(fmt.Sprintf("triggers/%d", triggerID), http.MethodGet, nil)
	if response.err != nil {
		return nil, response.err
	}

	responseDTO := &dto.TriggerDTO{}

	err := json.Unmarshal(response.data, responseDTO)
	if err != nil {
		return nil, err
	}

	return responseDTO, nil
}

func (s *TriggersService) GetTriggers() (*[]dto.TriggerDTO, error) {
	response := s.client.carmaRequest("triggers", http.MethodGet, nil)
	if response.err != nil {
		return nil, response.err
	}

	responseDTO := &[]dto.TriggerDTO{}

	err := json.Unmarshal(response.data, responseDTO)
	if err != nil {
		return nil, err
	}

	return responseDTO, nil
}

func (s *TriggersService) PostTriggersMessages(triggerID int64, triggerDTO dto.SendTriggerDTO) (*dto.MessageDTO, error) {
	response := s.client.carmaRequest(fmt.Sprintf("triggers/%d/messages", triggerID), http.MethodPost, triggerDTO)
	if response.err != nil {
		return nil, response.err
	}

	if response.statusCode != http.StatusOK {
		return nil, errors.New(http.StatusText(response.statusCode))
	}

	responseDTO := &dto.MessageDTO{}

	err := json.Unmarshal(response.data, responseDTO)
	if err != nil {
		return nil, err
	}

	return responseDTO, nil
}
