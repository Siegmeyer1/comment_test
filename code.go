import (
	"fmt"
)

func main() {
	s := "Hello, World!"

	fmt.Println(s)
	fmt.Println(s)
	fmt.Println(s)

	fmt.Println(`
		a1
		a2
		a3



		b1
		b2
		b3



		c1
		c2
		c3
	`)

	// TODO: remove
	return
}