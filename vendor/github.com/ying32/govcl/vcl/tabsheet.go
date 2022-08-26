
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

type TTabSheet struct {
    IWinControl
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与LCL没有太多关系。
    ptr unsafe.Pointer
}

// 创建一个新的对象。
// 
// Create a new object.
func NewTabSheet(owner IComponent) *TTabSheet {
    t := new(TTabSheet)
    t.instance = TabSheet_Create(CheckPtr(owner))
    t.ptr = unsafe.Pointer(t.instance)
    return t
}

// 动态转换一个已存在的对象实例。
// 
// Dynamically convert an existing object instance.
func AsTabSheet(obj interface{}) *TTabSheet {
    instance, ptr := getInstance(obj)
    if instance == 0 { return nil }
    return &TTabSheet{instance: instance, ptr: ptr}
}

// -------------------------- Deprecated begin --------------------------
// 新建一个对象来自已经存在的对象实例指针。
// 
// Create a new object from an existing object instance pointer.
// Deprecated: use AsTabSheet.
func TabSheetFromInst(inst uintptr) *TTabSheet {
    return AsTabSheet(inst)
}

// 新建一个对象来自已经存在的对象实例。
// 
// Create a new object from an existing object instance.
// Deprecated: use AsTabSheet.
func TabSheetFromObj(obj IObject) *TTabSheet {
    return AsTabSheet(obj)
}

// 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// 
// Create a new object from an unsecured address. Note: Using this function may cause some unclear situations and be used with caution..
// Deprecated: use AsTabSheet.
func TabSheetFromUnsafePointer(ptr unsafe.Pointer) *TTabSheet {
    return AsTabSheet(ptr)
}

// -------------------------- Deprecated end --------------------------
// 释放对象。
// 
// Free object.
func (t *TTabSheet) Free() {
    if t.instance != 0 {
        TabSheet_Free(t.instance)
        t.instance, t.ptr = 0, nullptr
    }
}

// 返回对象实例指针。
// 
// Return object instance pointer.
func (t *TTabSheet) Instance() uintptr {
    return t.instance
}

// 获取一个不安全的地址。
// 
// Get an unsafe address.
func (t *TTabSheet) UnsafeAddr() unsafe.Pointer {
    return t.ptr
}

// 检测地址是否为空。
// 
// Check if the address is empty.
func (t *TTabSheet) IsValid() bool {
    return t.instance != 0
}

// 检测当前对象是否继承自目标对象。
// 
// Checks whether the current object is inherited from the target object.
func (t *TTabSheet) Is() TIs {
    return TIs(t.instance)
}

// 动态转换当前对象为目标对象。
// 
// Dynamically convert the current object to the target object.
//func (t *TTabSheet) As() TAs {
//    return TAs(t.instance)
//}

// 获取类信息指针。
// 
// Get class information pointer.
func TTabSheetClass() TClass {
    return TabSheet_StaticClassType()
}

// 是否可以获得焦点。
func (t *TTabSheet) CanFocus() bool {
    return TabSheet_CanFocus(t.instance)
}

// 返回是否包含指定控件。
//
// it's contain a specified control.
func (t *TTabSheet) ContainsControl(Control IControl) bool {
    return TabSheet_ContainsControl(t.instance, CheckPtr(Control))
}

// 返回指定坐标及相关属性位置控件。
//
// Returns the specified coordinate and the relevant attribute position control..
func (t *TTabSheet) ControlAtPos(Pos TPoint, AllowDisabled bool, AllowWinControls bool, AllLevels bool) *TControl {
    return AsControl(TabSheet_ControlAtPos(t.instance, Pos , AllowDisabled , AllowWinControls , AllLevels))
}

// 禁用控件的对齐。
//
// Disable control alignment.
func (t *TTabSheet) DisableAlign() {
    TabSheet_DisableAlign(t.instance)
}

// 启用控件对齐。
//
// Enabled control alignment.
func (t *TTabSheet) EnableAlign() {
    TabSheet_EnableAlign(t.instance)
}

// 查找子控件。
//
// Find sub controls.
func (t *TTabSheet) FindChildControl(ControlName string) *TControl {
    return AsControl(TabSheet_FindChildControl(t.instance, ControlName))
}

func (t *TTabSheet) FlipChildren(AllLevels bool) {
    TabSheet_FlipChildren(t.instance, AllLevels)
}

// 返回是否获取焦点。
//
// Return to get focus.
func (t *TTabSheet) Focused() bool {
    return TabSheet_Focused(t.instance)
}

// 句柄是否已经分配。
//
// Is the handle already allocated.
func (t *TTabSheet) HandleAllocated() bool {
    return TabSheet_HandleAllocated(t.instance)
}

