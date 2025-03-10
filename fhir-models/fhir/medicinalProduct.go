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

// MedicinalProduct is documented here http://hl7.org/fhir/StructureDefinition/MedicinalProduct
type MedicinalProduct struct {
	Id                             *string                                          `bson:"id,omitempty" json:"id,omitempty"`
	Meta                           *Meta                                            `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules                  *string                                          `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language                       *string                                          `bson:"language,omitempty" json:"language,omitempty"`
	Text                           *Narrative                                       `bson:"text,omitempty" json:"text,omitempty"`
	Extension                      []Extension                                      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension              []Extension                                      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier                     []Identifier                                     `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Type                           *CodeableConcept                                 `bson:"type,omitempty" json:"type,omitempty"`
	Domain                         *Coding                                          `bson:"domain,omitempty" json:"domain,omitempty"`
	CombinedPharmaceuticalDoseForm *CodeableConcept                                 `bson:"combinedPharmaceuticalDoseForm,omitempty" json:"combinedPharmaceuticalDoseForm,omitempty"`
	LegalStatusOfSupply            *CodeableConcept                                 `bson:"legalStatusOfSupply,omitempty" json:"legalStatusOfSupply,omitempty"`
	AdditionalMonitoringIndicator  *CodeableConcept                                 `bson:"additionalMonitoringIndicator,omitempty" json:"additionalMonitoringIndicator,omitempty"`
	SpecialMeasures                []string                                         `bson:"specialMeasures,omitempty" json:"specialMeasures,omitempty"`
	PaediatricUseIndicator         *CodeableConcept                                 `bson:"paediatricUseIndicator,omitempty" json:"paediatricUseIndicator,omitempty"`
	ProductClassification          []CodeableConcept                                `bson:"productClassification,omitempty" json:"productClassification,omitempty"`
	MarketingStatus                []MarketingStatus                                `bson:"marketingStatus,omitempty" json:"marketingStatus,omitempty"`
	PharmaceuticalProduct          []Reference                                      `bson:"pharmaceuticalProduct,omitempty" json:"pharmaceuticalProduct,omitempty"`
	PackagedMedicinalProduct       []Reference                                      `bson:"packagedMedicinalProduct,omitempty" json:"packagedMedicinalProduct,omitempty"`
	AttachedDocument               []Reference                                      `bson:"attachedDocument,omitempty" json:"attachedDocument,omitempty"`
	MasterFile                     []Reference                                      `bson:"masterFile,omitempty" json:"masterFile,omitempty"`
	Contact                        []Reference                                      `bson:"contact,omitempty" json:"contact,omitempty"`
	ClinicalTrial                  []Reference                                      `bson:"clinicalTrial,omitempty" json:"clinicalTrial,omitempty"`
	Name                           []MedicinalProductName                           `bson:"name" json:"name"`
	CrossReference                 []Identifier                                     `bson:"crossReference,omitempty" json:"crossReference,omitempty"`
	ManufacturingBusinessOperation []MedicinalProductManufacturingBusinessOperation `bson:"manufacturingBusinessOperation,omitempty" json:"manufacturingBusinessOperation,omitempty"`
	SpecialDesignation             []MedicinalProductSpecialDesignation             `bson:"specialDesignation,omitempty" json:"specialDesignation,omitempty"`
}
type MedicinalProductName struct {
	Id                *string                               `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                           `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                           `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	ProductName       string                                `bson:"productName" json:"productName"`
	NamePart          []MedicinalProductNameNamePart        `bson:"namePart,omitempty" json:"namePart,omitempty"`
	CountryLanguage   []MedicinalProductNameCountryLanguage `bson:"countryLanguage,omitempty" json:"countryLanguage,omitempty"`
}
type MedicinalProductNameNamePart struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Part              string      `bson:"part" json:"part"`
	Type              Coding      `bson:"type" json:"type"`
}
type MedicinalProductNameCountryLanguage struct {
	Id                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Country           CodeableConcept  `bson:"country" json:"country"`
	Jurisdiction      *CodeableConcept `bson:"jurisdiction,omitempty" json:"jurisdiction,omitempty"`
	Language          CodeableConcept  `bson:"language" json:"language"`
}
type MedicinalProductManufacturingBusinessOperation struct {
	Id                           *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension                    []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension            []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	OperationType                *CodeableConcept `bson:"operationType,omitempty" json:"operationType,omitempty"`
	AuthorisationReferenceNumber *Identifier      `bson:"authorisationReferenceNumber,omitempty" json:"authorisationReferenceNumber,omitempty"`
	EffectiveDate                *string          `bson:"effectiveDate,omitempty" json:"effectiveDate,omitempty"`
	ConfidentialityIndicator     *CodeableConcept `bson:"confidentialityIndicator,omitempty" json:"confidentialityIndicator,omitempty"`
	Manufacturer                 []Reference      `bson:"manufacturer,omitempty" json:"manufacturer,omitempty"`
	Regulator                    *Reference       `bson:"regulator,omitempty" json:"regulator,omitempty"`
}
type MedicinalProductSpecialDesignation struct {
	Id                        *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension                 []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension         []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier                []Identifier     `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Type                      *CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	IntendedUse               *CodeableConcept `bson:"intendedUse,omitempty" json:"intendedUse,omitempty"`
	IndicationCodeableConcept *CodeableConcept `bson:"indicationCodeableConcept,omitempty" json:"indicationCodeableConcept,omitempty"`
	IndicationReference       *Reference       `bson:"indicationReference,omitempty" json:"indicationReference,omitempty"`
	Status                    *CodeableConcept `bson:"status,omitempty" json:"status,omitempty"`
	Date                      *string          `bson:"date,omitempty" json:"date,omitempty"`
	Species                   *CodeableConcept `bson:"species,omitempty" json:"species,omitempty"`
}
type OtherMedicinalProduct MedicinalProduct

// MarshalJSON marshals the given MedicinalProduct as JSON into a byte slice
func (r MedicinalProduct) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherMedicinalProduct
		ResourceType string `json:"resourceType"`
	}{
		OtherMedicinalProduct: OtherMedicinalProduct(r),
		ResourceType:          "MedicinalProduct",
	})
}

// UnmarshalMedicinalProduct unmarshals a MedicinalProduct.
func UnmarshalMedicinalProduct(b []byte) (MedicinalProduct, error) {
	var medicinalProduct MedicinalProduct
	if err := json.Unmarshal(b, &medicinalProduct); err != nil {
		return medicinalProduct, err
	}
	return medicinalProduct, nil
}
