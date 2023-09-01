package cluster

import (
	"go.temporal.io/api/workflowservice/v1"
	"go.temporal.io/server/common/headers"
)

// SystemInfo returns the current system info including version and
// capabilities.
func SystemInfo() *workflowservice.GetSystemInfoResponse {
	return &workflowservice.GetSystemInfoResponse{
		ServerVersion: headers.ServerVersion,
		// Capabilities should be added as needed. In most cases, capabilities are
		// hardcoded boolean true values since older servers will respond with a
		// form of this message without the field which is implied false. That a
		// server has a capability doesn't necessarily mean it's enabled.
		Capabilities: &workflowservice.GetSystemInfoResponse_Capabilities{
			SignalAndQueryHeader:            true,
			InternalErrorDifferentiation:    true,
			ActivityFailureIncludeHeartbeat: true,
			SupportsSchedules:               true,
			EncodedFailureAttributes:        true,
			UpsertMemo:                      true,
			EagerWorkflowStart:              true,
			SdkMetadata:                     true,
			BuildIdBasedVersioning:          true,
			CountGroupByExecutionStatus:     true,
		},
	}
}
