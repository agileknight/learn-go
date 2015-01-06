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
		expectedParam  int
		randResult     int
		expectedResult int
	}{
		{3, 0, 3},
		{3, 1, 4},
		{3, 2, 5},
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

func TestNoDuplicatesBoundedDiceRoll(t *testing.T) {
	rand := MockRandInt{}
	dice := NoDuplicatesBoundedDice{
		rand:     &rand,
		minValue: 3,
		maxValue: 5,
	}

	cases := []struct {
		expectedParams  []int
		randResults     []int
		expectedResults []int
	}{
		{[]int{3, 3, 3}, []int{0, 1, 2}, []int{3, 4, 5}},
	}

	for _, c := range cases {
		for i := range c.randResults {
			rand.Returns(c.randResults[i])
			got := dice.Roll()
			if got != c.expectedResults[i] {
				t.Errorf("Expected %d but found %d", c.expectedResults[i], got)
			}
			if param := rand.CalledWith(); param != c.expectedParams[i] {
				t.Errorf("Expected param %d but found %d", c.expectedParams[i], param)
			}	
		}
	}
}
