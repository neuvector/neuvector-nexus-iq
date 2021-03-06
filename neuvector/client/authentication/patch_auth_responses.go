// Code generated by go-swagger; DO NOT EDIT.

package authentication

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// PatchAuthReader is a Reader for the PatchAuth structure.
type PatchAuthReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchAuthReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchAuthOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPatchAuthOK creates a PatchAuthOK with default headers values
func NewPatchAuthOK() *PatchAuthOK {
	return &PatchAuthOK{}
}

/*PatchAuthOK handles this case with default header values.

Success
*/
type PatchAuthOK struct {
}

func (o *PatchAuthOK) Error() string {
	return fmt.Sprintf("[PATCH /auth][%d] patchAuthOK ", 200)
}

func (o *PatchAuthOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
