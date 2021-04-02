package config

func FromEnv() *Configuration {
	return &Configuration{
		DbConfig: DatabaseConfig{
			Hostname: GetenvOrDefault("DINNERDASH_DB_HOST", "localhost"),
			Username: GetenvOrDefault("DINNERDASH_DB_USER", "postgres"),
			Password: GetenvOrDefault("DINNERDASH_DB_PASS", "password"),
			Database: GetenvOrDefault("DINNERDASH_DB_DB", "dinnerdash_test"),
		},
		CookieHost: GetenvOrDefault("DINNERDASH_COOKIE_HOST", "localhost"),
		JwtCertificate: GetenvOrDefault("DINNERDASH_JWT_CERTIFICATE", `-----BEGIN CERTIFICATE-----
MIIDCzCCAfOgAwIBAgIJPLeCDeyWY9LKMA0GCSqGSIb3DQEBCwUAMCMxITAfBgNV
BAMTGGRpbm5lci1kYXNoLmV1LmF1dGgwLmNvbTAeFw0yMTAzMjMyMTQ1MjJaFw0z
NDExMzAyMTQ1MjJaMCMxITAfBgNVBAMTGGRpbm5lci1kYXNoLmV1LmF1dGgwLmNv
bTCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBAMxKD475YN3v1OEipq5v
YNBKeFX+PT26oZvjvF/dsQ0m0MUZTfL51ZEWDHGWEmZY6A+tpRLjBIT+4YmKbqKb
mMGWZKoG5k+QHHwTbS3RElKbqWvKgdg4qI/qM5l7bk0OWA41eHY4RCOV0/4UJ883
EodyUAYDLFFPCStou8gFov2by2uB459DmK/NbnQ/sdt6JyCi1w/PXmKIR3KY0nGN
0/Zp7b/s1BLQtt8gUPn5l2fXj69RuiMJ4CXJT7TXPcptYfgHia4OKB09meDcB1fB
lSr1YYWi29rRT7FCLYWafuaQXrQCDhUYEHmvvjtJJMW5v4ElgoauUjMu0kJ8ST9I
+X0CAwEAAaNCMEAwDwYDVR0TAQH/BAUwAwEB/zAdBgNVHQ4EFgQUtKx4vCE7gO98
z5zlcg7WrvDw/tgwDgYDVR0PAQH/BAQDAgKEMA0GCSqGSIb3DQEBCwUAA4IBAQCT
85wcLeOao8nZjuG8D8E+gAfXN5b4HD6fpDyB5QrOqomCMq66EH+jgXTsE9TDPO+V
vDYsx8jmRDJkAEBOxA0ZJ4xPqyl2T1W8cjP35M2LP7WvvgzikLQu8L8c5rKnaszt
BwqQUFoXCdIzh2ZMU1AM6iC4UfdAygo3GpLz93rbOe6QM8UQ5R81FGFwjj3KUxgX
4DwZunhIK/rzGuRRekY9PFnjrzNGXEgW8zwANmMfkqQGP3sYHjKT5z4Wh9aUjiWf
ACl9sVI9nlJGGag7o3YR576KPzgJaII8nfdzNuhPY6eBII5jra7ouEe1OnmDFPoL
8/4QPBkCXvsDQux6cOsY
-----END CERTIFICATE-----`),
	}
}

type Configuration struct {
	DbConfig       DatabaseConfig
	CookieHost     string
	JwtCertificate string
}

type DatabaseConfig struct {
	Hostname string
	Username string
	Password string
	Database string
}
