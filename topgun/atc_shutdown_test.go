package topgun_test

import (
	"regexp"
	"time"

	_ "github.com/lib/pq"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
	"github.com/onsi/gomega/gexec"
)

var _ = Describe("[#137641079] ATC Shutting down", func() {
	Context("with two atcs available", func() {
		var atcs []boshInstance
		var atc0URL string
		var atc1URL string

		BeforeEach(func() {
			Deploy(
				"deployments/concourse.yml",
				"-o", "operations/web-instances.yml",
				"-v", "web_instances=2",
				"-o", "operations/slow-tracking.yml",
			)

			waitForRunningWorker()

			atcs = JobInstances("web")
			atc0URL = "http://" + atcs[0].IP + ":8080"
			atc1URL = "http://" + atcs[1].IP + ":8080"

			fly.Login(atcUsername, atcPassword, atc0URL)
		})

		Context("when one of the ATCS is stopped", func() {
			var stopSession *gexec.Session
			BeforeEach(func() {
				By("stopping one of the web instances")
				stopSession = spawnBosh("stop", atcs[1].Name)
				Eventually(stopSession).Should(gexec.Exit(0))
			})

			AfterEach(func() {
				restartSession := spawnBosh("start", atcs[0].Name)
				<-restartSession.Exited
				Eventually(restartSession).Should(gexec.Exit(0))
			})

			Describe("workers registering with random TSA address", func() {
				It("recovers from the TSA going down by registering with a random TSA", func() {
					waitForRunningWorker()

					By("stopping the other web instance")
					stopSession = spawnBosh("stop", atcs[0].Name)
					Eventually(stopSession).Should(gexec.Exit(0))

					By("starting the stopped web instance")
					startSession := spawnBosh("start", atcs[1].Name)
					Eventually(startSession).Should(gexec.Exit(0))

					atcs = JobInstances("web")
					atc0URL = "http://" + atcs[0].IP + ":8080"
					atc1URL = "http://" + atcs[1].IP + ":8080"

					fly.Login(atcUsername, atcPassword, atc1URL)

					waitForRunningWorker()
				})
			})

			Describe("tracking builds previously tracked by shutdown ATC", func() {
				var buildID string

				BeforeEach(func() {
					By("executing a task")
					buildSession := fly.Start("execute", "-c", "tasks/wait.yml")
					Eventually(buildSession).Should(gbytes.Say("executing build"))

					buildRegex := regexp.MustCompile(`executing build (\d+)`)
					matches := buildRegex.FindSubmatch(buildSession.Out.Contents())
					buildID = string(matches[1])

					Eventually(buildSession).Should(gbytes.Say("waiting for /tmp/stop-waiting"))

					By("starting the stopped web instance")
					startSession := spawnBosh("start", atcs[1].Name)
					<-startSession.Exited
					Eventually(startSession).Should(gexec.Exit(0))

					fly.Login(atcUsername, atcPassword, atc1URL)
				})

				AfterEach(func() {
					restartSession := spawnBosh("start", atcs[0].Name)
					<-restartSession.Exited
					Eventually(restartSession).Should(gexec.Exit(0))
				})

				Context("when the atc tracking the build shuts down", func() {
					JustBeforeEach(func() {
						By("stopping the first web instance")
						landSession := spawnBosh("stop", atcs[0].Name)
						<-landSession.Exited
						Eventually(landSession).Should(gexec.Exit(0))
					})

					It("continues tracking the build progress", func() {
						By("waiting for another atc to attach to process")
						watchSession := fly.Start("watch", "-b", buildID)
						Eventually(watchSession).Should(gbytes.Say("waiting for /tmp/stop-waiting"))
						time.Sleep(10 * time.Second)

						By("hijacking the build to tell it to finish")
						hijackSession := fly.Start(
							"hijack",
							"-b", buildID,
							"-s", "one-off",
							"touch", "/tmp/stop-waiting",
						)
						<-hijackSession.Exited
						Expect(hijackSession.ExitCode()).To(Equal(0))

						By("waiting for the build to exit")
						Eventually(watchSession, 1*time.Minute).Should(gbytes.Say("done"))
						<-watchSession.Exited
						Expect(watchSession.ExitCode()).To(Equal(0))
					})
				})
			})
		})
	})
})
