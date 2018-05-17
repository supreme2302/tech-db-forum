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

// UserUpdateReader is a Reader for the UserUpdate structure.
type UserUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UserUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUserUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewUserUpdateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewUserUpdateConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUserUpdateOK creates a UserUpdateOK with default headers values
func NewUserUpdateOK() *UserUpdateOK {
	return &UserUpdateOK{}
}

/*UserUpdateOK handles this case with default header values.

Актуальная информация о пользователе после изменения профиля.

*/
type UserUpdateOK struct {
	Payload *models.User
}

func (o *UserUpdateOK) Error() string {
	return fmt.Sprintf("[POST /user/{nickname}/profile][%d] userUpdateOK  %+v", 200, o.Payload)
}

func (o *UserUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.User)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserUpdateNotFound creates a UserUpdateNotFound with default headers values
func NewUserUpdateNotFound() *UserUpdateNotFound {
	return &UserUpdateNotFound{}
}

/*UserUpdateNotFound handles this case with default header values.

Пользователь отсутсвует в системе.

*/
type UserUpdateNotFound struct {
	Payload *models.Error
}

func (o *UserUpdateNotFound) Error() string {
	return fmt.Sprintf("[POST /user/{nickname}/profile][%d] userUpdateNotFound  %+v", 404, o.Payload)
}

func (o *UserUpdateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserUpdateConflict creates a UserUpdateConflict with default headers values
func NewUserUpdateConflict() *UserUpdateConflict {
	return &UserUpdateConflict{}
}

/*UserUpdateConflict handles this case with default header values.

Новые данные профиля пользователя конфликтуют с имеющимися пользователями.

*/
type UserUpdateConflict struct {
	Payload *models.Error
}

func (o *UserUpdateConflict) Error() string {
	return fmt.Sprintf("[POST /user/{nickname}/profile][%d] userUpdateConflict  %+v", 409, o.Payload)
}

func (o *UserUpdateConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
