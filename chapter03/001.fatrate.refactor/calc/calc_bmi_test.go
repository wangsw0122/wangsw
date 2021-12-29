package calc

import "testing"

func TestCalcBmi(t *testing.T) {
	inputTall,inputWeight:=1.0,30.0
	expectOutput:=30.0
	t.Logf("开始计算,输入tall:%f,weight:%f,期望结果bmi:%f",inputTall,inputWeight,expectOutput)
	actualOutput,err:=CalcBmi(inputWeight,inputTall)
	if err!=nil{
		t.Fatalf("expecting yes,but got:%v",err)    //fatal  直接结束运行
	}
	if expectOutput!=actualOutput{
		t.Errorf("expecting %f,but got %f",expectOutput,actualOutput)
	}

}
