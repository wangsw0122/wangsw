package calc

import "testing"

func TestCalcFatRate(t *testing.T) {
	inputBmi:=24.691358
	inputAge:=25
	inputName:="小李"
	inputSex:="女"
	expectFatRate:=0.299796296
	t.Logf("开始计算fatRate，输入:bmi:%f,age:%d,sex:%s,期望输出:expextFatRate:%f",inputBmi,inputAge,inputSex,expectFatRate)
	actualOutputFatRate,err:=CalcFatRate(inputBmi,inputAge,inputName,inputSex)
	if err!=nil{
		t.Fatalf("expecting yes,but got:%v",err)    //fatal  直接结束运行
	}
	if expectFatRate!=actualOutputFatRate{
		t.Errorf("expecting %f,but got %f",expectFatRate,actualOutputFatRate)
	}
}
