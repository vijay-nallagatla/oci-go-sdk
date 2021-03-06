// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Data Transfer Service API
//
// Data Transfer Service API Specification
//

package dts

import (
	"github.com/oracle/oci-go-sdk/common"
)

// TransferApplianceEncryptionPassphrase The representation of TransferApplianceEncryptionPassphrase
type TransferApplianceEncryptionPassphrase struct {
	EncryptionPassphrase *string `mandatory:"false" json:"encryptionPassphrase"`
}

func (m TransferApplianceEncryptionPassphrase) String() string {
	return common.PointerString(m)
}
