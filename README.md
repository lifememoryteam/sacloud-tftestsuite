# sacloud-tftestsuite

Terraform for さくらのクラウドで記述するHCLファイルのテストスイートです。

## テスト内容

- サーバリソース
  - 提供されているコミットメント、CPU、メモリのペアが正しいかを確認します。

## テストの実行

```console
$ bazel test //tf/... --test_env=SAKURACLOUD_ZONE=is1b --test_env=SAKURACLOUD_ACCESS_TOKEN=token --test_env=SAKURACLOUD_ACCESS_TOKEN_SECRET=token_secret
```

```console
$ cat /private/var/tmp/<snipped>/test.log
exec ${PAGER:-/usr/bin/less} "$0" || exit 1
Executing tests from //tf:test
-----------------------------------------------------------------------------
testing: tf/hoge.tf
testing: tf/test.tf
validserver02: not serviced in 100Core/100GB
test failed
```