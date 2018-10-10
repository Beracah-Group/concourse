package testflight_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
)

var _ = Describe("Resource version", func() {
	Describe("when the version is not pinned on the resource", func() {
		Describe("version: latest", func() {
			BeforeEach(func() {
				setAndUnpausePipeline("fixtures/resource-version-latest.yml")
			})

			It("only runs builds with latest version", func() {
				guid1 := newMockVersion("some-resource", "guid1")
				watch := fly("trigger-job", "-j", inPipeline("some-passing-job"), "-w")
				Expect(watch).To(gbytes.Say(guid1))

				_ = newMockVersion("some-resource", "guid2")
				_ = newMockVersion("some-resource", "guid3")
				guid4 := newMockVersion("some-resource", "guid4")

				watch = fly("trigger-job", "-j", inPipeline("some-passing-job"), "-w")
				Expect(watch).To(gbytes.Say(guid4))
			})
		})

		Describe("version: every", func() {
			BeforeEach(func() {
				setAndUnpausePipeline("fixtures/resource-version-every.yml")
			})

			It("runs builds with every version", func() {
				guid1 := newMockVersion("some-resource", "guid1")
				watch := fly("trigger-job", "-j", inPipeline("some-passing-job"), "-w")
				Expect(watch).To(gbytes.Say(guid1))

				guid2 := newMockVersion("some-resource", "guid2")
				guid3 := newMockVersion("some-resource", "guid3")
				guid4 := newMockVersion("some-resource", "guid4")

				watch = fly("trigger-job", "-j", inPipeline("some-passing-job"), "-w")
				Expect(watch).To(gbytes.Say(guid2))

				watch = fly("trigger-job", "-j", inPipeline("some-passing-job"), "-w")
				Expect(watch).To(gbytes.Say(guid3))

				watch = fly("trigger-job", "-j", inPipeline("some-passing-job"), "-w")
				Expect(watch).To(gbytes.Say(guid4))
			})
		})

		Describe("version: pinned", func() {
			BeforeEach(func() {
				setAndUnpausePipeline("fixtures/resource-version-every.yml")
			})

			It("only runs builds with the pinned version", func() {
				guid1 := newMockVersion("some-resource", "guid1")

				watch := fly("trigger-job", "-j", inPipeline("some-passing-job"), "-w")
				Eventually(watch).Should(gbytes.Say(guid1))

				_ = newMockVersion("some-resource", "guid2")
				guid3 := newMockVersion("some-resource", "guid3")
				_ = newMockVersion("some-resource", "guid4")

				setPipeline("fixtures/pinned-version.yml", "-v", "pinned_version="+guid3)

				watch = fly("trigger-job", "-j", inPipeline("some-passing-job"), "-w")
				Expect(watch).To(gbytes.Say(guid3))
			})
		})
	})

	Describe("when the version is pinned on the resource", func() {
		var olderGUID string
		var pinnedGUID string
		var versionConfig string

		BeforeEach(func() {
			versionConfig = "nil"

			setAndUnpausePipeline(
				"fixtures/pinned-resource-simple-trigger.yml",
				"-v", "pinned_resource_version=bogus",
				"-y", "version_config="+versionConfig,
			)

			olderGUID = newMockVersion("some-resource", "older")
			pinnedGUID = newMockVersion("some-resource", "pinned")
			_ = newMockVersion("some-resource", "newer")
		})

		JustBeforeEach(func() {
			setPipeline(
				"fixtures/pinned-resource-simple-trigger.yml",
				"-v", "pinned_resource_version="+pinnedGUID,
				"-y", "version_config="+versionConfig,
			)
		})

		Describe("version: latest", func() {
			BeforeEach(func() {
				versionConfig = "latest"
			})

			It("only runs builds with pinned version", func() {
				watch := fly("trigger-job", "-j", inPipeline("some-passing-job"), "-w")
				Expect(watch).To(gbytes.Say(pinnedGUID))
			})
		})

		Describe("version: every", func() {
			BeforeEach(func() {
				versionConfig = "every"
			})

			It("only runs builds with pinned version", func() {
				watch := fly("trigger-job", "-j", inPipeline("some-passing-job"), "-w")
				Expect(watch).To(gbytes.Say(pinnedGUID))

				watch = fly("trigger-job", "-j", inPipeline("some-passing-job"), "-w")
				Expect(watch).To(gbytes.Say(pinnedGUID))
			})
		})

		Describe("version: pinned", func() {
			BeforeEach(func() {
				versionConfig = "version:" + olderGUID
			})

			It("only runs builds with the pinned version", func() {
				watch := fly("trigger-job", "-j", inPipeline("some-passing-job"), "-w")
				Expect(watch).To(gbytes.Say(pinnedGUID))
			})
		})
	})
})
