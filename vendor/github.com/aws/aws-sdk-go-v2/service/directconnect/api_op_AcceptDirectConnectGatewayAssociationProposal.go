// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package directconnect

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/directconnect-2012-10-25/AcceptDirectConnectGatewayAssociationProposalRequest
type AcceptDirectConnectGatewayAssociationProposalInput struct {
	_ struct{} `type:"structure"`

	// The ID of the AWS account that owns the virtual private gateway or transit
	// gateway.
	//
	// AssociatedGatewayOwnerAccount is a required field
	AssociatedGatewayOwnerAccount *string `locationName:"associatedGatewayOwnerAccount" type:"string" required:"true"`

	// The ID of the Direct Connect gateway.
	//
	// DirectConnectGatewayId is a required field
	DirectConnectGatewayId *string `locationName:"directConnectGatewayId" type:"string" required:"true"`

	// Overrides the Amazon VPC prefixes advertised to the Direct Connect gateway.
	//
	// For information about how to set the prefixes, see Allowed Prefixes (https://docs.aws.amazon.com/directconnect/latest/UserGuide/multi-account-associate-vgw.html#allowed-prefixes)
	// in the AWS Direct Connect User Guide.
	OverrideAllowedPrefixesToDirectConnectGateway []RouteFilterPrefix `locationName:"overrideAllowedPrefixesToDirectConnectGateway" type:"list"`

	// The ID of the request proposal.
	//
	// ProposalId is a required field
	ProposalId *string `locationName:"proposalId" type:"string" required:"true"`
}

// String returns the string representation
func (s AcceptDirectConnectGatewayAssociationProposalInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AcceptDirectConnectGatewayAssociationProposalInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "AcceptDirectConnectGatewayAssociationProposalInput"}

	if s.AssociatedGatewayOwnerAccount == nil {
		invalidParams.Add(aws.NewErrParamRequired("AssociatedGatewayOwnerAccount"))
	}

	if s.DirectConnectGatewayId == nil {
		invalidParams.Add(aws.NewErrParamRequired("DirectConnectGatewayId"))
	}

	if s.ProposalId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ProposalId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/directconnect-2012-10-25/AcceptDirectConnectGatewayAssociationProposalResult
type AcceptDirectConnectGatewayAssociationProposalOutput struct {
	_ struct{} `type:"structure"`

	// Information about an association between a Direct Connect gateway and a virtual
	// private gateway or transit gateway.
	DirectConnectGatewayAssociation *DirectConnectGatewayAssociation `locationName:"directConnectGatewayAssociation" type:"structure"`
}

// String returns the string representation
func (s AcceptDirectConnectGatewayAssociationProposalOutput) String() string {
	return awsutil.Prettify(s)
}

const opAcceptDirectConnectGatewayAssociationProposal = "AcceptDirectConnectGatewayAssociationProposal"

// AcceptDirectConnectGatewayAssociationProposalRequest returns a request value for making API operation for
// AWS Direct Connect.
//
// Accepts a proposal request to attach a virtual private gateway or transit
// gateway to a Direct Connect gateway.
//
//    // Example sending a request using AcceptDirectConnectGatewayAssociationProposalRequest.
//    req := client.AcceptDirectConnectGatewayAssociationProposalRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/directconnect-2012-10-25/AcceptDirectConnectGatewayAssociationProposal
func (c *Client) AcceptDirectConnectGatewayAssociationProposalRequest(input *AcceptDirectConnectGatewayAssociationProposalInput) AcceptDirectConnectGatewayAssociationProposalRequest {
	op := &aws.Operation{
		Name:       opAcceptDirectConnectGatewayAssociationProposal,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &AcceptDirectConnectGatewayAssociationProposalInput{}
	}

	req := c.newRequest(op, input, &AcceptDirectConnectGatewayAssociationProposalOutput{})
	return AcceptDirectConnectGatewayAssociationProposalRequest{Request: req, Input: input, Copy: c.AcceptDirectConnectGatewayAssociationProposalRequest}
}

// AcceptDirectConnectGatewayAssociationProposalRequest is the request type for the
// AcceptDirectConnectGatewayAssociationProposal API operation.
type AcceptDirectConnectGatewayAssociationProposalRequest struct {
	*aws.Request
	Input *AcceptDirectConnectGatewayAssociationProposalInput
	Copy  func(*AcceptDirectConnectGatewayAssociationProposalInput) AcceptDirectConnectGatewayAssociationProposalRequest
}

// Send marshals and sends the AcceptDirectConnectGatewayAssociationProposal API request.
func (r AcceptDirectConnectGatewayAssociationProposalRequest) Send(ctx context.Context) (*AcceptDirectConnectGatewayAssociationProposalResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &AcceptDirectConnectGatewayAssociationProposalResponse{
		AcceptDirectConnectGatewayAssociationProposalOutput: r.Request.Data.(*AcceptDirectConnectGatewayAssociationProposalOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// AcceptDirectConnectGatewayAssociationProposalResponse is the response type for the
// AcceptDirectConnectGatewayAssociationProposal API operation.
type AcceptDirectConnectGatewayAssociationProposalResponse struct {
	*AcceptDirectConnectGatewayAssociationProposalOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// AcceptDirectConnectGatewayAssociationProposal request.
func (r *AcceptDirectConnectGatewayAssociationProposalResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
