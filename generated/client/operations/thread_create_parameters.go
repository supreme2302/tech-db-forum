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

// NewThreadCreateParams creates a new ThreadCreateParams object
// with the default values initialized.
func NewThreadCreateParams() *ThreadCreateParams {
	var ()
	return &ThreadCreateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewThreadCreateParamsWithTimeout creates a new ThreadCreateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewThreadCreateParamsWithTimeout(timeout time.Duration) *ThreadCreateParams {
	var ()
	return &ThreadCreateParams{

		timeout: timeout,
	}
}

// NewThreadCreateParamsWithContext creates a new ThreadCreateParams object
// with the default values initialized, and the ability to set a context for a request
func NewThreadCreateParamsWithContext(ctx context.Context) *ThreadCreateParams {
	var ()
	return &ThreadCreateParams{

		Context: ctx,
	}
}

// NewThreadCreateParamsWithHTTPClient creates a new ThreadCreateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewThreadCreateParamsWithHTTPClient(client *http.Client) *ThreadCreateParams {
	var ()
	return &ThreadCreateParams{
		HTTPClient: client,
	}
}

/*ThreadCreateParams contains all the parameters to send to the API endpoint
for the thread create operation typically these are written to a http.Request
*/
type ThreadCreateParams struct {

	/*Slug
	  Идентификатор форума.

	*/
	Slug string
	/*Thread
	  Данные ветки обсуждения.

	*/
	Thread *models.Thread

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the thread create params
func (o *ThreadCreateParams) WithTimeout(timeout time.Duration) *ThreadCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the thread create params
func (o *ThreadCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the thread create params
func (o *ThreadCreateParams) WithContext(ctx context.Context) *ThreadCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the thread create params
func (o *ThreadCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the thread create params
func (o *ThreadCreateParams) WithHTTPClient(client *http.Client) *ThreadCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the thread create params
func (o *ThreadCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSlug adds the slug to the thread create params
func (o *ThreadCreateParams) WithSlug(slug string) *ThreadCreateParams {
	o.SetSlug(slug)
	return o
}

// SetSlug adds the slug to the thread create params
func (o *ThreadCreateParams) SetSlug(slug string) {
	o.Slug = slug
}

// WithThread adds the thread to the thread create params
func (o *ThreadCreateParams) WithThread(thread *models.Thread) *ThreadCreateParams {
	o.SetThread(thread)
	return o
}

// SetThread adds the thread to the thread create params
func (o *ThreadCreateParams) SetThread(thread *models.Thread) {
	o.Thread = thread
}

// WriteToRequest writes these params to a swagger request
func (o *ThreadCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param slug
	if err := r.SetPathParam("slug", o.Slug); err != nil {
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
