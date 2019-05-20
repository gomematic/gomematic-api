// Code generated by go-swagger; DO NOT EDIT.

package team

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/gomematic/gomematic-api/pkg/api/v1/models"
)

// CreateTeamOKCode is the HTTP code returned for type CreateTeamOK
const CreateTeamOKCode int = 200

/*CreateTeamOK The created team data

swagger:response createTeamOK
*/
type CreateTeamOK struct {

	/*
	  In: Body
	*/
	Payload *models.Team `json:"body,omitempty"`
}

// NewCreateTeamOK creates CreateTeamOK with default headers values
func NewCreateTeamOK() *CreateTeamOK {

	return &CreateTeamOK{}
}

// WithPayload adds the payload to the create team o k response
func (o *CreateTeamOK) WithPayload(payload *models.Team) *CreateTeamOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create team o k response
func (o *CreateTeamOK) SetPayload(payload *models.Team) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateTeamOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateTeamForbiddenCode is the HTTP code returned for type CreateTeamForbidden
const CreateTeamForbiddenCode int = 403

/*CreateTeamForbidden User is not authorized

swagger:response createTeamForbidden
*/
type CreateTeamForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewCreateTeamForbidden creates CreateTeamForbidden with default headers values
func NewCreateTeamForbidden() *CreateTeamForbidden {

	return &CreateTeamForbidden{}
}

// WithPayload adds the payload to the create team forbidden response
func (o *CreateTeamForbidden) WithPayload(payload *models.GeneralError) *CreateTeamForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create team forbidden response
func (o *CreateTeamForbidden) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateTeamForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateTeamPreconditionFailedCode is the HTTP code returned for type CreateTeamPreconditionFailed
const CreateTeamPreconditionFailedCode int = 412

/*CreateTeamPreconditionFailed Failed to parse request body

swagger:response createTeamPreconditionFailed
*/
type CreateTeamPreconditionFailed struct {

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewCreateTeamPreconditionFailed creates CreateTeamPreconditionFailed with default headers values
func NewCreateTeamPreconditionFailed() *CreateTeamPreconditionFailed {

	return &CreateTeamPreconditionFailed{}
}

// WithPayload adds the payload to the create team precondition failed response
func (o *CreateTeamPreconditionFailed) WithPayload(payload *models.GeneralError) *CreateTeamPreconditionFailed {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create team precondition failed response
func (o *CreateTeamPreconditionFailed) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateTeamPreconditionFailed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(412)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateTeamUnprocessableEntityCode is the HTTP code returned for type CreateTeamUnprocessableEntity
const CreateTeamUnprocessableEntityCode int = 422

/*CreateTeamUnprocessableEntity Failed to validate request

swagger:response createTeamUnprocessableEntity
*/
type CreateTeamUnprocessableEntity struct {

	/*
	  In: Body
	*/
	Payload *models.ValidationError `json:"body,omitempty"`
}

// NewCreateTeamUnprocessableEntity creates CreateTeamUnprocessableEntity with default headers values
func NewCreateTeamUnprocessableEntity() *CreateTeamUnprocessableEntity {

	return &CreateTeamUnprocessableEntity{}
}

// WithPayload adds the payload to the create team unprocessable entity response
func (o *CreateTeamUnprocessableEntity) WithPayload(payload *models.ValidationError) *CreateTeamUnprocessableEntity {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create team unprocessable entity response
func (o *CreateTeamUnprocessableEntity) SetPayload(payload *models.ValidationError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateTeamUnprocessableEntity) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(422)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*CreateTeamDefault Some error unrelated to the handler

swagger:response createTeamDefault
*/
type CreateTeamDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewCreateTeamDefault creates CreateTeamDefault with default headers values
func NewCreateTeamDefault(code int) *CreateTeamDefault {
	if code <= 0 {
		code = 500
	}

	return &CreateTeamDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the create team default response
func (o *CreateTeamDefault) WithStatusCode(code int) *CreateTeamDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the create team default response
func (o *CreateTeamDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the create team default response
func (o *CreateTeamDefault) WithPayload(payload *models.GeneralError) *CreateTeamDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create team default response
func (o *CreateTeamDefault) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateTeamDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}