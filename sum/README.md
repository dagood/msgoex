# sum

The program prints an example response to the console when it starts.
It can also host a web service to return similar responses.

```
$ go run .
fips140.Enabled(): false
{
  "Request": "/sum?v=value",
  "Version": "go1.26.0",
  "Backend": "none",
  "Type": "*sha256.Digest",
  "Value": "value",
  "Sum": "a622e911f76530690297989c755df7c90132b4177ba8081d7d5026ca818a9b91",
  "Value2": "value2026-03-16T17:40:00.413435842Z",
  "Sum2": "c56ba33d1a41883fc5a48d67efff8dc8ea2e39d534bb8be18e699958437ca1ed6b0d548cea1ce980aee7c56a0e4f3eecf58052fc83f0bd824f103dc2c8c9aa3a"
}

$ msgo1.26.0-1 run .
fips140.Enabled(): false
{
  "Request": "/sum?v=value",
  "Version": "go1.26.0",
  "Backend": "enabled",
  "Type": "*openssl.Hash",
  "Value": "value",
  "Sum": "a622e911f76530690297989c755df7c90132b4177ba8081d7d5026ca818a9b91",
  "Value2": "value2026-03-16T17:40:16.103056818Z",
  "Sum2": "80d6d9f325c18c8e04eafda54e60fd919167e0348001fa356777fcec66da1013c3fa9ead3b1b780a90a41079f5254eb7e3d356a49c43b42d50b0ad2124dcd2d6"
}
```

If you pass `-serve`, the program will start the HTTP server.
Go to <https://localhost:8080/sum?v=golang> to see something like this:

```json
{
  "Request": "/sum?v=golang",
  "Version": "go1.20.6 X:opensslcrypto",
  "Type": "*openssl.sha256Hash",
  "Value": "golang",
  "Sum": "856d7285aaa4797447bb5cc1e5d266fd0682141947e4394b921ad6618bdddd6f",
  "Value2": "golang2023-07-25T00:08:53.137249726Z",
  "Sum2": "e24bbe318ebb88d99f2d911cc1ce300e06e19ed9bf25a17be968275a261b826b"
}
```
