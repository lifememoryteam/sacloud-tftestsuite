# sacloud-tftestsuite

Terraform for さくらのクラウドで記述するHCLファイルのテストスイートです。

## テスト内容

- サーバリソース
  - 提供されているコミットメント、CPU、メモリのペアが正しいかを確認します。

## 起動

```console
$ bazel test //tf/... --test_env=SAKURACLOUD_ZONE=is1b --test_env=SAKURACLOUD_ACCESS_TOKEN=token --test_env=SAKURACLOUD_ACCESS_TOKEN_SECRET=token_secret
```
