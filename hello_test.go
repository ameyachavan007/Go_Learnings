package hello

import "testing"

func TestSayHello(t *testing.T) {
  subtests := []struct {
    items []string
    result string
  }{
    {
      result: "Hello, world!",
    },
    {
      items: []string{"Ameya"},
      result: "Hello, Ameya!",  
    },
    {
      items: []string{"Ameya", "Bhavika", "Aryan"},
      result: "Hello, Ameya, Bhavika, Aryan!",  
    },
  }

  for _, st := range subtests {
    if res:= Say(st.items); res != st.result {
      t.Errorf("wanted %s (%v), got %s", st.result, st.items, res)
    }
  }
} 