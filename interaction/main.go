package interaction

import "trpc.group/trpc-go/trpc-go"

func main() {
	s := trpc.NewServer()

	if err := s.Serve(); err != nil {
		panic(err)
	}
}
