// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Identity and Access Management Service API
//
// APIs for managing users, groups, compartments, and policies.
//

package identity

import (
	"time"
)

// Policy. A document that specifies the type of access a group has to the resources in a compartment. For information about
// policies and other IAM Service components, see
// [Overview of the IAM Service](/Content/Identity/Concepts/overview.htm). If you're new to policies, see
// [Getting Started with Policies](/Content/Identity/Concepts/policygetstarted.htm).
// The word "policy" is used by people in different ways:
//   * An individual statement written in the policy language
//   * A collection of statements in a single, named "policy" document (which has an Oracle Cloud ID (OCID) assigned to it)
//   * The overall body of policies your organization uses to control access to resources
// To use any of the API operations, you must be authorized in an IAM policy. If you're not authorized,
// talk to an administrator.
type Policy struct {

	// The OCID of the policy.
	ID string `mandatory:"true" json:"id,omitempty"`

	// The OCID of the compartment containing the policy (either the tenancy or another compartment).
	CompartmentID string `mandatory:"true" json:"compartmentId,omitempty"`

	// The name you assign to the policy during creation. The name must be unique across all policies
	// in the tenancy and cannot be changed.
	Name string `mandatory:"true" json:"name,omitempty"`

	// An array of one or more policy statements written in the policy language.
	Statements []string `mandatory:"true" json:"statements,omitempty"`

	// The description you assign to the policy. Does not have to be unique, and it's changeable.
	Description string `mandatory:"true" json:"description,omitempty"`

	// Date and time the policy was created, in the format defined by RFC3339.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated time.Time `mandatory:"true" json:"timeCreated,omitempty"`

	// The policy's current state. After creating a policy, make sure its `lifecycleState` changes from CREATING to
	// ACTIVE before using it.
	LifecycleState string `mandatory:"true" json:"lifecycleState,omitempty"`

	// The detailed status of INACTIVE lifecycleState.
	InactiveStatus int64 `mandatory:"false" json:"inactiveStatus,omitempty"`

	// The version of the policy. If null or set to an empty string, when a request comes in for authorization, the
	// policy will be evaluated according to the current behavior of the services at that moment. If set to a particular
	// date (YYYY-MM-DD), the policy will be evaluated according to the behavior of the services on that date.
	VersionDate time.Time `mandatory:"false" json:"versionDate,omitempty"`
}