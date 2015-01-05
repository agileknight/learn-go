package camelup

import "testing"

type MockRandInt struct {
	returnVal  int
	calledWith int
}

func (this *MockRandInt) Intn(n int) int {
	this.calledWith = n
	return this.returnVal
}

func (this *MockRandInt) Returns(n int) {
	this.returnVal = n
}

func (this *MockRandInt) CalledWith() int {
	return this.calledWith
}

func TestBoundedDiceRoll(t *testing.T) {
	rand := MockRandInt{}
	dice := BoundedDice{
		rand:     &rand,
		minValue: 3,
		maxValue: 5,
	}

	cases := []struct {
		randResult     int
		expectedParam  int
		expectedResult int
	}{
		{0, 3, 3},
		{1, 3, 4},
		{2, 3, 5},
	}

	for _, c := range cases {
		rand.Returns(c.randResult)
		got := dice.Roll()
		if got != c.expectedResult {
			t.Errorf("Expected %d but found %d", c.expectedResult, got)
		}
		if param := rand.CalledWith(); param != c.expectedParam {
			t.Errorf("Expected param %d but found %d", c.expectedParam, param)
		}
	}
}
