// Code generated by go-swagger; DO NOT EDIT.

package scan

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/neuvector/neuvector-nexus-iq/neuvector/models"
)

// GetScanImageIDReader is a Reader for the GetScanImageID structure.
type GetScanImageIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetScanImageIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetScanImageIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetScanImageIDOK creates a GetScanImageIDOK with default headers values
func NewGetScanImageIDOK() *GetScanImageIDOK {
	return &GetScanImageIDOK{}
}

/*GetScanImageIDOK handles this case with default header values.

Success
*/
type GetScanImageIDOK struct {
	Payload *models.RESTScanReportData
}

func (o *GetScanImageIDOK) Error() string {
	return fmt.Sprintf("[GET /scan/image/{id}][%d] getScanImageIdOK  %+v", 200, o.Payload)
}

func (o *GetScanImageIDOK) GetPayload() *models.RESTScanReportData {
	return o.Payload
}

func (o *GetScanImageIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RESTScanReportData)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
