// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/epilot-dev/terraform-provider-epilot-custom-variable/internal/sdk/internal/utils"
)

// CustomVariablesSearchParamsType - Variable type
type CustomVariablesSearchParamsType string

const (
	CustomVariablesSearchParamsTypeOrderTable CustomVariablesSearchParamsType = "order_table"
	CustomVariablesSearchParamsTypeCustom     CustomVariablesSearchParamsType = "custom"
)

func (e CustomVariablesSearchParamsType) ToPointer() *CustomVariablesSearchParamsType {
	return &e
}
func (e *CustomVariablesSearchParamsType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "order_table":
		fallthrough
	case "custom":
		*e = CustomVariablesSearchParamsType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CustomVariablesSearchParamsType: %v", v)
	}
}

type CustomVariablesSearchParams struct {
	// Variable type
	Type *CustomVariablesSearchParamsType `json:"type,omitempty"`
	// The tags of custom variable
	Tags []string `json:"tags,omitempty"`
	// Search string
	Query *string `json:"query,omitempty"`
	From  *int64  `default:"0" json:"from"`
	Size  *int64  `default:"25" json:"size"`
	// Sort by field
	SortBy *string `json:"sort_by,omitempty"`
	// Fields to return
	Fields []string `json:"fields,omitempty"`
}

func (c CustomVariablesSearchParams) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CustomVariablesSearchParams) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *CustomVariablesSearchParams) GetType() *CustomVariablesSearchParamsType {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *CustomVariablesSearchParams) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *CustomVariablesSearchParams) GetQuery() *string {
	if o == nil {
		return nil
	}
	return o.Query
}

func (o *CustomVariablesSearchParams) GetFrom() *int64 {
	if o == nil {
		return nil
	}
	return o.From
}

func (o *CustomVariablesSearchParams) GetSize() *int64 {
	if o == nil {
		return nil
	}
	return o.Size
}

func (o *CustomVariablesSearchParams) GetSortBy() *string {
	if o == nil {
		return nil
	}
	return o.SortBy
}

func (o *CustomVariablesSearchParams) GetFields() []string {
	if o == nil {
		return nil
	}
	return o.Fields
}
