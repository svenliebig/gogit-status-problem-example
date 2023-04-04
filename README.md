# reproduce

```bash
git clone https://github.com/svenliebig/gogit-status-problem-example.git
cd gogit-status-problem-example
git status
go run main.go
```

You'll see that both implementation return that no changes are made.

Let's run scripts.sh, this will create a `file.txt` in the directory `sub`, and add the
directory to the `.gitignore` in the root.

```bash
./scripts.sh
git status
go run main.go
```

`git status` now shows that we have a change in `.gitignore`, while the implemenation of
gogit also shows the `file.txt` in their status implementation.

This behauvior is produced in [this dir.go file](https://github.com/go-git/go-git/blob/master/plumbing/format/gitignore/dir.go#L50),
where we iterate over all directories regardless if they where previously ignored or not.
