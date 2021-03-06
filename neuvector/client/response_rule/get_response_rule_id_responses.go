// Code generated by go-swagger; DO NOT EDIT.

package response_rule

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/neuvector/neuvector-nexus-iq/neuvector/models"
)

// GetResponseRuleIDReader is a Reader for the GetResponseRuleID structure.
type GetResponseRuleIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetResponseRuleIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetResponseRuleIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetResponseRuleIDOK creates a GetResponseRuleIDOK with default headers values
func NewGetResponseRuleIDOK() *GetResponseRuleIDOK {
	return &GetResponseRuleIDOK{}
}

/*GetResponseRuleIDOK handles this case with default header values.

Success
*/
type GetResponseRuleIDOK struct {
	Payload *models.RESTResponseRuleData
}

func (o *GetResponseRuleIDOK) Error() string {
	return fmt.Sprintf("[GET /response/rule/{id}][%d] getResponseRuleIdOK  %+v", 200, o.Payload)
}

func (o *GetResponseRuleIDOK) GetPayload() *models.RESTResponseRuleData {
	return o.Payload
}

func (o *GetResponseRuleIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RESTResponseRuleData)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
