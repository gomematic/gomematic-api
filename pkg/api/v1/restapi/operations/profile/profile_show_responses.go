// Code generated by go-swagger; DO NOT EDIT.

package profile

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/gomematic/gomematic-api/pkg/api/v1/models"
)

// ProfileShowOKCode is the HTTP code returned for type ProfileShowOK
const ProfileShowOKCode int = 200

/*ProfileShowOK The current profile data

swagger:response profileShowOK
*/
type ProfileShowOK struct {

	/*
	  In: Body
	*/
	Payload *models.Profile `json:"body,omitempty"`
}

// NewProfileShowOK creates ProfileShowOK with default headers values
func NewProfileShowOK() *ProfileShowOK {

	return &ProfileShowOK{}
}

// WithPayload adds the payload to the profile show o k response
func (o *ProfileShowOK) WithPayload(payload *models.Profile) *ProfileShowOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the profile show o k response
func (o *ProfileShowOK) SetPayload(payload *models.Profile) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ProfileShowOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ProfileShowForbiddenCode is the HTTP code returned for type ProfileShowForbidden
const ProfileShowForbiddenCode int = 403

/*ProfileShowForbidden User is not authorized

swagger:response profileShowForbidden
*/
type ProfileShowForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewProfileShowForbidden creates ProfileShowForbidden with default headers values
func NewProfileShowForbidden() *ProfileShowForbidden {

	return &ProfileShowForbidden{}
}

// WithPayload adds the payload to the profile show forbidden response
func (o *ProfileShowForbidden) WithPayload(payload *models.GeneralError) *ProfileShowForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the profile show forbidden response
func (o *ProfileShowForbidden) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ProfileShowForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*ProfileShowDefault Some error unrelated to the handler

swagger:response profileShowDefault
*/
type ProfileShowDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewProfileShowDefault creates ProfileShowDefault with default headers values
func NewProfileShowDefault(code int) *ProfileShowDefault {
	if code <= 0 {
		code = 500
	}

	return &ProfileShowDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the profile show default response
func (o *ProfileShowDefault) WithStatusCode(code int) *ProfileShowDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the profile show default response
func (o *ProfileShowDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the profile show default response
func (o *ProfileShowDefault) WithPayload(payload *models.GeneralError) *ProfileShowDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the profile show default response
func (o *ProfileShowDefault) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ProfileShowDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
