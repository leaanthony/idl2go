package main

import (
	"github.com/alecthomas/participle/v2"
	"github.com/leaanthony/slicer"
	"idl2go/assets"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
)

func parseIDL(data string) error {
	parser := participle.MustBuild(&IDL{}, participle.UseLookahead(4))
	//	parser, err := participle.Build(&IDL{})
	//	if err != nil {
	//		return err
	//	}

	idl := &IDL{}
	err := parser.ParseString("", data, idl)
	if err != nil {
		return err
	}

	err = generateCode(idl)
	if err != nil {
		return err
	}
	//repr.Println(idl, repr.Indent("  "), repr.OmitEmpty(false))
	return nil
}

func generateCode(idl *IDL) error {

	var enums strings.Builder
	var structs strings.Builder
	var forewardDeclarations slicer.StringSlicer
	for _, library := range idl.Libraries {
		packageName := strings.ToLower(library.Name)
		err := os.RemoveAll(packageName)
		if err != nil {
			return err
		}
		err = os.Mkdir(packageName, 0755)
		if err != nil {
			return err
		}
		for _, statement := range library.Statements {
			if statement.InterfaceForewardDecl != "" {
				forewardDeclarations.Add(statement.InterfaceForewardDecl)
			}
			if statement.Enum != nil {
				generateEnum(&enums, statement.Enum)
			}
			if statement.Struct != nil {
				generateStruct(&structs, statement.Struct)
			}
			if statement.Interface != nil {
				//println("Interface")
				err := generateInterface(statement.Interface, packageName, forewardDeclarations)
				if err != nil {
					return err
				}
			}
		}

		err = writeFile("enums.go", &enums, packageName)
		if err != nil {
			return err
		}
		err = writeFile("structs.go", &structs, packageName)
		if err != nil {
			return err
		}

		cmd := exec.Command("go", "fmt", "./"+packageName)
		err = cmd.Run()
		if err != nil {
			return err
		}
		packageBytes := []byte("package " + packageName + "\n")
		err = os.WriteFile(filepath.Join(packageName, "com.go"), append(packageBytes, assets.ComGo...), 0644)
		if err != nil {
			return err
		}
	}

	return nil
}

func generateInterface(declaration *InterfaceDeclaration, packageName string, forewardDecls slicer.StringSlicer) error {

	var text strings.Builder

	text.WriteString(declaration.String(forewardDecls))

	err := writeFile(declaration.Name+".go", &text, packageName)
	if err != nil {
		return err
	}
	return nil
}

func writeFile(name string, s *strings.Builder, packageName string) error {
	println("Generating:", filepath.Join(packageName, name))
	output, err := os.Create(filepath.Join(packageName, name))
	if err != nil {
		return err
	}
	_, err = output.WriteString("package " + packageName + "\n" + s.String())
	if err != nil {
		return err
	}
	err = output.Close()
	if err != nil {
		return err
	}
	return nil

}

func generateStruct(s *strings.Builder, declaration *StructDeclaration) {
	name := declaration.Name
	s.WriteString("type " + name + " struct {\n")
	for _, field := range declaration.Fields {
		s.WriteString("\t" + field.Name + " " + idlTypeMap[field.Type] + "\n")
	}
	s.WriteString("}\n\n")
}

func generateEnum(builder *strings.Builder, enum *EnumDeclaration) {
	if len(enum.Values) == 0 {
		return
	}
	builder.WriteString("\ntype " + enum.Name + " uint32\n\n")
	builder.WriteString("const (\n")
	for idx, e := range enum.Values {
		builder.WriteString("\t" + e.Value + " " + enum.Name + " = " + strconv.Itoa(idx) + "\n")
	}
	builder.WriteString(")\n")
}
