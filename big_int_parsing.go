// https://stackoverflow.com/questions/55461131/there-was-an-error-math-big-cannot-unmarshal-into-a-big-int

package main

import (
	"encoding/json"
	"fmt"
	"math/big"
)

type Signature struct {
	R, S BigInt
	V, O uint8 // V is a reconstruction flag and O a multi sig order
}

type BigInt struct {
	big.Int
}

func (i *BigInt) UnmarshalJSON(b []byte) error {
	var val string
	err := json.Unmarshal(b, &val)
	if err != nil {
		return err
	}

	i.SetString(val, 10)

	return nil
}

func main() {
	string := []byte(`{"O":0,"R":"82794247871852158897004947856472973914188862150580220767211643167985440428259","S":"39475619887140601172207943363731402979187092853596849493781395367115389948109","V":0}`)

	var sig Signature

	err2 := json.Unmarshal([]byte(string), &sig)
	if err2 != nil {
		fmt.Println("There was an error:", err2)
	}

	fmt.Printf("r %s s %s o %d v %d", sig.R.String(), sig.S.String(), sig.O, sig.V)
}
