package topgun_test

import (
	"fmt"
	"time"

	_ "github.com/lib/pq"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Garbage collecting resource containers", func() {
	Describe("A container that is used by resource checking on freshly deployed worker", func() {
		BeforeEach(func() {
			Deploy(
				"deployments/concourse.yml",
				"-o", "operations/worker-instances.yml",
				"-v", "worker_instances=2",
			)
		})

		It("is recreated in database and worker", func() {
			By("setting pipeline that creates resource cache")
			fly.Run("set-pipeline", "-n", "-c", "pipelines/get-task-changing-resource.yml", "-p", "volume-gc-test")

			By("unpausing the pipeline")
			fly.Run("unpause-pipeline", "-p", "volume-gc-test")

			By("checking resource")
			fly.Run("check-resource", "-r", "volume-gc-test/tick-tock")

			By("getting the resource config container")
			containers := flyTable("containers")
			var checkContainerHandle string
			for _, container := range containers {
				if container["type"] == "check" {
					checkContainerHandle = container["handle"]
					break
				}
			}
			Expect(checkContainerHandle).NotTo(BeEmpty())

			By(fmt.Sprintf("eventually expiring the resource config container: %s", checkContainerHandle))
			Eventually(func() bool {
				containers := flyTable("containers")
				for _, container := range containers {
					if container["type"] == "check" && container["handle"] == checkContainerHandle {
						return true
					}
				}
				return false
			}, 10*time.Minute, 10*time.Second).Should(BeFalse())

			By("checking resource again")
			fly.Run("check-resource", "-r", "volume-gc-test/tick-tock")

			By("getting the resource config container")
			containers = flyTable("containers")
			var newCheckContainerHandle string
			for _, container := range containers {
				if container["type"] == "check" {
					newCheckContainerHandle = container["handle"]
					break
				}
			}
			Expect(newCheckContainerHandle).NotTo(Equal(checkContainerHandle))
		})
	})

	Describe("container for resource that is removed from pipeline", func() {
		BeforeEach(func() {
			Deploy("deployments/concourse.yml")
		})

		It("has its resource config, resource config uses and container removed", func() {
			By("setting pipeline that creates resource config")
			fly.Run("set-pipeline", "-n", "-c", "pipelines/get-task-changing-resource.yml", "-p", "resource-gc-test")

			By("unpausing the pipeline")
			fly.Run("unpause-pipeline", "-p", "resource-gc-test")

			By("checking resource")
			fly.Run("check-resource", "-r", "resource-gc-test/tick-tock")

			By("getting the resource config")
			var resourceConfigsNum int
			err := psql.Select("COUNT(id)").From("resource_configs").RunWith(dbConn).QueryRow().Scan(&resourceConfigsNum)
			Expect(err).ToNot(HaveOccurred())

			Expect(resourceConfigsNum).To(Equal(1))

			By("getting the resource config container")
			containers := flyTable("containers")
			var checkContainerHandle string
			for _, container := range containers {
				if container["type"] == "check" {
					checkContainerHandle = container["handle"]
					break
				}
			}
			Expect(checkContainerHandle).NotTo(BeEmpty())

			By("updating pipeline and removing resource")
			fly.Run("set-pipeline", "-n", "-c", "pipelines/task-waiting.yml", "-p", "resource-gc-test")

			By("eventually expiring the resource config")
			Eventually(func() int {
				var resourceConfigsNum int
				err := psql.Select("COUNT(id)").From("resource_configs").RunWith(dbConn).QueryRow().Scan(&resourceConfigsNum)
				Expect(err).ToNot(HaveOccurred())

				return resourceConfigsNum
			}, 5*time.Minute, 10*time.Second).Should(Equal(0))

			By(fmt.Sprintf("eventually expiring the resource config container: %s", checkContainerHandle))
			Eventually(func() bool {
				containers := flyTable("containers")
				for _, container := range containers {
					if container["type"] == "check" && container["handle"] == checkContainerHandle {
						return true
					}
				}
				return false
			}, 5*time.Minute, 10*time.Second).Should(BeFalse())
		})
	})

	Describe("container for resource when pipeline is paused", func() {
		BeforeEach(func() {
			Deploy("deployments/concourse.yml")
		})

		It("has its resource config, resource config uses and container removed", func() {
			By("setting pipeline that creates resource config")
			fly.Run("set-pipeline", "-n", "-c", "pipelines/get-task-changing-resource.yml", "-p", "resource-gc-test")

			By("unpausing the pipeline")
			fly.Run("unpause-pipeline", "-p", "resource-gc-test")

			By("checking resource")
			fly.Run("check-resource", "-r", "resource-gc-test/tick-tock")

			By("getting the resource config")
			var resourceConfigsNum int
			err := psql.Select("COUNT(id)").From("resource_configs").RunWith(dbConn).QueryRow().Scan(&resourceConfigsNum)
			Expect(err).ToNot(HaveOccurred())

			Expect(resourceConfigsNum).To(Equal(1))

			By("getting the resource config container")
			containers := flyTable("containers")
			var checkContainerHandle string
			for _, container := range containers {
				if container["type"] == "check" {
					checkContainerHandle = container["handle"]
					break
				}
			}
			Expect(checkContainerHandle).NotTo(BeEmpty())

			By("pausing the pipeline")
			fly.Run("pause-pipeline", "-p", "resource-gc-test")

			By("eventually expiring the resource config")
			Eventually(func() int {
				var resourceConfigsNum int
				err := psql.Select("COUNT(id)").From("resource_configs").RunWith(dbConn).QueryRow().Scan(&resourceConfigsNum)
				Expect(err).ToNot(HaveOccurred())

				return resourceConfigsNum
			}, 5*time.Minute, 10*time.Second).Should(Equal(0))

			By(fmt.Sprintf("eventually expiring the resource config container: %s", checkContainerHandle))
			Eventually(func() bool {
				containers := flyTable("containers")
				for _, container := range containers {
					if container["type"] == "check" && container["handle"] == checkContainerHandle {
						return true
					}
				}
				return false
			}, 5*time.Minute, 10*time.Second).Should(BeFalse())
		})
	})

	Describe("container for resource that is updated", func() {
		BeforeEach(func() {
			Deploy("deployments/concourse.yml")
		})

		It("has its resource config, resource config uses and container removed", func() {
			By("setting pipeline that creates resource config")
			fly.Run("set-pipeline", "-n", "-c", "pipelines/get-task.yml", "-p", "resource-gc-test")

			By("unpausing the pipeline")
			fly.Run("unpause-pipeline", "-p", "resource-gc-test")

			By("checking resource")
			fly.Run("check-resource", "-r", "resource-gc-test/tick-tock")

			By("getting the resource config")
			var originalResourceConfigID int
			err := psql.Select("id").From("resource_configs").RunWith(dbConn).QueryRow().Scan(&originalResourceConfigID)
			Expect(err).ToNot(HaveOccurred())

			Expect(originalResourceConfigID).NotTo(BeZero())

			By("getting the resource config container")
			containers := flyTable("containers")
			var originalCheckContainerHandle string
			for _, container := range containers {
				if container["type"] == "check" {
					originalCheckContainerHandle = container["handle"]
					break
				}
			}
			Expect(originalCheckContainerHandle).NotTo(BeEmpty())

			By("updating pipeline with new resource configuration")
			fly.Run("set-pipeline", "-n", "-c", "pipelines/get-task-changing-resource.yml", "-p", "resource-gc-test")

			By("eventually expiring the resource config")
			Eventually(func() int {
				var resourceConfigsNum int
				err := psql.Select("COUNT(id)").From("resource_configs").Where("id = $1", originalResourceConfigID).RunWith(dbConn).QueryRow().Scan(&resourceConfigsNum)
				Expect(err).ToNot(HaveOccurred())

				return resourceConfigsNum
			}, 5*time.Minute, 10*time.Second).Should(Equal(0))

			By(fmt.Sprintf("eventually expiring the resource config container: %s", originalCheckContainerHandle))
			Eventually(func() bool {
				containers := flyTable("containers")
				for _, container := range containers {
					if container["type"] == "check" && container["handle"] == originalCheckContainerHandle {
						return true
					}
				}
				return false
			}, 5*time.Minute, 10*time.Second).Should(BeFalse())
		})
	})

	Describe("container for resource checking", func() {
		BeforeEach(func() {
			Deploy("deployments/concourse.yml", "-o", "operations/fast-gc.yml")
		})

		It("is not immediately removed", func() {
			By("setting pipeline that creates resource config")
			fly.Run("set-pipeline", "-n", "-c", "pipelines/get-task.yml", "-p", "resource-gc-test")

			By("unpausing the pipeline")
			fly.Run("unpause-pipeline", "-p", "resource-gc-test")

			By("checking resource")
			fly.Run("check-resource", "-r", "resource-gc-test/tick-tock")

			Consistently(func() string {
				By("getting the resource config container")
				containers := flyTable("containers")
				for _, container := range containers {
					if container["type"] == "check" {
						return container["handle"]
					}
				}

				return ""
			}, 2*time.Minute).ShouldNot(BeEmpty())
		})

		Context("when two teams use identical configuration", func() {
			var teamName = "A-Team"

			It("doesn't create many containers for one resource check", func() {
				By("setting pipeline that creates resource config")
				fly.Run("set-pipeline", "-n", "-c", "pipelines/get-task.yml", "-p", "resource-gc-test")

				By("unpausing the pipeline")
				fly.Run("unpause-pipeline", "-p", "resource-gc-test")

				By("checking resource")
				fly.Run("check-resource", "-r", "resource-gc-test/tick-tock")

				By("creating another team")
				fly.Run("set-team", "--non-interactive", "--team-name", teamName, "--local-user", atcUsername)

				fly.Run("login", "-c", atcExternalURL, "-n", teamName, "-u", atcUsername, "-p", atcPassword)

				By("setting pipeline that creates an identical resource config")
				fly.Run("set-pipeline", "-n", "-c", "pipelines/get-task.yml", "-p", "resource-gc-test")

				By("unpausing the pipeline")
				fly.Run("unpause-pipeline", "-p", "resource-gc-test")

				By("checking resource excessively")
				for i := 0; i < 20; i++ {
					fly.Run("check-resource", "-r", "resource-gc-test/tick-tock")
				}

				otherTeamCheckCount := len(flyTable("containers"))
				Expect(otherTeamCheckCount).To(Equal(1))

				By("checking resource excessively")
				fly.Run("login", "-c", atcExternalURL, "-n", "main", "-u", atcUsername, "-p", atcPassword)
				for i := 0; i < 20; i++ {
					fly.Run("check-resource", "-r", "resource-gc-test/tick-tock")
				}

				mainTeamCheckCount := len(flyTable("containers"))
				Expect(mainTeamCheckCount).To(Equal(1))
			})
		})
	})
})
