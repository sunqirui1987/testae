package ae

type KeyFrame struct {
	from interface{}
	to   interface{}
	fun  func(val interface{}) interface{} //曲线函数
}

type KeyFrameManger struct {
	FrameKeyList []KeyFrame
}
