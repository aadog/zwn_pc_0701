
//----------------------------------------
// The code is automatically generated by the GenlibLcl tool.
// Copyright © ying32. All Rights Reserved.
// 
// Licensed under Apache License 2.0
//
//----------------------------------------


package vcl


import (
    . "github.com/ying32/govcl/vcl/api"
    . "github.com/ying32/govcl/vcl/types"
    "unsafe"
)

type TBevel struct {
    IControl
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与LCL没有太多关系。
    ptr unsafe.Pointer
}

// 创建一个新的对象。
// 
// Create a new object.
func NewBevel(owner IComponent) *TBevel {
    b := new(TBevel)
    b.instance = Bevel_Create(CheckPtr(owner))
    b.ptr = unsafe.Pointer(b.instance)
    return b
}

// 动态转换一个已存在的对象实例。
// 
// Dynamically convert an existing object instance.
func AsBevel(obj interface{}) *TBevel {
    instance, ptr := getInstance(obj)
    if instance == 0 { return nil }
    return &TBevel{instance: instance, ptr: ptr}
}

// -------------------------- Deprecated begin --------------------------
// 新建一个对象来自已经存在的对象实例指针。
// 
// Create a new object from an existing object instance pointer.
// Deprecated: use AsBevel.
func BevelFromInst(inst uintptr) *TBevel {
    return AsBevel(inst)
}

// 新建一个对象来自已经存在的对象实例。
// 
// Create a new object from an existing object instance.
// Deprecated: use AsBevel.
func BevelFromObj(obj IObject) *TBevel {
    return AsBevel(obj)
}

// 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// 
// Create a new object from an unsecured address. Note: Using this function may cause some unclear situations and be used with caution..
// Deprecated: use AsBevel.
func BevelFromUnsafePointer(ptr unsafe.Pointer) *TBevel {
    return AsBevel(ptr)
}

// -------------------------- Deprecated end --------------------------
// 释放对象。
// 
// Free object.
func (b *TBevel) Free() {
    if b.instance != 0 {
        Bevel_Free(b.instance)
        b.instance, b.ptr = 0, nullptr
    }
}

// 返回对象实例指针。
// 
// Return object instance pointer.
func (b *TBevel) Instance() uintptr {
    return b.instance
}

// 获取一个不安全的地址。
// 
// Get an unsafe address.
func (b *TBevel) UnsafeAddr() unsafe.Pointer {
    return b.ptr
}

// 检测地址是否为空。
// 
// Check if the address is empty.
func (b *TBevel) IsValid() bool {
    return b.instance != 0
}

// 检测当前对象是否继承自目标对象。
// 
// Checks whether the current object is inherited from the target object.
func (b *TBevel) Is() TIs {
    return TIs(b.instance)
}

// 动态转换当前对象为目标对象。
// 
// Dynamically convert the current object to the target object.
//func (b *TBevel) As() TAs {
//    return TAs(b.instance)
//}

// 获取类信息指针。
// 
// Get class information pointer.
func TBevelClass() TClass {
    return Bevel_StaticClassType()
}

// 将控件置于最前。
//
// Bring the control to the front.
func (b *TBevel) BringToFront() {
    Bevel_BringToFront(b.instance)
}

// 将客户端坐标转为绝对的屏幕坐标。
//
// Convert client coordinates to absolute screen coordinates.
func (b *TBevel) ClientToScreen(Point TPoint) TPoint {
    return Bevel_ClientToScreen(b.instance, Point)
}

// 将客户端坐标转为父容器坐标。
//
// Convert client coordinates to parent container coordinates.
func (b *TBevel) ClientToParent(Point TPoint, AParent IWinControl) TPoint {
    return Bevel_ClientToParent(b.instance, Point , CheckPtr(AParent))
}

// 是否在拖拽中。
//
// Is it in the middle of dragging.
func (b *TBevel) Dragging() bool {
    return Bevel_Dragging(b.instance)
}

// 是否有父容器。
//
// Is there a parent container.
func (b *TBevel) HasParent() bool {
    return Bevel_HasParent(b.instance)
}

// 隐藏控件。
//
// Hidden control.
func (b *TBevel) Hide() {
    Bevel_Hide(b.instance)
}

// 要求重绘。
//
// Redraw.
func (b *TBevel) Invalidate() {
    Bevel_Invalidate(b.instance)
}

// 发送一个消息。
//
// Send a message.
func (b *TBevel) Perform(Msg uint32, WParam uintptr, LParam int) int {
    return Bevel_Perform(b.instance, Msg , WParam , LParam)
}

