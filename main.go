package main

import (
	"bufio" //I/Oのバッファリング機能を提供するパッケージ
	"fmt"   //標準出力へのアウトプットに使用
	"io/ioutil"
	"os" //主にここでは標準入力用
	"strings"

	"github.com/fatih/color"
)

func main() {
	//対話型で「docker-compose.yml」を生成
	fmt.Printf("指定したディレクトリへの「docker-compose.yml」ファイルの生成コマンドです。\n")
	fmt.Printf("入力は全て半角英数字でお願いします。\nまた、保存先ディレクトリは作成しておいてください。\n")
	userInput, path := dialog() //ymlファイルに記載する内容の取得
	if userInput != "キャンセル" {
		fc, err := os.Create(path + "/docker-compose.yml") //指定ディレクトリにymlファイル生成
		//エラーハンドリング
		if err != nil {
			color.Red("Could not read template")
			fmt.Println(err)
			os.Exit(1)
		}
		//処理終了時にファイル閉じる
		defer fc.Close()
		//ファイルに書き込み
		fc.WriteString(userInput)
	} else {
		fmt.Println("ファイルの作成を中止しました。")
	}
}

func dialog() (ymlInput, path string) {
	f, err := os.Open("docker-tmp/docker-compose-tmp.yml") // ファイルをOpenする
	// 読み取り時の例外処理
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()                      //この関数の終了時にファイルを閉じる
	b, _ := ioutil.ReadAll(f)            //ファイルの内容を取得する。「_」はブランク変数。
	var composeYml = string(b)           //文字列型にキャスト
	fmt.Printf("データベース名を入力してください。\n")    //ガイドメッセージを標準出力にアウトプット
	uInput := bufio.NewScanner(os.Stdin) //標準入力の内容を取得
	var inputTxt string                  //数字がきた時の型推測がめんどくさそうなので明示的に指定しておく
	for uInput.Scan() {                  //どうやらユーザの入力を待つにはループさせる必要があるらしい。
		inputTxt = uInput.Text()
		if len(inputTxt) != 0 {
			break
		}
	}
	composeYml = strings.Replace(composeYml, "exampledb", inputTxt, -1)

	fmt.Printf("コンテナ名を入力してください。\n")         //ガイドメッセージを標準出力にアウトプット
	uInputCon := bufio.NewScanner(os.Stdin) //標準入力の内容を取得
	var inputTxtCon string                  //数字がきた時の型推測がめんどくさそうなので明示的に指定しておく
	for uInputCon.Scan() {                  //どうやらユーザの入力を待つにはループさせる必要があるらしい。
		inputTxtCon = uInputCon.Text()
		if len(inputTxtCon) != 0 {
			break
		}
	}
	composeYml = strings.Replace(composeYml, "examplecontainername", inputTxtCon, -1)

	fmt.Printf("「docker-compose.yml」の保存先をフルパスで入力してください。\n") //ガイドメッセージを標準出力にアウトプット
	uInputPlace := bufio.NewScanner(os.Stdin)               //標準入力の内容を取得
	var inputTxtPlace string                                //数字がきた時の型推測がめんどくさそうなので明示的に指定しておく
	for uInputPlace.Scan() {                                //どうやらユーザの入力を待つにはループさせる必要があるらしい。
		inputTxtPlace = uInputPlace.Text()
		if len(inputTxtPlace) != 0 {
			break
		}
	}

	fmt.Printf("コンテナ名:" + inputTxtCon + "\n")
	fmt.Printf("データベース名:" + inputTxt + "\n")
	fmt.Printf("ファイル保存先:" + inputTxtPlace + "/docker-compose.yml\n")
	//yesかnoにて生成確認を行う。
	fmt.Printf("で「docker-compose.yml」を生成します。よろしいですか？[y/n]:\n")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		i := scanner.Text()

		if i == "Y" || i == "y" {
			break
		} else if i == "N" || i == "n" {
			return "キャンセル", "キャンセル"
		} else {
			fmt.Println("yかnで答えてください。")
			fmt.Printf("コンテナ名:" + inputTxtCon + "\n")
			fmt.Printf("データベース名:" + inputTxt + "\n")
			fmt.Printf("ファイル保存先:" + inputTxtPlace + "/docker-compose.yml\n")
			fmt.Printf("で「docker-compose.yml」を生成します。よろしいですか？[y/n]:\n")
		}
	}

	return composeYml, inputTxtPlace
}
