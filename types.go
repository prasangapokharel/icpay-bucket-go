// Package icpaybucket is a client for the ICPay Bucket canister.
package icpaybucket

import (
	"github.com/aviate-labs/agent-go/candid/idl"
)

const DefaultCanisterID = "6vbhm-nqaaa-aaaan-q6muq-cai"
const SingleMaxBytes = 1_850_000

type ApiResult[T any] struct {
	Ok  *T      `ic:"ok,variant"`
	Err *string `ic:"err,variant"`
}

type FilePublic struct {
	ID          string   `ic:"id"`
	Path        string   `ic:"path"`
	Name        string   `ic:"name"`
	Size        idl.Nat  `ic:"size"`
	ContentType string   `ic:"contentType"`
	CreatedAt   idl.Int  `ic:"createdAt"`
	UpdatedAt   *idl.Int `ic:"updatedAt"`
	Metadata    *string  `ic:"metadata"`
	Tags        []string `ic:"tags"`
	PublicURL   *string  `ic:"publicUrl"`
}

type FileListPage struct {
	Items    []FilePublic `ic:"items"`
	Total    idl.Nat      `ic:"total"`
	Page     idl.Nat      `ic:"page"`
	PageSize idl.Nat      `ic:"pageSize"`
}

type FilePathOp struct {
	Source      string `ic:"source"`
	Destination string `ic:"destination"`
}

type UploadStatusPublic struct {
	UploadID     string  `ic:"uploadId"`
	BucketID     string  `ic:"bucketId"`
	Path         string  `ic:"path"`
	TotalSize    idl.Nat `ic:"totalSize"`
	UploadedSize idl.Nat `ic:"uploadedSize"`
	ChunkSize    idl.Nat `ic:"chunkSize"`
	Status       string  `ic:"status"`
	CreatedAt    idl.Int `ic:"createdAt"`
}

type ApiKeyPermissions struct {
	Read   bool `ic:"read"`
	Write  bool `ic:"write"`
	Delete bool `ic:"delete"`
}

type resultFileID = ApiResult[string]
type resultUnit = ApiResult[idl.Null]
type resultBlob = ApiResult[[]byte]
type resultString = ApiResult[string]
type resultFileList = ApiResult[FileListPage]
type resultFilePublic = ApiResult[FilePublic]
type resultBool = ApiResult[bool]
type resultNat = ApiResult[idl.Nat]
