package internal

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/go-clang/clang-v15/clang"
)

var (
	rTypePattern = regexp.MustCompile(`^/// ([a-zA-Z]+).+`)
)

type MacroGroup struct {
	baseDefine

	Defines []*MacroDefine
}

func (g *MacroGroup) Append(d *MacroDefine) {
	g.Defines = append(g.Defines, d)
}

type MacroDefine struct {
	baseDefine

	Token   string
	RefType *MacroGroup
}

func (m MacroDefine) String() string {
	buff := bytes.NewBufferString(m.Name)

	if m.RefType != nil {
		fmt.Fprintf(buff, " %s", m.RefType.Name)
	} else {
		fmt.Fprintf(buff, " byte")
	}

	fmt.Fprintf(buff, " = %s", m.Token)
	if len(m.Comments) > 0 {
		fmt.Fprintf(buff, " // %s", strings.Join(m.Comments, " "))
	}

	buff.WriteString("\n")

	return buff.String()
}

func (e *entry) ParseMacro(cursor *clang.Cursor, prefix ...string) (*MacroDefine, error) {
	macroName := cursor.Spelling()

	for _, p := range prefix {
		if strings.HasPrefix(macroName, p) {
			goto PARSE
		}
	}

	return nil, nil

PARSE:
	define := MacroDefine{}
	define.Name = macroName

	tu := cursor.TranslationUnit()

	for _, token := range tu.Tokenize(cursor.Extent()) {
		spelling := tu.TokenSpelling(token)
		if spelling != macroName {
			define.Token = spelling
		}
	}

	loc := cursor.Location()

	file, line, _, _ := loc.FileLocation()
	if e.files == nil {
		e.files = make(map[string]*os.File)
	}
	f, exist := e.files[file.Name()]

	if !exist {
		var err error
		f, err = os.OpenFile(file.Name(), os.O_RDONLY, os.ModePerm)
		if err != nil {
			return nil, err
		}
		e.files[file.Name()] = f
	} else if _, err := f.Seek(0, 0); err != nil {
		return nil, err
	}

	cataList := strings.Split(define.Name, "_")
	cataName := strings.Join(cataList[:len(cataList)-1], "_")

	reverse := 1

	var rType string

	if e.defineType == nil {
		e.defineType = map[string]string{}
		reverse += 3
	} else if rType, exist = e.defineType[cataName]; !exist || rType == "" {
		reverse += 3
	}

	if e.defineCache == nil {
		e.defineCache = make(map[string]*MacroGroup)
	}

	seeker := bufio.NewScanner(f)
	currLine := 0

	var rTypeComments []string

	for seeker.Scan() {
		currLine++

		if currLine < int(line)-reverse {
			continue
		}

		if currLine < int(line)-1 {
			c := seeker.Text()

			rTypeComments = append(rTypeComments, c)

			if rType == "" {
				m := rTypePattern.FindStringSubmatch(c)
				if len(m) > 1 {
					rType = m[1]
				}
			}

			continue
		}

		if comment := strings.TrimSpace(
			strings.ReplaceAll(seeker.Text(), "/", ""),
		); comment != "" {
			define.Comments = append(define.Comments, comment)
		}
		break
	}

	if err := seeker.Err(); err != nil {
		return nil, err
	}

	if rType == "" {
		return nil, errors.New("define ref type not found")
	}

	e.defineType[cataName] = rType
	g, exist := e.defineCache[rType]
	if !exist {
		g = &MacroGroup{
			baseDefine: baseDefine{
				Name:     rType,
				Comments: rTypeComments,
			},
		}
		e.defineCache[rType] = g
		g.trimComments()
	}
	g.Defines = append(g.Defines, &define)
	define.RefType = g
	define.trimComments()

	return &define, nil
}
