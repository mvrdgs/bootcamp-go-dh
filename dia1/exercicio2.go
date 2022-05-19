package dia1

import "fmt"

func Clima(temp, press float32, umid int) {
	fmt.Printf("Temperatura: %.1f°C\nUmidade: %d%%\nPressão atmosférica: %.2fmb\n", temp, umid, press)
}
