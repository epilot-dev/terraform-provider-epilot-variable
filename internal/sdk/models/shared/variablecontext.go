// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type VariableContext struct {
	UnsubscribeURL *string        `json:"unsubscribe_url,omitempty"`
	Main           map[string]any `json:"main,omitempty"`
	Contact        map[string]any `json:"contact,omitempty"`
	Brand          map[string]any `json:"brand,omitempty"`
}

func (o *VariableContext) GetUnsubscribeURL() *string {
	if o == nil {
		return nil
	}
	return o.UnsubscribeURL
}

func (o *VariableContext) GetMain() map[string]any {
	if o == nil {
		return nil
	}
	return o.Main
}

func (o *VariableContext) GetContact() map[string]any {
	if o == nil {
		return nil
	}
	return o.Contact
}

func (o *VariableContext) GetBrand() map[string]any {
	if o == nil {
		return nil
	}
	return o.Brand
}
