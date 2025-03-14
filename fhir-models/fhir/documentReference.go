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

// DocumentReference is documented here http://hl7.org/fhir/StructureDefinition/DocumentReference
type DocumentReference struct {
	Id                *string                      `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta                        `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string                      `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string                      `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative                   `bson:"text,omitempty" json:"text,omitempty"`
	Extension         []Extension                  `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                  `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	MasterIdentifier  *Identifier                  `bson:"masterIdentifier,omitempty" json:"masterIdentifier,omitempty"`
	Identifier        []Identifier                 `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Status            DocumentReferenceStatus      `bson:"status" json:"status"`
	DocStatus         *CompositionStatus           `bson:"docStatus,omitempty" json:"docStatus,omitempty"`
	Type              *CodeableConcept             `bson:"type,omitempty" json:"type,omitempty"`
	Category          []CodeableConcept            `bson:"category,omitempty" json:"category,omitempty"`
	Subject           *Reference                   `bson:"subject,omitempty" json:"subject,omitempty"`
	Date              *string                      `bson:"date,omitempty" json:"date,omitempty"`
	Author            []Reference                  `bson:"author,omitempty" json:"author,omitempty"`
	Authenticator     *Reference                   `bson:"authenticator,omitempty" json:"authenticator,omitempty"`
	Custodian         *Reference                   `bson:"custodian,omitempty" json:"custodian,omitempty"`
	RelatesTo         []DocumentReferenceRelatesTo `bson:"relatesTo,omitempty" json:"relatesTo,omitempty"`
	Description       *string                      `bson:"description,omitempty" json:"description,omitempty"`
	SecurityLabel     []CodeableConcept            `bson:"securityLabel,omitempty" json:"securityLabel,omitempty"`
	Content           []DocumentReferenceContent   `bson:"content" json:"content"`
	Context           *DocumentReferenceContext    `bson:"context,omitempty" json:"context,omitempty"`
}
type DocumentReferenceRelatesTo struct {
	Id                *string                  `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension              `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension              `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Code              DocumentRelationshipType `bson:"code" json:"code"`
	Target            Reference                `bson:"target" json:"target"`
}
type DocumentReferenceContent struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Attachment        Attachment  `bson:"attachment" json:"attachment"`
	Format            *Coding     `bson:"format,omitempty" json:"format,omitempty"`
}
type DocumentReferenceContext struct {
	Id                *string           `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension       `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension       `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Encounter         []Reference       `bson:"encounter,omitempty" json:"encounter,omitempty"`
	Event             []CodeableConcept `bson:"event,omitempty" json:"event,omitempty"`
	Period            *Period           `bson:"period,omitempty" json:"period,omitempty"`
	FacilityType      *CodeableConcept  `bson:"facilityType,omitempty" json:"facilityType,omitempty"`
	PracticeSetting   *CodeableConcept  `bson:"practiceSetting,omitempty" json:"practiceSetting,omitempty"`
	SourcePatientInfo *Reference        `bson:"sourcePatientInfo,omitempty" json:"sourcePatientInfo,omitempty"`
	Related           []Reference       `bson:"related,omitempty" json:"related,omitempty"`
}
type OtherDocumentReference DocumentReference

// MarshalJSON marshals the given DocumentReference as JSON into a byte slice
func (r DocumentReference) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherDocumentReference
		ResourceType string `json:"resourceType"`
	}{
		OtherDocumentReference: OtherDocumentReference(r),
		ResourceType:           "DocumentReference",
	})
}

// UnmarshalDocumentReference unmarshals a DocumentReference.
func UnmarshalDocumentReference(b []byte) (DocumentReference, error) {
	var documentReference DocumentReference
	if err := json.Unmarshal(b, &documentReference); err != nil {
		return documentReference, err
	}
	return documentReference, nil
}
