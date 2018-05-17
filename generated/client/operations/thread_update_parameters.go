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

// NewThreadUpdateParams creates a new ThreadUpdateParams object
// with the default values initialized.
func NewThreadUpdateParams() *ThreadUpdateParams {
	var ()
	return &ThreadUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewThreadUpdateParamsWithTimeout creates a new ThreadUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewThreadUpdateParamsWithTimeout(timeout time.Duration) *ThreadUpdateParams {
	var ()
	return &ThreadUpdateParams{

		timeout: timeout,
	}
}

// NewThreadUpdateParamsWithContext creates a new ThreadUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewThreadUpdateParamsWithContext(ctx context.Context) *ThreadUpdateParams {
	var ()
	return &ThreadUpdateParams{

		Context: ctx,
	}
}

// NewThreadUpdateParamsWithHTTPClient creates a new ThreadUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewThreadUpdateParamsWithHTTPClient(client *http.Client) *ThreadUpdateParams {
	var ()
	return &ThreadUpdateParams{
		HTTPClient: client,
	}
}

/*ThreadUpdateParams contains all the parameters to send to the API endpoint
for the thread update operation typically these are written to a http.Request
*/
type ThreadUpdateParams struct {

	/*SlugOrID
	  Идентификатор ветки обсуждения.

	*/
	SlugOrID string
	/*Thread
	  Данные ветки обсуждения.

	*/
	Thread *models.ThreadUpdate

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the thread update params
func (o *ThreadUpdateParams) WithTimeout(timeout time.Duration) *ThreadUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the thread update params
func (o *ThreadUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the thread update params
func (o *ThreadUpdateParams) WithContext(ctx context.Context) *ThreadUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the thread update params
func (o *ThreadUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the thread update params
func (o *ThreadUpdateParams) WithHTTPClient(client *http.Client) *ThreadUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the thread update params
func (o *ThreadUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSlugOrID adds the slugOrID to the thread update params
func (o *ThreadUpdateParams) WithSlugOrID(slugOrID string) *ThreadUpdateParams {
	o.SetSlugOrID(slugOrID)
	return o
}

// SetSlugOrID adds the slugOrId to the thread update params
func (o *ThreadUpdateParams) SetSlugOrID(slugOrID string) {
	o.SlugOrID = slugOrID
}

// WithThread adds the thread to the thread update params
func (o *ThreadUpdateParams) WithThread(thread *models.ThreadUpdate) *ThreadUpdateParams {
	o.SetThread(thread)
	return o
}

// SetThread adds the thread to the thread update params
func (o *ThreadUpdateParams) SetThread(thread *models.ThreadUpdate) {
	o.Thread = thread
}

// WriteToRequest writes these params to a swagger request
func (o *ThreadUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param slug_or_id
	if err := r.SetPathParam("slug_or_id", o.SlugOrID); err != nil {
		return err
	}

	if o.Thread != nil {
		if err := r.SetBodyParam(o.Thread); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