// 刷新控件。
//
// Refresh control.
func (b *TBevel) Refresh() {
    Bevel_Refresh(b.instance)
}

// 重绘。
//
// Repaint.
func (b *TBevel) Repaint() {
    Bevel_Repaint(b.instance)
}

// 将屏幕坐标转为客户端坐标。
//
// Convert screen coordinates to client coordinates.
func (b *TBevel) ScreenToClient(Point TPoint) TPoint {
    return Bevel_ScreenToClient(b.instance, Point)
}

// 将父容器坐标转为客户端坐标。
//
// Convert parent container coordinates to client coordinates.
func (b *TBevel) ParentToClient(Point TPoint, AParent IWinControl) TPoint {
    return Bevel_ParentToClient(b.instance, Point , CheckPtr(AParent))
}

// 控件至于最后面。
//
// The control is placed at the end.
func (b *TBevel) SendToBack() {
    Bevel_SendToBack(b.instance)
}

// 设置组件边界。
//
// Set component boundaries.
func (b *TBevel) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    Bevel_SetBounds(b.instance, ALeft , ATop , AWidth , AHeight)
}

// 显示控件。
//
// Show control.
func (b *TBevel) Show() {
    Bevel_Show(b.instance)
}

// 控件更新。
//
// Update.
func (b *TBevel) Update() {
    Bevel_Update(b.instance)
}

// 获取控件的字符，如果有。
//
// Get the characters of the control, if any.
func (b *TBevel) GetTextBuf(Buffer *string, BufSize int32) int32 {
    return Bevel_GetTextBuf(b.instance, Buffer , BufSize)
}

// 获取控件的字符长，如果有。
//
// Get the character length of the control, if any.
func (b *TBevel) GetTextLen() int32 {
    return Bevel_GetTextLen(b.instance)
}

// 设置控件字符，如果有。
//
// Set control characters, if any.
func (b *TBevel) SetTextBuf(Buffer string) {
    Bevel_SetTextBuf(b.instance, Buffer)
}

// 查找指定名称的组件。
//
// Find the component with the specified name.
func (b *TBevel) FindComponent(AName string) *TComponent {
    return AsComponent(Bevel_FindComponent(b.instance, AName))
}

// 获取类名路径。
//
// Get the class name path.
func (b *TBevel) GetNamePath() string {
    return Bevel_GetNamePath(b.instance)
}

// 复制一个对象，如果对象实现了此方法的话。
//
// Copy an object, if the object implements this method.
func (b *TBevel) Assign(Source IObject) {
    Bevel_Assign(b.instance, CheckPtr(Source))
}

// 获取类的类型信息。
//
// Get class type information.
func (b *TBevel) ClassType() TClass {
    return Bevel_ClassType(b.instance)
}

// 获取当前对象类名称。
//
// Get the current object class name.
func (b *TBevel) ClassName() string {
    return Bevel_ClassName(b.instance)
}

// 获取当前对象实例大小。
//
// Get the current object instance size.
func (b *TBevel) InstanceSize() int32 {
    return Bevel_InstanceSize(b.instance)
}

// 判断当前类是否继承自指定类。
//
// Determine whether the current class inherits from the specified class.
func (b *TBevel) InheritsFrom(AClass TClass) bool {
    return Bevel_InheritsFrom(b.instance, AClass)
}

// 与一个对象进行比较。
//
// Compare with an object.
func (b *TBevel) Equals(Obj IObject) bool {
    return Bevel_Equals(b.instance, CheckPtr(Obj))
}

// 获取类的哈希值。
//
// Get the hash value of the class.
func (b *TBevel) GetHashCode() int32 {
    return Bevel_GetHashCode(b.instance)
}

// 文本类信息。
//
// Text information.
func (b *TBevel) ToString() string {
    return Bevel_ToString(b.instance)
}

func (b *TBevel) AnchorToNeighbour(ASide TAnchorKind, ASpace int32, ASibling IControl) {
    Bevel_AnchorToNeighbour(b.instance, ASide , ASpace , CheckPtr(ASibling))
}

func (b *TBevel) AnchorParallel(ASide TAnchorKind, ASpace int32, ASibling IControl) {
    Bevel_AnchorParallel(b.instance, ASide , ASpace , CheckPtr(ASibling))
}

// 置于指定控件的横向中心。
func (b *TBevel) AnchorHorizontalCenterTo(ASibling IControl) {
    Bevel_AnchorHorizontalCenterTo(b.instance, CheckPtr(ASibling))
}

// 置于指定控件的纵向中心。
func (b *TBevel) AnchorVerticalCenterTo(ASibling IControl) {
    Bevel_AnchorVerticalCenterTo(b.instance, CheckPtr(ASibling))
}

