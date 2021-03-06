// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Application Migration Service API
//
// API for the Application Migration service. Use this API to migrate applications from Oracle Cloud Infrastructure - Classic to Oracle Cloud Infrastructure.
//

package applicationmigration

import (
	"github.com/oracle/oci-go-sdk/common"
)

// SourceApplicationSummary An application running in the source environment that is available for export.
type SourceApplicationSummary struct {

	// The name of the application
	Name *string `mandatory:"false" json:"name"`

	// The type of application
	Type MigrationTypesEnum `mandatory:"false" json:"type,omitempty"`

	// Unique identifier (OCID) for the Source to which the application belongs
	SourceId *string `mandatory:"false" json:"sourceId"`

	// The version of the application server
	Version *string `mandatory:"false" json:"version"`

	// The current application running state
	State *string `mandatory:"false" json:"state"`
}

func (m SourceApplicationSummary) String() string {
	return common.PointerString(m)
}
