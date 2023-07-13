package codegen

import (
	"os"

	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/enum"
	"github.com/llir/llvm/ir/types"
)

func genHelloWorld() {
	m := ir.NewModule()
	param := ir.NewParam("", types.I8Ptr)
	param.Attrs = append(param.Attrs, enum.ParamAttrNoCapture)
	printf := m.NewFunc("printf", types.I32, param)
	printf.FuncAttrs = append(printf.FuncAttrs, enum.FuncAttrNoUnwind)

	pv := m.NewGlobalDef(".v", constant.NewCharArrayFromString("Hello World!"))
	pv.IsConstant()
	pv.Linkage = enum.LinkagePrivate

	funcMain := m.NewFunc(
		"main",
		types.I32,
	)
	mb := funcMain.NewBlock("")
	mb.NewCall(printf, mb.NewGetElementPtr(types.I8Ptr, pv, constant.NewInt(types.I32, 0)))
	mb.NewRet(constant.NewInt(types.I32, 0))

	println(m.String())
	f, _ := os.Create("ir.ll")
	defer f.Close()
	m.WriteTo(f)
}
