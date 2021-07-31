package main

import (
	"github.com/leaanthony/slicer"
	"strings"
)

type InterfaceDeclaration struct {
	//Header *InterfaceHeader `parser:"'[' @@* ']'"`
	Name   string             `parser:"'interface' @Ident ':' 'IUnknown' '{' "`
	Method []*InterfaceMethod `parser:"@@+ '}'"`
}

func (i *InterfaceDeclaration) String(forewardDecls slicer.StringSlicer) string {
	var text strings.Builder

	interfaceName := "_" + i.Name
	implName := interfaceName + "Impl"
	vtblName := interfaceName + "Vtbl"
	invokeMethod := i.invokeMethod()
	interfaceGoMethodName := strings.TrimPrefix(i.Name, "ICoreWebView2")
	interfaceGoMethodName = strings.TrimSuffix(interfaceGoMethodName, "Handler")
	interfaceGoMethodName = strings.TrimSuffix(interfaceGoMethodName, "Event")

	// Write Vtbl
	text.WriteString("type " + vtblName + " struct {\n")
	text.WriteString("\t_IUnknownVtbl\n")
	for _, method := range i.Method {
		text.WriteString("\t" + method.MethodName() + "\tComProc\n")
	}
	text.WriteString("}\n\n")

	// Write main struct
	text.WriteString("type " + i.Name + " struct {\n")
	text.WriteString("\tvtbl *" + vtblName + "\n")
	if invokeMethod != nil {
		text.WriteString("\timpl " + implName + "\n")
	}
	text.WriteString("}\n\n")

	if invokeMethod != nil {
		invoke := strings.ReplaceAll(`
func $$IUnknownQueryInterface(this *$NAME$, refiid, object uintptr) uintptr {
	return this.impl.QueryInterface(refiid, object)
}

func $$IUnknownAddRef(this *$NAME$) uintptr {
	return this.impl.AddRef()
}

func $$IUnknownRelease(this *$NAME$) uintptr {
	return this.impl.Release()
}
`, "$$", interfaceName)
		invoke = strings.ReplaceAll(invoke, "$NAME$", i.Name)
		text.WriteString(invoke)
		// TODO: Dynamically create Invoke call
		text.WriteString("func " + interfaceName + "Invoke(this *" + i.Name)
		for _, param := range invokeMethod.Params {
			text.WriteString(", " + param.String(forewardDecls))
		}
		text.WriteString(") " + GoType(invokeMethod.ReturnType) + " {\n")
		text.WriteString("	return this.impl." + interfaceGoMethodName + "(" + invokeMethod.GoArgVars() + ")\n")
		text.WriteString("}")

		// TODO: Create Impl call
		implTemplate := `
type $$Impl interface {
	_IUnknownImpl
	$GOMETHOD$($ARGS$) $RETURN$
}

var $$Fn = $VTBL${
	_IUnknownVtbl{
		NewComProc($$IUnknownQueryInterface),
		NewComProc($$IUnknownAddRef),
		NewComProc($$IUnknownRelease),
	},
	NewComProc($$Invoke),
}

func New$NAME$(impl $$Impl) *$NAME$ {
	return &$NAME${
		vtbl: &$$Fn,
		impl: impl,
	}
}
`
		implTemplate = strings.ReplaceAll(implTemplate, "$$", interfaceName)
		implTemplate = strings.ReplaceAll(implTemplate, "$NAME$", i.Name)
		implTemplate = strings.ReplaceAll(implTemplate, "$VTBL$", vtblName)
		implTemplate = strings.ReplaceAll(implTemplate, "$RETURN$", GoType(invokeMethod.ReturnType))
		implTemplate = strings.ReplaceAll(implTemplate, "$GOMETHOD$", interfaceGoMethodName)
		implTemplate = strings.ReplaceAll(implTemplate, "$ARGS$", invokeMethod.GoArgs(forewardDecls))
		text.WriteString(implTemplate)

	}

	return text.String()
}

func (m *InterfaceDeclaration) invokeMethod() *InterfaceMethod {
	for _, m := range m.Method {
		if m == nil {
			continue
		}
		if *m.Name == "Invoke" {
			return m
		}
	}
	return nil
}

type Prop string

func (p *Prop) Capture(values []string) error {
	if len(values) == 0 {
		return nil
	}
	result := strings.Title(values[0][4:])
	*p = Prop(result)
	return nil
}

type MethodName string

func (m *MethodName) Capture(values []string) error {
	if len(values) == 0 {
		return nil
	}
	result := values[0]
	if strings.HasPrefix(values[0], "add_") {
		result = "Add" + result[4:]
	}
	if strings.HasPrefix(values[0], "remove_") {
		result = "Remove" + result[7:]
	}
	*m = MethodName(result)
	return nil
}

type InterfaceMethod struct {
	Prop       *Prop       `parser:"('[' @('propget'|'propput') ']')?"`
	ReturnType string      `parser:"@Ident"`
	Name       *MethodName `parser:"@Ident '('"`
	Params     []*Param    `parser:" @@* ')' ';'"`
}

func (m InterfaceMethod) MethodName() string {
	if m.Prop != nil {
		return string(*m.Prop) + string(*m.Name)
	}
	return string(*m.Name)
}

var idlTypeMap = map[string]string{
	"IUnknown": "_IUnknown",
	"HRESULT":  "uintptr",
	"LPWSTR":   "*uint16",
	"LPCWSTR":  "*uint16",
	"IStream":  "uintptr",
	"UINT64":   "uint64",
	"UINT32":   "uint32",
	"UINT":     "uint",
	"INT":      "int",
	"BOOL":     "bool",
}

func GoType(inputType string) string {
	if inputType == "" {
		return ""
	}

	return idlTypeMap[inputType]
}

func (m InterfaceMethod) GoArgs(forewardDecls slicer.StringSlicer) string {
	var result slicer.StringSlicer
	for _, param := range m.Params {
		result.Add(param.String(forewardDecls))
	}
	return result.Join(",")
}

func (m InterfaceMethod) GoArgVars() string {
	var result slicer.StringSlicer
	for _, param := range m.Params {
		result.Add(param.Name)
	}
	return result.Join(",")
}

type UUID string

func (u *UUID) Capture(values []string) error {
	//println("UUID =", values[0])
	*u = UUID(values[0])
	return nil
}

type InterfaceHeader struct {
	UUID *UUID `parser:"'uuid' '(' @~')' ')' ',' 'object' ',' 'pointer_default' '(' 'unique' ')'"`
}

type Boolean bool

func (b *Boolean) Capture(values []string) error {
	*b = values[0] == "true"
	return nil
}

type Direction struct {
	Dir    string `parser:"'[' @('out'|'in')"`
	Retval string `parser:"(',' @'retval')? ']'"`
}

type Param struct {
	Direction *Direction `parser:"@@?"`
	Type      string     `parser:"@Ident"`
	Const     string     `parser:"@('const')?"`
	Pointer   string     `parser:"@('*')*"`
	Name      string     `parser:"@Ident ','?"`
}

func (p *Param) String(forewardDecls slicer.StringSlicer) string {
	t := p.Type
	//println("Looking for", t)
	//println(forewardDecls.Join(" "))
	if !forewardDecls.Contains(t) {
		t = GoType(t)
	}
	return p.Name + " " + p.Pointer + t
}
