package class

var _shimClass = &Class{name:"~shim"}
var _returnCode = []byte{0xb1} // return
var _athrowCode = []byte{0xbf} // athrow

var _returnMethod = &Method{
    ClassMember: ClassMember{
        AccessFlags: AccessFlags{ACC_STATIC},
        name: "<return>",
        class: _shimClass,
    },
    code: _returnCode,
}

func ReturnShimMethod() (*Method) {
    return _returnMethod
}

func NewAthrowMethod(maxStack, maxLocals uint) (*Method) {
    method := &Method{}
    method.class = _shimClass
    method.name = "<athrow>"
    method.accessFlags = ACC_STATIC
    method.maxStack = maxStack
    method.maxLocals = maxLocals
    method.code = _athrowCode
    return method
}