func (b *TBevel) AnchorSame(ASide TAnchorKind, ASibling IControl) {
    Bevel_AnchorSame(b.instance, ASide , CheckPtr(ASibling))
}

func (b *TBevel) AnchorAsAlign(ATheAlign TAlign, ASpace int32) {
    Bevel_AnchorAsAlign(b.instance, ATheAlign , ASpace)
}

func (b *TBevel) AnchorClient(ASpace int32) {
    Bevel_AnchorClient(b.instance, ASpace)
}

func (b *TBevel) ScaleDesignToForm(ASize int32) int32 {
    return Bevel_ScaleDesignToForm(b.instance, ASize)
}

func (b *TBevel) ScaleFormToDesign(ASize int32) int32 {
    return Bevel_ScaleFormToDesign(b.instance, ASize)
}

func (b *TBevel) Scale96ToForm(ASize int32) int32 {
    return Bevel_Scale96ToForm(b.instance, ASize)
}

func (b *TBevel) ScaleFormTo96(ASize int32) int32 {
    return Bevel_ScaleFormTo96(b.instance, ASize)
}

func (b *TBevel) Scale96ToFont(ASize int32) int32 {
    return Bevel_Scale96ToFont(b.instance, ASize)
}

func (b *TBevel) ScaleFontTo96(ASize int32) int32 {
    return Bevel_ScaleFontTo96(b.instance, ASize)
}

func (b *TBevel) ScaleScreenToFont(ASize int32) int32 {
    return Bevel_ScaleScreenToFont(b.instance, ASize)
}

func (b *TBevel) ScaleFontToScreen(ASize int32) int32 {
    return Bevel_ScaleFontToScreen(b.instance, ASize)
}

func (b *TBevel) Scale96ToScreen(ASize int32) int32 {
    return Bevel_Scale96ToScreen(b.instance, ASize)
}

func (b *TBevel) ScaleScreenTo96(ASize int32) int32 {
    return Bevel_ScaleScreenTo96(b.instance, ASize)
}

func (b *TBevel) AutoAdjustLayout(AMode TLayoutAdjustmentPolicy, AFromPPI int32, AToPPI int32, AOldFormWidth int32, ANewFormWidth int32) {
    Bevel_AutoAdjustLayout(b.instance, AMode , AFromPPI , AToPPI , AOldFormWidth , ANewFormWidth)
}

func (b *TBevel) FixDesignFontsPPI(ADesignTimePPI int32) {
    Bevel_FixDesignFontsPPI(b.instance, ADesignTimePPI)
}

func (b *TBevel) ScaleFontsPPI(AToPPI int32, AProportion float64) {
    Bevel_ScaleFontsPPI(b.instance, AToPPI , AProportion)
}

// 获取控件自动调整。
//
// Get Control automatically adjusts.
func (b *TBevel) Align() TAlign {
    return Bevel_GetAlign(b.instance)
}

// 设置控件自动调整。
//
// Set Control automatically adjusts.
func (b *TBevel) SetAlign(value TAlign) {
    Bevel_SetAlign(b.instance, value)
}

// 获取四个角位置的锚点。
func (b *TBevel) Anchors() TAnchors {
    return Bevel_GetAnchors(b.instance)
}

// 设置四个角位置的锚点。
func (b *TBevel) SetAnchors(value TAnchors) {
    Bevel_SetAnchors(b.instance, value)
}

// 获取约束控件大小。
func (b *TBevel) Constraints() *TSizeConstraints {
    return AsSizeConstraints(Bevel_GetConstraints(b.instance))
}

// 设置约束控件大小。
func (b *TBevel) SetConstraints(value *TSizeConstraints) {
    Bevel_SetConstraints(b.instance, CheckPtr(value))
}

// 获取以父容器的ShowHint属性为准。
func (b *TBevel) ParentShowHint() bool {
    return Bevel_GetParentShowHint(b.instance)
}

// 设置以父容器的ShowHint属性为准。
func (b *TBevel) SetParentShowHint(value bool) {
    Bevel_SetParentShowHint(b.instance, value)
}

func (b *TBevel) Shape() TBevelShape {
    return Bevel_GetShape(b.instance)
}

func (b *TBevel) SetShape(value TBevelShape) {
    Bevel_SetShape(b.instance, value)
}

// 获取显示鼠标悬停提示。
//
// Get Show mouseover tips.
func (b *TBevel) ShowHint() bool {
    return Bevel_GetShowHint(b.instance)
}

