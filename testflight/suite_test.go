package testflight_test

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"regexp"
	"strings"
	"testing"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/concourse/concourse/go-concourse/concourse"
	"github.com/nu7hatch/gouuid"
	"github.com/onsi/gomega/gexec"
	"golang.org/x/oauth2"
)

const flyTarget = "tf"
const pipelinePrefix = "tf-pipeline"
const teamName = "testflight"

type suiteConfig struct {
	FlyBin      string `json:"fly"`
	ATCURL      string `json:"atc_url"`
	ATCUsername string `json:"atc_username"`
	ATCPassword string `json:"atc_password"`
}

var (
	config = suiteConfig{
		FlyBin:      "fly",
		ATCURL:      "http://localhost:8080",
		ATCUsername: "test",
		ATCPassword: "test",
	}

	client concourse.Client
	team   concourse.Team

	pipelineName string
	tmp          string
)

func TestTestflight(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "TestFlight Suite")
}

var _ = SynchronizedBeforeSuite(func() []byte {
	var err error
	config.FlyBin, err = gexec.Build("github.com/concourse/concourse/fly")
	Expect(err).ToNot(HaveOccurred())

	atcURL := os.Getenv("ATC_URL")
	if atcURL != "" {
		config.ATCURL = atcURL
	}

	atcUsername := os.Getenv("ATC_USERNAME")
	if atcUsername != "" {
		config.ATCUsername = atcUsername
	}

	atcPassword := os.Getenv("ATC_PASSWORD")
	if atcPassword != "" {
		config.ATCPassword = atcPassword
	}

	payload, err := json.Marshal(config)
	Expect(err).ToNot(HaveOccurred())

	Eventually(func() *gexec.Session {
		login := spawnFlyLogin()
		<-login.Exited
		return login
	}, 2*time.Minute, time.Second).Should(gexec.Exit(0))

	fly("set-team", "--non-interactive", "-n", teamName, "--local-user", config.ATCUsername)
	wait(spawnFlyLogin("-n", teamName))

	for _, ps := range flyTable("pipelines") {
		name := ps["name"]
		if strings.HasPrefix(name, pipelinePrefix) {
			fly("destroy-pipeline", "-n", "-p", name)
		}
	}

	return payload
}, func(data []byte) {
	err := json.Unmarshal(data, &config)
	Expect(err).ToNot(HaveOccurred())

	client = concourseClient()
	team = client.Team(teamName)
})

var _ = SynchronizedAfterSuite(func() {
}, func() {
	gexec.CleanupBuildArtifacts()
})

var _ = BeforeEach(func() {
	SetDefaultEventuallyTimeout(5 * time.Minute)
	SetDefaultEventuallyPollingInterval(time.Second)
	SetDefaultConsistentlyDuration(time.Minute)
	SetDefaultConsistentlyPollingInterval(time.Second)

	var err error
	tmp, err = ioutil.TempDir("", "testflight-tmp")
	Expect(err).ToNot(HaveOccurred())

	pipelineName = randomPipelineName()
})

var _ = AfterEach(func() {
	fly("destroy-pipeline", "-n", "-p", pipelineName)
	Expect(os.RemoveAll(tmp)).To(Succeed())
})

func randomPipelineName() string {
	guid, err := uuid.NewV4()
	Expect(err).ToNot(HaveOccurred())

	return fmt.Sprintf("%s-%d-%s", pipelinePrefix, GinkgoParallelNode(), guid)
}

func fly(argv ...string) *gexec.Session {
	sess := spawnFly(argv...)
	wait(sess)
	return sess
}

func flyIn(dir string, argv ...string) *gexec.Session {
	sess := spawnFlyIn(dir, argv...)
	wait(sess)
	return sess
}

func concourseClient() concourse.Client {
	token, err := fetchToken(config.ATCURL, config.ATCUsername, config.ATCPassword)
	Expect(err).NotTo(HaveOccurred())

	httpClient := &http.Client{
		Transport: &oauth2.Transport{
			Source: oauth2.StaticTokenSource(token),
			Base: &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			},
		},
	}

	return concourse.NewClient(config.ATCURL, httpClient, false)
}

func fetchToken(atcURL string, username, password string) (*oauth2.Token, error) {
	oauth2Config := oauth2.Config{
		ClientID:     "fly",
		ClientSecret: "Zmx5",
		Endpoint:     oauth2.Endpoint{TokenURL: atcURL + "/sky/token"},
		Scopes:       []string{"openid", "profile", "email", "federated:id"},
	}

	return oauth2Config.PasswordCredentialsToken(context.Background(), username, password)
}

