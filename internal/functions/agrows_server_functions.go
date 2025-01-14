/*
	Code generated by agrows. DO NOT EDIT :)
	This code was generated on 2024-07-16 at 22:11:54
	Any changes made to this file will be lost
*/
package functions

import (
	"agrows-usage-example/web"
	"context"
	"errors"
	"fmt"
	protocol "github.com/codeupdateandmodificationsystem/protocol"
	"math"
	"reflect"
	"strconv"
	"strings"
)

func agrows_SayHello(name string) string {
	buf := strings.Builder{}
	web.HelloResult(name).Render(context.Background(), &buf)
	return buf.String()
}

type CalcInput struct {
	A int
	B int
}

func agrows_CrazyMath(inp CalcInput) string {
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

func AgrowsReceive(data []byte) (string, error) {
	functionName, args, err := protocol.DecodeFunctionCall(data, protocol.Options())
	if err != nil {
		return "", err
	}
	switch functionName {

	case "SayHello":
		var (
			nameParam    string
			nameParamArg protocol.Argument
			ok           bool
		)
		if nameParamArg, ok = args["name"]; !ok {
			return "", errors.New(fmt.Sprintf("parameter %s is not in the received arguments", "name"))
		}
		if nameParamArg, ok = args["name"]; !ok {
			return "", errors.New(fmt.Sprintf("parameter %s is not in the received arguments", "name"))
		}
		if nameParam, ok = nameParamArg.Value.(string); !ok {
			return "", errors.New(fmt.Sprintf("failed to cast parameter '%s' to '%s'", nameParam, "string"))
		}
		str0 := agrows_SayHello(nameParam)
		return str0, nil

	case "CrazyMath":
		var (
			inpParam                CalcInput
			inpParamArg             protocol.Argument
			inpParamValue, inpValue reflect.Value
			ok                      bool
		)
		if inpParamArg, ok = args["inp"]; !ok {
			return "", errors.New(fmt.Sprintf("parameter %s is not in the received arguments", "inp"))
		}
		inpParamValue = reflect.ValueOf(inpParamArg.Value)
		inpValue = reflect.ValueOf(&inpParam).Elem()
		for i := 0; i < inpValue.NumField(); i++ {
			key := inpValue.Type().Field(i).Name
			paramValueField := inpParamValue.FieldByName(key)
			fieldValue := inpValue.Field(i)
			if fieldValue.CanSet() {
				inpValue.Field(i).Set(paramValueField)
			}
		}
		str0 := agrows_CrazyMath(inpParam)
		return str0, nil

	default:
		return "", errors.New(fmt.Sprintf("unknown function '%s'", functionName))
	}
}
