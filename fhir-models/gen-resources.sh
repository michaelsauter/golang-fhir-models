#!/usr/bin/env bash
set -eu

# Fix version to R4 for now.
# As R4B has been released in May 2022, we should update to this or at least
# allow the option to select the version.
wget -O definitions.zip https://www.hl7.org/fhir/R4/definitions.json.zip
unzip definitions.zip profiles-resources.json profiles-types.json valuesets.json -d fhir
rm definitions.zip

go generate ./fhir
