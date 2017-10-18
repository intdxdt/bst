package bst

import (
    "strings"
    "github.com/intdxdt/str"
)

//PrintBST - Print BST
func PrintBST(node *Node, keyfn func(interface{}) string) string {
    pretty, _, _ := printer(node, keyfn)
    return strings.Join(pretty, "\n")
}

//printer - print BST tree
func printer(node *Node, keyfn func(interface{}) string) ([]string, int, int) {
    //base case
    if IsNil(node) {
        return []string{}, 0, 0
    }

    var label = keyfn(node.Key)

    lstr, lpos, lwidth := printer(node.Left, keyfn)
    rstr, rpos, rwidth := printer(node.Right, keyfn)

    var mpos = max(rpos + lwidth - lpos + 1, len(label), 2)

    var pos = lpos + mpos / 2
    var width = lpos + mpos + rwidth - rpos
    var lines = make([]string, 0)

    for len(lstr) < len(rstr) {
        lstr = append(lstr, str.Repeat(" ", lwidth))
    }

    for len(rstr) < len(lstr) {
        rstr = append(rstr, str.Repeat(" ", rwidth))
    }

    bln := ((mpos - len(label)) % 2 == 1 && (node.Parent != nil) &&
            (node == node.Parent.Left) && (len(label) < mpos))
    if bln {
        label += "."
    }

    label = str.Center(label, mpos, ".")

    if first(label) == "." {
        label = " " + label[1:]
    }
    if last(label) == "." {
        label = label[:len(label) - 1] + " "
    }

    lines = append(lines,
        str.Repeat(" ", lpos) + label + str.Repeat(" ", (rwidth - rpos)),
    )
    lines = append(lines,
        str.Repeat(" ", lpos) + "/" + str.Repeat(" ", (mpos - 2)) + "\\" +
            str.Repeat(" ", (rwidth - rpos)),
    )

    for i := range lstr {
        l, r := lstr[i], rstr[i]
        lr := l + str.Repeat(" ", (width - lwidth - rwidth)) + r
        lines = append(lines, lr)
    }
    return lines, pos, width
}

//first str
func first(str string) string {
    return string(str[0])
}

//last str
func last(str string) string {
    return string(str[len(str) - 1])
}

// max value
func max(a ...int) int {
    var m = a[0]
    for _, v := range a {
        if v > m {
            m = v
        }
    }
    return m
}
