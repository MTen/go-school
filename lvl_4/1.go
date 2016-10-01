package main

func main() {

}

func functionalJumpingGopher {
  // not really maintainable
    gopher1Name := "Phil"
    gopher1Age := 30
    jump(gopher1Name, gopher1Age)

    gopher2Name := "Noodles"
    gopher2Age := 90
    functionalJump(gopher2Name, gopher2Age)
}
func functionalJump(name string, age int) string {
  if age < 65 {
    return name + "jumps high!"
  } else {
    return name + "can still kinda jump"
  }
}

type gopher struct {
  name string
  age int
}

func structuralJumpingGopher {
  gopher1 := gopher{ name: "Phil", age: 30 }
  gopher2 := gopher{ name: "Noodles", age: 90 }
}

func (g gopher) structuralJump() string {
  if g.age < 65 {
    return g.name + " jumps high!"
  }
  return g.name + " still got ups"
}