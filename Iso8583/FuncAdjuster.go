package Iso8583

type FuncAdjuster struct  {
	getFunc func (string) string
	setFunc func (string) string
}

func NewFuncAdjuster(gFunc func(string) string, sFunc func(string) string) *FuncAdjuster {
	return &FuncAdjuster{getFunc:gFunc, setFunc:sFunc}
}

func (fa *FuncAdjuster) Get(value string) string {
	if fa.getFunc != nil {
		return fa.getFunc(value)
	}

	return value
}

func (fa *FuncAdjuster) Set(value string) string {
	if fa.setFunc != nil {
		return fa.setFunc(value)
	}

	return value
}