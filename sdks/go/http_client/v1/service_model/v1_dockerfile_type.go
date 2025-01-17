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

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1DockerfileType Dockerfile type specification
//
// swagger:model v1DockerfileType
type V1DockerfileType struct {

	// An optional list of tuples for copying paths, translate to a COPY statements
	Copy map[string]string `json:"copy,omitempty"`

	// An optional list of tuples(key, value) for defining env vars, translate to an ENV statements
	Env map[string]string `json:"env,omitempty"`

	// A filename to give to the generated dockerfile
	Filename string `json:"filename,omitempty"`

	// A gid to use when creating the docker image
	Gid int32 `json:"gid,omitempty"`

	// Docker image to use as a base
	Image string `json:"image,omitempty"`

	// An optional string defining a language, e.g. en_US.UTF-8
	LangEnv string `json:"langEnv,omitempty"`

	// An optional list of tuples for exporting paths, translate to a PATH statements
	Path map[string]string `json:"path,omitempty"`

	// An optional list of strubg for executing Run commands, translate to a RUN statements
	Run []string `json:"run"`

	// An optional shell type, defaults to "/bin/bash"
	Shell string `json:"shell,omitempty"`

	// A uid to use when creating the docker image
	UID int32 `json:"uid,omitempty"`

	// A work dir to copy code to, default to /code
	Workdir string `json:"workdir,omitempty"`

	// An optional workdir path
	WorkdirPath string `json:"workdirPath,omitempty"`
}

// Validate validates this v1 dockerfile type
func (m *V1DockerfileType) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this v1 dockerfile type based on context it is used
func (m *V1DockerfileType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1DockerfileType) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1DockerfileType) UnmarshalBinary(b []byte) error {
	var res V1DockerfileType
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
