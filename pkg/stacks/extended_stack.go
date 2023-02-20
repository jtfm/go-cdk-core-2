package stacks

import (
	"fmt"
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

// ExtendedStack is a stack that adds some additional functionality to the standard stack

type ExtendedStackProps struct {
	awscdk.StackProps
	Version *string
	Build *string
}

func NewExtendedStack(
	scope constructs.Construct, 
	id *string, 
	props *ExtendedStackProps) awscdk.Stack {

	props.StackName = jsii.String(fmt.Sprintf(
		"%s-%s-%s", *props.StackName, *props.Version, *props.Build))

	stack := awscdk.NewStack(scope, id, &props.StackProps)

	return stack
}