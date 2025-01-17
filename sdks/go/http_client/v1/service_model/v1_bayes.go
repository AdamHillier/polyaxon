// Copyright 2018-2021 Polyaxon, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by go-swagger; DO NOT EDIT.

package service_model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1Bayes Matrix based on Bayesian Optimization
//
// swagger:model v1Bayes
type V1Bayes struct {

	// Number of concurrent runs
	Concurrency int32 `json:"concurrency,omitempty"`

	// Container to override
	Container V1Container `json:"container,omitempty"`

	// A list of Early stopping objects, accpets both metric and failure early stopping mechanisms
	EarlyStopping []interface{} `json:"earlyStopping"`

	// Kind of matrix, should be equal to "bayes"
	Kind *string `json:"kind,omitempty"`

	// Maximim number of iteration to produce new observations
	MaxIterations int32 `json:"maxIterations,omitempty"`

	// Metric to optimize during the iterations
	Metric *V1OptimizationMetric `json:"metric,omitempty"`

	// Number of intial random observations to create
	NumInitialRuns int32 `json:"numInitialRuns,omitempty"`

	// Hyperparams/Space definition of params to traverse
	Params map[string]interface{} `json:"params,omitempty"`

	// Seed for the random generator
	Seed int32 `json:"seed,omitempty"`

	// A utility function to use for the bayesian optimization
	UtilityFunction interface{} `json:"utilityFunction,omitempty"`
}

// Validate validates this v1 bayes
func (m *V1Bayes) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMetric(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1Bayes) validateMetric(formats strfmt.Registry) error {
	if swag.IsZero(m.Metric) { // not required
		return nil
	}

	if m.Metric != nil {
		if err := m.Metric.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("metric")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this v1 bayes based on the context it is used
func (m *V1Bayes) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateMetric(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1Bayes) contextValidateMetric(ctx context.Context, formats strfmt.Registry) error {

	if m.Metric != nil {
		if err := m.Metric.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("metric")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1Bayes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1Bayes) UnmarshalBinary(b []byte) error {
	var res V1Bayes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
