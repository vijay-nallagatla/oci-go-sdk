// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package ffsw

import (
    "github.com/oracle/oci-go-sdk/common"
    "net/http"
)

// DeleteMountTargetRequest wrapper for the DeleteMountTarget operation
type DeleteMountTargetRequest struct {
        
 // The OCID of the mount target. This feature is currently in preview and may change before public release. Do not use it for production workloads.
        MountTargetId *string `mandatory:"true" contributesTo:"path" name:"mountTargetId"`
        
 // For optimistic concurrency control. In the PUT or DELETE call
 // for a resource, set the `if-match` parameter to the value of the
 // etag from a previous GET or POST response for that resource.
 // The resource will be updated or deleted only if the etag you
 // provide matches the resource's current etag value. 
        IfMatch *string `mandatory:"false" contributesTo:"header" name:"if-match"`
}

func (request DeleteMountTargetRequest) String() string {
    return common.PointerString(request)
}

// DeleteMountTargetResponse wrapper for the DeleteMountTarget operation
type DeleteMountTargetResponse struct {

    // The underlying http response
    RawResponse *http.Response

    
 // Unique Oracle-assigned identifier for the request. If
 // you need to contact Oracle about a particular request,
 // please provide the request ID.
    OpcRequestId *string `presentIn:"header" name:"opc-request-id"`


}

func (response DeleteMountTargetResponse) String() string {
    return common.PointerString(response)
}
