package interactions

import "fmt"

func PrintGreeting() {
  fmt.Println("MONSTER SLAYER")
  fmt.Println("Starting a new game...")
  fmt.Println("Good luck!")
}

func ShowAvailableActions(specialAttackIsAvailable bool) {
  fmt.Println("Choose your course of action:")
  fmt.Println("-----------------------------")
  fmt.Println("(1) Smite your foe.")
  fmt.Println("(2) Imbibe a healing elixir and heal thyself.")
  if specialAttackIsAvailable {
    fmt.Println("(3) Unleash a devastating special attack.")
  }
}

