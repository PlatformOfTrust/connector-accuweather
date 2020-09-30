package keyutil

import "testing"

// Test parsing of the key type if it's a url or a file path
type parseKeyTypeTest struct {
	in  string
	out rsaKeyType
}

var parseKeyTypeTests = []parseKeyTypeTest{
	{in: "http://test.com/key.pub", out: rsaKeyURL},
	{in: "https://test.com/key.pub", out: rsaKeyURL},
	{in: "./key.pub", out: rsaKeyFile},
	{in: "test", out: rsaKeyFile},
}

func TestParseKeyType(t *testing.T) {
	for _, c := range parseKeyTypeTests {
		keyType := ParseKeyType(c.in)
		if keyType != c.out {
			t.Errorf("Result doesn't match: In: %s, Expected: %d, Got: %d", c.in, c.out, keyType)
		}
	}
}

func TestParsePEM(t *testing.T) {
	pem := `-----BEGIN PUBLIC KEY-----
MIICIjANBgkqhkiG9w0BAQEFAAOCAg8AMIICCgKCAgEA8BcyJF5BRRgyTrldmCxz
P5VzP/BzZ2cpqi4DJJ7N2YbNnV3jAgjRIWT8s39hXhbIru+Lcof3KQNod9ftvSu5
3EHMYiByLHWf0uUbuu7N48idc1RTZZT2LGZNeUKK4+APiHX2CSAn5YiPDC2bcnXL
z9jCXxrGzBdMBqq4Lyd2cTz7HD1yCp5Stlu8dlg1ZnbMRiBxfTJQRF0OwxSn0PPl
xL2pbxOuJxkBjmbET++Fi/UoBhGDnlIRUUg25JoZmgaqG2mX/da13PFeGc1Dwz6J
jooNSmkR7xNIYvHIHTdgFPIApFG3FwHQLtzWiFxHj+yoLZobvJDqiUxU+Yq67LFl
UedXFySbMK2K3xlWxMCH0oZ4KXPNBDnDbsQWbtgs7cjY3UxKL260QOo9SXSex0Kw
f5zrkEK08+pE1sYtbwBkjZvSMWZTuDhjs24VFkgnxIEFT0f1jTgivZJ/d6XgaVYZ
CMWlifY3zbs3P82TP6iHRsdoTvc5F9oV1+QISTusLx1ekJAxy2bPI0DuIKI4H3FX
WBESotUXiO5mj3GIYKKhZKCX+lteDSI4hTnXS+coayY7nxA4/J2TaBnx9zYd/u8r
m6ZP6UiQWIJ5HBgxoZoJlZmZ2ibD9a0fGnCMrZnBJS2D8ADdh2ICiamb6t2AFz/e
bEXGyjRgun0RTWOMXRdLHBMCAwEAAQ==
-----END PUBLIC KEY-----`

	_, err := ParsePEM([]byte(pem))
	if err != nil {
		t.Error(err)
	}
}

func TestParsePrivateKey(t *testing.T) {
	pem := `-----BEGIN PRIVATE KEY-----
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
LOYD9cmn9Pw1N56gOFL4uDSxybIY1/A=
-----END PRIVATE KEY-----`

	_, err := ParsePrivateKey([]byte(pem))
	if err != nil {
		t.Error(err)
	}
}

func TestLoadRsaKeyFile(t *testing.T) {
	path := "../test-data/public.pem"
	_, err := LoadRsaKeyFile(path)
	if err != nil {
		t.Error(err)
	}
}

func TestLoadRsaKeys(t *testing.T) {
	_, err := LoadRsaKeys([]string{"../test-data/public.pem", "https://static-test.oftrust.net/keys/translator.pub"})
	if err != nil {
		t.Error(err)
	}
}
