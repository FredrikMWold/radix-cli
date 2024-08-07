// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// HorizontalScalingSummary HorizontalScalingSummary describe the summary of horizontal scaling of a component
//
// swagger:model HorizontalScalingSummary
type HorizontalScalingSummary struct {

	// CooldownPeriod in seconds. From radixconfig.yaml
	// Example: 300
	CooldownPeriod int32 `json:"cooldownPeriod,omitempty"`

	// Deprecated: Component current average CPU utilization over all pods, represented as a percentage of requested CPU. Use Triggers instead. Will be removed from Radix API 2025-01-01.
	// Example: 70
	CurrentCPUUtilizationPercentage int32 `json:"currentCPUUtilizationPercentage,omitempty"`

	// Deprecated: Component current average memory utilization over all pods, represented as a percentage of requested memory. Use Triggers instead. Will be removed from Radix API 2025-01-01.
	// Example: 80
	CurrentMemoryUtilizationPercentage int32 `json:"currentMemoryUtilizationPercentage,omitempty"`

	// CurrentReplicas returns the current number of replicas
	// Example: 1
	// Required: true
	CurrentReplicas *int32 `json:"currentReplicas"`

	// DesiredReplicas returns the target number of replicas across all triggers
	// Example: 2
	// Required: true
	DesiredReplicas *int32 `json:"desiredReplicas"`

	// Component maximum replicas. From radixconfig.yaml
	// Example: 5
	MaxReplicas int32 `json:"maxReplicas,omitempty"`

	// Component minimum replicas. From radixconfig.yaml
	// Example: 2
	MinReplicas int32 `json:"minReplicas,omitempty"`

	// PollingInterval in seconds. From radixconfig.yaml
	// Example: 30
	PollingInterval int32 `json:"pollingInterval,omitempty"`

	// Deprecated: Component target average CPU utilization over all pods. Use Triggers instead. Will be removed from Radix API 2025-01-01.
	// Example: 80
	TargetCPUUtilizationPercentage int32 `json:"targetCPUUtilizationPercentage,omitempty"`

	// Deprecated: Component target average memory utilization over all pods. use Triggers instead. Will be removed from Radix API 2025-01-01.
	// Example: 80
	TargetMemoryUtilizationPercentage int32 `json:"targetMemoryUtilizationPercentage,omitempty"`

	// Triggers lists status of all triggers found in radixconfig.yaml
	// Example: 30
	// Required: true
	Triggers []*HorizontalScalingSummaryTriggerStatus `json:"triggers"`
}

// Validate validates this horizontal scaling summary
func (m *HorizontalScalingSummary) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCurrentReplicas(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDesiredReplicas(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTriggers(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HorizontalScalingSummary) validateCurrentReplicas(formats strfmt.Registry) error {

	if err := validate.Required("currentReplicas", "body", m.CurrentReplicas); err != nil {
		return err
	}

	return nil
}

func (m *HorizontalScalingSummary) validateDesiredReplicas(formats strfmt.Registry) error {

	if err := validate.Required("desiredReplicas", "body", m.DesiredReplicas); err != nil {
		return err
	}

	return nil
}

func (m *HorizontalScalingSummary) validateTriggers(formats strfmt.Registry) error {

	if err := validate.Required("triggers", "body", m.Triggers); err != nil {
		return err
	}

	for i := 0; i < len(m.Triggers); i++ {
		if swag.IsZero(m.Triggers[i]) { // not required
			continue
		}

		if m.Triggers[i] != nil {
			if err := m.Triggers[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("triggers" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("triggers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this horizontal scaling summary based on the context it is used
func (m *HorizontalScalingSummary) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateTriggers(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HorizontalScalingSummary) contextValidateTriggers(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Triggers); i++ {

		if m.Triggers[i] != nil {

			if swag.IsZero(m.Triggers[i]) { // not required
				return nil
			}

			if err := m.Triggers[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("triggers" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("triggers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *HorizontalScalingSummary) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HorizontalScalingSummary) UnmarshalBinary(b []byte) error {
	var res HorizontalScalingSummary
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
