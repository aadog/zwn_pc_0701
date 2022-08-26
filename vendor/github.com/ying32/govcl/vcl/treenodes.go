
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

type TTreeNodes struct {
    IObject
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与LCL没有太多关系。
    ptr unsafe.Pointer
}

// 创建一个新的对象。
// 
// Create a new object.
func NewTreeNodes(AOwner *TTreeView) *TTreeNodes {
    t := new(TTreeNodes)
    t.instance = TreeNodes_Create(CheckPtr(AOwner))
    t.ptr = unsafe.Pointer(t.instance)
    setFinalizer(t, (*TTreeNodes).Free)
    return t
}

// 动态转换一个已存在的对象实例。
// 
// Dynamically convert an existing object instance.
func AsTreeNodes(obj interface{}) *TTreeNodes {
    instance, ptr := getInstance(obj)
    if instance == 0 { return nil }
    return &TTreeNodes{instance: instance, ptr: ptr}
}

// -------------------------- Deprecated begin --------------------------
// 新建一个对象来自已经存在的对象实例指针。
// 
// Create a new object from an existing object instance pointer.
// Deprecated: use AsTreeNodes.
func TreeNodesFromInst(inst uintptr) *TTreeNodes {
    return AsTreeNodes(inst)
}

// 新建一个对象来自已经存在的对象实例。
// 
// Create a new object from an existing object instance.
// Deprecated: use AsTreeNodes.
func TreeNodesFromObj(obj IObject) *TTreeNodes {
    return AsTreeNodes(obj)
}

// 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// 
// Create a new object from an unsecured address. Note: Using this function may cause some unclear situations and be used with caution..
// Deprecated: use AsTreeNodes.
func TreeNodesFromUnsafePointer(ptr unsafe.Pointer) *TTreeNodes {
    return AsTreeNodes(ptr)
}

// -------------------------- Deprecated end --------------------------
// 释放对象。
// 
// Free object.
func (t *TTreeNodes) Free() {
    if t.instance != 0 {
        TreeNodes_Free(t.instance)
        t.instance, t.ptr = 0, nullptr
    }
}

// 返回对象实例指针。
// 
// Return object instance pointer.
func (t *TTreeNodes) Instance() uintptr {
    return t.instance
}

// 获取一个不安全的地址。
// 
// Get an unsafe address.
func (t *TTreeNodes) UnsafeAddr() unsafe.Pointer {
    return t.ptr
}

// 检测地址是否为空。
// 
// Check if the address is empty.
func (t *TTreeNodes) IsValid() bool {
    return t.instance != 0
}

// 检测当前对象是否继承自目标对象。
// 
// Checks whether the current object is inherited from the target object.
func (t *TTreeNodes) Is() TIs {
    return TIs(t.instance)
}

// 动态转换当前对象为目标对象。
// 
// Dynamically convert the current object to the target object.
//func (t *TTreeNodes) As() TAs {
//    return TAs(t.instance)
//}

// 获取类信息指针。
// 
// Get class information pointer.
func TTreeNodesClass() TClass {
    return TreeNodes_StaticClassType()
}

func (t *TTreeNodes) AddChildFirst(Parent *TTreeNode, S string) *TTreeNode {
    return AsTreeNode(TreeNodes_AddChildFirst(t.instance, CheckPtr(Parent), S))
}

func (t *TTreeNodes) AddChild(Parent *TTreeNode, S string) *TTreeNode {
    return AsTreeNode(TreeNodes_AddChild(t.instance, CheckPtr(Parent), S))
}

func (t *TTreeNodes) AddChildObjectFirst(Parent *TTreeNode, S string, Ptr uintptr) *TTreeNode {
    return AsTreeNode(TreeNodes_AddChildObjectFirst(t.instance, CheckPtr(Parent), S , Ptr))
}

func (t *TTreeNodes) AddChildObject(Parent *TTreeNode, S string, Ptr uintptr) *TTreeNode {
    return AsTreeNode(TreeNodes_AddChildObject(t.instance, CheckPtr(Parent), S , Ptr))
}

func (t *TTreeNodes) AddObjectFirst(Sibling *TTreeNode, S string, Ptr uintptr) *TTreeNode {
    return AsTreeNode(TreeNodes_AddObjectFirst(t.instance, CheckPtr(Sibling), S , Ptr))
}

