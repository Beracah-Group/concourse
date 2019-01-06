package topgun_test

import (
	"bytes"
	"os"
	"os/exec"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/onsi/gomega/gbytes"
	"github.com/onsi/gomega/gexec"
)

var _ = Describe("Database secrets encryption", func() {
	configurePipelineAndTeamAndTriggerJob := func() {
		By("setting a pipeline that contains secrets")
		fly.Run("set-pipeline", "-n", "-c", "pipelines/secrets.yml", "-p", "pipeline-secrets-test")
		fly.Run("unpause-pipeline", "-p", "pipeline-secrets-test")

		By("creating a team with auth")
		setTeamSession := fly.SpawnInteractive(
			bytes.NewBufferString("y\n"),
			"set-team",
			"--team-name", "victoria",
			"--github-user", "victorias_id",
			"--github-org", "victorias_secret_org",
		)
		<-setTeamSession.Exited

		buildSession := fly.Start("trigger-job", "-w", "-j", "pipeline-secrets-test/simple-job")
		<-buildSession.Exited
		Expect(buildSession.ExitCode()).To(Equal(0))
	}

	pgDump := func() *gexec.Session {
		dump := exec.Command("pg_dump", "-U", "atc", "-h", dbInstance.IP, "atc")
		dump.Env = append(os.Environ(), "PGPASSWORD=dummy-password")
		dump.Stdin = bytes.NewBufferString("dummy-password\n")
		session, err := gexec.Start(dump, GinkgoWriter, GinkgoWriter)
		Expect(err).ToNot(HaveOccurred())
		<-session.Exited
		Expect(session.ExitCode()).To(Equal(0))
		return session
	}

	getPipeline := func() *gexec.Session {
		session := fly.Start("get-pipeline", "-p", "pipeline-secrets-test")
		<-session.Exited
		Expect(session.ExitCode()).To(Equal(0))
		return session
	}

	Describe("A deployment with encryption enabled immediately", func() {
		BeforeEach(func() {
			Deploy("deployments/concourse.yml", "-o", "operations/encryption.yml")
		})

		It("encrypts pipeline credentials", func() {
			configurePipelineAndTeamAndTriggerJob()

			By("taking a dump")
			session := pgDump()
			Expect(session).ToNot(gbytes.Say("resource_secret"))
			Expect(session).ToNot(gbytes.Say("concourse/time-resource"))
			Expect(session).ToNot(gbytes.Say("job_secret"))
		})
	})

	Describe("A deployment with encryption initially not configured", func() {
		BeforeEach(func() {
			Deploy("deployments/concourse.yml")
		})

		Context("with credentials in plaintext", func() {
			BeforeEach(func() {
				configurePipelineAndTeamAndTriggerJob()

				By("taking a dump")
				session := pgDump()
				Expect(string(session.Out.Contents())).To(ContainSubstring("resource_secret"))
				Expect(string(session.Out.Contents())).To(ContainSubstring("concourse/time-resource"))
				Expect(string(session.Out.Contents())).To(ContainSubstring("job_secret"))
			})

			Context("when redeployed with encryption enabled", func() {
				BeforeEach(func() {
					Deploy("deployments/concourse.yml", "-o", "operations/encryption.yml")
				})

				It("encrypts pipeline credentials", func() {
					By("taking a dump")
					session := pgDump()
					Expect(session).ToNot(gbytes.Say("resource_secret"))
					Expect(session).ToNot(gbytes.Say("concourse/time-resource"))
					Expect(session).ToNot(gbytes.Say("job_secret"))

					By("getting the pipeline config")
					session = getPipeline()
					Expect(string(session.Out.Contents())).To(ContainSubstring("resource_secret"))
					Expect(string(session.Out.Contents())).To(ContainSubstring("concourse/time-resource"))
					Expect(string(session.Out.Contents())).To(ContainSubstring("job_secret"))
					Expect(string(session.Out.Contents())).To(ContainSubstring("busybox"))
				})

				Context("when the encryption key is rotated", func() {
					BeforeEach(func() {
						Deploy("deployments/concourse.yml", "-o", "operations/encryption-rotated.yml")
					})

					It("can still get and set pipelines", func() {
						By("taking a dump")
						session := pgDump()
						Expect(session).ToNot(gbytes.Say("resource_secret"))
						Expect(session).ToNot(gbytes.Say("concourse/time-resource"))
						Expect(session).ToNot(gbytes.Say("job_secret"))

						By("getting the pipeline config")
						session = getPipeline()
						Expect(string(session.Out.Contents())).To(ContainSubstring("resource_secret"))
						Expect(string(session.Out.Contents())).To(ContainSubstring("concourse/time-resource"))
						Expect(string(session.Out.Contents())).To(ContainSubstring("job_secret"))
						Expect(string(session.Out.Contents())).To(ContainSubstring("busybox"))

						By("setting the pipeline again")
						fly.Run("set-pipeline", "-n", "-c", "pipelines/secrets.yml", "-p", "pipeline-secrets-test")

						By("getting the pipeline config again")
						session = getPipeline()
						Expect(string(session.Out.Contents())).To(ContainSubstring("resource_secret"))
						Expect(string(session.Out.Contents())).To(ContainSubstring("concourse/time-resource"))
						Expect(string(session.Out.Contents())).To(ContainSubstring("job_secret"))
						Expect(string(session.Out.Contents())).To(ContainSubstring("busybox"))
					})
				})

				Context("when an old key is given but all the data is already using the new key", func() {
					BeforeEach(func() {
						Deploy("deployments/concourse.yml", "-o", "operations/encryption-already-rotated.yml")
					})

					It("can still get and set pipelines", func() {
						By("taking a dump")
						session := pgDump()
						Expect(session).ToNot(gbytes.Say("resource_secret"))
						Expect(session).ToNot(gbytes.Say("concourse/time-resource"))
						Expect(session).ToNot(gbytes.Say("job_secret"))

						By("getting the pipeline config")
						session = getPipeline()
						Expect(string(session.Out.Contents())).To(ContainSubstring("resource_secret"))
						Expect(string(session.Out.Contents())).To(ContainSubstring("concourse/time-resource"))
						Expect(string(session.Out.Contents())).To(ContainSubstring("job_secret"))
						Expect(string(session.Out.Contents())).To(ContainSubstring("busybox"))

						By("setting the pipeline again")
						fly.Run("set-pipeline", "-n", "-c", "pipelines/secrets.yml", "-p", "pipeline-secrets-test")

						By("getting the pipeline config again")
						session = getPipeline()
						Expect(string(session.Out.Contents())).To(ContainSubstring("resource_secret"))
						Expect(string(session.Out.Contents())).To(ContainSubstring("concourse/time-resource"))
						Expect(string(session.Out.Contents())).To(ContainSubstring("job_secret"))
						Expect(string(session.Out.Contents())).To(ContainSubstring("busybox"))
					})
				})

				Context("when an old key and new key are both given that do not match the key in use", func() {
					var deploy *gexec.Session
					var boshLogs *gexec.Session

					BeforeEach(func() {
						boshLogs = spawnBosh("logs", "-f")

						deploy = StartDeploy("deployments/concourse.yml", "-o", "operations/encryption-bogus.yml")
						<-deploy.Exited
						Expect(deploy.ExitCode()).To(Equal(1))
					})

					AfterEach(func() {
						boshLogs.Signal(os.Interrupt)
						<-boshLogs.Exited
					})

					AfterEach(func() {
						Deploy("deployments/concourse.yml", "-o", "operations/encryption.yml")
					})

					It("fails to deploy with a useful message", func() {
						Expect(deploy).To(gbytes.Say("Review logs for failed jobs: web"))
						Expect(boshLogs).To(gbytes.Say("row encrypted with neither old nor new key"))
					})
				})

				Context("when the encryption key is removed", func() {
					BeforeEach(func() {
						Deploy("deployments/concourse.yml", "-o", "operations/encryption-removed.yml")
					})

					It("decrypts pipeline credentials", func() {
						By("taking a dump")
						session := pgDump()
						Expect(string(session.Out.Contents())).To(ContainSubstring("resource_secret"))
						Expect(string(session.Out.Contents())).To(ContainSubstring("concourse/time-resource"))
						Expect(string(session.Out.Contents())).To(ContainSubstring("job_secret"))

						By("getting the pipeline config")
						session = getPipeline()
						Expect(string(session.Out.Contents())).To(ContainSubstring("resource_secret"))
						Expect(string(session.Out.Contents())).To(ContainSubstring("concourse/time-resource"))
						Expect(string(session.Out.Contents())).To(ContainSubstring("job_secret"))
						Expect(string(session.Out.Contents())).To(ContainSubstring("busybox"))
					})
				})
			})
		})
	})
})
