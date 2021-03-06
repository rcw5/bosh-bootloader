package certs_test

import (
	"errors"
	"fmt"
	"io/ioutil"

	"github.com/cloudfoundry/bosh-bootloader/certs"
	"github.com/cloudfoundry/bosh-bootloader/testhelpers"
	"github.com/cloudfoundry/multierror"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("CertificateValidator", func() {
	var (
		certificateValidator certs.Validator
		certFilePath         string
		keyFilePath          string
		chainFilePath        string
		certNonPEMFilePath   string
		keyNonPEMFilePath    string
		chainNonPEMFilePath  string
	)

	BeforeEach(func() {
		var err error
		certificateValidator = certs.NewValidator()
		chainFilePath, err = testhelpers.WriteContentsToTempFile(testhelpers.BBL_CHAIN)
		Expect(err).NotTo(HaveOccurred())

		certFilePath, err = testhelpers.WriteContentsToTempFile(testhelpers.BBL_CERT)
		Expect(err).NotTo(HaveOccurred())

		keyFilePath, err = testhelpers.WriteContentsToTempFile(testhelpers.BBL_KEY)
		Expect(err).NotTo(HaveOccurred())

		certNonPEMFilePath, err = testhelpers.WriteContentsToTempFile("not a cert")
		Expect(err).NotTo(HaveOccurred())

		keyNonPEMFilePath, err = testhelpers.WriteContentsToTempFile("not a key")
		Expect(err).NotTo(HaveOccurred())

		chainNonPEMFilePath, err = testhelpers.WriteContentsToTempFile("not a chain")
		Expect(err).NotTo(HaveOccurred())
	})

	Describe("ReadAndValidate", func() {
		Context("when using a PKCS#1 key", func() {
			Context("when cert and key are valid", func() {
				It("does not return an error", func() {
					_, err := certificateValidator.ReadAndValidate(certFilePath, keyFilePath, "")

					Expect(err).NotTo(HaveOccurred())
				})
			})

			Context("when cert, key, and chain are valid", func() {
				It("does not return an error", func() {
					_, err := certificateValidator.ReadAndValidate(certFilePath, keyFilePath, chainFilePath)

					Expect(err).NotTo(HaveOccurred())
				})
			})
		})
	})

	Describe("Read", func() {
		Context("when cert and key files exist and can be read", func() {
			It("returns cert and key data", func() {
				certData, err := certificateValidator.Read(certNonPEMFilePath, keyNonPEMFilePath, "")
				Expect(err).NotTo(HaveOccurred())

				Expect(string(certData.Cert)).To(Equal("not a cert"))
				Expect(string(certData.Key)).To(Equal("not a key"))
				Expect(string(certData.Chain)).To(Equal(""))
			})
		})

		Context("when cert, key, and chain files exist and can be read", func() {
			It("returns cert, key, and chain data", func() {
				certData, err := certificateValidator.Read(certNonPEMFilePath, keyNonPEMFilePath, chainNonPEMFilePath)
				Expect(err).NotTo(HaveOccurred())

				Expect(string(certData.Cert)).To(Equal("not a cert"))
				Expect(string(certData.Key)).To(Equal("not a key"))
				Expect(string(certData.Chain)).To(Equal("not a chain"))
			})
		})

		Context("if cert and key are not provided", func() {
			It("returns an error", func() {
				_, err := certificateValidator.Read("", "", "")
				expectedErr := multierror.NewMultiError("")
				expectedErr.Add(errors.New("--cert is required"))
				expectedErr.Add(errors.New("--key is required"))

				Expect(err).To(Equal(expectedErr))
			})
		})

		Context("if the cert key file does not exist", func() {
			It("returns an error", func() {
				_, err := certificateValidator.Read("/some/fake/cert/path", "/some/fake/key/path", "")
				expectedErr := multierror.NewMultiError("")
				expectedErr.Add(errors.New(`certificate file not found: "/some/fake/cert/path"`))
				expectedErr.Add(errors.New(`key file not found: "/some/fake/key/path"`))

				Expect(err).To(Equal(expectedErr))
			})
		})

		Context("if the cert and key are not regular files", func() {
			It("returns an error", func() {
				_, err := certificateValidator.Read("/dev/null", "/dev/null", "")
				expectedErr := multierror.NewMultiError("")
				expectedErr.Add(errors.New(`certificate is not a regular file: "/dev/null"`))
				expectedErr.Add(errors.New(`key is not a regular file: "/dev/null"`))

				Expect(err).To(Equal(expectedErr))
			})
		})

		Context("chain is provided", func() {
			Context("when chain file does not exist", func() {
				It("returns an error", func() {
					_, err := certificateValidator.Read(certFilePath, keyFilePath, "/some/fake/chain/path")
					expectedErr := multierror.NewMultiError("")
					expectedErr.Add(errors.New(`chain file not found: "/some/fake/chain/path"`))

					Expect(err).To(Equal(expectedErr))
				})
			})

			Context("when chain file is not a regular file", func() {
				It("returns an error", func() {
					_, err := certificateValidator.Read(certFilePath, keyFilePath, "/dev/null")
					expectedErr := multierror.NewMultiError("")
					expectedErr.Add(errors.New(`chain is not a regular file: "/dev/null"`))

					Expect(err).To(Equal(expectedErr))
				})
			})
		})
	})

	Describe("Validate", func() {
		Context("when using a PKCS#1 key", func() {
			var (
				realCert  []byte
				realKey   []byte
				realChain []byte

				otherKey   []byte
				otherChain []byte

				fakeCert  []byte
				fakeKey   []byte
				fakeChain []byte

				invalidKey  []byte
				invalidCert []byte
			)

			BeforeEach(func() {
				realCert = []byte(testhelpers.BBL_CERT)
				realKey = []byte(testhelpers.BBL_KEY)
				realChain = []byte(testhelpers.BBL_CHAIN)

				otherKey = []byte(testhelpers.OTHER_BBL_KEY)
				otherChain = []byte(testhelpers.OTHER_BBL_CHAIN)

				fakeCert = []byte("not a cert")
				fakeKey = []byte("not a key")
				fakeChain = []byte("not a chain")

				invalidKey = []byte("-----BEGIN RSA PRIVATE KEY-----\n-----END RSA PRIVATE KEY-----")
				invalidCert = []byte("-----BEGIN CERTIFICATE-----\n-----END CERTIFICATE-----")
			})

			Context("when cert and key are valid", func() {
				It("does not return an error", func() {
					err := certificateValidator.Validate(realCert, realKey, []byte{})

					Expect(err).NotTo(HaveOccurred())
				})
			})

			Context("when cert, key, and chain are valid", func() {
				It("does not return an error", func() {
					err := certificateValidator.Validate(realCert, realKey, realChain)

					Expect(err).NotTo(HaveOccurred())
				})
			})

			Context("if the cert and key are not PEM encoded", func() {
				It("returns an error", func() {
					err := certificateValidator.Validate(fakeCert, fakeKey, []byte{})

					expectedErr := multierror.NewMultiError("")
					expectedErr.Add(fmt.Errorf(`certificate is not PEM encoded: %q`, "not a cert"))
					expectedErr.Add(fmt.Errorf(`key is not PEM encoded: %q`, "not a key"))

					Expect(err).To(Equal(expectedErr))
				})
			})

			Context("if the key and cert are not compatible", func() {
				It("returns an error", func() {
					err := certificateValidator.Validate(realCert, otherKey, []byte{})

					expectedErr := multierror.NewMultiError("")
					expectedErr.Add(errors.New("tls: private key does not match public key"))
					Expect(err).To(Equal(expectedErr))
				})
			})

			Context("when the key is not valid", func() {
				It("returns an error", func() {
					err := certificateValidator.Validate(realCert, invalidKey, []byte{})

					expectedErr := multierror.NewMultiError("")
					expectedErr.Add(errors.New("tls: failed to parse private key"))
					Expect(err).To(Equal(expectedErr))
				})
			})

			Context("when the cert is not valid", func() {
				It("returns an error", func() {
					err := certificateValidator.Validate(invalidCert, realKey, []byte{})

					expectedErr := multierror.NewMultiError("")
					expectedErr.Add(errors.New("asn1: syntax error: sequence truncated"))
					Expect(err).To(Equal(expectedErr))
				})
			})

			Context("when chain is provided", func() {
				Context("if the cert, key and chain are incompatible", func() {
					It("returns multiple errors", func() {
						err := certificateValidator.Validate(realCert, otherKey, otherChain)
						expectedErr := multierror.NewMultiError("")
						expectedErr.Add(errors.New("tls: private key does not match public key"))
						expectedErr.Add(errors.New("certificate and chain mismatch: x509: certificate signed by unknown authority"))

						Expect(err).To(Equal(expectedErr))
					})
				})

				Context("if the chain and cert are not compatible", func() {
					It("returns an error", func() {
						err := certificateValidator.Validate(realCert, realKey, otherChain)

						expectedErr := multierror.NewMultiError("")
						expectedErr.Add(errors.New("certificate and chain mismatch: x509: certificate signed by unknown authority"))
						Expect(err).To(Equal(expectedErr))
					})
				})

				Context("if the chain is not PEM encoded", func() {
					It("returns an error", func() {
						err := certificateValidator.Validate(realCert, realKey, fakeChain)

						expectedErr := multierror.NewMultiError("")
						expectedErr.Add(fmt.Errorf(`chain is not PEM encoded: "not a chain"`))

						Expect(err).To(Equal(expectedErr))
					})
				})

				Context("when the chain is not valid", func() {
					It("returns an error", func() {
						err := certificateValidator.Validate(realCert, realKey, invalidCert)

						expectedErr := multierror.NewMultiError("")
						expectedErr.Add(fmt.Errorf("failed to parse chain"))

						Expect(err).To(Equal(expectedErr))
					})
				})
			})
		})

		Context("when using a PKCS#8 key", func() {
			var (
				cert []byte
				key  []byte
			)

			BeforeEach(func() {
				var err error
				cert, err = ioutil.ReadFile("fixtures/pkcs8.crt")
				Expect(err).NotTo(HaveOccurred())

				key, err = ioutil.ReadFile("fixtures/pkcs8.key")
				Expect(err).NotTo(HaveOccurred())
			})

			Context("when cert and key are valid", func() {
				It("does not return an error", func() {
					err := certificateValidator.Validate(cert, key, []byte{})

					Expect(err).NotTo(HaveOccurred())
				})
			})
		})
	})
})
