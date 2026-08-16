# icpay-bucket (Go)

Go client for [ICPay Bucket](https://icpay.app) — on-chain encrypted file storage on the Internet Computer.

**Module:** `github.com/prasangapokharel/icpay-bucket-go` · **Canister:** `6vbhm-nqaaa-aaaan-q6muq-cai` · **Version:** 1.1.1

**Docs:** [Package guide](https://icpay.app/icpay-bucket-sdk) · [API reference](https://icpay.app/bucket/docs)

## Install

```bash
go get github.com/prasangapokharel/icpay-bucket-go@v1.1.1
```

```go
import icpaybucket "github.com/prasangapokharel/icpay-bucket-go"
```

## Quick start

```go
client, err := icpaybucket.NewClient(icpaybucket.ClientOptions{
    APIKey: "icp_cloud_…",
})
if err != nil {
    panic(err)
}

up, err := client.UploadFile("icp", "/hello.txt", []byte("hello\n"), "text/plain", "")
if err != nil || up.Err != nil {
    panic(err)
}

fmt.Println(client.PublicURL("icp", "/hello.txt"))

list, err := client.ListFiles("icp", 0, 20, "")
if err != nil || list.Err != nil {
    panic(err)
}
```

## Authentication

Set `APIKey` in `ClientOptions`, or pass `apiKey` as the last argument on each method.

| Scope | Methods |
|---|---|
| Read key | `ListFiles`, `DownloadFile`, `GetFile`, `FileExists`, `ListFolder`, `SearchFiles`, `GetFileMetadata` |
| Write key | `UploadFile`, `MoveFile`, `CopyFile`, `SetFileTags`, `SetFileMetadata`, `BeginFileUpload`, chunks |
| Delete key | `DeleteFile`, `BulkDeleteFiles` |

## API reference

Results use `ApiResult[T]` with `Ok *T` and `Err *string`.

### Files (`client.go`, `files.go`)

| Method | Description |
|---|---|
| `UploadFile(bucketID, path, data, contentType, apiKey)` | Single upload (≤ ~1.85 MB) |
| `DeleteFile(bucketID, path, apiKey)` | Delete file |
| `ListFiles(bucketID, page, pageSize, apiKey)` | Paginated list |
| `DownloadFile(bucketID, path, apiKey)` | File bytes |
| `GetPublicFileURL(bucketID, path)` | Canister CDN URL |
| `GetFile(bucketID, path, apiKey)` | File record |
| `FileExists(bucketID, path, apiKey)` | Exists check |
| `ListFolder(bucketID, prefix, page, pageSize, apiKey)` | List by prefix |
| `SearchFiles(bucketID, query, page, pageSize, apiKey)` | Search |
| `GetFileMetadata(bucketID, path, apiKey)` | Metadata JSON |
| `MoveFile(bucketID, source, dest, apiKey)` | Move |
| `CopyFile(bucketID, source, dest, apiKey)` | Copy |
| `SetFileTags(bucketID, path, tags, apiKey)` | Set tags |
| `SetFileMetadata(bucketID, path, metadata, apiKey)` | Set metadata |
| `BulkDeleteFiles(bucketID, paths, apiKey)` | Bulk delete |
| `BulkMoveFiles(bucketID, ops, apiKey)` | `FilePathOp{Source, Destination}` |
| `BulkCopyFiles(bucketID, ops, apiKey)` | `FilePathOp{Source, Destination}` |
| `BeginFileUpload(bucketID, path, contentType, totalSize, apiKey)` | Chunked session |
| `UploadFileChunkIndexed(uploadID, chunkIndex, data)` | Upload chunk |
| `CompleteFileUpload(uploadID, apiKey)` | Finalize upload |
| `GetUpload(uploadID)` | Session status (owner principal) |
| `CancelUpload(uploadID)` | Cancel session (owner principal) |
| `GetBucketCycleStatus()` | Cycle health |

### Helpers (`url.go`)

| Function | Description |
|---|---|
| `PublicURL(bucketName, path)` | CDN URL on client |
| `PublicFileURL(canisterID, bucketName, path)` | Standalone CDN URL |

Bucket admin and API-key CRUD are not wrapped yet — call the canister directly or use the npm/Python client.

## Public CDN

```
https://6vbhm-nqaaa-aaaan-q6muq-cai.raw.icp0.io/cloud/{bucketName}{path}
```

## Notes

- Destination paths must use allowed extensions (`.txt`, `.webp`, …).
- `GetUpload` / `CancelUpload` require the bucket owner's Internet Identity principal.

## License

MIT
