package executehelpers

import (
	"fmt"

	"github.com/concourse/concourse/fly/ui"
	"github.com/concourse/concourse/go-concourse/concourse"
	"github.com/concourse/go-archive/tgzfs"
)

func Download(client concourse.Client, buildID int, output Output) {
	out, found, err := client.ReadOutputFromBuildPlan(buildID, output.Plan.ID)
	if err != nil {
		fmt.Fprintf(ui.Stderr, "failed to download output '%s': %s", output.Name, err)
		return
	}

	if !found {
		fmt.Fprintf(ui.Stderr, "build disappeared while downloading '%s'", output.Name)
		return
	}

	defer out.Close()

	err = tgzfs.Extract(out, output.Path)
	if err != nil {
		fmt.Fprintf(ui.Stderr, "failed to extract output '%s': %s", output.Name, err)
		return
	}
}
