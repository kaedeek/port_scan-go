# Port Scanner

シンプルな並行処理対応のポートスキャナーです。指定された IP アドレスに対して、ポート 0 から 65535 までの TCP ポートスキャンを実行します。

[English](README-en.md)

## 機能

- **指定した IP アドレスに対する TCP ポートスキャン**
- **並行処理による高速なスキャン実行**
- **よく知られているサービスの自動識別（FTP, SSH, HTTP など）**

> [!CAUTION]
> 禁止事項：
>
> - このツールは教育目的で作成されています
> - ポートスキャンを実行する前に、対象システムの所有者から許可を得てください
> - 不適切な使用は法的問題を引き起こす可能性があります

## スキャン結果

スキャン結果は以下の形式で出力されます

```
[+] Please enter an IP address to scan:
000.000.0.0
[+] Port scan started | 000.000.0.0 (Press 'e' to stop)


[-] 80 | HTTP Open
[-] 53 | DNS Open
[-] 27998 | Unknown Open
[-] 37444 | Unknown Open
[-] 37443 | Unknown Open
[-] 49652 | Unknown Open
[-] 49653 | Unknown Open
[+] Scan completed :)


[+] Press Enter to exit...
```

## EXE FILE LIST

[EXE FILE](exe/portscanner.exe)

## メッセージ一覧

| メッセージ     | 説明           |
| -------------- | -------------- |
| `Unknown Open` | サービス名なし |
