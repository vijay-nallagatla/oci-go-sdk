// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package ffsw

import (
    "github.com/oracle/oci-go-sdk/common"
    "net/http"
)

// UpdateMountTargetRequest wrapper for the UpdateMountTarget operation
type UpdateMountTargetRequest struct {
        
 // The OCID of the mount target. This feature is currently in preview and may change before public release. Do not use it for production workloads.
        MountTargetId *string `mandatory:"true" contributesTo:"path" name:"mountTargetId"`
        
 // Details object for updating a mount target. 
         UpdateMountTargetDetails `contributesTo:"body"`
        
 // For optimistic concurrency control. In the PUT or DELETE call
 // for a resource, set the `if-match` parameter to the value of the
 // etag from a previous GET or POST response for that resource.
 // The resource will be updated or deleted only if the etag you
 // provide matches the resource's current etag value. 
        IfMatch *string `mandatory:"false" contributesTo:"header" name:"if-match"`
}

func (request UpdateMountTargetRequest) String() string {
    return common.PointerString(request)
}

// UpdateMountTargetResponse wrapper for the UpdateMountTarget operation
type UpdateMountTargetResponse struct {

    // The underlying http response
    RawResponse *http.Response
    
 // The MountTarget instance
     MountTarget `presentIn:"body"`

    
 // For optimistic concurrency control. See `if-match`.
    Etag *string `presentIn:"header" name:"etag"`
    
 // Unique Oracle-assigned identifier for the request. If
 // you need to contact Oracle about a particular request,
 // please provide the request ID.
    OpcRequestId *string `presentIn:"header" name:"opc-request-id"`


}

func (response UpdateMountTargetResponse) String() string {
    return common.PointerString(response)
}

