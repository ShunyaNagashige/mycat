package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

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
		//１行分を出力する
		fmt.Println(bsf.Text())
	}

	//纏めてエラー処理をする
	if err := bsf.Err(); err != nil {
		return err
	}

	return nil
}

func main() {
	//オプションの設定
	flag.Parse()

	for _, fn := range flag.Args() {
		if err := readFile(fn); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}