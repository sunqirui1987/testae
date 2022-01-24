package ae

import (
	"github.com/goplus/ae/math32"
)

//材质器
type Material struct {
}

//灯光
type Light struct {
}

//网格
type Geometry struct{}

//矩形网格
type RectGeometry struct{}

type Mesh struct {
	Material Material
	Light    Light
	Uv       Geometry
}

//舞台元素对象
type Ele struct {
	//关键帧系统
	KeyFrameManger

	//位置
	Pos math32.Vector3
	//旋转
	Rotate math32.Vector3
	//缩放
	Scale math32.Vector3
	//旋转中心
	RotateCenter math32.Vector3

	//透明度
	Alpha float64

	//融合模式
	blendType int

	//子元素列表
	Eles []Ele

	//网格模型
	Mesh Mesh

	//特效合成器
	EffectComposer EffectComposer
}

//组
type GroupEle struct {
	Ele
}

//添加元素
func (e *GroupEle) Add(ele Ele) {
}

//删除元素
func (e *GroupEle) Remove(ele Ele) {
}

//钢笔
type PenEle struct {
	Ele
}

//文字
type FontEle struct {
	Ele
}

//图片视频
type ImageEle struct {
	Ele
}

//相机
type CameraEle struct {
	Ele
}

//切换全局相机还是用户相机 0 全局相机 1用户相机
func (*CameraEle) ChangeCameraType(modeltype int) {}

// * model  0: 透视相机 1：正交相机
//改变相机类型
func (*CameraEle) SetCameraModel(model int) {}

//相机视图 0正视图 1顶视图 2左视图
func (*CameraEle) SetCameraView(viewtype int) {}

//重置相机
func (*CameraEle) ResetCamera() {}

//粒子
type ParicleELe struct {
}
