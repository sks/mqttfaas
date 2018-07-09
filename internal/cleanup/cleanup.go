package cleanup

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/docker/docker/api/types/filters"
)

//Cleaner ...
type Cleaner struct {
	dockerCLI DockerCLI
	errChan   chan error
}

//New ...
func New(d time.Duration, dockerCLI DockerCLI, errChan chan error) {
	cleaner := &Cleaner{
		dockerCLI,
		errChan,
	}
	ticker := time.NewTicker(d)
	go func() {
		for {
			<-ticker.C
			cleaner.Cleanup()
		}
	}()
}

//Cleanup ...
func (c *Cleaner) Cleanup() {
	ctx := context.Background()
	filter := filters.NewArgs()
	createdBefore := time.Now().Add(-5 * time.Minute)
	filter.Add("until", fmt.Sprintf("%d", createdBefore.Unix()))
	report, err := c.dockerCLI.ContainersPrune(ctx, filter)
	if err != nil {
		c.errChan <- err
		return
	}
	log.Printf("Removing %d containers created before %s \n", len(report.ContainersDeleted), createdBefore)
}
