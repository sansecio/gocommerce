package phpcfg

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/node/expr"
	"github.com/z7zmey/php-parser/node/name"
	"github.com/z7zmey/php-parser/node/scalar"
	"github.com/z7zmey/php-parser/node/stmt"
	"github.com/z7zmey/php-parser/parser"
)

// Finds the "return" statement under the root node.
func findReturnNode(root *node.Root) *stmt.Return {
	for _, s := range root.Stmts {
		if s, ok := s.(*stmt.Return); ok {
			return s
		}
	}

	return nil
}

// Finds the "array" or "short array" expression under the "return" statement.
func getReturnedArray(ret *stmt.Return) node.Node {
	ex := ret.Expr

	switch ex := ex.(type) {
	case *expr.Array:
		return ex
	case *expr.ShortArray:
		return ex
	}

	return nil
}

// Removes single quotes if present.
func unquote(s string) string {
	if s[0] == '\'' && s[len(s)-1] == '\'' {
		return s[1 : len(s)-1]
	} else {
		return s
	}
}

// Converts the node into a string.
func nodeToString(n node.Node) string {
	ret := fmt.Sprintf("<unknown node %T>", n)

	switch n := n.(type) {
	case *scalar.String:
		ret = n.Value
	case *scalar.Dnumber:
		ret = n.Value
	case *scalar.Lnumber:
		ret = n.Value
	case *name.NamePart:
		ret = n.Value
	case *expr.ConstFetch:
		switch n := n.Constant.(type) {
		case *name.Name:
			parts := make([]string, len(n.Parts))
			for _, part := range n.Parts {
				parts = append(parts, nodeToString(part))
			}
			ret = strings.Join(parts, "")
		default:
			ret = "<unknown constant>"
		}
	}

	return unquote(ret)
}

func getArrayItems(array node.Node) ([]node.Node, error) {
	switch array := array.(type) {
	case *expr.Array:
		return array.Items, nil
	case *expr.ShortArray:
		return array.Items, nil
	default:
		return nil, errors.New("node is not an array")
	}
}

func isArray(n node.Node) bool {
	switch n.(type) {
	case *expr.Array:
		return true
	case *expr.ShortArray:
		return true
	default:
		return false
	}
}

// Recursively parses an array and populates the "out" map with parsed values.
func parseArray(path []string, array node.Node, out map[string]string) error {
	ord := 0

	items, err := getArrayItems(array)
	if err != nil {
		return err
	}

	for _, n := range items {
		if item, ok := n.(*expr.ArrayItem); ok {
			if item.Val == nil {
				// Items with empty values are caused by
				// trailing commas after the last array item:
				//
				// array(
				//  'date' => 'Tue, 08 Aug 2017 20:08:01 +0000',
				// )
				continue
			}

			itemKey := strconv.Itoa(ord)
			if item.Key != nil {
				itemKey = nodeToString(item.Key)
			} else {
				ord += 1
			}

			itemPath := append(path, itemKey)

			if isArray(item.Val) {
				if err := parseArray(itemPath, item.Val, out); err != nil {
					return err
				}
			} else {
				out[strings.Join(itemPath, ".")] = nodeToString(item.Val)
			}
		} else {
			return errors.New("array item is not an expr.ArrayItem")
		}
	}

	return nil
}

func Parse(src []byte) (map[string]string, error) {
	p, err := parser.NewParser(src, "7.0") // FIXME: hardcoded version
	if err != nil {
		return nil, err
	}

	p.Parse()
	if len(p.GetErrors()) != 0 {
		err0 := p.GetErrors()[0]
		return nil, errors.New(fmt.Sprintf("PHP parse failed: %v", err0.Msg))
	}

	nodes := p.GetRootNode()

	retNode := findReturnNode(nodes.(*node.Root))
	if retNode == nil {
		return nil, errors.New("'return' node not found")
	}

	itemsArray := getReturnedArray(retNode)
	if itemsArray == nil {
		return nil, errors.New("'return [...] | array(...)' expected")
	}

	ret := make(map[string]string)
	if err := parseArray([]string{"root"}, itemsArray, ret); err != nil {
		return nil, err
	}

	return ret, nil
}
