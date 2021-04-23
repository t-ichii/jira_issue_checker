# jira_issue_checker

## how to build

```bash
make build
```

## execute example

事前に環境変数を設定すること。環境変数は [`environment/environment.go`](https://github.com/t-ichii/jira_issue_checker/blob/f73dac3e2c0c8d49162637e2dea56ab3e70b9a4a/environment/environment.go#L10-L12) を参照すること

```bash
# rank checker mode
./bin/jira_issue_checker -m rank_checker -b 710 -s 3933
```
