package ae

import "image"

//glfw窗体
type IWindow interface {
}

//舞台
type Stage struct {
	Res  string //资源文件
	Root Ele    //根节点
}

func InitStage(res string) *Stage {
	return &Stage{}
}

//创建舞台
func (s *Stage) Create(canvas IWindow) {}

//舞台更改宽高
func (s *Stage) Resize(width, height int) {}

//设置质量 0 标准 1 极速
func (s *Stage) SetQuality(quality int) {}

//获取舞台图像(用来给视频编码的RGBA数据或YUV数据)
func (s *Stage) getImage() image.Image { return nil }

//添加元素
func (s *Stage) AddEle(e *Ele) {}

//删除元素
func (s *Stage) RemoveEle(e *Ele) {}

//隐藏元素
func (s *Stage) HideEle(e *Ele) {}

//显示元素
func (s *Stage) ShowEle(e *Ele) {}

//设置根渲染元素
func (s *Stage) SetRootEle(e *Ele) {}

//舞台Draw
func (s *Stage) Draw() {}