// 设置显示鼠标悬停提示。
//
// Set Show mouseover tips.
func (b *TBevel) SetShowHint(value bool) {
    Bevel_SetShowHint(b.instance, value)
}

func (b *TBevel) Style() TBevelStyle {
    return Bevel_GetStyle(b.instance)
}

func (b *TBevel) SetStyle(value TBevelStyle) {
    Bevel_SetStyle(b.instance, value)
}

// 获取控件可视。
//
// Get the control visible.
func (b *TBevel) Visible() bool {
    return Bevel_GetVisible(b.instance)
}

// 设置控件可视。
//
// Set the control visible.
func (b *TBevel) SetVisible(value bool) {
    Bevel_SetVisible(b.instance, value)
}

// 获取控件启用。
//
// Get the control enabled.
func (b *TBevel) Enabled() bool {
    return Bevel_GetEnabled(b.instance)
}

// 设置控件启用。
//
// Set the control enabled.
func (b *TBevel) SetEnabled(value bool) {
    Bevel_SetEnabled(b.instance, value)
}

func (b *TBevel) Action() *TAction {
    return AsAction(Bevel_GetAction(b.instance))
}

func (b *TBevel) SetAction(value IComponent) {
    Bevel_SetAction(b.instance, CheckPtr(value))
}

func (b *TBevel) BiDiMode() TBiDiMode {
    return Bevel_GetBiDiMode(b.instance)
}

func (b *TBevel) SetBiDiMode(value TBiDiMode) {
    Bevel_SetBiDiMode(b.instance, value)
}

func (b *TBevel) BoundsRect() TRect {
    return Bevel_GetBoundsRect(b.instance)
}

func (b *TBevel) SetBoundsRect(value TRect) {
    Bevel_SetBoundsRect(b.instance, value)
}

// 获取客户区高度。
//
// Get client height.
func (b *TBevel) ClientHeight() int32 {
    return Bevel_GetClientHeight(b.instance)
}

// 设置客户区高度。
//
// Set client height.
func (b *TBevel) SetClientHeight(value int32) {
    Bevel_SetClientHeight(b.instance, value)
}

func (b *TBevel) ClientOrigin() TPoint {
    return Bevel_GetClientOrigin(b.instance)
}

// 获取客户区矩形。
//
// Get client rectangle.
func (b *TBevel) ClientRect() TRect {
    return Bevel_GetClientRect(b.instance)
}

// 获取客户区宽度。
//
// Get client width.
func (b *TBevel) ClientWidth() int32 {
    return Bevel_GetClientWidth(b.instance)
}

// 设置客户区宽度。
//
// Set client width.
func (b *TBevel) SetClientWidth(value int32) {
    Bevel_SetClientWidth(b.instance, value)
}

// 获取控件状态。
//
// Get control state.
func (b *TBevel) ControlState() TControlState {
    return Bevel_GetControlState(b.instance)
}

// 设置控件状态。
//
// Set control state.
func (b *TBevel) SetControlState(value TControlState) {
    Bevel_SetControlState(b.instance, value)
}

// 获取控件样式。
//
// Get control style.
func (b *TBevel) ControlStyle() TControlStyle {
    return Bevel_GetControlStyle(b.instance)
}

// 设置控件样式。
//
// Set control style.
func (b *TBevel) SetControlStyle(value TControlStyle) {
    Bevel_SetControlStyle(b.instance, value)
}

func (b *TBevel) Floating() bool {
    return Bevel_GetFloating(b.instance)
}

// 获取控件父容器。
//
// Get control parent container.
func (b *TBevel) Parent() *TWinControl {
    return AsWinControl(Bevel_GetParent(b.instance))
}

// 设置控件父容器。
//
// Set control parent container.
func (b *TBevel) SetParent(value IWinControl) {
    Bevel_SetParent(b.instance, CheckPtr(value))
}

// 获取左边位置。
//
// Get Left position.
func (b *TBevel) Left() int32 {
    return Bevel_GetLeft(b.instance)
}

// 设置左边位置。
//
// Set Left position.
func (b *TBevel) SetLeft(value int32) {
    Bevel_SetLeft(b.instance, value)
}

// 获取顶边位置。
//
// Get Top position.
func (b *TBevel) Top() int32 {
    return Bevel_GetTop(b.instance)
}

// 设置顶边位置。
//
// Set Top position.
func (b *TBevel) SetTop(value int32) {
    Bevel_SetTop(b.instance, value)
}

// 获取宽度。
//
// Get width.
func (b *TBevel) Width() int32 {
    return Bevel_GetWidth(b.instance)
}

// 设置宽度。
//
// Set width.
func (b *TBevel) SetWidth(value int32) {
    Bevel_SetWidth(b.instance, value)
}

