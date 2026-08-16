// Package icpaybucket is a minimal client for the ICPay Bucket canister —
// on-chain encrypted file storage on the Internet Computer.
package icpaybucket

import (
	"github.com/aviate-labs/agent-go/candid/idl"
)

// DefaultCanisterID is the live ICPay backend canister on mainnet.
const DefaultCanisterID = "6vbhm-nqaaa-aaaan-q6muq-cai"

// ApiResult is a canister variant { ok: T } | { err: Text }.
type ApiResult[T any] struct {
	Ok  *T      `ic:"ok,variant"`
	Err *string `ic:"err,variant"`
}

// FilePublic is the canister FilePublic record.
type FilePublic struct {
	ID          string  `ic:"id"`
	Path        string  `ic:"path"`
	Size        idl.Nat `ic:"size"`
	ContentType string  `ic:"contentType"`
	CreatedAt   idl.Int `ic:"createdAt"`
	PublicURL   *string `ic:"publicUrl"`
}

// FileListPage is the canister FileListPage record.
type FileListPage struct {
	Items    []FilePublic `ic:"items"`
	Total    idl.Nat      `ic:"total"`
	Page     idl.Nat      `ic:"page"`
	PageSize idl.Nat      `ic:"pageSize"`
}

type resultFileID = ApiResult[string]
type resultUnit = ApiResult[idl.Null]
type resultBlob = ApiResult[[]byte]
type resultString = ApiResult[string]
type resultFileList = ApiResult[FileListPage]

// SingleMaxBytes is the ceiling for a single uploadFile call (~1.85 MB).
const SingleMaxBytes = 1_850_000
