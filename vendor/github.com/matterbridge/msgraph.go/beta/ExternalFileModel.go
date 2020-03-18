// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "time"

// ExternalFile undocumented
type ExternalFile struct {
	// ExternalItem is the base model of ExternalFile
	ExternalItem
	// CreatedDateTime undocumented
	CreatedDateTime *time.Time `json:"createdDateTime,omitempty"`
	// ModifiedDateTime undocumented
	ModifiedDateTime *time.Time `json:"modifiedDateTime,omitempty"`
	// CreatedBy undocumented
	CreatedBy *string `json:"createdBy,omitempty"`
	// LastModifiedBy undocumented
	LastModifiedBy *string `json:"lastModifiedBy,omitempty"`
	// Title undocumented
	Title *string `json:"title,omitempty"`
	// URL undocumented
	URL *string `json:"url,omitempty"`
	// Name undocumented
	Name *string `json:"name,omitempty"`
	// Extension undocumented
	Extension *string `json:"extension,omitempty"`
	// Size undocumented
	Size *int `json:"size,omitempty"`
}