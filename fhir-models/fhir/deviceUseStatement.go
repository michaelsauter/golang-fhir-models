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

import "encoding/json"

// THIS FILE IS GENERATED BY https://github.com/lifebox-healthcare/golang-fhir-models
// PLEASE DO NOT EDIT BY HAND

// DeviceUseStatement is documented here http://hl7.org/fhir/StructureDefinition/DeviceUseStatement
type DeviceUseStatement struct {
	Id                *string                  `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta                    `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string                  `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string                  `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative               `bson:"text,omitempty" json:"text,omitempty"`
	Extension         []Extension              `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension              `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier        []Identifier             `bson:"identifier,omitempty" json:"identifier,omitempty"`
	BasedOn           []Reference              `bson:"basedOn,omitempty" json:"basedOn,omitempty"`
	Status            DeviceUseStatementStatus `bson:"status" json:"status"`
	Subject           Reference                `bson:"subject" json:"subject"`
	DerivedFrom       []Reference              `bson:"derivedFrom,omitempty" json:"derivedFrom,omitempty"`
	TimingTiming      *Timing                  `bson:"timingTiming,omitempty" json:"timingTiming,omitempty"`
	TimingPeriod      *Period                  `bson:"timingPeriod,omitempty" json:"timingPeriod,omitempty"`
	TimingDateTime    *string                  `bson:"timingDateTime,omitempty" json:"timingDateTime,omitempty"`
	RecordedOn        *string                  `bson:"recordedOn,omitempty" json:"recordedOn,omitempty"`
	Source            *Reference               `bson:"source,omitempty" json:"source,omitempty"`
	Device            Reference                `bson:"device" json:"device"`
	ReasonCode        []CodeableConcept        `bson:"reasonCode,omitempty" json:"reasonCode,omitempty"`
	ReasonReference   []Reference              `bson:"reasonReference,omitempty" json:"reasonReference,omitempty"`
	BodySite          *CodeableConcept         `bson:"bodySite,omitempty" json:"bodySite,omitempty"`
	Note              []Annotation             `bson:"note,omitempty" json:"note,omitempty"`
}
type OtherDeviceUseStatement DeviceUseStatement

// MarshalJSON marshals the given DeviceUseStatement as JSON into a byte slice
func (r DeviceUseStatement) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherDeviceUseStatement
		ResourceType string `json:"resourceType"`
	}{
		OtherDeviceUseStatement: OtherDeviceUseStatement(r),
		ResourceType:            "DeviceUseStatement",
	})
}

// UnmarshalDeviceUseStatement unmarshals a DeviceUseStatement.
func UnmarshalDeviceUseStatement(b []byte) (DeviceUseStatement, error) {
	var deviceUseStatement DeviceUseStatement
	if err := json.Unmarshal(b, &deviceUseStatement); err != nil {
		return deviceUseStatement, err
	}
	return deviceUseStatement, nil
}
