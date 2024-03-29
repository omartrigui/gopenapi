// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// GetPasswordsGenerateRandomOKCode is the HTTP code returned for type GetPasswordsGenerateRandomOK
const GetPasswordsGenerateRandomOKCode int = 200

/*GetPasswordsGenerateRandomOK Returns the randomly generated password.

swagger:response getPasswordsGenerateRandomOK
*/
type GetPasswordsGenerateRandomOK struct {

	/*
	  In: Body
	*/
	Payload *GetPasswordsGenerateRandomOKBody `json:"body,omitempty"`
}

// NewGetPasswordsGenerateRandomOK creates GetPasswordsGenerateRandomOK with default headers values
func NewGetPasswordsGenerateRandomOK() *GetPasswordsGenerateRandomOK {

	return &GetPasswordsGenerateRandomOK{}
}

// WithPayload adds the payload to the get passwords generate random o k response
func (o *GetPasswordsGenerateRandomOK) WithPayload(payload *GetPasswordsGenerateRandomOKBody) *GetPasswordsGenerateRandomOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get passwords generate random o k response
func (o *GetPasswordsGenerateRandomOK) SetPayload(payload *GetPasswordsGenerateRandomOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetPasswordsGenerateRandomOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetPasswordsGenerateRandomBadRequestCode is the HTTP code returned for type GetPasswordsGenerateRandomBadRequest
const GetPasswordsGenerateRandomBadRequestCode int = 400

/*GetPasswordsGenerateRandomBadRequest Invalid length was provided.

swagger:response getPasswordsGenerateRandomBadRequest
*/
type GetPasswordsGenerateRandomBadRequest struct {
}

// NewGetPasswordsGenerateRandomBadRequest creates GetPasswordsGenerateRandomBadRequest with default headers values
func NewGetPasswordsGenerateRandomBadRequest() *GetPasswordsGenerateRandomBadRequest {

	return &GetPasswordsGenerateRandomBadRequest{}
}

// WriteResponse to the client
func (o *GetPasswordsGenerateRandomBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}