// 插入一个控件。
//
// Insert a control.
func (t *TTabSheet) InsertControl(AControl IControl) {
    TabSheet_InsertControl(t.instance, CheckPtr(AControl))
}

// 要求重绘。
//
// Redraw.
func (t *TTabSheet) Invalidate() {
    TabSheet_Invalidate(t.instance)
}

// 绘画至指定DC。
//
// Painting to the specified DC.
func (t *TTabSheet) PaintTo(DC HDC, X int32, Y int32) {
    TabSheet_PaintTo(t.instance, DC , X , Y)
}

// 移除一个控件。
//
// Remove a control.
func (t *TTabSheet) RemoveControl(AControl IControl) {
    TabSheet_RemoveControl(t.instance, CheckPtr(AControl))
}

// 重新对齐。
//
// Realign.
func (t *TTabSheet) Realign() {
    TabSheet_Realign(t.instance)
}

// 重绘。
//
// Repaint.
func (t *TTabSheet) Repaint() {
    TabSheet_Repaint(t.instance)
}

// 按比例缩放。
//
// Scale by.
func (t *TTabSheet) ScaleBy(M int32, D int32) {
    TabSheet_ScaleBy(t.instance, M , D)
}

// 滚动至指定位置。
//
// Scroll by.
func (t *TTabSheet) ScrollBy(DeltaX int32, DeltaY int32) {
    TabSheet_ScrollBy(t.instance, DeltaX , DeltaY)
}

// 设置组件边界。
//
// Set component boundaries.
func (t *TTabSheet) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    TabSheet_SetBounds(t.instance, ALeft , ATop , AWidth , AHeight)
}

// 设置控件焦点。
//
// Set control focus.
func (t *TTabSheet) SetFocus() {
    TabSheet_SetFocus(t.instance)
}

// 控件更新。
//
// Update.
func (t *TTabSheet) Update() {
    TabSheet_Update(t.instance)
}

// 将控件置于最前。
//
// Bring the control to the front.
func (t *TTabSheet) BringToFront() {
    TabSheet_BringToFront(t.instance)
}

// 将客户端坐标转为绝对的屏幕坐标。
//
// Convert client coordinates to absolute screen coordinates.
func (t *TTabSheet) ClientToScreen(Point TPoint) TPoint {
    return TabSheet_ClientToScreen(t.instance, Point)
}

// 将客户端坐标转为父容器坐标。
//
// Convert client coordinates to parent container coordinates.
func (t *TTabSheet) ClientToParent(Point TPoint, AParent IWinControl) TPoint {
    return TabSheet_ClientToParent(t.instance, Point , CheckPtr(AParent))
}

// 是否在拖拽中。
//
// Is it in the middle of dragging.
func (t *TTabSheet) Dragging() bool {
    return TabSheet_Dragging(t.instance)
}

// 是否有父容器。
//
// Is there a parent container.
func (t *TTabSheet) HasParent() bool {
    return TabSheet_HasParent(t.instance)
}

// 隐藏控件。
//
// Hidden control.
func (t *TTabSheet) Hide() {
    TabSheet_Hide(t.instance)
}

// 发送一个消息。
//
// Send a message.
func (t *TTabSheet) Perform(Msg uint32, WParam uintptr, LParam int) int {
    return TabSheet_Perform(t.instance, Msg , WParam , LParam)
}

// 刷新控件。
//
// Refresh control.
func (t *TTabSheet) Refresh() {
    TabSheet_Refresh(t.instance)
}

// 将屏幕坐标转为客户端坐标。
//
// Convert screen coordinates to client coordinates.
func (t *TTabSheet) ScreenToClient(Point TPoint) TPoint {
    return TabSheet_ScreenToClient(t.instance, Point)
}

// 将父容器坐标转为客户端坐标。
//
// Convert parent container coordinates to client coordinates.
func (t *TTabSheet) ParentToClient(Point TPoint, AParent IWinControl) TPoint {
    return TabSheet_ParentToClient(t.instance, Point , CheckPtr(AParent))
}

// 控件至于最后面。
//
// The control is placed at the end.
func (t *TTabSheet) SendToBack() {
    TabSheet_SendToBack(t.instance)
}

// 显示控件。
//
// Show control.
func (t *TTabSheet) Show() {
    TabSheet_Show(t.instance)
}

// 获取控件的字符，如果有。
//
// Get the characters of the control, if any.
func (t *TTabSheet) GetTextBuf(Buffer *string, BufSize int32) int32 {
    return TabSheet_GetTextBuf(t.instance, Buffer , BufSize)
}

