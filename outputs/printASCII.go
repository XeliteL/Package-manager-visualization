package outputs

import "fmt"

func PrintASCII(deps []string) {
	fmt.Println("Дерево зависимостей (ASCII):")
	for _, d := range deps {
		fmt.Printf("├── %s\n", d)
	}
	fmt.Println("└── (конец списка)")
}
