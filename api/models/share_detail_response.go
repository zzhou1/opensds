package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ShareDetailResponse share detail response
// swagger:model ShareDetailResponse
type ShareDetailResponse struct {

	// access rules status
	AccessRulesStatus string `json:"accessRulesStatus,omitempty"`

	// availability zone
	AvailabilityZone string `json:"availabilityZone,omitempty"`

	// consistency group Id
	ConsistencyGroupID string `json:"consistencyGroupId,omitempty"`

	// created at
	CreatedAt strfmt.DateTime `json:"createdAt,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// export location
	ExportLocation string `json:"exportLocation,omitempty"`

	// export locations
	ExportLocations []string `json:"exportLocations"`

	// has replicas
	HasReplicas *bool `json:"hasReplicas,omitempty"`

	// host
	Host string `json:"host,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// is public
	IsPublic *bool `json:"isPublic,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// project Id
	ProjectID string `json:"projectId,omitempty"`

	// replication type
	ReplicationType string `json:"replicationType,omitempty"`

	// share network Id
	ShareNetworkID string `json:"shareNetworkId,omitempty"`

	// share proto
	ShareProto string `json:"shareProto,omitempty"`

	// share server Id
	ShareServerID string `json:"shareServerId,omitempty"`

	// share type
	ShareType string `json:"shareType,omitempty"`

	// share type name
	ShareTypeName string `json:"shareTypeName,omitempty"`

	// size
	Size int32 `json:"size,omitempty"`

	// snapshot support
	SnapshotSupport *bool `json:"snapshotSupport,omitempty"`

	// source cgsnapshot member Id
	SourceCgsnapshotMemberID string `json:"sourceCgsnapshotMemberId,omitempty"`

	// status
	Status string `json:"status,omitempty"`

	// task state
	TaskState string `json:"taskState,omitempty"`

	// volume type
	VolumeType string `json:"volumeType,omitempty"`
}

// Validate validates this share detail response
func (m *ShareDetailResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateExportLocations(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ShareDetailResponse) validateExportLocations(formats strfmt.Registry) error {

	if swag.IsZero(m.ExportLocations) { // not required
		return nil
	}

	return nil
}