// 获取控件的字符长，如果有。
//
// Get the character length of the control, if any.
func (t *TTabSheet) GetTextLen() int32 {
    return TabSheet_GetTextLen(t.instance)
}

// 设置控件字符，如果有。
//
// Set control characters, if any.
func (t *TTabSheet) SetTextBuf(Buffer string) {
    TabSheet_SetTextBuf(t.instance, Buffer)
}

// 查找指定名称的组件。
//
// Find the component with the specified name.
func (t *TTabSheet) FindComponent(AName string) *TComponent {
    return AsComponent(TabSheet_FindComponent(t.instance, AName))
}

// 获取类名路径。
//
// Get the class name path.
func (t *TTabSheet) GetNamePath() string {
    return TabSheet_GetNamePath(t.instance)
}

// 复制一个对象，如果对象实现了此方法的话。
//
// Copy an object, if the object implements this method.
func (t *TTabSheet) Assign(Source IObject) {
    TabSheet_Assign(t.instance, CheckPtr(Source))
}

// 获取类的类型信息。
//
// Get class type information.
func (t *TTabSheet) ClassType() TClass {
    return TabSheet_ClassType(t.instance)
}

// 获取当前对象类名称。
//
// Get the current object class name.
func (t *TTabSheet) ClassName() string {
    return TabSheet_ClassName(t.instance)
}

// 获取当前对象实例大小。
//
// Get the current object instance size.
func (t *TTabSheet) InstanceSize() int32 {
    return TabSheet_InstanceSize(t.instance)
}

// 判断当前类是否继承自指定类。
//
// Determine whether the current class inherits from the specified class.
func (t *TTabSheet) InheritsFrom(AClass TClass) bool {
    return TabSheet_InheritsFrom(t.instance, AClass)
}

// 与一个对象进行比较。
//
// Compare with an object.
func (t *TTabSheet) Equals(Obj IObject) bool {
    return TabSheet_Equals(t.instance, CheckPtr(Obj))
}

// 获取类的哈希值。
//
// Get the hash value of the class.
func (t *TTabSheet) GetHashCode() int32 {
    return TabSheet_GetHashCode(t.instance)
}

// 文本类信息。
//
// Text information.
func (t *TTabSheet) ToString() string {
    return TabSheet_ToString(t.instance)
}

func (t *TTabSheet) AnchorToNeighbour(ASide TAnchorKind, ASpace int32, ASibling IControl) {
    TabSheet_AnchorToNeighbour(t.instance, ASide , ASpace , CheckPtr(ASibling))
}

func (t *TTabSheet) AnchorParallel(ASide TAnchorKind, ASpace int32, ASibling IControl) {
    TabSheet_AnchorParallel(t.instance, ASide , ASpace , CheckPtr(ASibling))
}

// 置于指定控件的横向中心。
func (t *TTabSheet) AnchorHorizontalCenterTo(ASibling IControl) {
    TabSheet_AnchorHorizontalCenterTo(t.instance, CheckPtr(ASibling))
}

// 置于指定控件的纵向中心。
func (t *TTabSheet) AnchorVerticalCenterTo(ASibling IControl) {
    TabSheet_AnchorVerticalCenterTo(t.instance, CheckPtr(ASibling))
}

func (t *TTabSheet) AnchorSame(ASide TAnchorKind, ASibling IControl) {
    TabSheet_AnchorSame(t.instance, ASide , CheckPtr(ASibling))
}

func (t *TTabSheet) AnchorAsAlign(ATheAlign TAlign, ASpace int32) {
    TabSheet_AnchorAsAlign(t.instance, ATheAlign , ASpace)
}

func (t *TTabSheet) AnchorClient(ASpace int32) {
    TabSheet_AnchorClient(t.instance, ASpace)
}

func (t *TTabSheet) ScaleDesignToForm(ASize int32) int32 {
    return TabSheet_ScaleDesignToForm(t.instance, ASize)
}

func (t *TTabSheet) ScaleFormToDesign(ASize int32) int32 {
    return TabSheet_ScaleFormToDesign(t.instance, ASize)
}

func (t *TTabSheet) Scale96ToForm(ASize int32) int32 {
    return TabSheet_Scale96ToForm(t.instance, ASize)
}

func (t *TTabSheet) ScaleFormTo96(ASize int32) int32 {
    return TabSheet_ScaleFormTo96(t.instance, ASize)
}

func (t *TTabSheet) Scale96ToFont(ASize int32) int32 {
    return TabSheet_Scale96ToFont(t.instance, ASize)
}

