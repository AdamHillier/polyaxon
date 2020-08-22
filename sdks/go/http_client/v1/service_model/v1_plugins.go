// Copyright 2018-2020 Polyaxon, Inc.
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
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1Plugins Plugins specification
//
// swagger:model v1Plugins
type V1Plugins struct {

	// Optional flag to tell Polyaxon if it should set an auth context for the run, default true
	Auth bool `json:"auth,omitempty"`

	// Auto resume a run's artifacts (applies to resume and retries), works if collects_artifacts is enabled
	AutoResume bool `json:"auto_resume,omitempty"`

	// Optional flag to tell Polyaxon to collect articats and outputs
	CollectArtifacts bool `json:"collect_artifacts,omitempty"`

	// Optional flag to tell Polyaxon to collect logs
	CollectLogs bool `json:"collect_logs,omitempty"`

	// Optional flag to tell Polyaxon to collect container resouces (cpu/memory/gpu)
	CollectResources string `json:"collect_resources,omitempty"`

	// Optional flag to tell Polyaxon if it should set a docker socket context for the run, default false
	Docker bool `json:"docker,omitempty"`

	// Optional log level
	LogLevel string `json:"log_level,omitempty"`

	// Option Notifications: Deprecated
	Notifications []*V1Notification `json:"notifications"`

	// Optional flag to tell Polyaxon if it should set a shm context for the run, default false
	Shm bool `json:"shm,omitempty"`

	// Optional flag to tell Polyaxon to sync statuses
	SyncStatuses bool `json:"sync_statuses,omitempty"`
}

// Validate validates this v1 plugins
func (m *V1Plugins) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNotifications(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1Plugins) validateNotifications(formats strfmt.Registry) error {

	if swag.IsZero(m.Notifications) { // not required
		return nil
	}

	for i := 0; i < len(m.Notifications); i++ {
		if swag.IsZero(m.Notifications[i]) { // not required
			continue
		}

		if m.Notifications[i] != nil {
			if err := m.Notifications[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("notifications" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1Plugins) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1Plugins) UnmarshalBinary(b []byte) error {
	var res V1Plugins
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
