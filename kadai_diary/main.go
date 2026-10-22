package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type Entry struct {
	Timestamp time.Time // 登録日時
	Title     string    // タイトル
	Content   string    // 本文
}

func saveDiary(entry []Entry) any {
	fmt.Println("save開始")
	jsonEntry, err := json.Marshal(entry)
	if err != nil {
		fmt.Println("Json変換に失敗しました：", err)
		return err
	}

	err = os.WriteFile("kadai_diary.json", jsonEntry, 0644)
	if err != nil {
		fmt.Println("ファイルへの書き込みに失敗しました：", err)
		return err
	}
	fmt.Println("ファイル書き込み完了")
	return nil
}

// 課題には引数設定があるが、使用しなかった、、？
func loadDiary() ([]Entry, error) {
	jsonData, readErr := os.ReadFile("kadai_diary.json")
	if os.IsNotExist(readErr) {
		// 初回はファイルがないため、空のスライスを返す
		return []Entry{}, nil
	}

	if readErr != nil {
		fmt.Println("Jsonデータ読み込みに失敗しました：", readErr)
		return nil, readErr
	}

	var sliceData []Entry
	sliceErr := json.Unmarshal(jsonData, &sliceData)
	if sliceErr != nil {
		fmt.Println("Jsonデータの変換に失敗しました：", sliceErr)
		return nil, sliceErr
	}

	return sliceData, nil
}

func main() {
	// コマンドライン引数の数をチェック
	if len(os.Args) < 2 {
		fmt.Println("使い方：go run main.go [add|list] ...")
		return
	}

	// 1番目の引数（コマンド）を取得
	command := os.Args[1]

	// コマンドの内容に応じた処理
	switch command {
	case "add": // 日記エントリーの記録
		// addコマンドの処理
		if len(os.Args) < 4 {
			fmt.Println("addコマンドを使用する際はタイトルと本文も入力してください")
			return
		} else {
			title := os.Args[2]
			content := os.Args[3]
			fmt.Println("タイトル：", title, "本文：", content)
			data, err := loadDiary()
			if err != nil {
				fmt.Println("ファイル読み込みに失敗しました：", err)
				return
			}
			newEntry := Entry{
				Timestamp: time.Now(),
				Title:     title,
				Content:   content,
			}
			data = append(data, newEntry)
			saveDiary(data)
		}

	case "list": // 日記の一覧表示
		// listコマンドの処理
		data, err := loadDiary()
		if err != nil {
			fmt.Println("ファイル読み込みに失敗しました：", err)
			return
		}

		for _, v := range data {
			fmt.Println("--- 日記データ ---")
			formatTimestamp := v.Timestamp.Format("2006-01-02 15:04:05")
			fmt.Println("日時：", formatTimestamp)
			fmt.Println("タイトル：", v.Title)
			fmt.Printf("本文：\n%s\n", v.Content)
		}
	default:
		fmt.Println("不明なコマンドです：", command)
	}
}
