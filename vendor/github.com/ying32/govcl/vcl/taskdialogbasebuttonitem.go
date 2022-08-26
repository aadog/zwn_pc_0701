
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

type TTaskDialogBaseButtonItem struct {
    IObject
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与LCL没有太多关系。
    ptr unsafe.Pointer
}

// 创建一个新的对象。
// 
// Create a new object.
func NewTaskDialogBaseButtonItem(AOwner *TCollection) *TTaskDialogBaseButtonItem {
    t := new(TTaskDialogBaseButtonItem)
    t.instance = TaskDialogBaseButtonItem_Create(CheckPtr(AOwner))
    t.ptr = unsafe.Pointer(t.instance)
    setFinalizer(t, (*TTaskDialogBaseButtonItem).Free)
    return t
}

// 动态转换一个已存在的对象实例。
// 
// Dynamically convert an existing object instance.
func AsTaskDialogBaseButtonItem(obj interface{}) *TTaskDialogBaseButtonItem {
    instance, ptr := getInstance(obj)
    if instance == 0 { return nil }
    return &TTaskDialogBaseButtonItem{instance: instance, ptr: ptr}
}

// -------------------------- Deprecated begin --------------------------
// 新建一个对象来自已经存在的对象实例指针。
// 
// Create a new object from an existing object instance pointer.
// Deprecated: use AsTaskDialogBaseButtonItem.
func TaskDialogBaseButtonItemFromInst(inst uintptr) *TTaskDialogBaseButtonItem {
    return AsTaskDialogBaseButtonItem(inst)
}

// 新建一个对象来自已经存在的对象实例。
// 
// Create a new object from an existing object instance.
// Deprecated: use AsTaskDialogBaseButtonItem.
func TaskDialogBaseButtonItemFromObj(obj IObject) *TTaskDialogBaseButtonItem {
    return AsTaskDialogBaseButtonItem(obj)
}

// 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// 
// Create a new object from an unsecured address. Note: Using this function may cause some unclear situations and be used with caution..
// Deprecated: use AsTaskDialogBaseButtonItem.
func TaskDialogBaseButtonItemFromUnsafePointer(ptr unsafe.Pointer) *TTaskDialogBaseButtonItem {
    return AsTaskDialogBaseButtonItem(ptr)
}

// -------------------------- Deprecated end --------------------------
// 释放对象。
// 
// Free object.
func (t *TTaskDialogBaseButtonItem) Free() {
    if t.instance != 0 {
        TaskDialogBaseButtonItem_Free(t.instance)
        t.instance, t.ptr = 0, nullptr
    }
}

// 返回对象实例指针。
// 
// Return object instance pointer.
func (t *TTaskDialogBaseButtonItem) Instance() uintptr {
    return t.instance
}

// 获取一个不安全的地址。
// 
// Get an unsafe address.
func (t *TTaskDialogBaseButtonItem) UnsafeAddr() unsafe.Pointer {
    return t.ptr
}

// 检测地址是否为空。
// 
// Check if the address is empty.
func (t *TTaskDialogBaseButtonItem) IsValid() bool {
    return t.instance != 0
}

// 检测当前对象是否继承自目标对象。
// 
// Checks whether the current object is inherited from the target object.
func (t *TTaskDialogBaseButtonItem) Is() TIs {
    return TIs(t.instance)
}

// 动态转换当前对象为目标对象。
// 
// Dynamically convert the current object to the target object.
//func (t *TTaskDialogBaseButtonItem) As() TAs {
//    return TAs(t.instance)
//}

// 获取类信息指针。
// 
// Get class information pointer.
func TTaskDialogBaseButtonItemClass() TClass {
    return TaskDialogBaseButtonItem_StaticClassType()
}

// 获取类名路径。
//
// Get the class name path.
func (t *TTaskDialogBaseButtonItem) GetNamePath() string {
    return TaskDialogBaseButtonItem_GetNamePath(t.instance)
}

