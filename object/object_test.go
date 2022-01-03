package object

import "testing"

func TestStringHashKey(t *testing.T) {
	hello1 := &String{Value: "hello world"}
	hello2 := &String{Value: "hello world"}
	diff1 := &String{Value: "Johny Johny"}
	diff2 := &String{Value: "Johny Johny"}

	a := &Integer{Value: 1}
	b := &Boolean{Value: true}

	if a.HashKey() == b.HashKey() {
		t.Errorf("interger 1 and boolean true have same hash keys")
	}

	if hello1.HashKey() != hello2.HashKey() {
		t.Errorf("strings with same content have different hash keys")
	}

	if diff1.HashKey() != diff2.HashKey() {
		t.Errorf("strings with same content have different hash keys")
	}

	if hello1.HashKey() == diff1.HashKey() {
		t.Errorf("strings with different content have same hash keys")
	}
}
