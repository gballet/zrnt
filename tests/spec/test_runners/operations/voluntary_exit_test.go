package operations

import (
	"github.com/protolambda/zrnt/eth2/beacon/exits"
	"github.com/protolambda/zrnt/tests/spec/test_util"
	"testing"
)

type VoluntaryExitTestCase struct {
	test_util.BaseTransitionTest
	VoluntaryExit exits.SignedVoluntaryExit
}

func (c *VoluntaryExitTestCase) Load(t *testing.T, readPart test_util.TestPartReader) {
	c.BaseTransitionTest.Load(t, readPart)
	test_util.LoadSSZ(t, "voluntary_exit", &c.VoluntaryExit, exits.SignedVoluntaryExitSSZ, readPart)
}

func (c *VoluntaryExitTestCase) Run() error {
	state := c.Prepare()
	return state.ProcessVoluntaryExit(&c.VoluntaryExit)
}

func TestVoluntaryExit(t *testing.T) {
	test_util.RunTransitionTest(t, "operations", "voluntary_exit",
		func() test_util.TransitionTest { return new(VoluntaryExitTestCase) })
}