func (t *TTreeNodes) AddObject(Sibling *TTreeNode, S string, Ptr uintptr) *TTreeNode {
    return AsTreeNode(TreeNodes_AddObject(t.instance, CheckPtr(Sibling), S , Ptr))
}

func (t *TTreeNodes) AddNode(Node *TTreeNode, Relative *TTreeNode, S string, Ptr uintptr, Method TNodeAttachMode) *TTreeNode {
    return AsTreeNode(TreeNodes_AddNode(t.instance, CheckPtr(Node), CheckPtr(Relative), S , Ptr , Method))
}

func (t *TTreeNodes) AddFirst(Sibling *TTreeNode, S string) *TTreeNode {
    return AsTreeNode(TreeNodes_AddFirst(t.instance, CheckPtr(Sibling), S))
}

func (t *TTreeNodes) Add(Sibling *TTreeNode, S string) *TTreeNode {
    return AsTreeNode(TreeNodes_Add(t.instance, CheckPtr(Sibling), S))
}

// 复制一个对象，如果对象实现了此方法的话。
//
// Copy an object, if the object implements this method.
func (t *TTreeNodes) Assign(Source IObject) {
    TreeNodes_Assign(t.instance, CheckPtr(Source))
}

func (t *TTreeNodes) BeginUpdate() {
    TreeNodes_BeginUpdate(t.instance)
}

// 清除。
func (t *TTreeNodes) Clear() {
    TreeNodes_Clear(t.instance)
}

func (t *TTreeNodes) Delete(Node *TTreeNode) {
    TreeNodes_Delete(t.instance, CheckPtr(Node))
}

func (t *TTreeNodes) EndUpdate() {
    TreeNodes_EndUpdate(t.instance)
}

func (t *TTreeNodes) GetFirstNode() *TTreeNode {
    return AsTreeNode(TreeNodes_GetFirstNode(t.instance))
}

func (t *TTreeNodes) Insert(Sibling *TTreeNode, S string) *TTreeNode {
    return AsTreeNode(TreeNodes_Insert(t.instance, CheckPtr(Sibling), S))
}

func (t *TTreeNodes) InsertObject(Sibling *TTreeNode, S string, Ptr uintptr) *TTreeNode {
    return AsTreeNode(TreeNodes_InsertObject(t.instance, CheckPtr(Sibling), S , Ptr))
}

func (t *TTreeNodes) CustomSort(SortProc PFNTVCOMPARE, Data int, ARecurse bool) bool {
    return TreeNodes_CustomSort(t.instance, SortProc , Data , ARecurse)
}

// 获取类名路径。
//
// Get the class name path.
func (t *TTreeNodes) GetNamePath() string {
    return TreeNodes_GetNamePath(t.instance)
}

// 获取类的类型信息。
//
// Get class type information.
func (t *TTreeNodes) ClassType() TClass {
    return TreeNodes_ClassType(t.instance)
}

// 获取当前对象类名称。
//
// Get the current object class name.
func (t *TTreeNodes) ClassName() string {
    return TreeNodes_ClassName(t.instance)
}

// 获取当前对象实例大小。
//
// Get the current object instance size.
func (t *TTreeNodes) InstanceSize() int32 {
    return TreeNodes_InstanceSize(t.instance)
}

// 判断当前类是否继承自指定类。
//
// Determine whether the current class inherits from the specified class.
func (t *TTreeNodes) InheritsFrom(AClass TClass) bool {
    return TreeNodes_InheritsFrom(t.instance, AClass)
}

// 与一个对象进行比较。
//
// Compare with an object.
func (t *TTreeNodes) Equals(Obj IObject) bool {
    return TreeNodes_Equals(t.instance, CheckPtr(Obj))
}

// 获取类的哈希值。
//
// Get the hash value of the class.
func (t *TTreeNodes) GetHashCode() int32 {
    return TreeNodes_GetHashCode(t.instance)
}

// 文本类信息。
//
// Text information.
func (t *TTreeNodes) ToString() string {
    return TreeNodes_ToString(t.instance)
}

func (t *TTreeNodes) Count() int32 {
    return TreeNodes_GetCount(t.instance)
}

// 获取组件所有者。
//
// Get component owner.
func (t *TTreeNodes) Owner() *TWinControl {
    return AsWinControl(TreeNodes_GetOwner(t.instance))
}

func (t *TTreeNodes) Item(Index int32) *TTreeNode {
    return AsTreeNode(TreeNodes_GetItem(t.instance, Index))
}

