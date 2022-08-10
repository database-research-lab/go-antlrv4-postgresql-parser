package main

import (
	"fmt"
	"github.com/CC11001100/go-antlrv4-postgresql/parser"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type TreeShapeListener struct {
	*parser.BasePostgreSQLParserListener
}

func NewTreeShapeListener() *TreeShapeListener {
	return new(TreeShapeListener)
}

func (s *TreeShapeListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	fmt.Println(ctx.GetText())
}

func main() {

	sql := "SELECT a,b,a<b AS \"a<b\",a<=b AS \"a<=b\",a=b AS \"a=b\",\n        a>=b AS \"a>=b\",a>b AS \"a>b\",a<>b AS \"a<>b\" FROM bit_table;"

	input := antlr.NewInputStream(sql)
	lexer := parser.NewPostgreSQLLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewPostgreSQLParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.BuildParseTrees = true
	antlr.ParseTreeWalkerDefault.Walk(NewTreeShapeListener(), p.Root())

}
