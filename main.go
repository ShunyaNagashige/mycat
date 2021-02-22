package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

var isFlagNSelected bool
var lineNum int = 1

func init(){
	usage:="行番号を表示させたい場合は、このオプションを選択してください。"
	flag.BoolVar(&isFlagNSelected,"n",false,usage)
}

func main() {
	//オプションの値を、コマンドから取得し設定する
	flag.Parse()

	for _, fn := range flag.Args() {
		if err := readFile(fn); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}

func readFile(fn string) error {
	//sf=sourcefile(送り元、つまり読み込むファイル),fn=filename
	sf, err := os.Open(fn)
	defer sf.Close()
	if err != nil {
		return err
	}
	bsf := bufio.NewScanner(sf)

	//１行ずつ読み込んで繰り返す
	for bsf.Scan() {
		//オプションnを選択した場合、行番号を表示
		if(isFlagNSelected){
			fmt.Print(lineNum,": ")
			lineNum++
		}
		//１行分を出力する
		fmt.Println(bsf.Text())
	}

	//纏めてエラー処理をする
	if err := bsf.Err(); err != nil {
		return err
	}

	return nil
}
