package ecs

import "github.com/jiangshengwu/aliyun-sdk-for-go/util"

type ImageService interface {
	DescribeImages(params map[string]string) (DescribeImagesResponse, error)
	CreateImage(params map[string]string) (CreateImageResponse, error)
	ModifyImageAttribute(params map[string]string) (ModifyImageAttributeResponse, error)
	DeleteImage(params map[string]string) (DeleteImageResponse, error)
	CopyImage(params map[string]string) (CopyImageResponse, error)
	CancelCopyImage(params map[string]string) (CancelCopyImageResponse, error)
	ModifyImageSharePermission(params map[string]string) (ModifyImageSharePermissionResponse, error)
	DescribeImageSharePermission(params map[string]string) (DescribeImageSharePermissionResponse, error)
}

type ImageOperator struct {
	Common *CommonParam
}

// Response struct for DescribeImages
type DescribeImagesResponse struct {
	util.ErrorResponse
	util.PageResponse
	AllImages Images `json:"Images"`
}

type Images struct {
	AllImage []ImageType `json:"Image"`
}

// See http://docs.aliyun.com/?spm=5176.775974174.2.4.BYfRJ2#/ecs/open-api/datatype&imagetype
type ImageType struct {
	ImageId               string             `json:"ImageId"`
	ImageVersion          string             `json:"ImageVersion"`
	Architecture          string             `json:"Architecture"`
	ImageName             string             `json:"ImageName"`
	Description           string             `json:"Description"`
	Size                  int                `json:"Size"`
	ImageOwnerAlias       string             `json:"ImageOwnerAlias"`
	OSName                string             `json:"OSName"`
	AllDiskDeviceMappings DiskDeviceMappings `json:"DiskDeviceMappings"`
	ProductCode           string             `json:"ProductCode"`
	IsSubscribed          string             `json:"IsSubscribed"`
	Progress              string             `json:"Progress"`
	Status                string             `json:"Status"`
	CreationTime          string             `json:"CreationTime"`
}

type DiskDeviceMappings struct {
	AllDiskDeviceMapping []DiskDeviceMapping `json:"DiskDeviceMapping"`
}

// Response struct for CreateImage
type CreateImageResponse struct {
	util.ErrorResponse
	ImageId string `json:"ImageId"`
}

// Response struct for ModifyImageAttribute
type ModifyImageAttributeResponse struct {
	util.ErrorResponse
}

// Response struct for DeleteImage
type DeleteImageResponse struct {
	util.ErrorResponse
}

// Response struct for CopyImage
type CopyImageResponse struct {
	util.ErrorResponse
	ImageId string `json:"ImageId"`
}

// See http://docs.aliyun.com/?spm=5176.775974174.2.4.BYfRJ2#/ecs/open-api/datatype&diskdevicemapping
type DiskDeviceMapping struct {
	SnapshotId string `json:"SnapshotId"`
	Size       string `json:"Size"`
	Device     string `json:"Device"`
}

// Response struct for CancelCopyImage
type CancelCopyImageResponse struct {
	util.ErrorResponse
}

// Response struct for ModifyImageSharePermission
type ModifyImageSharePermissionResponse struct {
	util.ErrorResponse
}

// Response struct for DescribeImageSharePermission
type DescribeImageSharePermissionResponse struct {
	util.ErrorResponse
	util.PageResponse
	ImageId        string      `json:"ImageId"`
	RegionId       string      `json:"RegionId"`
	AllShareGroups ShareGroups `json:"ShareGroups"`
	AllAccounts    Accounts    `json:"Accounts"`
}

type ShareGroups struct {
	AllShareGroup []ShareGroupType `json:"ShareGroup"`
}

// See http://docs.aliyun.com/?spm=5176.775974174.2.4.BYfRJ2#/ecs/open-api/datatype&shareGroupType
type ShareGroupType struct {
	Group string `json:"Group"`
}

type Accounts struct {
	AllAccount []AccountType `json:"Account"`
}

//See http://docs.aliyun.com/?spm=5176.775974174.2.4.BYfRJ2#/ecs/open-api/datatype&accountType
type AccountType struct {
	AliyunId string `json:"AliyunId"`
}

func (op *ImageOperator) DescribeImages(params map[string]string) (DescribeImagesResponse, error) {
	var resp DescribeImagesResponse
	err := op.Common.Request(GetFuncName(1), params, &resp)
	return resp, err
}

func (op *ImageOperator) CreateImage(params map[string]string) (CreateImageResponse, error) {
	var resp CreateImageResponse
	err := op.Common.Request(GetFuncName(1), params, &resp)
	return resp, err
}

func (op *ImageOperator) ModifyImageAttribute(params map[string]string) (ModifyImageAttributeResponse, error) {
	var resp ModifyImageAttributeResponse
	err := op.Common.Request(GetFuncName(1), params, &resp)
	return resp, err
}

func (op *ImageOperator) DeleteImage(params map[string]string) (DeleteImageResponse, error) {
	var resp DeleteImageResponse
	err := op.Common.Request(GetFuncName(1), params, &resp)
	return resp, err
}

func (op *ImageOperator) CopyImage(params map[string]string) (CopyImageResponse, error) {
	var resp CopyImageResponse
	err := op.Common.Request(GetFuncName(1), params, &resp)
	return resp, err
}

func (op *ImageOperator) CancelCopyImage(params map[string]string) (CancelCopyImageResponse, error) {
	var resp CancelCopyImageResponse
	err := op.Common.Request(GetFuncName(1), params, &resp)
	return resp, err
}

func (op *ImageOperator) ModifyImageSharePermission(params map[string]string) (ModifyImageSharePermissionResponse, error) {
	var resp ModifyImageSharePermissionResponse
	err := op.Common.Request(GetFuncName(1), params, &resp)
	return resp, err
}

func (op *ImageOperator) DescribeImageSharePermission(params map[string]string) (DescribeImageSharePermissionResponse, error) {
	var resp DescribeImageSharePermissionResponse
	err := op.Common.Request(GetFuncName(1), params, &resp)
	return resp, err
}