func (t *TTabSheet) ScaleFontTo96(ASize int32) int32 {
    return TabSheet_ScaleFontTo96(t.instance, ASize)
}

func (t *TTabSheet) ScaleScreenToFont(ASize int32) int32 {
    return TabSheet_ScaleScreenToFont(t.instance, ASize)
}

func (t *TTabSheet) ScaleFontToScreen(ASize int32) int32 {
    return TabSheet_ScaleFontToScreen(t.instance, ASize)
}

func (t *TTabSheet) Scale96ToScreen(ASize int32) int32 {
    return TabSheet_Scale96ToScreen(t.instance, ASize)
}

func (t *TTabSheet) ScaleScreenTo96(ASize int32) int32 {
    return TabSheet_ScaleScreenTo96(t.instance, ASize)
}

func (t *TTabSheet) AutoAdjustLayout(AMode TLayoutAdjustmentPolicy, AFromPPI int32, AToPPI int32, AOldFormWidth int32, ANewFormWidth int32) {
    TabSheet_AutoAdjustLayout(t.instance, AMode , AFromPPI , AToPPI , AOldFormWidth , ANewFormWidth)
}

func (t *TTabSheet) FixDesignFontsPPI(ADesignTimePPI int32) {
    TabSheet_FixDesignFontsPPI(t.instance, ADesignTimePPI)
}

func (t *TTabSheet) ScaleFontsPPI(AToPPI int32, AProportion float64) {
    TabSheet_ScaleFontsPPI(t.instance, AToPPI , AProportion)
}

func (t *TTabSheet) PageControl() *TPageControl {
    return AsPageControl(TabSheet_GetPageControl(t.instance))
}

func (t *TTabSheet) SetPageControl(value IWinControl) {
    TabSheet_SetPageControl(t.instance, CheckPtr(value))
}

func (t *TTabSheet) TabIndex() int32 {
    return TabSheet_GetTabIndex(t.instance)
}

// 获取边框的宽度。
func (t *TTabSheet) BorderWidth() int32 {
    return TabSheet_GetBorderWidth(t.instance)
}

// 设置边框的宽度。
func (t *TTabSheet) SetBorderWidth(value int32) {
    TabSheet_SetBorderWidth(t.instance, value)
}

// 获取控件标题。
//
// Get the control title.
func (t *TTabSheet) Caption() string {
    return TabSheet_GetCaption(t.instance)
}

// 设置控件标题。
//
// Set the control title.
func (t *TTabSheet) SetCaption(value string) {
    TabSheet_SetCaption(t.instance, value)
}

// 获取设置控件双缓冲。
//
// Get Set control double buffering.
func (t *TTabSheet) DoubleBuffered() bool {
    return TabSheet_GetDoubleBuffered(t.instance)
}

// 设置设置控件双缓冲。
//
// Set Set control double buffering.
func (t *TTabSheet) SetDoubleBuffered(value bool) {
    TabSheet_SetDoubleBuffered(t.instance, value)
}

// 获取控件启用。
//
// Get the control enabled.
func (t *TTabSheet) Enabled() bool {
    return TabSheet_GetEnabled(t.instance)
}

// 设置控件启用。
//
// Set the control enabled.
func (t *TTabSheet) SetEnabled(value bool) {
    TabSheet_SetEnabled(t.instance, value)
}

// 获取字体。
//
// Get Font.
func (t *TTabSheet) Font() *TFont {
    return AsFont(TabSheet_GetFont(t.instance))
}

// 设置字体。
//
// Set Font.
func (t *TTabSheet) SetFont(value *TFont) {
    TabSheet_SetFont(t.instance, CheckPtr(value))
}

// 获取高度。
//
// Get height.
func (t *TTabSheet) Height() int32 {
    return TabSheet_GetHeight(t.instance)
}

// 设置高度。
//
// Set height.
func (t *TTabSheet) SetHeight(value int32) {
    TabSheet_SetHeight(t.instance, value)
}

// 获取图像在images中的索引。
func (t *TTabSheet) ImageIndex() int32 {
    return TabSheet_GetImageIndex(t.instance)
}

// 设置图像在images中的索引。
func (t *TTabSheet) SetImageIndex(value int32) {
    TabSheet_SetImageIndex(t.instance, value)
}

// 获取左边位置。
//
// Get Left position.
func (t *TTabSheet) Left() int32 {
    return TabSheet_GetLeft(t.instance)
}

// 设置左边位置。
//
// Set Left position.
func (t *TTabSheet) SetLeft(value int32) {
    TabSheet_SetLeft(t.instance, value)
}

