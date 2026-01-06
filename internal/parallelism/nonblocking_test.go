package parallelism

import (
	"testing"
	"yoshiyoshifujii/fpingolang/internal/common"
	testing2 "yoshiyoshifujii/fpingolang/internal/testing"

	"github.com/stretchr/testify/assert"
)

var (
	genParBoolean = testing2.Map(testing2.GenBoolean, func(a bool) Par[bool] {
		return ParUnit(a)
	})
	genParInt = testing2.Map(common.GenShortNumber(), func(a int) Par[int] {
		return ParUnit(a)
	})
	es = GoExecutor{}
)

func checkChoice(t *testing.T, t1 Par[int], t2 Par[int], t3 Par[bool], choice Par[int]) {
	actual := choice.Run(es)
	expected := t2.Run(es)
	if t3.Run(es) {
		expected = t1.Run(es)
	}
	assert.Equal(t, expected, actual)
}

func TestParChoice(t *testing.T) {
	common.PropSuiteTest(t, "ParChoice", testing2.GenProduct3(genParInt, genParInt, genParBoolean), func(p3 testing2.Pair3[Par[int], Par[int], Par[bool]]) {
		t1, t2, t3 := p3.First.First, p3.First.Second, p3.Second
		checkChoice(t, t1, t2, t3, ParChoice[int](t3, t1, t2))
	})
}
