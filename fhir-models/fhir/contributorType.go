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

// ContributorType is documented here http://hl7.org/fhir/ValueSet/contributor-type
type ContributorType int

const (
	ContributorTypeAuthor ContributorType = iota
	ContributorTypeEditor
	ContributorTypeReviewer
	ContributorTypeEndorser
)

func (code ContributorType) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *ContributorType) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "author":
		*code = ContributorTypeAuthor
	case "editor":
		*code = ContributorTypeEditor
	case "reviewer":
		*code = ContributorTypeReviewer
	case "endorser":
		*code = ContributorTypeEndorser
	default:
		return fmt.Errorf("unknown ContributorType code `%s`", s)
	}
	return nil
}
func (code ContributorType) String() string {
	return code.Code()
}
func (code ContributorType) Code() string {
	switch code {
	case ContributorTypeAuthor:
		return "author"
	case ContributorTypeEditor:
		return "editor"
	case ContributorTypeReviewer:
		return "reviewer"
	case ContributorTypeEndorser:
		return "endorser"
	}
	return "<unknown>"
}
func (code ContributorType) Display() string {
	switch code {
	case ContributorTypeAuthor:
		return "Author"
	case ContributorTypeEditor:
		return "Editor"
	case ContributorTypeReviewer:
		return "Reviewer"
	case ContributorTypeEndorser:
		return "Endorser"
	}
	return "<unknown>"
}
func (code ContributorType) Definition() string {
	switch code {
	case ContributorTypeAuthor:
		return "An author of the content of the module."
	case ContributorTypeEditor:
		return "An editor of the content of the module."
	case ContributorTypeReviewer:
		return "A reviewer of the content of the module."
	case ContributorTypeEndorser:
		return "An endorser of the content of the module."
	}
	return "<unknown>"
}
