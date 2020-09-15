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

// GetScanRegistryNameImageIDReader is a Reader for the GetScanRegistryNameImageID structure.
type GetScanRegistryNameImageIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetScanRegistryNameImageIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetScanRegistryNameImageIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetScanRegistryNameImageIDOK creates a GetScanRegistryNameImageIDOK with default headers values
func NewGetScanRegistryNameImageIDOK() *GetScanRegistryNameImageIDOK {
	return &GetScanRegistryNameImageIDOK{}
}

/*GetScanRegistryNameImageIDOK handles this case with default header values.

Success
*/
type GetScanRegistryNameImageIDOK struct {
	Payload *models.RESTScanReportData
}

func (o *GetScanRegistryNameImageIDOK) Error() string {
	return fmt.Sprintf("[GET /scan/registry/{name}/image/{id}][%d] getScanRegistryNameImageIdOK  %+v", 200, o.Payload)
}

func (o *GetScanRegistryNameImageIDOK) GetPayload() *models.RESTScanReportData {
	return o.Payload
}

func (o *GetScanRegistryNameImageIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RESTScanReportData)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
