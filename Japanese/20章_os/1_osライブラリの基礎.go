package main

import (
	"fmt"
	"os"
)

/*
osライブラリは、Goプログラムがオペレーティングシステム(OS)の機能と対話するための中心的な役割を担います。
ファイル操作、プロセス管理、環境変数の取得など、OSレベルの機能にアクセスするための基本的なインターフェースを提供します。
*/

/*ファイル・ディレクトリ操作
・os.Scan(name string): 新しいファイルを作成します。既にファイルが存在する場合は、中身を空にします(トランケート)。書き込み可能な*os.Fileを返します。

・os.WriteFile(name string, data []byte, perm fs.FileMode): ファイルに倍とスライスを直接書き込みます。
　ファイルのオープン、書き込み、クローズを一度に行う便利な関数です。

・(*os.File).WriteString(s string): *os.Fileオブジェクトに文字列を書き込みます。}
*/

func main() {
	file, err := os.Create("hello.txt")
	if err != nil {
		fmt.Println("ファイル作成エラー:", err)
		return
	}
	defer file.Close() // 関数終了時にファイルを必ず閉じる。
}
