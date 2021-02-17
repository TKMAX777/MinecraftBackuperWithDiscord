# GoogleDriveBackuper
## 概要
GoogleDriveにディレクトリを毎日バックアップしつつ、その進捗をDiscordに上げるプログラム。

## 目次
<!-- TOC -->

- [GoogleDriveBackuper](#googledrivebackuper)
    - [概要](#概要)
    - [目次](#目次)
    - [使い方](#使い方)

<!-- /TOC -->

## 使い方
- git clone
- GoogleDriveAPIのcredentioal.jsonをレポジトリルートに配置
- 次のsettings.jsonを同位置に配置

```json
{
    "DiscordHookURI": "https://discord.com/api/webhooks/****",
    "FilePath": "/path/to/minecraft/world/dir",
    "UploadTime": {
        "Hour": 4,
        "Minute": 0
    },
    "DriveParentDir": "Google Drive ParentDir's ID",
    "MessageInfo": {
        "AvaterURI": "",
        "UserName": "GoogleDriveBackup - Info -"
    },
    "ErrorMessageInfo": {
        "AvaterURI": "",
        "UserName": "GoogleDriveBackup - Error - "
    }
}
```

- Compile and run
