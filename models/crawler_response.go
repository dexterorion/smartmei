// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// CrawlerResponse crawler response
// swagger:model CrawlerResponse
type CrawlerResponse struct {

	// date
	Date string `json:"date,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// rates
	Rates CrawlerResponseRates `json:"rates,omitempty"`
}

// Validate validates this crawler response
func (m *CrawlerResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRates(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CrawlerResponse) validateRates(formats strfmt.Registry) error {

	if swag.IsZero(m.Rates) { // not required
		return nil
	}

	if err := m.Rates.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("rates")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CrawlerResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CrawlerResponse) UnmarshalBinary(b []byte) error {
	var res CrawlerResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// CrawlerResponseRates crawler response rates
// swagger:model CrawlerResponseRates
type CrawlerResponseRates struct {

	// brl
	Brl float64 `json:"brl,omitempty"`

	// eur
	Eur float64 `json:"eur,omitempty"`

	// usd
	Usd float64 `json:"usd,omitempty"`
}

// Validate validates this crawler response rates
func (m *CrawlerResponseRates) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CrawlerResponseRates) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CrawlerResponseRates) UnmarshalBinary(b []byte) error {
	var res CrawlerResponseRates
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
