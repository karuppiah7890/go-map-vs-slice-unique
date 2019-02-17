package mapvsslice

type uniqueTest struct {
	input          []string
	expectedOutput []string
}

var uniqueTestCases = []uniqueTest{
	{
		input:          []string{"this", "is", "all", "unique"},
		expectedOutput: []string{"this", "is", "all", "unique"},
	},
	{
		input:          []string{"there", "are", "some", "duplicates", "duplicates", "here"},
		expectedOutput: []string{"there", "are", "some", "duplicates", "here"},
	},
	{
		input:          []string{"some", "some", "some", "some", "words", "just", "are", "cool", "some", "some", "some"},
		expectedOutput: []string{"some", "words", "are", "just", "cool"},
	},
}
