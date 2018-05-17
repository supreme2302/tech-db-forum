// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/supreme2302/tech-db-forum/generated/models"
)

// ThreadGetOneReader is a Reader for the ThreadGetOne structure.
type ThreadGetOneReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ThreadGetOneReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewThreadGetOneOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewThreadGetOneNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewThreadGetOneOK creates a ThreadGetOneOK with default headers values
func NewThreadGetOneOK() *ThreadGetOneOK {
	return &ThreadGetOneOK{}
}

/*ThreadGetOneOK handles this case with default header values.

Информация о ветке обсуждения.

*/
type ThreadGetOneOK struct {
	Payload *models.Thread
}

func (o *ThreadGetOneOK) Error() string {
	return fmt.Sprintf("[GET /thread/{slug_or_id}/details][%d] threadGetOneOK  %+v", 200, o.Payload)
}

func (o *ThreadGetOneOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Thread)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewThreadGetOneNotFound creates a ThreadGetOneNotFound with default headers values
func NewThreadGetOneNotFound() *ThreadGetOneNotFound {
	return &ThreadGetOneNotFound{}
}

/*ThreadGetOneNotFound handles this case with default header values.

Ветка обсуждения отсутсвует в форуме.

*/
type ThreadGetOneNotFound struct {
	Payload *models.Error
}

func (o *ThreadGetOneNotFound) Error() string {
	return fmt.Sprintf("[GET /thread/{slug_or_id}/details][%d] threadGetOneNotFound  %+v", 404, o.Payload)
}

func (o *ThreadGetOneNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
