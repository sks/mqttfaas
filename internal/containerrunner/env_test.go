package containerrunner_test

import (
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/sks/mqttfaas/internal/containerrunner"
)

var _ = Describe("Env", func() {
	Describe("with proxy set", func() {
		BeforeEach(func() {
			os.Setenv("http_proxy", "proxy.example.com")
			os.Setenv("no_proxy", "localhost")
			os.Setenv("https_proxy", "proxy.example.com")
		})
		AfterEach(func() {
			os.Unsetenv("http_proxy")
			os.Unsetenv("https_proxy")
			os.Unsetenv("no_proxy")
		})
		It("copies over the proxy settings", func() {
			env := containerrunner.GetDefaultEnvs()
			Expect(env).To(Equal([]string{
				"http_proxy=proxy.example.com",
				"https_proxy=proxy.example.com",
				"no_proxy=localhost",
			}))
		})
		It("caches the property", func() {
			os.Setenv("http_proxy", "proxy.example.com")
			env := containerrunner.GetDefaultEnvs()
			cachedEnv := containerrunner.GetDefaultEnvs()
			Expect(cachedEnv).To(Equal(env))
			os.Unsetenv("http_proxy")
		})
	})

	Describe("It copies over the env variables with prefix mqtt_faas over", func() {
		BeforeEach(func() {
			os.Setenv("MQTT_FAAS_KEY_1", "VALUE_1")
		})
		AfterEach(func() {
			os.Unsetenv("MQTT_FAAS_KEY_1")
		})
		It("carries over those env values", func() {
			env := containerrunner.GetDefaultEnvs()
			Expect(env).To(ContainElement("MQTT_FAAS_KEY_1=VALUE_1"))
		})
	})
})
