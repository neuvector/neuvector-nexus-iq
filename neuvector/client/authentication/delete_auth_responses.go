// Code generated by go-swagger; DO NOT EDIT.

package authentication

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/neuvector/neuvector-nexus-iq/neuvector/models"
)

// DeleteAuthReader is a Reader for the DeleteAuth structure.
type DeleteAuthReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteAuthReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteAuthOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 408:
		result := NewDeleteAuthRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteAuthOK creates a DeleteAuthOK with default headers values
func NewDeleteAuthOK() *DeleteAuthOK {
	return &DeleteAuthOK{}
}

/*DeleteAuthOK handles this case with default header values.

Success
*/
type DeleteAuthOK struct {
}

func (o *DeleteAuthOK) Error() string {
	return fmt.Sprintf("[DELETE /auth][%d] deleteAuthOK ", 200)
}

func (o *DeleteAuthOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteAuthRequestTimeout creates a DeleteAuthRequestTimeout with default headers values
func NewDeleteAuthRequestTimeout() *DeleteAuthRequestTimeout {
	return &DeleteAuthRequestTimeout{}
}

/*DeleteAuthRequestTimeout handles this case with default header values.

Authentication failed
*/
type DeleteAuthRequestTimeout struct {
	Payload *models.RESTError
}

func (o *DeleteAuthRequestTimeout) Error() string {
	return fmt.Sprintf("[DELETE /auth][%d] deleteAuthRequestTimeout  %+v", 408, o.Payload)
}

func (o *DeleteAuthRequestTimeout) GetPayload() *models.RESTError {
	return o.Payload
}

func (o *DeleteAuthRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RESTError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