func spawnFlyLogin(args ...string) *gexec.Session {
	return spawnFly(append([]string{"login", "-c", config.ATCURL, "-u", config.ATCUsername, "-p", config.ATCPassword}, args...)...)
}

func spawnFly(argv ...string) *gexec.Session {
	return spawn(config.FlyBin, append([]string{"-t", flyTarget}, argv...)...)
}

func spawnFlyIn(dir string, argv ...string) *gexec.Session {
	return spawnIn(dir, config.FlyBin, append([]string{"-t", flyTarget}, argv...)...)
}

func spawnFlyInteractive(stdin io.Reader, argv ...string) *gexec.Session {
	return spawnInteractive(stdin, config.FlyBin, append([]string{"-t", flyTarget}, argv...)...)
}

func run(argc string, argv ...string) {
	wait(spawn(argc, argv...))
}

func spawn(argc string, argv ...string) *gexec.Session {
	By("running: " + argc + " " + strings.Join(argv, " "))
	cmd := exec.Command(argc, argv...)
	session, err := gexec.Start(cmd, GinkgoWriter, GinkgoWriter)
	Expect(err).ToNot(HaveOccurred())
	return session
}

func spawnIn(dir string, argc string, argv ...string) *gexec.Session {
	By("running in " + dir + ": " + argc + " " + strings.Join(argv, " "))
	cmd := exec.Command(argc, argv...)
	cmd.Dir = dir
	session, err := gexec.Start(cmd, GinkgoWriter, GinkgoWriter)
	Expect(err).ToNot(HaveOccurred())
	return session
}

func spawnInteractive(stdin io.Reader, argc string, argv ...string) *gexec.Session {
	By("interactively running: " + argc + " " + strings.Join(argv, " "))
	cmd := exec.Command(argc, argv...)
	cmd.Stdin = stdin
	session, err := gexec.Start(cmd, GinkgoWriter, GinkgoWriter)
	Expect(err).ToNot(HaveOccurred())
	return session
}

func wait(session *gexec.Session) {
	<-session.Exited
	Expect(session.ExitCode()).To(Equal(0))
}

var colSplit = regexp.MustCompile(`\s{2,}`)

func flyTable(argv ...string) []map[string]string {
	session := spawnFly(append([]string{"--print-table-headers"}, argv...)...)
	<-session.Exited
	Expect(session.ExitCode()).To(Equal(0))

	result := []map[string]string{}
	var headers []string

	rows := strings.Split(string(session.Out.Contents()), "\n")
	for i, row := range rows {
		columns := colSplit.Split(strings.TrimSpace(row), -1)

		if i == 0 {
			headers = columns
			continue
		}

		if row == "" {
			continue
		}

		result = append(result, map[string]string{})

		Expect(columns).To(HaveLen(len(headers)))

		for j, header := range headers {
			if header == "" || columns[j] == "" {
				continue
			}

			result[i-1][header] = columns[j]
		}
	}

	return result
}

func setAndUnpausePipeline(config string, args ...string) {
	setPipeline(config, args...)
	fly("unpause-pipeline", "-p", pipelineName)
}

func setPipeline(config string, args ...string) {
	sp := []string{"set-pipeline", "-n", "-p", pipelineName, "-c", config}
	fly(append(sp, args...)...)
}

func inPipeline(thing string) string {
	return pipelineName + "/" + thing
}

func newMockVersion(resourceName string, tag string) string {
	guid, err := uuid.NewV4()
	Expect(err).ToNot(HaveOccurred())

	version := guid.String() + "-" + tag

	fly("check-resource", "-r", inPipeline(resourceName), "-f", "version:"+version)

	return version
}

func waitForBuildAndWatch(jobName string, buildName ...string) *gexec.Session {
	args := []string{"watch", "-j", inPipeline(jobName)}

	if len(buildName) > 0 {
		args = append(args, "-b", buildName[0])
	}

	keepPollingCheck := regexp.MustCompile("job has no builds|build not found|failed to get build")
	for {
		session := spawnFly(args...)
		<-session.Exited

		if session.ExitCode() == 1 {
			output := strings.TrimSpace(string(session.Err.Contents()))
			if keepPollingCheck.MatchString(output) {
				// build hasn't started yet; keep polling
				time.Sleep(time.Second)
				continue
			}
		}

		return session
	}
}
