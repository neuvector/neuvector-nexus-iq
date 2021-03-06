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

// GetResponseWorkloadRulesIDReader is a Reader for the GetResponseWorkloadRulesID structure.
type GetResponseWorkloadRulesIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetResponseWorkloadRulesIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetResponseWorkloadRulesIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetResponseWorkloadRulesIDOK creates a GetResponseWorkloadRulesIDOK with default headers values
func NewGetResponseWorkloadRulesIDOK() *GetResponseWorkloadRulesIDOK {
	return &GetResponseWorkloadRulesIDOK{}
}

/*GetResponseWorkloadRulesIDOK handles this case with default header values.

Success
*/
type GetResponseWorkloadRulesIDOK struct {
	Payload *models.RESTResponseRulesData
}

func (o *GetResponseWorkloadRulesIDOK) Error() string {
	return fmt.Sprintf("[GET /response/workload_rules/{id}][%d] getResponseWorkloadRulesIdOK  %+v", 200, o.Payload)
}

func (o *GetResponseWorkloadRulesIDOK) GetPayload() *models.RESTResponseRulesData {
	return o.Payload
}

func (o *GetResponseWorkloadRulesIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RESTResponseRulesData)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
