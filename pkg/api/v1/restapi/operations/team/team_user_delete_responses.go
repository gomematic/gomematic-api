// Code generated by go-swagger; DO NOT EDIT.

package team

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/gomematic/gomematic-api/pkg/api/v1/models"
)

// TeamUserDeleteOKCode is the HTTP code returned for type TeamUserDeleteOK
const TeamUserDeleteOKCode int = 200

/*TeamUserDeleteOK Plain success message

swagger:response teamUserDeleteOK
*/
type TeamUserDeleteOK struct {

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewTeamUserDeleteOK creates TeamUserDeleteOK with default headers values
func NewTeamUserDeleteOK() *TeamUserDeleteOK {

	return &TeamUserDeleteOK{}
}

// WithPayload adds the payload to the team user delete o k response
func (o *TeamUserDeleteOK) WithPayload(payload *models.GeneralError) *TeamUserDeleteOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the team user delete o k response
func (o *TeamUserDeleteOK) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *TeamUserDeleteOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// TeamUserDeleteForbiddenCode is the HTTP code returned for type TeamUserDeleteForbidden
const TeamUserDeleteForbiddenCode int = 403

/*TeamUserDeleteForbidden User is not authorized

swagger:response teamUserDeleteForbidden
*/
type TeamUserDeleteForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewTeamUserDeleteForbidden creates TeamUserDeleteForbidden with default headers values
func NewTeamUserDeleteForbidden() *TeamUserDeleteForbidden {

	return &TeamUserDeleteForbidden{}
}

// WithPayload adds the payload to the team user delete forbidden response
func (o *TeamUserDeleteForbidden) WithPayload(payload *models.GeneralError) *TeamUserDeleteForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the team user delete forbidden response
func (o *TeamUserDeleteForbidden) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *TeamUserDeleteForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// TeamUserDeletePreconditionFailedCode is the HTTP code returned for type TeamUserDeletePreconditionFailed
const TeamUserDeletePreconditionFailedCode int = 412

/*TeamUserDeletePreconditionFailed Failed to parse request body

swagger:response teamUserDeletePreconditionFailed
*/
type TeamUserDeletePreconditionFailed struct {

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewTeamUserDeletePreconditionFailed creates TeamUserDeletePreconditionFailed with default headers values
func NewTeamUserDeletePreconditionFailed() *TeamUserDeletePreconditionFailed {

	return &TeamUserDeletePreconditionFailed{}
}

// WithPayload adds the payload to the team user delete precondition failed response
func (o *TeamUserDeletePreconditionFailed) WithPayload(payload *models.GeneralError) *TeamUserDeletePreconditionFailed {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the team user delete precondition failed response
func (o *TeamUserDeletePreconditionFailed) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *TeamUserDeletePreconditionFailed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(412)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// TeamUserDeleteUnprocessableEntityCode is the HTTP code returned for type TeamUserDeleteUnprocessableEntity
const TeamUserDeleteUnprocessableEntityCode int = 422

/*TeamUserDeleteUnprocessableEntity User is not assigned

swagger:response teamUserDeleteUnprocessableEntity
*/
type TeamUserDeleteUnprocessableEntity struct {

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewTeamUserDeleteUnprocessableEntity creates TeamUserDeleteUnprocessableEntity with default headers values
func NewTeamUserDeleteUnprocessableEntity() *TeamUserDeleteUnprocessableEntity {

	return &TeamUserDeleteUnprocessableEntity{}
}

// WithPayload adds the payload to the team user delete unprocessable entity response
func (o *TeamUserDeleteUnprocessableEntity) WithPayload(payload *models.GeneralError) *TeamUserDeleteUnprocessableEntity {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the team user delete unprocessable entity response
func (o *TeamUserDeleteUnprocessableEntity) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *TeamUserDeleteUnprocessableEntity) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(422)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*TeamUserDeleteDefault Some error unrelated to the handler

swagger:response teamUserDeleteDefault
*/
type TeamUserDeleteDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewTeamUserDeleteDefault creates TeamUserDeleteDefault with default headers values
func NewTeamUserDeleteDefault(code int) *TeamUserDeleteDefault {
	if code <= 0 {
		code = 500
	}

	return &TeamUserDeleteDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the team user delete default response
func (o *TeamUserDeleteDefault) WithStatusCode(code int) *TeamUserDeleteDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the team user delete default response
func (o *TeamUserDeleteDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the team user delete default response
func (o *TeamUserDeleteDefault) WithPayload(payload *models.GeneralError) *TeamUserDeleteDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the team user delete default response
func (o *TeamUserDeleteDefault) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *TeamUserDeleteDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
