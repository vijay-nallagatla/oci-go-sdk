// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Application Migration Service API
//
// API for the Application Migration service. Use this API to migrate applications from Oracle Cloud Infrastructure - Classic to Oracle Cloud Infrastructure.
//

package applicationmigration

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/common"
)

// OcicSourceDetails Specifies configuration specific to the source environment.
type OcicSourceDetails struct {

	// The Oracle Cloud Infrastructure - Classic region name (e.g. us2-z11 or uscom-central-1)
	Region *string `mandatory:"true" json:"region"`

	// The compute account id
	ComputeAccount *string `mandatory:"true" json:"computeAccount"`
}

func (m OcicSourceDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m OcicSourceDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeOcicSourceDetails OcicSourceDetails
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeOcicSourceDetails
	}{
		"OCIC",
		(MarshalTypeOcicSourceDetails)(m),
	}

	return json.Marshal(&s)
}
