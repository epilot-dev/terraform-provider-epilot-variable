// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// CustomVariableType - Custom variable type
type CustomVariableType string

const (
	CustomVariableTypeOrderTable  CustomVariableType = "order_table"
	CustomVariableTypeCustom      CustomVariableType = "custom"
	CustomVariableTypeJourneyLink CustomVariableType = "journey_link"
)

func (e CustomVariableType) ToPointer() *CustomVariableType {
	return &e
}
func (e *CustomVariableType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "order_table":
		fallthrough
	case "custom":
		fallthrough
	case "journey_link":
		*e = CustomVariableType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CustomVariableType: %v", v)
	}
}

type CustomVariable struct {
	// ID
	ID *string `json:"id,omitempty"`
	// Custom variable type
	Type *CustomVariableType `json:"type,omitempty"`
	// Custom variable name
	Name *string `json:"name,omitempty"`
	// The key which is used for Handlebar variable syntax {{"{{"}}key}}
	Key *string `json:"key,omitempty"`
	// The tags of custom variable
	Tags []string `json:"_tags,omitempty"`
	// The helper function parameter's names
	HelperParams []string `json:"helper_params,omitempty"`
	// The helper function logic
	HelperLogic *string `json:"helper_logic,omitempty"`
	// Handlebar template that used to generate the variable content
	Template *string `json:"template,omitempty"`
	// Creation time
	CreatedAt *string `json:"created_at,omitempty"`
	// Created by
	CreatedBy *string `json:"created_by,omitempty"`
	// Last update time
	UpdatedAt *string `json:"updated_at,omitempty"`
	// Updated by
	UpdatedBy *string `json:"updated_by,omitempty"`
	Config    any     `json:"config,omitempty"`
}

func (o *CustomVariable) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *CustomVariable) GetType() *CustomVariableType {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *CustomVariable) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *CustomVariable) GetKey() *string {
	if o == nil {
		return nil
	}
	return o.Key
}

func (o *CustomVariable) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *CustomVariable) GetHelperParams() []string {
	if o == nil {
		return nil
	}
	return o.HelperParams
}

func (o *CustomVariable) GetHelperLogic() *string {
	if o == nil {
		return nil
	}
	return o.HelperLogic
}

func (o *CustomVariable) GetTemplate() *string {
	if o == nil {
		return nil
	}
	return o.Template
}

func (o *CustomVariable) GetCreatedAt() *string {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *CustomVariable) GetCreatedBy() *string {
	if o == nil {
		return nil
	}
	return o.CreatedBy
}

func (o *CustomVariable) GetUpdatedAt() *string {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

func (o *CustomVariable) GetUpdatedBy() *string {
	if o == nil {
		return nil
	}
	return o.UpdatedBy
}

func (o *CustomVariable) GetConfig() any {
	if o == nil {
		return nil
	}
	return o.Config
}
