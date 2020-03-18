// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// TeamsTemplateRequestBuilder is request builder for TeamsTemplate
type TeamsTemplateRequestBuilder struct{ BaseRequestBuilder }

// Request returns TeamsTemplateRequest
func (b *TeamsTemplateRequestBuilder) Request() *TeamsTemplateRequest {
	return &TeamsTemplateRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// TeamsTemplateRequest is request for TeamsTemplate
type TeamsTemplateRequest struct{ BaseRequest }

// Get performs GET request for TeamsTemplate
func (r *TeamsTemplateRequest) Get(ctx context.Context) (resObj *TeamsTemplate, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for TeamsTemplate
func (r *TeamsTemplateRequest) Update(ctx context.Context, reqObj *TeamsTemplate) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for TeamsTemplate
func (r *TeamsTemplateRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}