
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

type TReplaceDialog struct {
    IComponent
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与LCL没有太多关系。
    ptr unsafe.Pointer
}

// 创建一个新的对象。
// 
// Create a new object.
func NewReplaceDialog(owner IComponent) *TReplaceDialog {
    r := new(TReplaceDialog)
    r.instance = ReplaceDialog_Create(CheckPtr(owner))
    r.ptr = unsafe.Pointer(r.instance)
    return r
}

// 动态转换一个已存在的对象实例。
// 
// Dynamically convert an existing object instance.
func AsReplaceDialog(obj interface{}) *TReplaceDialog {
    instance, ptr := getInstance(obj)
    if instance == 0 { return nil }
    return &TReplaceDialog{instance: instance, ptr: ptr}
}

// -------------------------- Deprecated begin --------------------------
// 新建一个对象来自已经存在的对象实例指针。
// 
// Create a new object from an existing object instance pointer.
// Deprecated: use AsReplaceDialog.
func ReplaceDialogFromInst(inst uintptr) *TReplaceDialog {
    return AsReplaceDialog(inst)
}

// 新建一个对象来自已经存在的对象实例。
// 
// Create a new object from an existing object instance.
// Deprecated: use AsReplaceDialog.
func ReplaceDialogFromObj(obj IObject) *TReplaceDialog {
    return AsReplaceDialog(obj)
}

// 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// 
// Create a new object from an unsecured address. Note: Using this function may cause some unclear situations and be used with caution..
// Deprecated: use AsReplaceDialog.
func ReplaceDialogFromUnsafePointer(ptr unsafe.Pointer) *TReplaceDialog {
    return AsReplaceDialog(ptr)
}

// -------------------------- Deprecated end --------------------------
// 释放对象。
// 
// Free object.
func (r *TReplaceDialog) Free() {
    if r.instance != 0 {
        ReplaceDialog_Free(r.instance)
        r.instance, r.ptr = 0, nullptr
    }
}

// 返回对象实例指针。
// 
// Return object instance pointer.
func (r *TReplaceDialog) Instance() uintptr {
    return r.instance
}

// 获取一个不安全的地址。
// 
// Get an unsafe address.
func (r *TReplaceDialog) UnsafeAddr() unsafe.Pointer {
    return r.ptr
}

// 检测地址是否为空。
// 
// Check if the address is empty.
func (r *TReplaceDialog) IsValid() bool {
    return r.instance != 0
}

// 检测当前对象是否继承自目标对象。
// 
// Checks whether the current object is inherited from the target object.
func (r *TReplaceDialog) Is() TIs {
    return TIs(r.instance)
}

// 动态转换当前对象为目标对象。
// 
// Dynamically convert the current object to the target object.
//func (r *TReplaceDialog) As() TAs {
//    return TAs(r.instance)
//}

// 获取类信息指针。
// 
// Get class information pointer.
func TReplaceDialogClass() TClass {
    return ReplaceDialog_StaticClassType()
}

func (r *TReplaceDialog) CloseDialog() {
    ReplaceDialog_CloseDialog(r.instance)
}

// 执行。
func (r *TReplaceDialog) Execute() bool {
    return ReplaceDialog_Execute(r.instance)
}

// 查找指定名称的组件。
//
// Find the component with the specified name.
func (r *TReplaceDialog) FindComponent(AName string) *TComponent {
    return AsComponent(ReplaceDialog_FindComponent(r.instance, AName))
}

// 获取类名路径。
//
// Get the class name path.
func (r *TReplaceDialog) GetNamePath() string {
    return ReplaceDialog_GetNamePath(r.instance)
}

// 是否有父容器。
//
// Is there a parent container.
func (r *TReplaceDialog) HasParent() bool {
    return ReplaceDialog_HasParent(r.instance)
}

// 复制一个对象，如果对象实现了此方法的话。
//
// Copy an object, if the object implements this method.
func (r *TReplaceDialog) Assign(Source IObject) {
    ReplaceDialog_Assign(r.instance, CheckPtr(Source))
}

// 获取类的类型信息。
//
// Get class type information.
func (r *TReplaceDialog) ClassType() TClass {
    return ReplaceDialog_ClassType(r.instance)
}

// 获取当前对象类名称。
//
// Get the current object class name.
func (r *TReplaceDialog) ClassName() string {
    return ReplaceDialog_ClassName(r.instance)
}

// 获取当前对象实例大小。
//
// Get the current object instance size.
func (r *TReplaceDialog) InstanceSize() int32 {
    return ReplaceDialog_InstanceSize(r.instance)
}

// 判断当前类是否继承自指定类。
//
// Determine whether the current class inherits from the specified class.
func (r *TReplaceDialog) InheritsFrom(AClass TClass) bool {
    return ReplaceDialog_InheritsFrom(r.instance, AClass)
}

