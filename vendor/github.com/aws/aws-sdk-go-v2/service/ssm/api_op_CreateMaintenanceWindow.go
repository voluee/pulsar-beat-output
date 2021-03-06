// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ssm

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ssm-2014-11-06/CreateMaintenanceWindowRequest
type CreateMaintenanceWindowInput struct {
	_ struct{} `type:"structure"`

	// Enables a maintenance window task to run on managed instances, even if you
	// have not registered those instances as targets. If enabled, then you must
	// specify the unregistered instances (by instance ID) when you register a task
	// with the maintenance window.
	//
	// If you don't enable this option, then you must specify previously-registered
	// targets when you register a task with the maintenance window.
	//
	// AllowUnassociatedTargets is a required field
	AllowUnassociatedTargets *bool `type:"boolean" required:"true"`

	// User-provided idempotency token.
	ClientToken *string `min:"1" type:"string" idempotencyToken:"true"`

	// The number of hours before the end of the maintenance window that Systems
	// Manager stops scheduling new tasks for execution.
	//
	// Cutoff is a required field
	Cutoff *int64 `type:"integer" required:"true"`

	// An optional description for the maintenance window. We recommend specifying
	// a description to help you organize your maintenance windows.
	Description *string `min:"1" type:"string"`

	// The duration of the maintenance window in hours.
	//
	// Duration is a required field
	Duration *int64 `min:"1" type:"integer" required:"true"`

	// The date and time, in ISO-8601 Extended format, for when you want the maintenance
	// window to become inactive. EndDate allows you to set a date and time in the
	// future when the maintenance window will no longer run.
	EndDate *string `type:"string"`

	// The name of the maintenance window.
	//
	// Name is a required field
	Name *string `min:"3" type:"string" required:"true"`

	// The schedule of the maintenance window in the form of a cron or rate expression.
	//
	// Schedule is a required field
	Schedule *string `min:"1" type:"string" required:"true"`

	// The time zone that the scheduled maintenance window executions are based
	// on, in Internet Assigned Numbers Authority (IANA) format. For example: "America/Los_Angeles",
	// "etc/UTC", or "Asia/Seoul". For more information, see the Time Zone Database
	// (https://www.iana.org/time-zones) on the IANA website.
	ScheduleTimezone *string `type:"string"`

	// The date and time, in ISO-8601 Extended format, for when you want the maintenance
	// window to become active. StartDate allows you to delay activation of the
	// maintenance window until the specified future date.
	StartDate *string `type:"string"`

	// Optional metadata that you assign to a resource. Tags enable you to categorize
	// a resource in different ways, such as by purpose, owner, or environment.
	// For example, you might want to tag a maintenance window to identify the type
	// of tasks it will run, the types of targets, and the environment it will run
	// in. In this case, you could specify the following key name/value pairs:
	//
	//    * Key=TaskType,Value=AgentUpdate
	//
	//    * Key=OS,Value=Windows
	//
	//    * Key=Environment,Value=Production
	//
	// To add tags to an existing maintenance window, use the AddTagsToResource
	// action.
	Tags []Tag `type:"list"`
}

// String returns the string representation
func (s CreateMaintenanceWindowInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateMaintenanceWindowInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateMaintenanceWindowInput"}

	if s.AllowUnassociatedTargets == nil {
		invalidParams.Add(aws.NewErrParamRequired("AllowUnassociatedTargets"))
	}
	if s.ClientToken != nil && len(*s.ClientToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ClientToken", 1))
	}

	if s.Cutoff == nil {
		invalidParams.Add(aws.NewErrParamRequired("Cutoff"))
	}
	if s.Description != nil && len(*s.Description) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Description", 1))
	}

	if s.Duration == nil {
		invalidParams.Add(aws.NewErrParamRequired("Duration"))
	}
	if s.Duration != nil && *s.Duration < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("Duration", 1))
	}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}
	if s.Name != nil && len(*s.Name) < 3 {
		invalidParams.Add(aws.NewErrParamMinLen("Name", 3))
	}

	if s.Schedule == nil {
		invalidParams.Add(aws.NewErrParamRequired("Schedule"))
	}
	if s.Schedule != nil && len(*s.Schedule) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Schedule", 1))
	}
	if s.Tags != nil {
		for i, v := range s.Tags {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Tags", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ssm-2014-11-06/CreateMaintenanceWindowResult
type CreateMaintenanceWindowOutput struct {
	_ struct{} `type:"structure"`

	// The ID of the created maintenance window.
	WindowId *string `min:"20" type:"string"`
}

// String returns the string representation
func (s CreateMaintenanceWindowOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateMaintenanceWindow = "CreateMaintenanceWindow"

// CreateMaintenanceWindowRequest returns a request value for making API operation for
// Amazon Simple Systems Manager (SSM).
//
// Creates a new maintenance window.
//
//    // Example sending a request using CreateMaintenanceWindowRequest.
//    req := client.CreateMaintenanceWindowRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ssm-2014-11-06/CreateMaintenanceWindow
func (c *Client) CreateMaintenanceWindowRequest(input *CreateMaintenanceWindowInput) CreateMaintenanceWindowRequest {
	op := &aws.Operation{
		Name:       opCreateMaintenanceWindow,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateMaintenanceWindowInput{}
	}

	req := c.newRequest(op, input, &CreateMaintenanceWindowOutput{})
	return CreateMaintenanceWindowRequest{Request: req, Input: input, Copy: c.CreateMaintenanceWindowRequest}
}

// CreateMaintenanceWindowRequest is the request type for the
// CreateMaintenanceWindow API operation.
type CreateMaintenanceWindowRequest struct {
	*aws.Request
	Input *CreateMaintenanceWindowInput
	Copy  func(*CreateMaintenanceWindowInput) CreateMaintenanceWindowRequest
}

// Send marshals and sends the CreateMaintenanceWindow API request.
func (r CreateMaintenanceWindowRequest) Send(ctx context.Context) (*CreateMaintenanceWindowResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateMaintenanceWindowResponse{
		CreateMaintenanceWindowOutput: r.Request.Data.(*CreateMaintenanceWindowOutput),
		response:                      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateMaintenanceWindowResponse is the response type for the
// CreateMaintenanceWindow API operation.
type CreateMaintenanceWindowResponse struct {
	*CreateMaintenanceWindowOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateMaintenanceWindow request.
func (r *CreateMaintenanceWindowResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