// 获取高度。
//
// Get height.
func (b *TBevel) Height() int32 {
    return Bevel_GetHeight(b.instance)
}

// 设置高度。
//
// Set height.
func (b *TBevel) SetHeight(value int32) {
    Bevel_SetHeight(b.instance, value)
}

// 获取控件光标。
//
// Get control cursor.
func (b *TBevel) Cursor() TCursor {
    return Bevel_GetCursor(b.instance)
}

// 设置控件光标。
//
// Set control cursor.
func (b *TBevel) SetCursor(value TCursor) {
    Bevel_SetCursor(b.instance, value)
}

// 获取组件鼠标悬停提示。
//
// Get component mouse hints.
func (b *TBevel) Hint() string {
    return Bevel_GetHint(b.instance)
}

// 设置组件鼠标悬停提示。
//
// Set component mouse hints.
func (b *TBevel) SetHint(value string) {
    Bevel_SetHint(b.instance, value)
}

// 获取组件总数。
//
// Get the total number of components.
func (b *TBevel) ComponentCount() int32 {
    return Bevel_GetComponentCount(b.instance)
}

// 获取组件索引。
//
// Get component index.
func (b *TBevel) ComponentIndex() int32 {
    return Bevel_GetComponentIndex(b.instance)
}

// 设置组件索引。
//
// Set component index.
func (b *TBevel) SetComponentIndex(value int32) {
    Bevel_SetComponentIndex(b.instance, value)
}

// 获取组件所有者。
//
// Get component owner.
func (b *TBevel) Owner() *TComponent {
    return AsComponent(Bevel_GetOwner(b.instance))
}

// 获取组件名称。
//
// Get the component name.
func (b *TBevel) Name() string {
    return Bevel_GetName(b.instance)
}

// 设置组件名称。
//
// Set the component name.
func (b *TBevel) SetName(value string) {
    Bevel_SetName(b.instance, value)
}

// 获取对象标记。
//
// Get the control tag.
func (b *TBevel) Tag() int {
    return Bevel_GetTag(b.instance)
}

// 设置对象标记。
//
// Set the control tag.
func (b *TBevel) SetTag(value int) {
    Bevel_SetTag(b.instance, value)
}

// 获取左边锚点。
func (b *TBevel) AnchorSideLeft() *TAnchorSide {
    return AsAnchorSide(Bevel_GetAnchorSideLeft(b.instance))
}

// 设置左边锚点。
func (b *TBevel) SetAnchorSideLeft(value *TAnchorSide) {
    Bevel_SetAnchorSideLeft(b.instance, CheckPtr(value))
}

// 获取顶边锚点。
func (b *TBevel) AnchorSideTop() *TAnchorSide {
    return AsAnchorSide(Bevel_GetAnchorSideTop(b.instance))
}

// 设置顶边锚点。
func (b *TBevel) SetAnchorSideTop(value *TAnchorSide) {
    Bevel_SetAnchorSideTop(b.instance, CheckPtr(value))
}

// 获取右边锚点。
func (b *TBevel) AnchorSideRight() *TAnchorSide {
    return AsAnchorSide(Bevel_GetAnchorSideRight(b.instance))
}

// 设置右边锚点。
func (b *TBevel) SetAnchorSideRight(value *TAnchorSide) {
    Bevel_SetAnchorSideRight(b.instance, CheckPtr(value))
}

// 获取底边锚点。
func (b *TBevel) AnchorSideBottom() *TAnchorSide {
    return AsAnchorSide(Bevel_GetAnchorSideBottom(b.instance))
}

// 设置底边锚点。
func (b *TBevel) SetAnchorSideBottom(value *TAnchorSide) {
    Bevel_SetAnchorSideBottom(b.instance, CheckPtr(value))
}

// 获取边框间距。
func (b *TBevel) BorderSpacing() *TControlBorderSpacing {
    return AsControlBorderSpacing(Bevel_GetBorderSpacing(b.instance))
}

// 设置边框间距。
func (b *TBevel) SetBorderSpacing(value *TControlBorderSpacing) {
    Bevel_SetBorderSpacing(b.instance, CheckPtr(value))
}

// 获取指定索引组件。
//
// Get the specified index component.
func (b *TBevel) Components(AIndex int32) *TComponent {
    return AsComponent(Bevel_GetComponents(b.instance, AIndex))
}

// 获取锚侧面。
func (b *TBevel) AnchorSide(AKind TAnchorKind) *TAnchorSide {
    return AsAnchorSide(Bevel_GetAnchorSide(b.instance, AKind))
}