// 获取约束控件大小。
func (t *TTabSheet) Constraints() *TSizeConstraints {
    return AsSizeConstraints(TabSheet_GetConstraints(t.instance))
}

// 设置约束控件大小。
func (t *TTabSheet) SetConstraints(value *TSizeConstraints) {
    TabSheet_SetConstraints(t.instance, CheckPtr(value))
}

func (t *TTabSheet) PageIndex() int32 {
    return TabSheet_GetPageIndex(t.instance)
}

func (t *TTabSheet) SetPageIndex(value int32) {
    TabSheet_SetPageIndex(t.instance, value)
}

// 获取使用父容器双缓冲。
//
// Get Parent container double buffering.
func (t *TTabSheet) ParentDoubleBuffered() bool {
    return TabSheet_GetParentDoubleBuffered(t.instance)
}

// 设置使用父容器双缓冲。
//
// Set Parent container double buffering.
func (t *TTabSheet) SetParentDoubleBuffered(value bool) {
    TabSheet_SetParentDoubleBuffered(t.instance, value)
}

// 获取使用父容器字体。
//
// Get Parent container font.
func (t *TTabSheet) ParentFont() bool {
    return TabSheet_GetParentFont(t.instance)
}

// 设置使用父容器字体。
//
// Set Parent container font.
func (t *TTabSheet) SetParentFont(value bool) {
    TabSheet_SetParentFont(t.instance, value)
}

// 获取以父容器的ShowHint属性为准。
func (t *TTabSheet) ParentShowHint() bool {
    return TabSheet_GetParentShowHint(t.instance)
}

// 设置以父容器的ShowHint属性为准。
func (t *TTabSheet) SetParentShowHint(value bool) {
    TabSheet_SetParentShowHint(t.instance, value)
}

// 获取右键菜单。
//
// Get Right click menu.
func (t *TTabSheet) PopupMenu() *TPopupMenu {
    return AsPopupMenu(TabSheet_GetPopupMenu(t.instance))
}

// 设置右键菜单。
//
// Set Right click menu.
func (t *TTabSheet) SetPopupMenu(value IComponent) {
    TabSheet_SetPopupMenu(t.instance, CheckPtr(value))
}

// 获取显示鼠标悬停提示。
//
// Get Show mouseover tips.
func (t *TTabSheet) ShowHint() bool {
    return TabSheet_GetShowHint(t.instance)
}

// 设置显示鼠标悬停提示。
//
// Set Show mouseover tips.
func (t *TTabSheet) SetShowHint(value bool) {
    TabSheet_SetShowHint(t.instance, value)
}

func (t *TTabSheet) TabVisible() bool {
    return TabSheet_GetTabVisible(t.instance)
}

func (t *TTabSheet) SetTabVisible(value bool) {
    TabSheet_SetTabVisible(t.instance, value)
}

// 获取顶边位置。
//
// Get Top position.
func (t *TTabSheet) Top() int32 {
    return TabSheet_GetTop(t.instance)
}

// 设置顶边位置。
//
// Set Top position.
func (t *TTabSheet) SetTop(value int32) {
    TabSheet_SetTop(t.instance, value)
}

// 获取控件可视。
//
// Get the control visible.
func (t *TTabSheet) Visible() bool {
    return TabSheet_GetVisible(t.instance)
}

// 设置控件可视。
//
// Set the control visible.
func (t *TTabSheet) SetVisible(value bool) {
    TabSheet_SetVisible(t.instance, value)
}

// 获取宽度。
//
// Get width.
func (t *TTabSheet) Width() int32 {
    return TabSheet_GetWidth(t.instance)
}

// 设置宽度。
//
// Set width.
func (t *TTabSheet) SetWidth(value int32) {
    TabSheet_SetWidth(t.instance, value)
}

// 设置上下文弹出事件，一般是右键时弹出。
//
// Set Context popup event, usually pop up when right click.
func (t *TTabSheet) SetOnContextPopup(fn TContextPopupEvent) {
    TabSheet_SetOnContextPopup(t.instance, fn)
}

// 设置拖拽下落事件。
//
// Set Drag and drop event.
func (t *TTabSheet) SetOnDragDrop(fn TDragDropEvent) {
    TabSheet_SetOnDragDrop(t.instance, fn)
}

// 设置拖拽完成事件。
//
// Set Drag and drop completion event.
func (t *TTabSheet) SetOnDragOver(fn TDragOverEvent) {
    TabSheet_SetOnDragOver(t.instance, fn)
}

// 设置拖拽结束。
//
// Set End of drag.
func (t *TTabSheet) SetOnEndDrag(fn TEndDragEvent) {
    TabSheet_SetOnEndDrag(t.instance, fn)
}

