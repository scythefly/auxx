package utility

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"

	"github.com/spf13/cobra"
)

const (
	numberExpr     = `^(\s)*第(\s)*[0-9]+(\s)*章 [^\n]*\n$`
	chineseExpr    = `^(\s)*第[一二三四五六七八九十百千零]+章[^\n]*\n$`
	numberOnlyExpr = `^(\s)*[0-9]+(\s)+[^\n]+\n$`
)

type splitFileOption struct {
	splitBySize     bool
	splitByChapter  bool
	splitSize       int64
	splitChapterCnt int
	splitExprType   int
	expr            string
}

type ut struct {
	regs []*regexp.Regexp
	reg  *regexp.Regexp
}

var opt *splitFileOption
var _ut = newUT()

func (u *ut) init() {
	u.regs = append(u.regs, regexp.MustCompile(numberExpr))
	u.regs = append(u.regs, regexp.MustCompile(numberOnlyExpr))
	u.regs = append(u.regs, regexp.MustCompile(chineseExpr))
}

func newUT() *ut {
	u := &ut{}
	u.init()

	return u
}

func (u *ut) match(line []byte) bool {
	for idx := range u.regs {
		if u.regs[idx].Match(line) {
			return true
		}
	}
	return false
}

func (u *ut) match1(line []byte) bool {
	return u.reg.Match(line)
}

func newsplitFileCommand() *cobra.Command {
	opt = &splitFileOption{}
	cmd := &cobra.Command{
		Use:   "split-file",
		Short: "Split file",
		RunE: func(cmd *cobra.Command, args []string) error {
			return splitFile(opt, args)
		},
	}
	flags := cmd.Flags()
	flags.BoolVar(&opt.splitByChapter, "split-by-chapter", true, "split file by chapter")
	flags.BoolVarP(&opt.splitBySize, "split-by-size", "c", false, "split file by size")
	flags.Int64VarP(&opt.splitSize, "split-size", "s", 512*1024, "split by size")
	flags.IntVar(&opt.splitChapterCnt, "split-chapter", 100, "split by chapter")
	flags.StringVar(&opt.expr, "expr", "", "regexp expression")

	if opt.expr != "" {
		_ut.reg = regexp.MustCompile(opt.expr)
	}
	return cmd
}

func splitFile(opt *splitFileOption, args []string) error {
	for _, fileName := range args {
		str := fileName
		g.Go(func() error {
			return split(opt, str)
		})
	}
	return g.Wait()
}

func split(opt *splitFileOption, fileName string) error {
	f, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer f.Close()
	if opt.splitBySize {
		var idx int = 1
		for {
			str := fmt.Sprintf("%s_%03d.txt", fileName, idx)
			ff, err := os.Create(str)
			if err != nil {
				return err
			}
			if _, err = io.CopyN(ff, f, opt.splitSize); err != nil {
				if err == io.EOF {
					err = nil
				}
				ff.Close()
				return err
			}
			ff.Close()
			idx++
		}
	} else {
		fn := _ut.match
		if opt.expr != "" {
			fn = _ut.match1
		}
		r := bufio.NewReader(f)
		var idx, cnt int
		str := fmt.Sprintf("%s_%03d.txt", fileName, idx)
		ff, err := os.Create(str)
		if err != nil {
			return err
		}
		for {
			line, rerr := r.ReadBytes('\n')
			if rerr != nil && rerr != io.EOF {
				return rerr
			}
			if fn(line) {
				fmt.Printf("[%d] - %s", cnt, line)
				cnt++
				if cnt >= opt.splitChapterCnt {
					ff.Close()
					idx++
					str := fmt.Sprintf("%s_%03d.txt", fileName, idx)
					fmt.Println(str)
					if ff, err = os.Create(str); err != nil {
						return err
					}
					cnt = 0
				}
			}
			if _, err = ff.Write(line); err != nil {
				return err
			}
			if rerr == io.EOF {
				ff.Close()
				return nil
			}
		}
	}
}
