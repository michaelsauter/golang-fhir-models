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

// Account is documented here http://hl7.org/fhir/StructureDefinition/Account
type Account struct {
	Id                *string            `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta              `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string            `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string            `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative         `bson:"text,omitempty" json:"text,omitempty"`
	Extension         []Extension        `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension        `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier        []Identifier       `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Status            AccountStatus      `bson:"status" json:"status"`
	Type              *CodeableConcept   `bson:"type,omitempty" json:"type,omitempty"`
	Name              *string            `bson:"name,omitempty" json:"name,omitempty"`
	Subject           []Reference        `bson:"subject,omitempty" json:"subject,omitempty"`
	ServicePeriod     *Period            `bson:"servicePeriod,omitempty" json:"servicePeriod,omitempty"`
	Coverage          []AccountCoverage  `bson:"coverage,omitempty" json:"coverage,omitempty"`
	Owner             *Reference         `bson:"owner,omitempty" json:"owner,omitempty"`
	Description       *string            `bson:"description,omitempty" json:"description,omitempty"`
	Guarantor         []AccountGuarantor `bson:"guarantor,omitempty" json:"guarantor,omitempty"`
	PartOf            *Reference         `bson:"partOf,omitempty" json:"partOf,omitempty"`
}
type AccountCoverage struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Coverage          Reference   `bson:"coverage" json:"coverage"`
	Priority          *int        `bson:"priority,omitempty" json:"priority,omitempty"`
}
type AccountGuarantor struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Party             Reference   `bson:"party" json:"party"`
	OnHold            *bool       `bson:"onHold,omitempty" json:"onHold,omitempty"`
	Period            *Period     `bson:"period,omitempty" json:"period,omitempty"`
}
type OtherAccount Account

// MarshalJSON marshals the given Account as JSON into a byte slice
func (r Account) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherAccount
		ResourceType string `json:"resourceType"`
	}{
		OtherAccount: OtherAccount(r),
		ResourceType: "Account",
	})
}

// UnmarshalAccount unmarshals a Account.
func UnmarshalAccount(b []byte) (Account, error) {
	var account Account
	if err := json.Unmarshal(b, &account); err != nil {
		return account, err
	}
	return account, nil
}
