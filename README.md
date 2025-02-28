# Port Scanner

シンプルな並行処理対応のポートスキャナーです。指定されたIPアドレスに対して、ポート1から1024までのTCPポートスキャンを実行します。

[English](README-en.md)

## 機能

- 指定したIPアドレスに対するTCPポートスキャン
- 並行処理による高速なスキャン実行
- よく知られているサービスの自動識別（FTP, SSH, HTTP など）
- スキャン結果のログファイル出力


## インストール

```bash
git clone https://github.com/kaedeek/port_scan-go.git
cd port_scan-go
```

## 使用方法

```bash
go run main.go -ip <target_ip>
```

### 例

```bash
go run main.go -ip 000.000.0.0
```

スキャン結果は `scan.log` ファイルに保存されます。

## スキャン結果

スキャン結果は以下の形式でログファイルに出力されます：

```
2024/03/XX HH:MM:SS Starting port scan on 000.000.0.0 (ports 1-1024)
2024/03/XX HH:MM:SS Port 22 is open (SSH)
2024/03/XX HH:MM:SS Port 80 is open (HTTP)
2024/03/XX HH:MM:SS Port 443 is open (HTTPS)
2024/03/XX HH:MM:SS Scan completed
```

## 注意事項

- このツールは教育目的で作成されています
- ポートスキャンを実行する前に、対象システムの所有者から許可を得てください
- 不適切な使用は法的問題を引き起こす可能性があります

## ライセンス

MIT License with additional restrictions.
Copyright (c) 2025 kaedeek

詳細は LICENSE ファイルをご確認ください。