package farm

import "testing"

func TestCow_NumberOfLegs(t *testing.T) {
	cow := Cow{Name: "Betsy"}
	expected := 4

	actual := cow.NumberOfLegs()

	if actual != expected {
		t.Errorf("Expected %d legs for Cow, but got %d", expected, actual)
	}
}

func TestChicken_Speak(t *testing.T) {
	chicken := Chicken{Name: "Clucky"}
	expected := "Cluck!"

	actual := chicken.Speak()

	if actual != expected {
		t.Errorf("Expected Chicken to say '%s', but it said '%s'", expected, actual)
	}
}

func TestFarmer_FeedAnimal(t *testing.T) {
	farmer := Farmer{Name: "Dooby"}
	animal := Cow{Name: "Moon"}

	str := farmer.FeedAnimal(&animal)

	expected := "Dooby is feeding animal with 4 legs"

	if str != expected {
		t.Errorf("Expected Chicken to say '%s', but it said '%s'", expected, str)
	}
}
