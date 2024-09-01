package welcome

import "testing"

func TestSayWelcome(t * testing.T) {

	subtests := []struct {
		items []string
		result string
	}{
		{
			// items:[]string{"Adi"},
			result: "Welcome, Poorva!",
		},
	}

	for _, st := range subtests {
		// got := SayWelcome(st.items)
		// want := st.result
		// if want != got {
		// 	t.Errorf("wanted %s, got %s", want, got)
		// }
		if s:= SayWelcome(st.items); s!= st.result {
			t.Errorf("wanted %s [%v], got %s", st.result, st.items , s)
		}
	}


	// want := "Welcome, test!"
	// got := SayWelcome([]string{"test"})
	// if want != got {
	// 	t.Errorf("wanted %s, got %s", want, got)
	// }
}