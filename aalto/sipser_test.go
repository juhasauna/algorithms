package aalto

import "testing"

func Test_sipser(t *testing.T) {
	modexp0713Test(t)
}

func modexp0713Test(t *testing.T) {
	// t.Log(modexp0713(3, 10, 4, 5))
	t.Log(modexp0713(7, 8, 1, 4))
}
