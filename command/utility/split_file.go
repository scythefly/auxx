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
	numberExpr  = `^(\s)*第(\s)*[0-9]+(\s)*章 [^\n]*\n$`
	chineseExpr = `^(\s)*第[一二三四五六七八九十百千零]+章[^\n]*\n$`
)

type splitFileOption struct {
	splitBySize     bool
	splitByChapter  bool
	splitSize       int64
	splitChapterCnt int
	splitExprType   int
	expr            string
}

var opt *splitFileOption

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
	flags.IntVarP(&opt.splitExprType, "split-expr-type", "e", 0, "expr\n\t0: 第123章\n\t1: 第一百二十三章")
	flags.StringVar(&opt.expr, "expr", "", "regexp expression")
	return cmd
}

func splitFile(opt *splitFileOption, args []string) error {
	if opt.expr == "" {
		if opt.splitExprType == 1 {
			opt.expr = chineseExpr
		} else {
			opt.expr = numberExpr
		}
	}
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
		r := bufio.NewReader(f)
		reg := regexp.MustCompile(opt.expr)
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
			if reg.Match(line) {
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
