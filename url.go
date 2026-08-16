package icpaybucket

// PublicFileURL builds the raw HTTP gateway URL for a public file. No canister
// call is involved — this is the URL served directly by the gateway.
func PublicFileURL(canisterID, bucketName, path string) string {
	filePath := path
	if len(path) == 0 || path[0] != '/' {
		filePath = "/" + path
	}
	return "https://" + canisterID + ".raw.icp0.io/cloud/" + bucketName + filePath
}
