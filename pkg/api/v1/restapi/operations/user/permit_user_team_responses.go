// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/gomematic/gomematic-api/pkg/api/v1/models"
)

// PermitUserTeamOKCode is the HTTP code returned for type PermitUserTeamOK
const PermitUserTeamOKCode int = 200

/*PermitUserTeamOK Plain success message

swagger:response permitUserTeamOK
*/
type PermitUserTeamOK struct {

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewPermitUserTeamOK creates PermitUserTeamOK with default headers values
func NewPermitUserTeamOK() *PermitUserTeamOK {

	return &PermitUserTeamOK{}
}

// WithPayload adds the payload to the permit user team o k response
func (o *PermitUserTeamOK) WithPayload(payload *models.GeneralError) *PermitUserTeamOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the permit user team o k response
func (o *PermitUserTeamOK) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PermitUserTeamOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PermitUserTeamForbiddenCode is the HTTP code returned for type PermitUserTeamForbidden
const PermitUserTeamForbiddenCode int = 403

/*PermitUserTeamForbidden User is not authorized

swagger:response permitUserTeamForbidden
*/
type PermitUserTeamForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewPermitUserTeamForbidden creates PermitUserTeamForbidden with default headers values
func NewPermitUserTeamForbidden() *PermitUserTeamForbidden {

	return &PermitUserTeamForbidden{}
}

// WithPayload adds the payload to the permit user team forbidden response
func (o *PermitUserTeamForbidden) WithPayload(payload *models.GeneralError) *PermitUserTeamForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the permit user team forbidden response
func (o *PermitUserTeamForbidden) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PermitUserTeamForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PermitUserTeamPreconditionFailedCode is the HTTP code returned for type PermitUserTeamPreconditionFailed
const PermitUserTeamPreconditionFailedCode int = 412

/*PermitUserTeamPreconditionFailed Failed to parse request body

swagger:response permitUserTeamPreconditionFailed
*/
type PermitUserTeamPreconditionFailed struct {

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewPermitUserTeamPreconditionFailed creates PermitUserTeamPreconditionFailed with default headers values
func NewPermitUserTeamPreconditionFailed() *PermitUserTeamPreconditionFailed {

	return &PermitUserTeamPreconditionFailed{}
}

// WithPayload adds the payload to the permit user team precondition failed response
func (o *PermitUserTeamPreconditionFailed) WithPayload(payload *models.GeneralError) *PermitUserTeamPreconditionFailed {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the permit user team precondition failed response
func (o *PermitUserTeamPreconditionFailed) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PermitUserTeamPreconditionFailed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(412)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PermitUserTeamUnprocessableEntityCode is the HTTP code returned for type PermitUserTeamUnprocessableEntity
const PermitUserTeamUnprocessableEntityCode int = 422

/*PermitUserTeamUnprocessableEntity Team is not assigned

swagger:response permitUserTeamUnprocessableEntity
*/
type PermitUserTeamUnprocessableEntity struct {

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewPermitUserTeamUnprocessableEntity creates PermitUserTeamUnprocessableEntity with default headers values
func NewPermitUserTeamUnprocessableEntity() *PermitUserTeamUnprocessableEntity {

	return &PermitUserTeamUnprocessableEntity{}
}

// WithPayload adds the payload to the permit user team unprocessable entity response
func (o *PermitUserTeamUnprocessableEntity) WithPayload(payload *models.GeneralError) *PermitUserTeamUnprocessableEntity {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the permit user team unprocessable entity response
func (o *PermitUserTeamUnprocessableEntity) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PermitUserTeamUnprocessableEntity) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(422)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*PermitUserTeamDefault Some error unrelated to the handler

swagger:response permitUserTeamDefault
*/
type PermitUserTeamDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewPermitUserTeamDefault creates PermitUserTeamDefault with default headers values
func NewPermitUserTeamDefault(code int) *PermitUserTeamDefault {
	if code <= 0 {
		code = 500
	}

	return &PermitUserTeamDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the permit user team default response
func (o *PermitUserTeamDefault) WithStatusCode(code int) *PermitUserTeamDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the permit user team default response
func (o *PermitUserTeamDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the permit user team default response
func (o *PermitUserTeamDefault) WithPayload(payload *models.GeneralError) *PermitUserTeamDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the permit user team default response
func (o *PermitUserTeamDefault) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PermitUserTeamDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}