// 设置焦点进入。
//
// Set Focus entry.
func (t *TTabSheet) SetOnEnter(fn TNotifyEvent) {
    TabSheet_SetOnEnter(t.instance, fn)
}

// 设置焦点退出。
//
// Set Focus exit.
func (t *TTabSheet) SetOnExit(fn TNotifyEvent) {
    TabSheet_SetOnExit(t.instance, fn)
}

// 设置隐藏事件。
func (t *TTabSheet) SetOnHide(fn TNotifyEvent) {
    TabSheet_SetOnHide(t.instance, fn)
}

// 设置鼠标按下事件。
//
// Set Mouse down event.
func (t *TTabSheet) SetOnMouseDown(fn TMouseEvent) {
    TabSheet_SetOnMouseDown(t.instance, fn)
}

// 设置鼠标进入事件。
//
// Set Mouse entry event.
func (t *TTabSheet) SetOnMouseEnter(fn TNotifyEvent) {
    TabSheet_SetOnMouseEnter(t.instance, fn)
}

// 设置鼠标离开事件。
//
// Set Mouse leave event.
func (t *TTabSheet) SetOnMouseLeave(fn TNotifyEvent) {
    TabSheet_SetOnMouseLeave(t.instance, fn)
}

// 设置鼠标移动事件。
func (t *TTabSheet) SetOnMouseMove(fn TMouseMoveEvent) {
    TabSheet_SetOnMouseMove(t.instance, fn)
}

// 设置鼠标抬起事件。
//
// Set Mouse lift event.
func (t *TTabSheet) SetOnMouseUp(fn TMouseEvent) {
    TabSheet_SetOnMouseUp(t.instance, fn)
}

// 设置大小被改变事件。
func (t *TTabSheet) SetOnResize(fn TNotifyEvent) {
    TabSheet_SetOnResize(t.instance, fn)
}

// 设置显示事件。
func (t *TTabSheet) SetOnShow(fn TNotifyEvent) {
    TabSheet_SetOnShow(t.instance, fn)
}

// 获取依靠客户端总数。
func (t *TTabSheet) DockClientCount() int32 {
    return TabSheet_GetDockClientCount(t.instance)
}

// 获取停靠站点。
//
// Get Docking site.
func (t *TTabSheet) DockSite() bool {
    return TabSheet_GetDockSite(t.instance)
}

// 设置停靠站点。
//
// Set Docking site.
func (t *TTabSheet) SetDockSite(value bool) {
    TabSheet_SetDockSite(t.instance, value)
}

// 获取鼠标是否在客户端，仅VCL有效。
//
// Get Whether the mouse is on the client, only VCL is valid.
func (t *TTabSheet) MouseInClient() bool {
    return TabSheet_GetMouseInClient(t.instance)
}

// 获取当前停靠的可视总数。
//
// Get The total number of visible calls currently docked.
func (t *TTabSheet) VisibleDockClientCount() int32 {
    return TabSheet_GetVisibleDockClientCount(t.instance)
}

// 获取画刷对象。
//
// Get Brush.
func (t *TTabSheet) Brush() *TBrush {
    return AsBrush(TabSheet_GetBrush(t.instance))
}

// 获取子控件数。
//
// Get Number of child controls.
func (t *TTabSheet) ControlCount() int32 {
    return TabSheet_GetControlCount(t.instance)
}

// 获取控件句柄。
//
// Get Control handle.
func (t *TTabSheet) Handle() HWND {
    return TabSheet_GetHandle(t.instance)
}

// 获取父容器句柄。
//
// Get Parent container handle.
func (t *TTabSheet) ParentWindow() HWND {
    return TabSheet_GetParentWindow(t.instance)
}

// 设置父容器句柄。
//
// Set Parent container handle.
func (t *TTabSheet) SetParentWindow(value HWND) {
    TabSheet_SetParentWindow(t.instance, value)
}

func (t *TTabSheet) Showing() bool {
    return TabSheet_GetShowing(t.instance)
}

// 获取Tab切换顺序序号。
//
// Get Tab switching sequence number.
func (t *TTabSheet) TabOrder() TTabOrder {
    return TabSheet_GetTabOrder(t.instance)
}

// 设置Tab切换顺序序号。
//
// Set Tab switching sequence number.
func (t *TTabSheet) SetTabOrder(value TTabOrder) {
    TabSheet_SetTabOrder(t.instance, value)
}

// 获取Tab可停留。
//
// Get Tab can stay.
func (t *TTabSheet) TabStop() bool {
    return TabSheet_GetTabStop(t.instance)
}

