package icpaybucket

import (
	"fmt"

	"github.com/aviate-labs/agent-go/candid/idl"
)

func (c *Client) UploadFile(bucketID, path string, data []byte, contentType, apiKey string) (resultFileID, error) {
	if len(data) > SingleMaxBytes {
		return resultFileID{}, fmt.Errorf("file too large for single upload")
	}
	out := resultFileID{}
	err := c.agent.Call(
		c.canisterID, "uploadFile",
		[]any{bucketID, path, data, contentType, c.resolveAPIKey(apiKey)},
		[]any{&out},
	)
	return out, err
}

func (c *Client) DeleteFile(bucketID, path, apiKey string) (resultUnit, error) {
	out := resultUnit{}
	err := c.agent.Call(
		c.canisterID, "deleteFile",
		[]any{bucketID, path, c.resolveAPIKey(apiKey)},
		[]any{&out},
	)
	return out, err
}

func (c *Client) ListFiles(bucketID string, page, pageSize uint64, apiKey string) (resultFileList, error) {
	out := resultFileList{}
	err := c.agent.Query(
		c.canisterID, "listFiles",
		[]any{bucketID, nat(page), nat(pageSize), c.resolveAPIKey(apiKey)},
		[]any{&out},
	)
	return out, err
}

func (c *Client) DownloadFile(bucketID, path, apiKey string) (resultBlob, error) {
	out := resultBlob{}
	err := c.agent.Query(
		c.canisterID, "downloadFile",
		[]any{bucketID, path, c.resolveAPIKey(apiKey)},
		[]any{&out},
	)
	return out, err
}

func (c *Client) GetPublicFileURL(bucketID, path string) (resultString, error) {
	out := resultString{}
	err := c.agent.Query(
		c.canisterID, "getPublicFileUrl",
		[]any{bucketID, path},
		[]any{&out},
	)
	return out, err
}

func (c *Client) GetFile(bucketID, path, apiKey string) (resultFilePublic, error) {
	out := resultFilePublic{}
	err := c.agent.Query(
		c.canisterID, "getFile",
		[]any{bucketID, path, c.resolveAPIKey(apiKey)},
		[]any{&out},
	)
	return out, err
}

func (c *Client) FileExists(bucketID, path, apiKey string) (resultBool, error) {
	out := resultBool{}
	err := c.agent.Query(
		c.canisterID, "fileExists",
		[]any{bucketID, path, c.resolveAPIKey(apiKey)},
		[]any{&out},
	)
	return out, err
}

func (c *Client) SearchFiles(bucketID, query string, page, pageSize uint64, apiKey string) (resultFileList, error) {
	out := resultFileList{}
	err := c.agent.Query(
		c.canisterID, "searchFiles",
		[]any{bucketID, query, nat(page), nat(pageSize), c.resolveAPIKey(apiKey)},
		[]any{&out},
	)
	return out, err
}

func (c *Client) ListFolder(bucketID, prefix string, page, pageSize uint64, apiKey string) (resultFileList, error) {
	out := resultFileList{}
	err := c.agent.Query(
		c.canisterID, "listFolder",
		[]any{bucketID, prefix, nat(page), nat(pageSize), c.resolveAPIKey(apiKey)},
		[]any{&out},
	)
	return out, err
}

func (c *Client) GetFileMetadata(bucketID, path, apiKey string) (resultString, error) {
	out := resultString{}
	err := c.agent.Query(
		c.canisterID, "getFileMetadata",
		[]any{bucketID, path, c.resolveAPIKey(apiKey)},
		[]any{&out},
	)
	return out, err
}

func (c *Client) MoveFile(bucketID, source, destination, apiKey string) (resultFilePublic, error) {
	out := resultFilePublic{}
	err := c.agent.Call(
		c.canisterID, "moveFile",
		[]any{bucketID, source, destination, c.resolveAPIKey(apiKey)},
		[]any{&out},
	)
	return out, err
}

