# samples 追加案（未実装）

実用性と「考察が残る」観点で、現状の `samples` に追加するとよさそうな項目。

## HTTP/ネットワーク
- http client 基本
  - `http.Client` のタイムアウト設定
  - 簡易リトライ
  - context 連携

## DB
- database/sql 基本
  - 接続/クエリ/トランザクション
  - `defer rows.Close()` の定形

## JSON
- encoding/json 実装例
  - Marshal/Unmarshal
  - `omitempty`
  - `Decoder` ストリーム読み取り

## 設定
- config 例
  - env + flag + config ファイル（軽量版）

## エラー
- errors 実装例
  - `errors.Is` / `errors.As`
  - `fmt.Errorf("%w")`
  - カスタムエラー型

## テスト
- testing 例
  - テーブル駆動テスト
  - `t.Run`
  - テストヘルパー
  - `httptest`

## context（実戦）
- HTTP ハンドラで context を渡す
- キャンセルや期限処理

## logging
- 標準 log の構成
- 可能なら `slog` の基本

## time
- Timer/Ticker
- タイムゾーンや期限の注意点

## concurrency 応用
- worker pool
- fan-in / fan-out
- errgroup パターン