// 复制一个对象，如果对象实现了此方法的话。
//
// Copy an object, if the object implements this method.
func (t *TTaskDialogBaseButtonItem) Assign(Source IObject) {
    TaskDialogBaseButtonItem_Assign(t.instance, CheckPtr(Source))
}

// 获取类的类型信息。
//
// Get class type information.
func (t *TTaskDialogBaseButtonItem) ClassType() TClass {
    return TaskDialogBaseButtonItem_ClassType(t.instance)
}

// 获取当前对象类名称。
//
// Get the current object class name.
func (t *TTaskDialogBaseButtonItem) ClassName() string {
    return TaskDialogBaseButtonItem_ClassName(t.instance)
}

// 获取当前对象实例大小。
//
// Get the current object instance size.
func (t *TTaskDialogBaseButtonItem) InstanceSize() int32 {
    return TaskDialogBaseButtonItem_InstanceSize(t.instance)
}

// 判断当前类是否继承自指定类。
//
// Determine whether the current class inherits from the specified class.
func (t *TTaskDialogBaseButtonItem) InheritsFrom(AClass TClass) bool {
    return TaskDialogBaseButtonItem_InheritsFrom(t.instance, AClass)
}

// 与一个对象进行比较。
//
// Compare with an object.
func (t *TTaskDialogBaseButtonItem) Equals(Obj IObject) bool {
    return TaskDialogBaseButtonItem_Equals(t.instance, CheckPtr(Obj))
}

// 获取类的哈希值。
//
// Get the hash value of the class.
func (t *TTaskDialogBaseButtonItem) GetHashCode() int32 {
    return TaskDialogBaseButtonItem_GetHashCode(t.instance)
}

// 文本类信息。
//
// Text information.
func (t *TTaskDialogBaseButtonItem) ToString() string {
    return TaskDialogBaseButtonItem_ToString(t.instance)
}

// 获取模态对话框显示结果。
func (t *TTaskDialogBaseButtonItem) ModalResult() TModalResult {
    return TaskDialogBaseButtonItem_GetModalResult(t.instance)
}

// 设置模态对话框显示结果。
func (t *TTaskDialogBaseButtonItem) SetModalResult(value TModalResult) {
    TaskDialogBaseButtonItem_SetModalResult(t.instance, value)
}

// 获取控件标题。
//
// Get the control title.
func (t *TTaskDialogBaseButtonItem) Caption() string {
    return TaskDialogBaseButtonItem_GetCaption(t.instance)
}

// 设置控件标题。
//
// Set the control title.
func (t *TTaskDialogBaseButtonItem) SetCaption(value string) {
    TaskDialogBaseButtonItem_SetCaption(t.instance, value)
}

func (t *TTaskDialogBaseButtonItem) Default() bool {
    return TaskDialogBaseButtonItem_GetDefault(t.instance)
}

func (t *TTaskDialogBaseButtonItem) SetDefault(value bool) {
    TaskDialogBaseButtonItem_SetDefault(t.instance, value)
}

func (t *TTaskDialogBaseButtonItem) Collection() *TCollection {
    return AsCollection(TaskDialogBaseButtonItem_GetCollection(t.instance))
}

func (t *TTaskDialogBaseButtonItem) SetCollection(value *TCollection) {
    TaskDialogBaseButtonItem_SetCollection(t.instance, CheckPtr(value))
}

func (t *TTaskDialogBaseButtonItem) Index() int32 {
    return TaskDialogBaseButtonItem_GetIndex(t.instance)
}

func (t *TTaskDialogBaseButtonItem) SetIndex(value int32) {
    TaskDialogBaseButtonItem_SetIndex(t.instance, value)
}

func (t *TTaskDialogBaseButtonItem) DisplayName() string {
    return TaskDialogBaseButtonItem_GetDisplayName(t.instance)
}

func (t *TTaskDialogBaseButtonItem) SetDisplayName(value string) {
    TaskDialogBaseButtonItem_SetDisplayName(t.instance, value)
}

