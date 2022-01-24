package ae

//特效通道
type EffectPass struct {
	Enable bool
}

func (e *EffectPass) render() {
}
func (e *EffectPass) setSize() {
}

//透明度
//AlphaEffectPass
//发光
//BloomEffectPass
//......

//FBO帧缓冲对象
type renderTarget struct {
}

//特效 渲染
type EffectComposer struct {
	Enable bool
}

//增加特效
func (e *EffectComposer) AddPass(effect EffectPass) {

}

//插入指定位置
func (e *EffectComposer) InsertPass(effect EffectPass, index int) {

}

//删除特效
func (e *EffectComposer) RemovePass(effect EffectPass) {

}

//设置输出纹理的大小
func (e *EffectComposer) SetSize(width, height int) {

}

//绘制
func (e *EffectComposer) Render(deltaTime float64) {

}
