// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tls

import (
	"crypto"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"errors"
	"hash"
)

// Split a premain secret in two as specified in RFC 4346, section 5.
func splitPreMainSecret(secret []byte) (s1, s2 []byte) {
	s1 = secret[0 : (len(secret)+1)/2]
	s2 = secret[len(secret)/2:]
	return
}

// pHash implements the P_hash function, as defined in RFC 4346, section 5.
func pHash(result, secret, seed []byte, hash func() hash.Hash) {
	h := hmac.New(hash, secret)
	h.Write(seed)
	a := h.Sum(nil)

	j := 0
	for j < len(result) {
		h.Reset()
		h.Write(a)
		h.Write(seed)
		b := h.Sum(nil)
		todo := len(b)
		if j+todo > len(result) {
			todo = len(result) - j
		}
		copy(result[j:j+todo], b)
		j += todo

		h.Reset()
		h.Write(a)
		a = h.Sum(nil)
	}
}

// prf10 implements the TLS 1.0 pseudo-random function, as defined in RFC 2246, section 5.
func prf10(result, secret, label, seed []byte) {
	hashSHA1 := sha1.New
	hashMD5 := md5.New

	labelAndSeed := make([]byte, len(label)+len(seed))
	copy(labelAndSeed, label)
	copy(labelAndSeed[len(label):], seed)

	s1, s2 := splitPreMainSecret(secret)
	pHash(result, s1, labelAndSeed, hashMD5)
	result2 := make([]byte, len(result))
	pHash(result2, s2, labelAndSeed, hashSHA1)

	for i, b := range result2 {
		result[i] ^= b
	}
}

// prf12 implements the TLS 1.2 pseudo-random function, as defined in RFC 5246, section 5.
func prf12(hashFunc func() hash.Hash) func(result, secret, label, seed []byte) {
	return func(result, secret, label, seed []byte) {
		labelAndSeed := make([]byte, len(label)+len(seed))
		copy(labelAndSeed, label)
		copy(labelAndSeed[len(label):], seed)

		pHash(result, secret, labelAndSeed, hashFunc)
	}
}

// prf30 implements the SSL 3.0 pseudo-random function, as defined in
// www.mozilla.org/projects/security/pki/nss/ssl/draft302.txt section 6.
func prf30(result, secret, label, seed []byte) {
	hashSHA1 := sha1.New()
	hashMD5 := md5.New()

	done := 0
	i := 0
	// RFC 5246 section 6.3 says that the largest PRF output needed is 128
	// bytes. Since no more ciphersuites will be added to SSLv3, this will
	// remain true. Each iteration gives us 16 bytes so 10 iterations will
	// be sufficient.
	var b [11]byte
	for done < len(result) {
		for j := 0; j <= i; j++ {
			b[j] = 'A' + byte(i)
		}

		hashSHA1.Reset()
		hashSHA1.Write(b[:i+1])
		hashSHA1.Write(secret)
		hashSHA1.Write(seed)
		digest := hashSHA1.Sum(nil)

		hashMD5.Reset()
		hashMD5.Write(secret)
		hashMD5.Write(digest)

		done += copy(result[done:], hashMD5.Sum(nil))
		i++
	}
}

const (
	tlsRandomLength      = 32 // Length of a random nonce in TLS 1.1.
	mainSecretLength   = 48 // Length of a main secret in TLS 1.1.
	finishedVerifyLength = 12 // Length of verify_data in a Finished message.
)

var mainSecretLabel = []byte("main secret")
var keyExpansionLabel = []byte("key expansion")
var clientFinishedLabel = []byte("client finished")
var serverFinishedLabel = []byte("server finished")

func prfAndHashForVersion(version uint16, suite *cipherSuite) (func(result, secret, label, seed []byte), crypto.Hash) {
	switch version {
	case VersionSSL30:
		return prf30, crypto.Hash(0)
	case VersionTLS10, VersionTLS11:
		return prf10, crypto.Hash(0)
	case VersionTLS12:
		if suite.flags&suiteSHA384 != 0 {
			return prf12(sha512.New384), crypto.SHA384
		}
		return prf12(sha256.New), crypto.SHA256
	default:
		panic("unknown version")
	}
}

func prfForVersion(version uint16, suite *cipherSuite) func(result, secret, label, seed []byte) {
	prf, _ := prfAndHashForVersion(version, suite)
	return prf
}

