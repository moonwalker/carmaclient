package dto

type PropertyDTO struct {
	ID                int64        `json:"id"`
	PropertyTypeID    PropertyType `json:"propertyTypeId"` // type of property
	Name              string       `json:"name"`
	Description       string       `json:"description,omitempty"`
	DataType          DataType     `json:"dataType,omitempty"`          // Data type to use for this property in Business Analytics export. Null means 'default', i.e. interpret value as a string
	IncludeInBaExport bool         `json:"includeInBaExport,omitempty"` // Whether this property should be included in Business Analytics export. Defaults to false
	DateCreated       string       `json:"dateCreated"`                 // read only, populated by server
	DateModified      string       `json:"dateModified"`                // read only, populated by server
	UserCreated       string       `json:"userCreated"`                 // read only, populated by server
	UserModified      string       `json:"userModified"`                // read only, populated by server
	StaticProperty    bool         `json:"staticProperty"`              // is this a static property
	FileHeaders       []string     `json:"fileHeaders,omitempty"`
}
