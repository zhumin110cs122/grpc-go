# gRPC-Go Retry Context Fix

## Fix: Client-side transparent retries preserve request context deadline

### Problem
When client-side transparent retries were triggered in gRPC-Go, the retry
used a new context instead of preserving the original request context's
deadline. This meant retries could exceed the caller's intended timeout.

### Fix
Ensures retry RPCs use the original context (with its deadline) rather than
creating a new context, so the caller's timeout is properly respected across
retry attempts.

### Test
```bash
go run main.go
```
