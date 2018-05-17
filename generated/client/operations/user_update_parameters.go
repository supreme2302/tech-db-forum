// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/supreme2302/tech-db-forum/generated/models"
)

// NewUserUpdateParams creates a new UserUpdateParams object
// with the default values initialized.
func NewUserUpdateParams() *UserUpdateParams {
	var ()
	return &UserUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUserUpdateParamsWithTimeout creates a new UserUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUserUpdateParamsWithTimeout(timeout time.Duration) *UserUpdateParams {
	var ()
	return &UserUpdateParams{

		timeout: timeout,
	}
}

// NewUserUpdateParamsWithContext creates a new UserUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewUserUpdateParamsWithContext(ctx context.Context) *UserUpdateParams {
	var ()
	return &UserUpdateParams{

		Context: ctx,
	}
}

// NewUserUpdateParamsWithHTTPClient creates a new UserUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUserUpdateParamsWithHTTPClient(client *http.Client) *UserUpdateParams {
	var ()
	return &UserUpdateParams{
		HTTPClient: client,
	}
}

/*UserUpdateParams contains all the parameters to send to the API endpoint
for the user update operation typically these are written to a http.Request
*/
type UserUpdateParams struct {

	/*Nickname
	  Идентификатор пользователя.

	*/
	Nickname string
	/*Profile
	  Изменения профиля пользователя.

	*/
	Profile *models.UserUpdate

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the user update params
func (o *UserUpdateParams) WithTimeout(timeout time.Duration) *UserUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the user update params
func (o *UserUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the user update params
func (o *UserUpdateParams) WithContext(ctx context.Context) *UserUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the user update params
func (o *UserUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the user update params
func (o *UserUpdateParams) WithHTTPClient(client *http.Client) *UserUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the user update params
func (o *UserUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNickname adds the nickname to the user update params
func (o *UserUpdateParams) WithNickname(nickname string) *UserUpdateParams {
	o.SetNickname(nickname)
	return o
}

// SetNickname adds the nickname to the user update params
func (o *UserUpdateParams) SetNickname(nickname string) {
	o.Nickname = nickname
}

// WithProfile adds the profile to the user update params
func (o *UserUpdateParams) WithProfile(profile *models.UserUpdate) *UserUpdateParams {
	o.SetProfile(profile)
	return o
}

// SetProfile adds the profile to the user update params
func (o *UserUpdateParams) SetProfile(profile *models.UserUpdate) {
	o.Profile = profile
}

// WriteToRequest writes these params to a swagger request
func (o *UserUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param nickname
	if err := r.SetPathParam("nickname", o.Nickname); err != nil {
		return err
	}

	if o.Profile != nil {
		if err := r.SetBodyParam(o.Profile); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
