// Code generated by go-swagger; DO NOT EDIT.

package team

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/gomematic/gomematic-api/pkg/api/v1/models"
)

// TeamDeleteOKCode is the HTTP code returned for type TeamDeleteOK
const TeamDeleteOKCode int = 200

/*TeamDeleteOK Plain success message

swagger:response teamDeleteOK
*/
type TeamDeleteOK struct {

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewTeamDeleteOK creates TeamDeleteOK with default headers values
func NewTeamDeleteOK() *TeamDeleteOK {

	return &TeamDeleteOK{}
}

// WithPayload adds the payload to the team delete o k response
func (o *TeamDeleteOK) WithPayload(payload *models.GeneralError) *TeamDeleteOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the team delete o k response
func (o *TeamDeleteOK) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *TeamDeleteOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// TeamDeleteBadRequestCode is the HTTP code returned for type TeamDeleteBadRequest
const TeamDeleteBadRequestCode int = 400

/*TeamDeleteBadRequest Failed to delete the team

swagger:response teamDeleteBadRequest
*/
type TeamDeleteBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewTeamDeleteBadRequest creates TeamDeleteBadRequest with default headers values
func NewTeamDeleteBadRequest() *TeamDeleteBadRequest {

	return &TeamDeleteBadRequest{}
}

// WithPayload adds the payload to the team delete bad request response
func (o *TeamDeleteBadRequest) WithPayload(payload *models.GeneralError) *TeamDeleteBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the team delete bad request response
func (o *TeamDeleteBadRequest) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *TeamDeleteBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// TeamDeleteForbiddenCode is the HTTP code returned for type TeamDeleteForbidden
const TeamDeleteForbiddenCode int = 403

/*TeamDeleteForbidden User is not authorized

swagger:response teamDeleteForbidden
*/
type TeamDeleteForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewTeamDeleteForbidden creates TeamDeleteForbidden with default headers values
func NewTeamDeleteForbidden() *TeamDeleteForbidden {

	return &TeamDeleteForbidden{}
}

// WithPayload adds the payload to the team delete forbidden response
func (o *TeamDeleteForbidden) WithPayload(payload *models.GeneralError) *TeamDeleteForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the team delete forbidden response
func (o *TeamDeleteForbidden) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *TeamDeleteForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*TeamDeleteDefault Some error unrelated to the handler

swagger:response teamDeleteDefault
*/
type TeamDeleteDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewTeamDeleteDefault creates TeamDeleteDefault with default headers values
func NewTeamDeleteDefault(code int) *TeamDeleteDefault {
	if code <= 0 {
		code = 500
	}

	return &TeamDeleteDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the team delete default response
func (o *TeamDeleteDefault) WithStatusCode(code int) *TeamDeleteDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the team delete default response
func (o *TeamDeleteDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the team delete default response
func (o *TeamDeleteDefault) WithPayload(payload *models.GeneralError) *TeamDeleteDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the team delete default response
func (o *TeamDeleteDefault) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *TeamDeleteDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
