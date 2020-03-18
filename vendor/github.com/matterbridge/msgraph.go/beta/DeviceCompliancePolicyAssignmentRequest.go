// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// DeviceCompliancePolicyAssignmentRequestBuilder is request builder for DeviceCompliancePolicyAssignment
type DeviceCompliancePolicyAssignmentRequestBuilder struct{ BaseRequestBuilder }

// Request returns DeviceCompliancePolicyAssignmentRequest
func (b *DeviceCompliancePolicyAssignmentRequestBuilder) Request() *DeviceCompliancePolicyAssignmentRequest {
	return &DeviceCompliancePolicyAssignmentRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// DeviceCompliancePolicyAssignmentRequest is request for DeviceCompliancePolicyAssignment
type DeviceCompliancePolicyAssignmentRequest struct{ BaseRequest }

// Get performs GET request for DeviceCompliancePolicyAssignment
func (r *DeviceCompliancePolicyAssignmentRequest) Get(ctx context.Context) (resObj *DeviceCompliancePolicyAssignment, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for DeviceCompliancePolicyAssignment
func (r *DeviceCompliancePolicyAssignmentRequest) Update(ctx context.Context, reqObj *DeviceCompliancePolicyAssignment) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for DeviceCompliancePolicyAssignment
func (r *DeviceCompliancePolicyAssignmentRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}