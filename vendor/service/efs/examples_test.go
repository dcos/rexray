// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package efs_test

import (
	"bytes"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/efs"
)

var _ time.Duration
var _ bytes.Buffer

func ExampleEFS_CreateFileSystem() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := efs.New(sess)

	params := &efs.CreateFileSystemInput{
		CreationToken:   aws.String("CreationToken"), // Required
		PerformanceMode: aws.String("PerformanceMode"),
	}
	resp, err := svc.CreateFileSystem(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleEFS_CreateMountTarget() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := efs.New(sess)

	params := &efs.CreateMountTargetInput{
		FileSystemId: aws.String("FileSystemId"), // Required
		SubnetId:     aws.String("SubnetId"),     // Required
		IpAddress:    aws.String("IpAddress"),
		SecurityGroups: []*string{
			aws.String("SecurityGroup"), // Required
			// More values...
		},
	}
	resp, err := svc.CreateMountTarget(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleEFS_CreateTags() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := efs.New(sess)

	params := &efs.CreateTagsInput{
		FileSystemId: aws.String("FileSystemId"), // Required
		Tags: []*efs.Tag{ // Required
			{ // Required
				Key:   aws.String("TagKey"),   // Required
				Value: aws.String("TagValue"), // Required
			},
			// More values...
		},
	}
	resp, err := svc.CreateTags(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleEFS_DeleteFileSystem() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := efs.New(sess)

	params := &efs.DeleteFileSystemInput{
		FileSystemId: aws.String("FileSystemId"), // Required
	}
	resp, err := svc.DeleteFileSystem(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleEFS_DeleteMountTarget() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := efs.New(sess)

	params := &efs.DeleteMountTargetInput{
		MountTargetId: aws.String("MountTargetId"), // Required
	}
	resp, err := svc.DeleteMountTarget(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleEFS_DeleteTags() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := efs.New(sess)

	params := &efs.DeleteTagsInput{
		FileSystemId: aws.String("FileSystemId"), // Required
		TagKeys: []*string{ // Required
			aws.String("TagKey"), // Required
			// More values...
		},
	}
	resp, err := svc.DeleteTags(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleEFS_DescribeFileSystems() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := efs.New(sess)

	params := &efs.DescribeFileSystemsInput{
		CreationToken: aws.String("CreationToken"),
		FileSystemId:  aws.String("FileSystemId"),
		Marker:        aws.String("Marker"),
		MaxItems:      aws.Int64(1),
	}
	resp, err := svc.DescribeFileSystems(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleEFS_DescribeMountTargetSecurityGroups() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := efs.New(sess)

	params := &efs.DescribeMountTargetSecurityGroupsInput{
		MountTargetId: aws.String("MountTargetId"), // Required
	}
	resp, err := svc.DescribeMountTargetSecurityGroups(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleEFS_DescribeMountTargets() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := efs.New(sess)

	params := &efs.DescribeMountTargetsInput{
		FileSystemId:  aws.String("FileSystemId"),
		Marker:        aws.String("Marker"),
		MaxItems:      aws.Int64(1),
		MountTargetId: aws.String("MountTargetId"),
	}
	resp, err := svc.DescribeMountTargets(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleEFS_DescribeTags() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := efs.New(sess)

	params := &efs.DescribeTagsInput{
		FileSystemId: aws.String("FileSystemId"), // Required
		Marker:       aws.String("Marker"),
		MaxItems:     aws.Int64(1),
	}
	resp, err := svc.DescribeTags(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleEFS_ModifyMountTargetSecurityGroups() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := efs.New(sess)

	params := &efs.ModifyMountTargetSecurityGroupsInput{
		MountTargetId: aws.String("MountTargetId"), // Required
		SecurityGroups: []*string{
			aws.String("SecurityGroup"), // Required
			// More values...
		},
	}
	resp, err := svc.ModifyMountTargetSecurityGroups(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}
