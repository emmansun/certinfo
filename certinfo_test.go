package certinfo

import (
	"bytes"
	"encoding/pem"
	"os"
	"testing"

	"github.com/emmansun/gmsm/smx509"
)

type InputType int

const (
	tCertificate InputType = iota
	tCertificateRequest
)

// Compares a PEM-encoded certificate to a reference file.
func testPair(t *testing.T, certFile, refFile string, inputType InputType) {
	// Read and parse the certificate
	pemData, err := os.ReadFile(certFile)
	if err != nil {
		t.Fatal(err)
	}
	block, rest := pem.Decode([]byte(pemData))
	if block == nil || len(rest) > 0 {
		t.Fatal("Certificate decoding error")
	}
	var result string
	switch inputType {
	case tCertificate:
		cert, err := smx509.ParseCertificate(block.Bytes)
		if err != nil {
			t.Fatal(err)
		}
		result, err = CertificateText(cert.ToX509())
		if err != nil {
			t.Fatal(err)
		}
	case tCertificateRequest:
		cert, err := smx509.ParseCertificateRequest(block.Bytes)
		if err != nil {
			t.Fatal(err)
		}
		result, err = CertificateRequestText(cert.ToX509())
		if err != nil {
			t.Fatal(err)
		}
	}
	resultData := []byte(result)

	// Read the reference output
	refData, err := os.ReadFile(refFile)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(resultData, refData) {
		t.Logf("'%s' did not match reference '%s'\n", certFile, refFile)
		t.Errorf("Dump follows:\n%s\n", result)
	}
}

// Compares a PEM-encoded certificate to a reference file.
func testPairShort(t *testing.T, certFile, refFile string, inputType InputType) {
	// Read and parse the certificate
	pemData, err := os.ReadFile(certFile)
	if err != nil {
		t.Fatal(err)
	}
	block, rest := pem.Decode([]byte(pemData))
	if block == nil || len(rest) > 0 {
		t.Fatal("Certificate decoding error")
	}
	var result string
	switch inputType {
	case tCertificate:
		cert, err := smx509.ParseCertificate(block.Bytes)
		if err != nil {
			t.Fatal(err)
		}
		result, err = CertificateShortText(cert.ToX509())
		if err != nil {
			t.Fatal(err)
		}
	case tCertificateRequest:
		cert, err := smx509.ParseCertificateRequest(block.Bytes)
		if err != nil {
			t.Fatal(err)
		}
		result, err = CertificateRequestShortText(cert.ToX509())
		if err != nil {
			t.Fatal(err)
		}
	}
	resultData := []byte(result)

	// Read the reference output
	refData, err := os.ReadFile(refFile)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(resultData, refData) {
		t.Logf("'%s' did not match reference '%s'\n", certFile, refFile)
		t.Errorf("Dump follows:\n%s\n", result)
	}
}

// Test the root CA certificate
func TestCertInfoRoot(t *testing.T) {
	testPair(t, "test_certs/sm2.rca.pem", "test_certs/sm2.rca.text", tCertificate)
	testPair(t, "test_certs/sm2.oca.pem", "test_certs/sm2.oca.text", tCertificate)
	testPair(t, "test_certs/root1.cert.pem", "test_certs/root1.cert.text", tCertificate)
	testPair(t, "test_certs/root1.csr.pem", "test_certs/root1.csr.text", tCertificateRequest)
	testPairShort(t, "test_certs/sm2.rca.pem", "test_certs/sm2.rca.short", tCertificate)
	testPairShort(t, "test_certs/sm2.oca.pem", "test_certs/sm2.oca.short", tCertificate)
	testPairShort(t, "test_certs/root1.cert.pem", "test_certs/root1.cert.short", tCertificate)
	testPairShort(t, "test_certs/root1.csr.pem", "test_certs/root1.csr.short", tCertificateRequest)
}

// Test the leaf (user) RSA certificate
func TestCertInfoLeaf1(t *testing.T) {
	testPair(t, "test_certs/leaf1.cert.pem", "test_certs/leaf1.cert.text", tCertificate)
	testPair(t, "test_certs/leaf1.csr.pem", "test_certs/leaf1.csr.text", tCertificateRequest)
	testPairShort(t, "test_certs/leaf1.cert.pem", "test_certs/leaf1.cert.short", tCertificate)
	testPairShort(t, "test_certs/leaf1.csr.pem", "test_certs/leaf1.csr.short", tCertificateRequest)
}

// Test the leaf (user) DSA certificate
func TestCertInfoLeaf2(t *testing.T) {
	testPair(t, "test_certs/leaf2.cert.pem", "test_certs/leaf2.cert.text", tCertificate)
	testPair(t, "test_certs/leaf2.csr.pem", "test_certs/leaf2.csr.text", tCertificateRequest)
	testPairShort(t, "test_certs/leaf2.cert.pem", "test_certs/leaf2.cert.short", tCertificate)
	testPairShort(t, "test_certs/leaf2.csr.pem", "test_certs/leaf2.csr.short", tCertificateRequest)
}

// Test the leaf (user) ECDSA certificate
func TestCertInfoLeaf3(t *testing.T) {
	testPair(t, "test_certs/leaf3.cert.pem", "test_certs/leaf3.cert.text", tCertificate)
	testPair(t, "test_certs/leaf3.csr.pem", "test_certs/leaf3.csr.text", tCertificateRequest)
	testPairShort(t, "test_certs/leaf3.cert.pem", "test_certs/leaf3.cert.short", tCertificate)
	testPairShort(t, "test_certs/leaf3.csr.pem", "test_certs/leaf3.csr.short", tCertificateRequest)
}

// Test the leaf (user) with multiple sans
func TestCertInfoLeaf4(t *testing.T) {
	testPair(t, "test_certs/leaf4.cert.pem", "test_certs/leaf4.cert.text", tCertificate)
	testPair(t, "test_certs/leaf4.csr.pem", "test_certs/leaf4.csr.text", tCertificateRequest)
	testPairShort(t, "test_certs/leaf4.cert.pem", "test_certs/leaf4.cert.short", tCertificate)
	testPairShort(t, "test_certs/leaf4.csr.pem", "test_certs/leaf4.csr.short", tCertificateRequest)
}

func TestCertInfoLeaf5(t *testing.T) {
	testPair(t, "test_certs/leaf5.cert.pem", "test_certs/leaf5.cert.text", tCertificate)
	testPair(t, "test_certs/leaf5.csr.pem", "test_certs/leaf5.csr.text", tCertificateRequest)
	testPairShort(t, "test_certs/leaf5.cert.pem", "test_certs/leaf5.cert.short", tCertificate)
	testPairShort(t, "test_certs/leaf5.csr.pem", "test_certs/leaf5.csr.short", tCertificateRequest)
}

func TestCsrInfoWackyExtensions(t *testing.T) {
	testPair(t, "test_certs/x509WackyExtensions.pem", "test_certs/x509WackyExtensions.text", tCertificateRequest)
}

func TestNoCN(t *testing.T) {
	testPair(t, "test_certs/noCN.csr", "test_certs/noCN.csr.text", tCertificateRequest)
	testPairShort(t, "test_certs/noCN.csr", "test_certs/noCN.csr.text.short", tCertificateRequest)
}
