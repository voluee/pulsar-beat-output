// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ssm

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ssm-2014-11-06/UpdateDocumentRequest
type UpdateDocumentInput struct {
	_ struct{} `type:"structure"`

	// A list of key and value pairs that describe attachments to a version of a
	// document.
	Attachments []AttachmentsSource `type:"list"`

	// A valid JSON or YAML string.
	//
	// Content is a required field
	Content *string `min:"1" type:"string" required:"true"`

	// Specify the document format for the new document version. Systems Manager
	// supports JSON and YAML documents. JSON is the default format.
	DocumentFormat DocumentFormat `type:"string" enum:"true"`

	// The version of the document that you want to update.
	DocumentVersion *string `type:"string"`

	// The name of the document that you want to update.
	//
	// Name is a required field
	Name *string `type:"string" required:"true"`

	// Specify a new target type for the document.
	TargetType *string `type:"string"`

	// An optional field specifying the version of the artifact you are updating
	// with the document. For example, "Release 12, Update 6". This value is unique
	// across all versions of a document, and cannot be changed.
	VersionName *string `type:"string"`
}

// String returns the string representation
func (s UpdateDocumentInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateDocumentInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateDocumentInput"}

	if s.Content == nil {
		invalidParams.Add(aws.NewErrParamRequired("Content"))
	}
	if s.Content != nil && len(*s.Content) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Content", 1))
	}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}
	if s.Attachments != nil {
		for i, v := range s.Attachments {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Attachments", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ssm-2014-11-06/UpdateDocumentResult
type UpdateDocumentOutput struct {
	_ struct{} `type:"structure"`

	// A description of the document that was updated.
	DocumentDescription *DocumentDescription `type:"structure"`
}

// String returns the string representation
func (s UpdateDocumentOutput) String() string {
	return awsutil.Prettify(s)
}

const opUpdateDocument = "UpdateDocument"

// UpdateDocumentRequest returns a request value for making API operation for
// Amazon Simple Systems Manager (SSM).
//
// Updates one or more values for an SSM document.
//
//    // Example sending a request using UpdateDocumentRequest.
//    req := client.UpdateDocumentRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ssm-2014-11-06/UpdateDocument
func (c *Client) UpdateDocumentRequest(input *UpdateDocumentInput) UpdateDocumentRequest {
	op := &aws.Operation{
		Name:       opUpdateDocument,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateDocumentInput{}
	}

	req := c.newRequest(op, input, &UpdateDocumentOutput{})
	return UpdateDocumentRequest{Request: req, Input: input, Copy: c.UpdateDocumentRequest}
}

// UpdateDocumentRequest is the request type for the
// UpdateDocument API operation.
type UpdateDocumentRequest struct {
	*aws.Request
	Input *UpdateDocumentInput
	Copy  func(*UpdateDocumentInput) UpdateDocumentRequest
}

// Send marshals and sends the UpdateDocument API request.
func (r UpdateDocumentRequest) Send(ctx context.Context) (*UpdateDocumentResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateDocumentResponse{
		UpdateDocumentOutput: r.Request.Data.(*UpdateDocumentOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateDocumentResponse is the response type for the
// UpdateDocument API operation.
type UpdateDocumentResponse struct {
	*UpdateDocumentOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateDocument request.
func (r *UpdateDocumentResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}