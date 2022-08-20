package eval

import (
	"fmt"
	"strconv"
	"strings"
)

func (v Var) String() string {
	return string(v)
}

func (l literal) String() string {
	return strconv.FormatFloat(float64(l), 'f', -1, 64)
}

//!-Eval1

//!+Eval2

func (u unary) String() string {
	return string(u.op) + u.x.String()
}

func (b binary) String() string {
	return b.x.String() + " " + string(b.op) + " " + b.y.String()
}

func (c call) String() string {
	toReturn := fmt.Sprintf("%s(", c.fn)
	for _, arg := range c.args {
		toReturn += fmt.Sprintf("%s, ", arg)
	}
	toReturn = strings.TrimSuffix(toReturn, ", ")
	toReturn += ")"
	return toReturn
}

func (u postFixUnary) String() string {
	return u.x.String() + string(u.op)
}