// mainFromPreMainSecret generates the main secret from the pre-main
// secret. See http://tools.ietf.org/html/rfc5246#section-8.1
func mainFromPreMainSecret(version uint16, suite *cipherSuite, preMainSecret, clientRandom, serverRandom []byte) []byte {
	seed := make([]byte, 0, len(clientRandom)+len(serverRandom))
	seed = append(seed, clientRandom...)
	seed = append(seed, serverRandom...)

	mainSecret := make([]byte, mainSecretLength)
	prfForVersion(version, suite)(mainSecret, preMainSecret, mainSecretLabel, seed)
	return mainSecret
}

// keysFromMainSecret generates the connection keys from the main
// secret, given the lengths of the MAC key, cipher key and IV, as defined in
// RFC 2246, section 6.3.
func keysFromMainSecret(version uint16, suite *cipherSuite, mainSecret, clientRandom, serverRandom []byte, macLen, keyLen, ivLen int) (clientMAC, serverMAC, clientKey, serverKey, clientIV, serverIV []byte) {
	seed := make([]byte, 0, len(serverRandom)+len(clientRandom))
	seed = append(seed, serverRandom...)
	seed = append(seed, clientRandom...)

	n := 2*macLen + 2*keyLen + 2*ivLen
	keyMaterial := make([]byte, n)
	prfForVersion(version, suite)(keyMaterial, mainSecret, keyExpansionLabel, seed)
	clientMAC = keyMaterial[:macLen]
	keyMaterial = keyMaterial[macLen:]
	serverMAC = keyMaterial[:macLen]
	keyMaterial = keyMaterial[macLen:]
	clientKey = keyMaterial[:keyLen]
	keyMaterial = keyMaterial[keyLen:]
	serverKey = keyMaterial[:keyLen]
	keyMaterial = keyMaterial[keyLen:]
	clientIV = keyMaterial[:ivLen]
	keyMaterial = keyMaterial[ivLen:]
	serverIV = keyMaterial[:ivLen]
	return
}

// lookupTLSHash looks up the corresponding crypto.Hash for a given
// TLS hash identifier.
func lookupTLSHash(hash uint8) (crypto.Hash, error) {
	switch hash {
	case hashSHA1:
		return crypto.SHA1, nil
	case hashSHA256:
		return crypto.SHA256, nil
	case hashSHA384:
		return crypto.SHA384, nil
	default:
		return 0, errors.New("tls: unsupported hash algorithm")
	}
}

func newFinishedHash(version uint16, cipherSuite *cipherSuite) finishedHash {
	var buffer []byte
	if version == VersionSSL30 || version >= VersionTLS12 {
		buffer = []byte{}
	}

	prf, hash := prfAndHashForVersion(version, cipherSuite)
	if hash != 0 {
		return finishedHash{hash.New(), hash.New(), nil, nil, buffer, version, prf}
	}

	return finishedHash{sha1.New(), sha1.New(), md5.New(), md5.New(), buffer, version, prf}
}

// A finishedHash calculates the hash of a set of handshake messages suitable
// for including in a Finished message.
type finishedHash struct {
	client hash.Hash
	server hash.Hash

	// Prior to TLS 1.2, an additional MD5 hash is required.
	clientMD5 hash.Hash
	serverMD5 hash.Hash

	// In TLS 1.2, a full buffer is sadly required.
	buffer []byte

	version uint16
	prf     func(result, secret, label, seed []byte)
}

func (h *finishedHash) Write(msg []byte) (n int, err error) {
	h.client.Write(msg)
	h.server.Write(msg)

	if h.version < VersionTLS12 {
		h.clientMD5.Write(msg)
		h.serverMD5.Write(msg)
	}

	if h.buffer != nil {
		h.buffer = append(h.buffer, msg...)
	}

	return len(msg), nil
}

func (h finishedHash) Sum() []byte {
	if h.version >= VersionTLS12 {
		return h.client.Sum(nil)
	}

	out := make([]byte, 0, md5.Size+sha1.Size)
	out = h.clientMD5.Sum(out)
	return h.client.Sum(out)
}

