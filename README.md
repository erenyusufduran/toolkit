## Modules

If you look at explorer, you have a toolkit folder with a Go module file and you have an app folder with the Go module file. Let`s look at workspaces:

So inside the toolkit-project folder, workspace let me do is work with multiple projects, all referring to local instances of particular modules.

So the command to make this happen

```
go work init toolkit app
```