# icpay-bucket

Full Go client for [ICPay Bucket](https://icpay.app).

```
go get github.com/icpay/icpay-bucket
```

See [docs/pkg/go/README.md](../../docs/pkg/go/README.md) for the complete API reference.

## Quick start

```go
client, _ := icpaybucket.NewClient(icpaybucket.ClientOptions{APIKey: "icp_cloud_…"})
client.UploadFile("icp", "/hello.txt", []byte("hello\n"), "text/plain", "")
fmt.Println(client.PublicURL("icp", "/hello.txt"))
```

## v1.1.0 — file API + bulk + chunked upload

Methods live in `files.go`. Bucket admin and API key methods use the canister actor directly until wrapped.

## License

MIT
