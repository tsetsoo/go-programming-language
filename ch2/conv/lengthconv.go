package conv

import "fmt"

type Feet float64
type Meter float64

func (m Meter) String() string { return fmt.Sprintf("%gm", m) }
func (f Feet) String() string { return fmt.Sprintf("%gf", f)}
