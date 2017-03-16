package usages

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/swag"

	strfmt "github.com/go-swagger/go-swagger/strfmt"
)

// NewGetAccountParams creates a new GetAccountParams object
// with the default values initialized.
func NewGetAccountParams() *GetAccountParams {
	var ()
	return &GetAccountParams{}
}

/*GetAccountParams contains all the parameters to send to the API endpoint
for the get account operation typically these are written to a http.Request
*/
type GetAccountParams struct {

	/*Cloud*/
	Cloud *string
	/*Filterenddate*/
	Filterenddate *int64
	/*Since*/
	Since *int64
	/*User*/
	User *string
	/*Zone*/
	Zone *string
}

// WithCloud adds the cloud to the get account params
func (o *GetAccountParams) WithCloud(cloud *string) *GetAccountParams {
	o.Cloud = cloud
	return o
}

// WithFilterenddate adds the filterenddate to the get account params
func (o *GetAccountParams) WithFilterenddate(filterenddate *int64) *GetAccountParams {
	o.Filterenddate = filterenddate
	return o
}

// WithSince adds the since to the get account params
func (o *GetAccountParams) WithSince(since *int64) *GetAccountParams {
	o.Since = since
	return o
}

// WithUser adds the user to the get account params
func (o *GetAccountParams) WithUser(user *string) *GetAccountParams {
	o.User = user
	return o
}

// WithZone adds the zone to the get account params
func (o *GetAccountParams) WithZone(zone *string) *GetAccountParams {
	o.Zone = zone
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *GetAccountParams) WriteToRequest(r client.Request, reg strfmt.Registry) error {

	var res []error

	if o.Cloud != nil {

		// query param cloud
		var qrCloud string
		if o.Cloud != nil {
			qrCloud = *o.Cloud
		}
		qCloud := qrCloud
		if qCloud != "" {
			if err := r.SetQueryParam("cloud", qCloud); err != nil {
				return err
			}
		}

	}

	if o.Filterenddate != nil {

		// query param filterenddate
		var qrFilterenddate int64
		if o.Filterenddate != nil {
			qrFilterenddate = *o.Filterenddate
		}
		qFilterenddate := swag.FormatInt64(qrFilterenddate)
		if qFilterenddate != "" {
			if err := r.SetQueryParam("filterenddate", qFilterenddate); err != nil {
				return err
			}
		}

	}

	if o.Since != nil {

		// query param since
		var qrSince int64
		if o.Since != nil {
			qrSince = *o.Since
		}
		qSince := swag.FormatInt64(qrSince)
		if qSince != "" {
			if err := r.SetQueryParam("since", qSince); err != nil {
				return err
			}
		}

	}

	if o.User != nil {

		// query param user
		var qrUser string
		if o.User != nil {
			qrUser = *o.User
		}
		qUser := qrUser
		if qUser != "" {
			if err := r.SetQueryParam("user", qUser); err != nil {
				return err
			}
		}

	}

	if o.Zone != nil {

		// query param zone
		var qrZone string
		if o.Zone != nil {
			qrZone = *o.Zone
		}
		qZone := qrZone
		if qZone != "" {
			if err := r.SetQueryParam("zone", qZone); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}