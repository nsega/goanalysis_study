package channel_close

import (
	"fmt"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/buildssa"
)

const doc = "channel_close is ..."

// Analyzer is ...
var Analyzer = &analysis.Analyzer{
	Name: "channel_close",
	Doc:  doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		buildssa.Analyzer,
	},
}

func run(pass *analysis.Pass) (interface{}, error) {
	s := pass.ResultOf[buildssa.Analyzer].(*buildssa.SSA)
	for _, f := range s.SrcFuncs {
		fmt.Println(f)
		for _, b := range f.Blocks {
			fmt.Println("\tBlock %d\n", b.Index)
			for _, instr := range b.Instrs {
				fmt.Println("\t\t%[1]v(%[1]p)\n", instr)
				for _, v := range instr.Operands(nil) {
					if v != nil {
						fmt.Printf("\t\t\t%[1]T\t%[1]v(%[1]p)\n", *v)
					}
				}
			}
		}
	}
	return nil, nil
}
