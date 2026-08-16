# icpay-bucket

Minimal Go client for [ICPay Bucket](https://icpay.app) — on-chain encrypted
file storage on the Internet Computer.

```
go get github.com/prasangapokharel/icpay-bucket-go
```

## Quick start

```go
package main

import (
	"fmt"

	icpaybucket "github.com/prasangapokharel/icpay-bucket-go"
)

func main() {
	client, err := icpaybucket.NewClient(icpaybucket.ClientOptions{
		APIKey: "icp_cloud_…", // authorizes writes
	})
	if err != nil {
		panic(err)
	}

	// Upload (single call, files up to ~1.85 MB)
	up, err := client.UploadFile("icp", "/logo.webp", data, "image/webp", "")
	if err != nil || up.Err != nil {
		panic(err)
	}
	fmt.Println("uploaded:", *up.Ok)

	// Public CDN URL — no call needed
	url := client.PublicURL("icp", "/logo.webp")

	// List files
	list, err := client.ListFiles("icp", 0, 20)
	if err != nil || list.Err != nil {
		panic(err)
	}
	for _, f := range list.Ok.Items {
		fmt.Println(f.Path)
	}

	// Delete with the API key
	del, err := client.DeleteFile("icp", "/logo.webp", "")
	if err != nil || del.Err != nil {
		panic(err)
	}
}
```

## API

- `UploadFile(bucketID, path, data, contentType, apiKey)` → `{ Ok } | { Err }`.
- `DeleteFile(bucketID, path, apiKey)` → `{ Ok } | { Err }`.
- `ListFiles(bucketID, page, pageSize)` → paginated file list.
- `DownloadFile(bucketID, path)` → file bytes (public + private buckets).
- `GetPublicFileURL(bucketID, path)` → canister-provided CDN URL.
- `PublicURL(bucketName, path)` — CDN URL computed locally (no call).
- `PublicFileURL(canisterID, bucketName, path)` — same, standalone.

Writes are authorized by a bucket API key or an Internet Identity. Reads on
public buckets work for any caller.

> **Note:** `DownloadFile` takes the bucket **id**, not the bucket name. Other
> endpoints (`UploadFile`, `DeleteFile`, `ListFiles`, `GetPublicFileURL`)
> accept the bucket name (e.g. `"icp"`). To download, resolve the id first via
> the canister's `getBucket(name)` query.

## License

MIT