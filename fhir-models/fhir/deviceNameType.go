// Copyright 2019 The Samply Development Community
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package fhir

import (
	"encoding/json"
	"fmt"
	"strings"
)

// THIS FILE IS GENERATED BY https://github.com/lifebox-healthcare/golang-fhir-models
// PLEASE DO NOT EDIT BY HAND

// DeviceNameType is documented here http://hl7.org/fhir/ValueSet/device-nametype
type DeviceNameType int

const (
	DeviceNameTypeUdiLabelName DeviceNameType = iota
	DeviceNameTypeUserFriendlyName
	DeviceNameTypePatientReportedName
	DeviceNameTypeManufacturerName
	DeviceNameTypeModelName
	DeviceNameTypeOther
)

func (code DeviceNameType) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *DeviceNameType) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "udi-label-name":
		*code = DeviceNameTypeUdiLabelName
	case "user-friendly-name":
		*code = DeviceNameTypeUserFriendlyName
	case "patient-reported-name":
		*code = DeviceNameTypePatientReportedName
	case "manufacturer-name":
		*code = DeviceNameTypeManufacturerName
	case "model-name":
		*code = DeviceNameTypeModelName
	case "other":
		*code = DeviceNameTypeOther
	default:
		return fmt.Errorf("unknown DeviceNameType code `%s`", s)
	}
	return nil
}
func (code DeviceNameType) String() string {
	return code.Code()
}
func (code DeviceNameType) Code() string {
	switch code {
	case DeviceNameTypeUdiLabelName:
		return "udi-label-name"
	case DeviceNameTypeUserFriendlyName:
		return "user-friendly-name"
	case DeviceNameTypePatientReportedName:
		return "patient-reported-name"
	case DeviceNameTypeManufacturerName:
		return "manufacturer-name"
	case DeviceNameTypeModelName:
		return "model-name"
	case DeviceNameTypeOther:
		return "other"
	}
	return "<unknown>"
}
func (code DeviceNameType) Display() string {
	switch code {
	case DeviceNameTypeUdiLabelName:
		return "UDI Label name"
	case DeviceNameTypeUserFriendlyName:
		return "User Friendly name"
	case DeviceNameTypePatientReportedName:
		return "Patient Reported name"
	case DeviceNameTypeManufacturerName:
		return "Manufacturer name"
	case DeviceNameTypeModelName:
		return "Model name"
	case DeviceNameTypeOther:
		return "other"
	}
	return "<unknown>"
}
func (code DeviceNameType) Definition() string {
	switch code {
	case DeviceNameTypeUdiLabelName:
		return "UDI Label name."
	case DeviceNameTypeUserFriendlyName:
		return "User Friendly name."
	case DeviceNameTypePatientReportedName:
		return "Patient Reported name."
	case DeviceNameTypeManufacturerName:
		return "Manufacturer name."
	case DeviceNameTypeModelName:
		return "Model name."
	case DeviceNameTypeOther:
		return "other."
	}
	return "<unknown>"
}