func (c *Client) CopyFile(bucketID, source, destination, apiKey string) (resultFilePublic, error) {
	out := resultFilePublic{}
	err := c.agent.Call(
		c.canisterID, "copyFile",
		[]any{bucketID, source, destination, c.resolveAPIKey(apiKey)},
		[]any{&out},
	)
	return out, err
}

func (c *Client) SetFileTags(bucketID, path string, tags []string, apiKey string) (resultFilePublic, error) {
	out := resultFilePublic{}
	err := c.agent.Call(
		c.canisterID, "setFileTags",
		[]any{bucketID, path, tags, c.resolveAPIKey(apiKey)},
		[]any{&out},
	)
	return out, err
}

func (c *Client) SetFileMetadata(bucketID, path, metadata, apiKey string) (resultFilePublic, error) {
	out := resultFilePublic{}
	err := c.agent.Call(
		c.canisterID, "setFileMetadata",
		[]any{bucketID, path, metadata, c.resolveAPIKey(apiKey)},
		[]any{&out},
	)
	return out, err
}

func (c *Client) BulkDeleteFiles(bucketID string, paths []string, apiKey string) (resultNat, error) {
	out := resultNat{}
	err := c.agent.Call(
		c.canisterID, "bulkDeleteFiles",
		[]any{bucketID, paths, c.resolveAPIKey(apiKey)},
		[]any{&out},
	)
	return out, err
}

func (c *Client) BulkMoveFiles(bucketID string, ops []FilePathOp, apiKey string) (resultNat, error) {
	out := resultNat{}
	err := c.agent.Call(
		c.canisterID, "bulkMoveFiles",
		[]any{bucketID, ops, c.resolveAPIKey(apiKey)},
		[]any{&out},
	)
	return out, err
}

func (c *Client) BulkCopyFiles(bucketID string, ops []FilePathOp, apiKey string) (resultNat, error) {
	out := resultNat{}
	err := c.agent.Call(
		c.canisterID, "bulkCopyFiles",
		[]any{bucketID, ops, c.resolveAPIKey(apiKey)},
		[]any{&out},
	)
	return out, err
}

func (c *Client) BeginFileUpload(bucketID, path, contentType string, totalSize uint64, apiKey string) (resultFileID, error) {
	out := resultFileID{}
	err := c.agent.Call(
		c.canisterID, "beginFileUpload",
		[]any{bucketID, path, contentType, nat(totalSize), c.resolveAPIKey(apiKey)},
		[]any{&out},
	)
	return out, err
}

func (c *Client) UploadFileChunkIndexed(uploadID string, chunkIndex uint64, data []byte) (resultNat, error) {
	out := resultNat{}
	err := c.agent.Call(
		c.canisterID, "uploadFileChunkIndexed",
		[]any{uploadID, nat(chunkIndex), data},
		[]any{&out},
	)
	return out, err
}

func (c *Client) CompleteFileUpload(uploadID, apiKey string) (resultFileID, error) {
	out := resultFileID{}
	err := c.agent.Call(
		c.canisterID, "completeFileUpload",
		[]any{uploadID, c.resolveAPIKey(apiKey)},
		[]any{&out},
	)
	return out, err
}

func (c *Client) GetBucketCycleStatus() (ApiResult[struct {
	Balance                idl.Nat `ic:"balance"`
	Status                 string  `ic:"status"`
	CanAcceptNewBuckets    bool    `ic:"canAcceptNewBuckets"`
	EstimatedDaysRemaining idl.Nat `ic:"estimatedDaysRemaining"`
	DailyBurn              idl.Nat `ic:"dailyBurn"`
}], error) {
	out := ApiResult[struct {
		Balance                idl.Nat `ic:"balance"`
		Status                 string  `ic:"status"`
		CanAcceptNewBuckets    bool    `ic:"canAcceptNewBuckets"`
		EstimatedDaysRemaining idl.Nat `ic:"estimatedDaysRemaining"`
		DailyBurn              idl.Nat `ic:"dailyBurn"`
	}]{}
	err := c.agent.Query(c.canisterID, "getBucketCycleStatus", []any{}, []any{&out})
	return out, err
}
