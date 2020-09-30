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
