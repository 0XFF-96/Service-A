package tests

import (
	"fmt"
	"testing"

	"github.com/0XFF-96/Service-A/business/data/dbtest"
	"github.com/0XFF-96/Service-A/foundation/docker"
)

var c *docker.Container

func TestMain(m *testing.M) {
	var err error
	c, err = dbtest.StartDB()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer dbtest.StopDB(c)

	m.Run()
}
