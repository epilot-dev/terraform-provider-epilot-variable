// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/epilot-dev/terraform-provider-epilot-custom-variable/internal/sdk/models/shared"
	"net/http"
)

type GetBluePrintTableConfigResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Success
	CustomVariable *shared.CustomVariable
}

func (o *GetBluePrintTableConfigResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetBluePrintTableConfigResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetBluePrintTableConfigResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetBluePrintTableConfigResponse) GetCustomVariable() *shared.CustomVariable {
	if o == nil {
		return nil
	}
	return o.CustomVariable
}
