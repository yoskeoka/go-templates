# go-templates

Go templates for gonew command.
About `gonew`, read [the article](https://go.dev/blog/gonew) about this command.

## cli/subcommand

Replace `github.com/yourname/mycli` with your repository path.

```sh
gonew github.com/yoskeoka/go-templates/cli/subcommand github.com/yourname/mycli 
cd mycli
git init
git add .
git commit -m "initial commit"
make build
./bin/example hello -msg "I'm happy with subcommand template!"
```
