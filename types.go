package main

type IDL struct {
	Imports   []*Import  `parser:"@@*"`
	Libraries []*Library `parser:"@@*"`
}

type Import struct {
	Name string `parser:"'import' @(!';')* ';'"`
}

type Library struct {
	Name       string         `parser:"'library' @Ident"`
	Statements []*Declaration `parser:"'{' @@* '}'"`
}

type Declaration struct {
	InterfaceForewardDecl string                `parser:"'interface' @Ident ';'"`
	Enum                  *EnumDeclaration      `parser:"| '[' 'v1_enum' ']' @@"`
	Struct                *StructDeclaration    `parser:"| @@"`
	Interface             *InterfaceDeclaration `parser:"| @@"`
	CppQuote              string                `parser:"| 'cpp_quote' '(' @String ')'"`
}

type EnumDeclaration struct {
	Name   string       `parser:"'typedef' 'enum' @Ident"`
	Values []*EnumValue `parser:"'{' (@@)+ '}' Ident ';'"`
}

type EnumValue struct {
	Value string `parser:"@Ident ','?"`
}

type StructDeclaration struct {
	Name   string         `parser:"'typedef' 'struct' @Ident '{' "`
	Fields []*StructField `parser:" (@@)+ '}' Ident ';'"`
}

type StructField struct {
	Type string `parser:"@('UINT32' | 'BOOL')"`
	Name string `parser:"@Ident ';'"`
}
