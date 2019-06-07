package operations

import (
	"github.com/protolambda/zrnt/eth2/beacon"
	"github.com/protolambda/zrnt/eth2/beacon/block_processing"
	. "github.com/protolambda/zrnt/tests/spec/test_util"
	"testing"
)

type VoluntaryExitTestCase struct {
	VoluntaryExit     *beacon.VoluntaryExit
	StateTransitionTestBase `mapstructure:",squash"`
}

func (testCase *VoluntaryExitTestCase) Process() error {
	return block_processing.ProcessVoluntaryExit(testCase.Pre, testCase.VoluntaryExit)
}

func (testCase *VoluntaryExitTestCase) Run(t *testing.T) {
	RunTest(t, testCase)
}

func TestVoluntaryExit(t *testing.T) {
	RunSuitesInPath("operations/voluntary_exit/",
		func(raw interface{}) (interface{}, interface {}) { return new(VoluntaryExitTestCase), raw }, t)
}
