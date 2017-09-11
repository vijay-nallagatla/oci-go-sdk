// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package identity

// CreateOrResetUIPasswordRequest wrapper for the CreateOrResetUIPassword operation
type CreateOrResetUIPasswordRequest struct {

	// The OCID of the user.
	UserID string `mandatory:"true" contributesTo:"path"`

	// A token that uniquely identifies a request so it can be retried in case of a timeout or
	// server error without risk of executing that same action again. Retry tokens expire after 24
	// hours, but can be invalidated before then due to conflicting operations (e.g., if a resource
	// has been deleted and purged from the system, then a retry of the original creation request
	// may be rejected).
	OpcRetryToken string `mandatory:"false" contributesTo:"header" name:"opc-retry-token"`
}

// CreateOrResetUIPasswordResponse wrapper for the CreateOrResetUIPassword operation
type CreateOrResetUIPasswordResponse struct {

	// The UiPassword instance
	UiPassword

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestID string

	// For optimistic concurrency control. See `if-match`.
	Etag string
}
