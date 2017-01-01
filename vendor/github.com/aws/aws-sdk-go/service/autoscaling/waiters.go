// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package autoscaling

import (
	"github.com/aws/aws-sdk-go/private/waiter"
)

// WaitUntilGroupExists uses the Auto Scaling API operation
// DescribeAutoScalingGroups to wait for a condition to be met before returning.
// If the condition is not meet within the max attempt window an error will
// be returned.
func (c *AutoScaling) WaitUntilGroupExists(input *DescribeAutoScalingGroupsInput) error {
	waiterCfg := waiter.Config{
		Operation:   "DescribeAutoScalingGroups",
		Delay:       5,
		MaxAttempts: 10,
		Acceptors: []waiter.WaitAcceptor{
			{
				State:    "success",
				Matcher:  "path",
				Argument: "length(AutoScalingGroups) > `0`",
				Expected: true,
			},
			{
				State:    "retry",
				Matcher:  "path",
				Argument: "length(AutoScalingGroups) > `0`",
				Expected: false,
			},
		},
	}

	w := waiter.Waiter{
		Client: c,
		Input:  input,
		Config: waiterCfg,
	}
	return w.Wait()
}

// WaitUntilGroupInService uses the Auto Scaling API operation
// DescribeAutoScalingGroups to wait for a condition to be met before returning.
// If the condition is not meet within the max attempt window an error will
// be returned.
func (c *AutoScaling) WaitUntilGroupInService(input *DescribeAutoScalingGroupsInput) error {
	waiterCfg := waiter.Config{
		Operation:   "DescribeAutoScalingGroups",
		Delay:       15,
		MaxAttempts: 40,
		Acceptors: []waiter.WaitAcceptor{
			{
				State:    "success",
				Matcher:  "path",
				Argument: "contains(AutoScalingGroups[].[length(Instances[?LifecycleState=='InService']) >= MinSize][], `false`)",
				Expected: false,
			},
			{
				State:    "retry",
				Matcher:  "path",
				Argument: "contains(AutoScalingGroups[].[length(Instances[?LifecycleState=='InService']) >= MinSize][], `false`)",
				Expected: true,
			},
		},
	}

	w := waiter.Waiter{
		Client: c,
		Input:  input,
		Config: waiterCfg,
	}
	return w.Wait()
}

// WaitUntilGroupNotExists uses the Auto Scaling API operation
// DescribeAutoScalingGroups to wait for a condition to be met before returning.
// If the condition is not meet within the max attempt window an error will
// be returned.
func (c *AutoScaling) WaitUntilGroupNotExists(input *DescribeAutoScalingGroupsInput) error {
	waiterCfg := waiter.Config{
		Operation:   "DescribeAutoScalingGroups",
		Delay:       15,
		MaxAttempts: 40,
		Acceptors: []waiter.WaitAcceptor{
			{
				State:    "success",
				Matcher:  "path",
				Argument: "length(AutoScalingGroups) > `0`",
				Expected: false,
			},
			{
				State:    "retry",
				Matcher:  "path",
				Argument: "length(AutoScalingGroups) > `0`",
				Expected: true,
			},
		},
	}

	w := waiter.Waiter{
		Client: c,
		Input:  input,
		Config: waiterCfg,
	}
	return w.Wait()
}