// 设置Tab可停留。
//
// Set Tab can stay.
func (t *TTabSheet) SetTabStop(value bool) {
    TabSheet_SetTabStop(t.instance, value)
}

// 获取使用停靠管理。
func (t *TTabSheet) UseDockManager() bool {
    return TabSheet_GetUseDockManager(t.instance)
}

// 设置使用停靠管理。
func (t *TTabSheet) SetUseDockManager(value bool) {
    TabSheet_SetUseDockManager(t.instance, value)
}

func (t *TTabSheet) Action() *TAction {
    return AsAction(TabSheet_GetAction(t.instance))
}

func (t *TTabSheet) SetAction(value IComponent) {
    TabSheet_SetAction(t.instance, CheckPtr(value))
}

// 获取控件自动调整。
//
// Get Control automatically adjusts.
func (t *TTabSheet) Align() TAlign {
    return TabSheet_GetAlign(t.instance)
}

// 设置控件自动调整。
//
// Set Control automatically adjusts.
func (t *TTabSheet) SetAlign(value TAlign) {
    TabSheet_SetAlign(t.instance, value)
}

// 获取四个角位置的锚点。
func (t *TTabSheet) Anchors() TAnchors {
    return TabSheet_GetAnchors(t.instance)
}

// 设置四个角位置的锚点。
func (t *TTabSheet) SetAnchors(value TAnchors) {
    TabSheet_SetAnchors(t.instance, value)
}

func (t *TTabSheet) BiDiMode() TBiDiMode {
    return TabSheet_GetBiDiMode(t.instance)
}

func (t *TTabSheet) SetBiDiMode(value TBiDiMode) {
    TabSheet_SetBiDiMode(t.instance, value)
}

func (t *TTabSheet) BoundsRect() TRect {
    return TabSheet_GetBoundsRect(t.instance)
}

func (t *TTabSheet) SetBoundsRect(value TRect) {
    TabSheet_SetBoundsRect(t.instance, value)
}

// 获取客户区高度。
//
// Get client height.
func (t *TTabSheet) ClientHeight() int32 {
    return TabSheet_GetClientHeight(t.instance)
}

// 设置客户区高度。
//
// Set client height.
func (t *TTabSheet) SetClientHeight(value int32) {
    TabSheet_SetClientHeight(t.instance, value)
}

func (t *TTabSheet) ClientOrigin() TPoint {
    return TabSheet_GetClientOrigin(t.instance)
}

// 获取客户区矩形。
//
// Get client rectangle.
func (t *TTabSheet) ClientRect() TRect {
    return TabSheet_GetClientRect(t.instance)
}

// 获取客户区宽度。
//
// Get client width.
func (t *TTabSheet) ClientWidth() int32 {
    return TabSheet_GetClientWidth(t.instance)
}

// 设置客户区宽度。
//
// Set client width.
func (t *TTabSheet) SetClientWidth(value int32) {
    TabSheet_SetClientWidth(t.instance, value)
}

// 获取控件状态。
//
// Get control state.
func (t *TTabSheet) ControlState() TControlState {
    return TabSheet_GetControlState(t.instance)
}

// 设置控件状态。
//
// Set control state.
func (t *TTabSheet) SetControlState(value TControlState) {
    TabSheet_SetControlState(t.instance, value)
}

// 获取控件样式。
//
// Get control style.
func (t *TTabSheet) ControlStyle() TControlStyle {
    return TabSheet_GetControlStyle(t.instance)
}

// 设置控件样式。
//
// Set control style.
func (t *TTabSheet) SetControlStyle(value TControlStyle) {
    TabSheet_SetControlStyle(t.instance, value)
}

func (t *TTabSheet) Floating() bool {
    return TabSheet_GetFloating(t.instance)
}

// 获取控件父容器。
//
// Get control parent container.
func (t *TTabSheet) Parent() *TWinControl {
    return AsWinControl(TabSheet_GetParent(t.instance))
}

// 设置控件父容器。
//
// Set control parent container.
func (t *TTabSheet) SetParent(value IWinControl) {
    TabSheet_SetParent(t.instance, CheckPtr(value))
}

// 获取控件光标。
//
// Get control cursor.
func (t *TTabSheet) Cursor() TCursor {
    return TabSheet_GetCursor(t.instance)
}

// 设置控件光标。
//
// Set control cursor.
func (t *TTabSheet) SetCursor(value TCursor) {
    TabSheet_SetCursor(t.instance, value)
}

// 获取组件鼠标悬停提示。
//
// Get component mouse hints.
func (t *TTabSheet) Hint() string {
    return TabSheet_GetHint(t.instance)
}

