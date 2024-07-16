//go:build exclude

package functions

import (
	"agrows-usage-example/web"
	"context"
	"math"
	"strconv"
	"strings"
)

func SayHello(name string) string {
	buf := strings.Builder{}
	web.HelloResult(name).Render(context.Background(), &buf)
	return buf.String()
}

type CalcInput struct {
	A int
	B int
}

func CrazyMath(inp CalcInput) string {
	a := inp.A
	b := inp.B
	trigResult := int(math.Sin(float64(a)) * math.Cos(float64(b)) * 1000)

	bitwiseResult := (a ^ b) & (a | b)

	expLogResult := int(math.Exp(float64(a%10)) - math.Log(float64(b%10+1)))

	complexResult := int(real(complex(float64(a), float64(b)) * complex(float64(b), float64(a))))

	result := math.Abs(float64(trigResult + bitwiseResult + expLogResult + complexResult))

	buf := strings.Builder{}
	web.MathResult(strconv.Itoa(int(result))).Render(context.Background(), &buf)

	return buf.String()
}