// finishedSum30 calculates the contents of the verify_data member of a SSLv3
// Finished message given the MD5 and SHA1 hashes of a set of handshake
// messages.
func finishedSum30(md5, sha1 hash.Hash, mainSecret []byte, magic []byte) []byte {
	md5.Write(magic)
	md5.Write(mainSecret)
	md5.Write(ssl30Pad1[:])
	md5Digest := md5.Sum(nil)

	md5.Reset()
	md5.Write(mainSecret)
	md5.Write(ssl30Pad2[:])
	md5.Write(md5Digest)
	md5Digest = md5.Sum(nil)

	sha1.Write(magic)
	sha1.Write(mainSecret)
	sha1.Write(ssl30Pad1[:40])
	sha1Digest := sha1.Sum(nil)

	sha1.Reset()
	sha1.Write(mainSecret)
	sha1.Write(ssl30Pad2[:40])
	sha1.Write(sha1Digest)
	sha1Digest = sha1.Sum(nil)

	ret := make([]byte, len(md5Digest)+len(sha1Digest))
	copy(ret, md5Digest)
	copy(ret[len(md5Digest):], sha1Digest)
	return ret
}

var ssl3ClientFinishedMagic = [4]byte{0x43, 0x4c, 0x4e, 0x54}
var ssl3ServerFinishedMagic = [4]byte{0x53, 0x52, 0x56, 0x52}

// clientSum returns the contents of the verify_data member of a client's
// Finished message.
func (h finishedHash) clientSum(mainSecret []byte) []byte {
	if h.version == VersionSSL30 {
		return finishedSum30(h.clientMD5, h.client, mainSecret, ssl3ClientFinishedMagic[:])
	}

	out := make([]byte, finishedVerifyLength)
	h.prf(out, mainSecret, clientFinishedLabel, h.Sum())
	return out
}

// serverSum returns the contents of the verify_data member of a server's
// Finished message.
func (h finishedHash) serverSum(mainSecret []byte) []byte {
	if h.version == VersionSSL30 {
		return finishedSum30(h.serverMD5, h.server, mainSecret, ssl3ServerFinishedMagic[:])
	}

	out := make([]byte, finishedVerifyLength)
	h.prf(out, mainSecret, serverFinishedLabel, h.Sum())
	return out
}

// selectClientCertSignatureAlgorithm returns a signatureAndHash to sign a
// client's CertificateVerify with, or an error if none can be found.
func (h finishedHash) selectClientCertSignatureAlgorithm(serverList []signatureAndHash, sigType uint8) (signatureAndHash, error) {
	if h.version < VersionTLS12 {
		// Nothing to negotiate before TLS 1.2.
		return signatureAndHash{signature: sigType}, nil
	}

	for _, v := range serverList {
		if v.signature == sigType && isSupportedSignatureAndHash(v, supportedSignatureAlgorithms) {
			return v, nil
		}
	}
	return signatureAndHash{}, errors.New("tls: no supported signature algorithm found for signing client certificate")
}

// hashForClientCertificate returns a digest, hash function, and TLS 1.2 hash
// id suitable for signing by a TLS client certificate.
func (h finishedHash) hashForClientCertificate(signatureAndHash signatureAndHash, mainSecret []byte) ([]byte, crypto.Hash, error) {
	if (h.version == VersionSSL30 || h.version >= VersionTLS12) && h.buffer == nil {
		panic("a handshake hash for a client-certificate was requested after discarding the handshake buffer")
	}

	if h.version == VersionSSL30 {
		if signatureAndHash.signature != signatureRSA {
			return nil, 0, errors.New("tls: unsupported signature type for client certificate")
		}

		md5Hash := md5.New()
		md5Hash.Write(h.buffer)
		sha1Hash := sha1.New()
		sha1Hash.Write(h.buffer)
		return finishedSum30(md5Hash, sha1Hash, mainSecret, nil), crypto.MD5SHA1, nil
	}
	if h.version >= VersionTLS12 {
		hashAlg, err := lookupTLSHash(signatureAndHash.hash)
		if err != nil {
			return nil, 0, err
		}
		hash := hashAlg.New()
		hash.Write(h.buffer)
		return hash.Sum(nil), hashAlg, nil
	}

	if signatureAndHash.signature == signatureECDSA {
		return h.server.Sum(nil), crypto.SHA1, nil
	}

	return h.Sum(), crypto.MD5SHA1, nil
}

// discardHandshakeBuffer is called when there is no more need to
// buffer the entirety of the handshake messages.
func (h *finishedHash) discardHandshakeBuffer() {
	h.buffer = nil
}
