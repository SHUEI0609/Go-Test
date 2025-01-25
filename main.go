package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// HTMLテンプレートをパース
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "テンプレートを読み込めませんでした", http.StatusInternalServerError)
		return
	}

	// テンプレートに渡すデータ
	data := struct {
		Title      string
		Background string
	}{
		Title:      "SHUEI's PORTFOLIO",
		Background: "static/image/Thelaat.png", // 背景画像のURL（静的ファイル用）
	}

	// HTMLをレンダリング
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "テンプレートのレンダリングに失敗しました", http.StatusInternalServerError)
	}
}

func main() {
	// 静的ファイルを提供する設定
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// ルートハンドラ
	http.HandleFunc("/", handler)

	// サーバーを起動
	fmt.Println("サーバーを起動中... http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("サーバーエラー:", err)
	}
}