// 与一个对象进行比较。
//
// Compare with an object.
func (r *TReplaceDialog) Equals(Obj IObject) bool {
    return ReplaceDialog_Equals(r.instance, CheckPtr(Obj))
}

// 获取类的哈希值。
//
// Get the hash value of the class.
func (r *TReplaceDialog) GetHashCode() int32 {
    return ReplaceDialog_GetHashCode(r.instance)
}

// 文本类信息。
//
// Text information.
func (r *TReplaceDialog) ToString() string {
    return ReplaceDialog_ToString(r.instance)
}

func (r *TReplaceDialog) ReplaceText() string {
    return ReplaceDialog_GetReplaceText(r.instance)
}

func (r *TReplaceDialog) SetReplaceText(value string) {
    ReplaceDialog_SetReplaceText(r.instance, value)
}

func (r *TReplaceDialog) SetOnReplace(fn TNotifyEvent) {
    ReplaceDialog_SetOnReplace(r.instance, fn)
}

// 获取左边位置。
//
// Get Left position.
func (r *TReplaceDialog) Left() int32 {
    return ReplaceDialog_GetLeft(r.instance)
}

// 设置左边位置。
//
// Set Left position.
func (r *TReplaceDialog) SetLeft(value int32) {
    ReplaceDialog_SetLeft(r.instance, value)
}

func (r *TReplaceDialog) Position() TPoint {
    return ReplaceDialog_GetPosition(r.instance)
}

func (r *TReplaceDialog) SetPosition(value TPoint) {
    ReplaceDialog_SetPosition(r.instance, value)
}

// 获取顶边位置。
//
// Get Top position.
func (r *TReplaceDialog) Top() int32 {
    return ReplaceDialog_GetTop(r.instance)
}

// 设置顶边位置。
//
// Set Top position.
func (r *TReplaceDialog) SetTop(value int32) {
    ReplaceDialog_SetTop(r.instance, value)
}

func (r *TReplaceDialog) FindText() string {
    return ReplaceDialog_GetFindText(r.instance)
}

func (r *TReplaceDialog) SetFindText(value string) {
    ReplaceDialog_SetFindText(r.instance, value)
}

func (r *TReplaceDialog) Options() TFindOptions {
    return ReplaceDialog_GetOptions(r.instance)
}

func (r *TReplaceDialog) SetOptions(value TFindOptions) {
    ReplaceDialog_SetOptions(r.instance, value)
}

func (r *TReplaceDialog) SetOnFind(fn TNotifyEvent) {
    ReplaceDialog_SetOnFind(r.instance, fn)
}

// 获取控件句柄。
//
// Get Control handle.
func (r *TReplaceDialog) Handle() HWND {
    return ReplaceDialog_GetHandle(r.instance)
}

func (r *TReplaceDialog) SetOnClose(fn TNotifyEvent) {
    ReplaceDialog_SetOnClose(r.instance, fn)
}

// 设置显示事件。
func (r *TReplaceDialog) SetOnShow(fn TNotifyEvent) {
    ReplaceDialog_SetOnShow(r.instance, fn)
}

// 获取组件总数。
//
// Get the total number of components.
func (r *TReplaceDialog) ComponentCount() int32 {
    return ReplaceDialog_GetComponentCount(r.instance)
}

// 获取组件索引。
//
// Get component index.
func (r *TReplaceDialog) ComponentIndex() int32 {
    return ReplaceDialog_GetComponentIndex(r.instance)
}

// 设置组件索引。
//
// Set component index.
func (r *TReplaceDialog) SetComponentIndex(value int32) {
    ReplaceDialog_SetComponentIndex(r.instance, value)
}

// 获取组件所有者。
//
// Get component owner.
func (r *TReplaceDialog) Owner() *TComponent {
    return AsComponent(ReplaceDialog_GetOwner(r.instance))
}

// 获取组件名称。
//
// Get the component name.
func (r *TReplaceDialog) Name() string {
    return ReplaceDialog_GetName(r.instance)
}

// 设置组件名称。
//
// Set the component name.
func (r *TReplaceDialog) SetName(value string) {
    ReplaceDialog_SetName(r.instance, value)
}

// 获取对象标记。
//
// Get the control tag.
func (r *TReplaceDialog) Tag() int {
    return ReplaceDialog_GetTag(r.instance)
}

// 设置对象标记。
//
// Set the control tag.
func (r *TReplaceDialog) SetTag(value int) {
    ReplaceDialog_SetTag(r.instance, value)
}

// 获取指定索引组件。
//
// Get the specified index component.
func (r *TReplaceDialog) Components(AIndex int32) *TComponent {
    return AsComponent(ReplaceDialog_GetComponents(r.instance, AIndex))
}

