package main

import (
	"github.com/fsouza/go-dockerclient"
)
type dockerManager struct {
	addr   string
	parent string
}

func newDockerManager(addr, parent string) *dockerManager {
	return &dockerManager{addr: addr, parent: parent}
}

// Return a list of all running containers on the system
func (m *dockerManager) Containers() ([]*container, error) {
	client, err := docker.NewClient(m.addr)
	if err != nil {
		return nil, err
	}

	client.SetTimeout(time.Minute)
	// Get all *running* containers
	containers, err := client.ListContainers(docker.ListContainersOptions{All: false})
	if err != nil {
		return nil, err
	}

	cl := []*container{}
	for _, c := range containers {
		cl = append(cl, &container{
			ID:    c.ID,
			Name:  c.Names[0][1:], // FIXME: This isn't a very good solution but the best I could think of.
			Image: c.Image,
		})
	}
	return cl, nil
}

func (m *dockerManager) Parent() string {
	return m.parent
}
