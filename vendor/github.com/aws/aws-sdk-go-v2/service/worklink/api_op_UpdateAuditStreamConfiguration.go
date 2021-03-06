// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package worklink

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/worklink-2018-09-25/UpdateAuditStreamConfigurationRequest
type UpdateAuditStreamConfigurationInput struct {
	_ struct{} `type:"structure"`

	// The ARN of the Amazon Kinesis data stream that receives the audit events.
	AuditStreamArn *string `type:"string"`

	// The ARN of the fleet.
	//
	// FleetArn is a required field
	FleetArn *string `min:"20" type:"string" required:"true"`
}

// String returns the string representation
func (s UpdateAuditStreamConfigurationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateAuditStreamConfigurationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateAuditStreamConfigurationInput"}

	if s.FleetArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("FleetArn"))
	}
	if s.FleetArn != nil && len(*s.FleetArn) < 20 {
		invalidParams.Add(aws.NewErrParamMinLen("FleetArn", 20))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateAuditStreamConfigurationInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.AuditStreamArn != nil {
		v := *s.AuditStreamArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "AuditStreamArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.FleetArn != nil {
		v := *s.FleetArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "FleetArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/worklink-2018-09-25/UpdateAuditStreamConfigurationResponse
type UpdateAuditStreamConfigurationOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s UpdateAuditStreamConfigurationOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateAuditStreamConfigurationOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opUpdateAuditStreamConfiguration = "UpdateAuditStreamConfiguration"

// UpdateAuditStreamConfigurationRequest returns a request value for making API operation for
// Amazon WorkLink.
//
// Updates the audit stream configuration for the fleet.
//
//    // Example sending a request using UpdateAuditStreamConfigurationRequest.
//    req := client.UpdateAuditStreamConfigurationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/worklink-2018-09-25/UpdateAuditStreamConfiguration
func (c *Client) UpdateAuditStreamConfigurationRequest(input *UpdateAuditStreamConfigurationInput) UpdateAuditStreamConfigurationRequest {
	op := &aws.Operation{
		Name:       opUpdateAuditStreamConfiguration,
		HTTPMethod: "POST",
		HTTPPath:   "/updateAuditStreamConfiguration",
	}

	if input == nil {
		input = &UpdateAuditStreamConfigurationInput{}
	}

	req := c.newRequest(op, input, &UpdateAuditStreamConfigurationOutput{})
	return UpdateAuditStreamConfigurationRequest{Request: req, Input: input, Copy: c.UpdateAuditStreamConfigurationRequest}
}

// UpdateAuditStreamConfigurationRequest is the request type for the
// UpdateAuditStreamConfiguration API operation.
type UpdateAuditStreamConfigurationRequest struct {
	*aws.Request
	Input *UpdateAuditStreamConfigurationInput
	Copy  func(*UpdateAuditStreamConfigurationInput) UpdateAuditStreamConfigurationRequest
}

// Send marshals and sends the UpdateAuditStreamConfiguration API request.
func (r UpdateAuditStreamConfigurationRequest) Send(ctx context.Context) (*UpdateAuditStreamConfigurationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateAuditStreamConfigurationResponse{
		UpdateAuditStreamConfigurationOutput: r.Request.Data.(*UpdateAuditStreamConfigurationOutput),
		response:                             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateAuditStreamConfigurationResponse is the response type for the
// UpdateAuditStreamConfiguration API operation.
type UpdateAuditStreamConfigurationResponse struct {
	*UpdateAuditStreamConfigurationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateAuditStreamConfiguration request.
func (r *UpdateAuditStreamConfigurationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
