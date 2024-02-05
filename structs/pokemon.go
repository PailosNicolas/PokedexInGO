package structs

import (
	"bufio"
	"fmt"
	"os"
)

type Pokemon struct {
	NickName string
	Name     string
	Height   int
	Weight   int
	Stats    struct {
		Hp         int
		Attack     int
		Defense    int
		SpecialAtk int
		SpecialDef int
		Speed      int
	}
	Types   []string
	IsShiny bool
}

func (p Pokemon) GetInfo() {
	fmt.Printf("Name: %s", p.NickName)
	if p.IsShiny {
		fmt.Print(" (Shiny)")
	}
	print("\n")
	if p.NickName != p.Name {
		fmt.Printf("Species: %s\n", p.Name)
	}
	fmt.Printf("Height: %d\n", p.Height)
	fmt.Printf("Weight: %d\n", p.Weight)
	fmt.Println("Stats:")
	fmt.Printf("  -Hp: %d\n", p.Stats.Hp)
	fmt.Printf("  -Attack: %d\n", p.Stats.Attack)
	fmt.Printf("  -Defense: %d\n", p.Stats.Defense)
	fmt.Printf("  -SpecialAtk: %d\n", p.Stats.SpecialAtk)
	fmt.Printf("  -SpecialDef: %d\n", p.Stats.SpecialDef)
	fmt.Printf("  -Speed: %d\n", p.Stats.Speed)
	println("Types:")
	for _, pokemon_type := range p.Types {
		fmt.Printf("  - %s\n", pokemon_type)
	}
}

func (p *Pokemon) SetNickname() error {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("Do you want to give %s a nickname? y/n\n", p.Name)
		scanner.Scan()
		answer := scanner.Text()

		switch answer {
		case "y":
			fmt.Printf("%s nickname: ", p.Name)
			scanner.Scan()
			p.NickName = scanner.Text()
			return nil

		case "n":
			p.NickName = p.Name
			return nil
		default:
			println("option not recognized")
		}
	}
}
