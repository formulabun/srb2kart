package network

import "testing"

func TestFileScanner(t *testing.T) {
	testData := []struct {
		data          filesNeeded
		expectedNames []string
	}{
		{
			filesNeeded{
				0,
				2,
				0,
				[915]byte{
					1, 0, 90, 224, 183, 0, 86, 83, 95, 83, 82, 66, 50, 84, 104, 105, 99, 99, 46, 112, 107, 51, 0,
					90, 10, 80, 195, 31, 160, 113, 229, 102, 218, 100, 199, 125, 186, 237, 205, 1, 0,
				},
			},
			[]string{},
		},
	}

	for _, data := range testData {
		_, err := scanFiles(data.data)
		if err != nil {
			t.Fatalf("Got error while parsing files: %s\n", err)
		}

	}
}