// 设置组件鼠标悬停提示。
//
// Set component mouse hints.
func (t *TTabSheet) SetHint(value string) {
    TabSheet_SetHint(t.instance, value)
}

// 获取组件总数。
//
// Get the total number of components.
func (t *TTabSheet) ComponentCount() int32 {
    return TabSheet_GetComponentCount(t.instance)
}

// 获取组件索引。
//
// Get component index.
func (t *TTabSheet) ComponentIndex() int32 {
    return TabSheet_GetComponentIndex(t.instance)
}

// 设置组件索引。
//
// Set component index.
func (t *TTabSheet) SetComponentIndex(value int32) {
    TabSheet_SetComponentIndex(t.instance, value)
}

// 获取组件所有者。
//
// Get component owner.
func (t *TTabSheet) Owner() *TComponent {
    return AsComponent(TabSheet_GetOwner(t.instance))
}

// 获取组件名称。
//
// Get the component name.
func (t *TTabSheet) Name() string {
    return TabSheet_GetName(t.instance)
}

// 设置组件名称。
//
// Set the component name.
func (t *TTabSheet) SetName(value string) {
    TabSheet_SetName(t.instance, value)
}

// 获取对象标记。
//
// Get the control tag.
func (t *TTabSheet) Tag() int {
    return TabSheet_GetTag(t.instance)
}

// 设置对象标记。
//
// Set the control tag.
func (t *TTabSheet) SetTag(value int) {
    TabSheet_SetTag(t.instance, value)
}

// 获取左边锚点。
func (t *TTabSheet) AnchorSideLeft() *TAnchorSide {
    return AsAnchorSide(TabSheet_GetAnchorSideLeft(t.instance))
}

// 设置左边锚点。
func (t *TTabSheet) SetAnchorSideLeft(value *TAnchorSide) {
    TabSheet_SetAnchorSideLeft(t.instance, CheckPtr(value))
}

// 获取顶边锚点。
func (t *TTabSheet) AnchorSideTop() *TAnchorSide {
    return AsAnchorSide(TabSheet_GetAnchorSideTop(t.instance))
}

// 设置顶边锚点。
func (t *TTabSheet) SetAnchorSideTop(value *TAnchorSide) {
    TabSheet_SetAnchorSideTop(t.instance, CheckPtr(value))
}

// 获取右边锚点。
func (t *TTabSheet) AnchorSideRight() *TAnchorSide {
    return AsAnchorSide(TabSheet_GetAnchorSideRight(t.instance))
}

// 设置右边锚点。
func (t *TTabSheet) SetAnchorSideRight(value *TAnchorSide) {
    TabSheet_SetAnchorSideRight(t.instance, CheckPtr(value))
}

// 获取底边锚点。
func (t *TTabSheet) AnchorSideBottom() *TAnchorSide {
    return AsAnchorSide(TabSheet_GetAnchorSideBottom(t.instance))
}

// 设置底边锚点。
func (t *TTabSheet) SetAnchorSideBottom(value *TAnchorSide) {
    TabSheet_SetAnchorSideBottom(t.instance, CheckPtr(value))
}

func (t *TTabSheet) ChildSizing() *TControlChildSizing {
    return AsControlChildSizing(TabSheet_GetChildSizing(t.instance))
}

func (t *TTabSheet) SetChildSizing(value *TControlChildSizing) {
    TabSheet_SetChildSizing(t.instance, CheckPtr(value))
}

// 获取边框间距。
func (t *TTabSheet) BorderSpacing() *TControlBorderSpacing {
    return AsControlBorderSpacing(TabSheet_GetBorderSpacing(t.instance))
}

// 设置边框间距。
func (t *TTabSheet) SetBorderSpacing(value *TControlBorderSpacing) {
    TabSheet_SetBorderSpacing(t.instance, CheckPtr(value))
}

// 获取指定索引停靠客户端。
func (t *TTabSheet) DockClients(Index int32) *TControl {
    return AsControl(TabSheet_GetDockClients(t.instance, Index))
}

// 获取指定索引子控件。
func (t *TTabSheet) Controls(Index int32) *TControl {
    return AsControl(TabSheet_GetControls(t.instance, Index))
}

// 获取指定索引组件。
//
// Get the specified index component.
func (t *TTabSheet) Components(AIndex int32) *TComponent {
    return AsComponent(TabSheet_GetComponents(t.instance, AIndex))
}

// 获取锚侧面。
func (t *TTabSheet) AnchorSide(AKind TAnchorKind) *TAnchorSide {
    return AsAnchorSide(TabSheet_GetAnchorSide(t.instance, AKind))
}

