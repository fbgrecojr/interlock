package client

import (
	"os"
	"testing"
)

const (
	TestCACert = `-----BEGIN CERTIFICATE-----
MIIC0jCCAbygAwIBAgIRAKmQU1XP3XW8ONTT7HQ9x+gwCwYJKoZIhvcNAQELMBQx
EjAQBgNVBAoTCWludGVybG9jazAeFw0xNjAyMjYyMjU0MDBaFw0xOTAyMTAyMjU0
MDBaMBQxEjAQBgNVBAoTCWludGVybG9jazCCASIwDQYJKoZIhvcNAQEBBQADggEP
ADCCAQoCggEBALlWkKybRVQT3eRuuG+XuSz68dFZ8PN/NRIhJV8+nlwG2vu3Dy9S
hjGBPYyoMaIBYOahHuGhNMQuMgkforobeyJL2XUc31kgL3Beb7q4CqpUiWcwmI1w
zxw4Nfc1u1cORLo8nJvFIZ7V2qqJJp0bv/uZlb6Liuf8vtykmA7Qmr3Nixod71jY
jkZIxqz9U7q5bFBDxUw4oEb+UFA4kUdf71N+1cyn5IXN2QrFftpl31s5Xz0yrQlv
FUzEBJOof9jAF3ntHAus7IVoL9DrpKwmJ1w6zhilItku42jNZRPYDLQcf/wugkZS
JREN7x8VSz6hTpeUd5KlWvHB4ng+OCfdBQ8CAwEAAaMjMCEwDgYDVR0PAQH/BAQD
AgCsMA8GA1UdEwEB/wQFMAMBAf8wCwYJKoZIhvcNAQELA4IBAQAeLsLVHWBcfRUO
6zMOGR2yv3L40RMmwGTyAh/9VMvkx9zSFvEa7ilrsD6In5HLTCVtIc//9/UaCgA2
xmnAWDiWbNcAHRn2ZOOvXkqVS7u7kdvSL+IGqEwQxdxx/WPkMpFBUKxsEaYGHmQc
L6cVDa7wg71DhpUbYelE/0x/28UTkswV4sLanXPJGlX4ZYvJ/POHrFlqZeJaYioP
TLOzQuZya72rhXmsElY0SlSVWBFaDo2+wJ688tfAAc3T+3GqZh17ArNpSUB/rsNE
TnEEMq7EDu3adFkocNW2L4Jp3Ny0oSBVoTX3NwF/fnld8qsNbW6qICCxTHJQyKRH
KyxQJWIx
-----END CERTIFICATE-----`

	TestCert = `-----BEGIN CERTIFICATE-----
MIIC7jCCAdigAwIBAgIRAMfciQv+IxRI4to5sDqUZpcwCwYJKoZIhvcNAQELMBQx
EjAQBgNVBAoTCWludGVybG9jazAeFw0xNjAyMjYyMjU0MDBaFw0xOTAyMTAyMjU0
MDBaMBQxEjAQBgNVBAoTCWludGVybG9jazCCASIwDQYJKoZIhvcNAQEBBQADggEP
ADCCAQoCggEBAMJ1gkztg+FFdEMrp0KlCe7nyFfEycYxW1O5PMmqe3o1SijQM4qi
/NHUQfLYWiFBWXRkVTCuqPQsp5lZqhmD/Tzd5EjrvzyORrOZ2lVJLLnt+omoAsVh
V/lioUdKWQD5Mw1mFg6l/kjil//4VfWd3P0SdDHOqtriKQruwu07cdL+MuQ+/xjA
SKGh6FWAUlb5FZHZxi9WubUc4jOKAbsA/WAKvRDDNVc93oy0925ummeOv5DsF2wn
X+ptNrFVFqP6pMteSSyEQFZROFkLjp5VfWoc2b9WeIrHs2OU/fM7knAkAwm3+KWA
GWiD3J9GVpG+vLnhReBj0heEGV8mg43/IlECAwEAAaM/MD0wDgYDVR0PAQH/BAQD
AgCoMB0GA1UdJQQWMBQGCCsGAQUFBwMCBggrBgEFBQcDATAMBgNVHRMBAf8EAjAA
MAsGCSqGSIb3DQEBCwOCAQEAlJWZO0HVrUHiRzI4PsBtq2yABLoqt4/c1V0FxkSL
EpsQPVSkpi/infvYojXo8XumivOKXCgs/4WwZoXYyhI81zppwCfy6CRYwJ8GvVpW
Gkj2iuw7VHYf9QPkQfmXAidnfChJwt7OSqkBVPS9KSM5zjWqwQUfZEQvnRbm+qgG
frYX3ikJm/Xs7heIjZyJdqV3o5eHcGpA6oK9aSpA283brrmen4cxUkW8Mt2KRccC
1l6ZZUEm73CMgWBYxnPZhIzTbbLTpMBGtdF0bJ16aEvD+92WSambRuJHhPOPpD9k
oUmZainyr6eRxb6bzE7qf5YfB+3uFZrPsEmH/rmZiNS3Cw==
-----END CERTIFICATE-----`

	TestKey = `-----BEGIN RSA PRIVATE KEY-----
MIIEpAIBAAKCAQEAwnWCTO2D4UV0QyunQqUJ7ufIV8TJxjFbU7k8yap7ejVKKNAz
iqL80dRB8thaIUFZdGRVMK6o9CynmVmqGYP9PN3kSOu/PI5Gs5naVUksue36iagC
xWFX+WKhR0pZAPkzDWYWDqX+SOKX//hV9Z3c/RJ0Mc6q2uIpCu7C7Ttx0v4y5D7/
GMBIoaHoVYBSVvkVkdnGL1a5tRziM4oBuwD9YAq9EMM1Vz3ejLT3bm6aZ46/kOwX
bCdf6m02sVUWo/qky15JLIRAVlE4WQuOnlV9ahzZv1Z4isezY5T98zuScCQDCbf4
pYAZaIPcn0ZWkb68ueFF4GPSF4QZXyaDjf8iUQIDAQABAoIBADOhscQtOFwC7fi7
yYBXg8isQDSVqqF2D3Kud2ZwXrK3HYayqUzBM/GesxgAvAWibVcLINd5OKEEjkeY
WCLIOeAEZo26Ep/IgxtC2YbVlAuWFXShaILx8sLjnkDoi7NHd3eySF8BUgAWMhej
32cE0F1dnf5vikvtysn9VUJaC0HjCDJXQaRj/D9y99rJ2SkW77ahymBy1QSAXEm3
Hg3q79vAqPczski7rRcoeKrnFiD0Y+4ho86NHyKFE/3O/e7GjpTWPdSBeNRcIMoI
xaqpmsMZp0s7Nzdk/2Qd8GmVYwncw/4WDebVIvmUlYQNcySRo2l6Flml2zwedviy
/0JFWRUCgYEA342mqQQstGlB1IM0T43x+BJ4AjOn0vGDPUrEvwu5zOT3IWggIIxS
RJYlcUIazN0KNmziY613v+GWPTC8NnWjQF3f/9++GVatAptzFyTk1IAKBX7oPu/d
JJFtwgtr66v5NqiGQyIsWuX19DLy32E5TVmZWKx9Tgg0xH71z3YXgbMCgYEA3q7U
a2Nkn86FTvsvdL5nigwuXMa/QQujZrTnzqMzJ7tYRauZXUG0Ns7cJU4fvMTfLUIG
4Nz2xqiEume0flnaBx6kAySe/9LUdmHBmvPtMF7GsJPHtdZZSDPdOdLOPqHrxnAz
E9E5vI62d4e4TCC5Vqxc8AQ+80B+8AYh528BIesCgYEAxf5zKSalYXQH9evunLcf
I5NX7rtJXC7DCbn63ynHeY0gw9mw+qLNCinhJ5pgmij7LpDpQVcVxEBMDA3p5GH0
IMID7l9/wnld6f07xbfLY9mzBoMLtxJCTmzvRPlQr/40TxCbOUI+/pLFb27gZK97
TOKaSksa/82MzquTkhcJYrkCgYEAzI/2eyA3U8a4F7IQCkLPgrVl8bxx/SLf3H3b
ZKvvVlR35qiYnl65Wo/1FCAMb7C7BCxffTn/SMeOBl82I8wOyfOP34NIvOHEY2uy
GtJx1bl69MMM9zINmpJqa7AH3umIWibABThyvZCsdmmrF+QH3mNAjQyZd4SMF5rK
knwaojkCgYAkhVol4h1dLhMaY/K6ttAAg/d6RuBVovsKnbXFzxUgHYUWm6pjWryK
iaKXFTXXkeGoX+NuQwL98CmLgFOgWo3P90r0SDsN9xqi4+MLiY9yzHAa+I5Qzt/k
eT0pvpz6WapVa/sXPmjZF2N/3AvCUGvrwCiOMV+yeYReJIRC0jVcSQ==
-----END RSA PRIVATE KEY-----`
)

func TestGetTLSConfig(t *testing.T) {
	cfg, err := GetTLSConfig(
		[]byte(TestCACert),
		[]byte(TestCert),
		[]byte(TestKey),
		true,
	)

	if err != nil {
		t.Fatal(err)
	}

	if cfg == nil {
		t.Fatal("unexpected nil TLS config")
	}

	if len(cfg.Certificates) != 1 {
		t.Fatal("expected certificate in TLS config")
	}
}

func TestGetDockerClientCleanEnvironment(t *testing.T) {
	c, err := GetDockerClient("", "", "", "", true)
	if err != nil {
		t.Fatal(err)
	}

	if c.TLSConfig != nil {
		t.Fatal("expected nil TLS config")
	}
}

func TestGetDockerClientEnvironment(t *testing.T) {
	testHost := "1.2.3.4:2375"
	os.Setenv("DOCKER_HOST", "tcp://"+testHost)
	c, err := GetDockerClient("", "", "", "", true)
	if err != nil {
		t.Fatal(err)
	}

	if c.TLSConfig != nil {
		t.Fatal("expected nil TLS config")
	}

	if c.URL.Host != testHost {
		t.Fatalf("expected Docker Host: %s; received %s", testHost, c.URL.Host)
	}
}
