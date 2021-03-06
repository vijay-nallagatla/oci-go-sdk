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

// PcsDiscoveryDetails Specifies the credentials to access the source PCS instance
type PcsDiscoveryDetails struct {

	// The PCS instance admin user
	ServiceInstanceUser *string `mandatory:"true" json:"serviceInstanceUser"`

	// The PCS instance admin password
	ServiceInstancePassword *string `mandatory:"true" json:"serviceInstancePassword"`
}

func (m PcsDiscoveryDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m PcsDiscoveryDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypePcsDiscoveryDetails PcsDiscoveryDetails
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypePcsDiscoveryDetails
	}{
		"PCS",
		(MarshalTypePcsDiscoveryDetails)(m),
	}

	return json.Marshal(&s)
}
