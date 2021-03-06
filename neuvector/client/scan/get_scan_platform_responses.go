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

// GetScanPlatformReader is a Reader for the GetScanPlatform structure.
type GetScanPlatformReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetScanPlatformReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetScanPlatformOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetScanPlatformOK creates a GetScanPlatformOK with default headers values
func NewGetScanPlatformOK() *GetScanPlatformOK {
	return &GetScanPlatformOK{}
}

/*GetScanPlatformOK handles this case with default header values.

Success
*/
type GetScanPlatformOK struct {
	Payload *models.RESTScanPlatformSummaryData
}

func (o *GetScanPlatformOK) Error() string {
	return fmt.Sprintf("[GET /scan/platform][%d] getScanPlatformOK  %+v", 200, o.Payload)
}

func (o *GetScanPlatformOK) GetPayload() *models.RESTScanPlatformSummaryData {
	return o.Payload
}

func (o *GetScanPlatformOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RESTScanPlatformSummaryData)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
