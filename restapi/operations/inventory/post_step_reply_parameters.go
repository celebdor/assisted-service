// Code generated by go-swagger; DO NOT EDIT.

package inventory

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"

	"github.com/filanov/bm-inventory/models"
)

// NewPostStepReplyParams creates a new PostStepReplyParams object
// no default values defined in spec.
func NewPostStepReplyParams() PostStepReplyParams {

	return PostStepReplyParams{}
}

// PostStepReplyParams contains all the bound params for the post step reply operation
// typically these are obtained from a http.Request
//
// swagger:parameters PostStepReply
type PostStepReplyParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*ID of node
	  Required: true
	  In: path
	*/
	NodeID strfmt.UUID
	/*
	  In: body
	*/
	Reply *models.StepReply
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewPostStepReplyParams() beforehand.
func (o *PostStepReplyParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rNodeID, rhkNodeID, _ := route.Params.GetOK("node_id")
	if err := o.bindNodeID(rNodeID, rhkNodeID, route.Formats); err != nil {
		res = append(res, err)
	}

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.StepReply
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			res = append(res, errors.NewParseError("reply", "body", "", err))
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.Reply = &body
			}
		}
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindNodeID binds and validates parameter NodeID from path.
func (o *PostStepReplyParams) bindNodeID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	// Format: uuid
	value, err := formats.Parse("uuid", raw)
	if err != nil {
		return errors.InvalidType("node_id", "path", "strfmt.UUID", raw)
	}
	o.NodeID = *(value.(*strfmt.UUID))

	if err := o.validateNodeID(formats); err != nil {
		return err
	}

	return nil
}

// validateNodeID carries on validations for parameter NodeID
func (o *PostStepReplyParams) validateNodeID(formats strfmt.Registry) error {

	if err := validate.FormatOf("node_id", "path", "uuid", o.NodeID.String(), formats); err != nil {
		return err
	}
	return nil
}