This project is about preparing modules, but also includes small projects related to using these modules. If you want to examine / run it locally, move the folder names `usage-of-toolkit` to a another folder. Afterwards, add the folder with the remaining files into the `usage-of-toolkit`. This way, you will be able to operate it without any problems.

## Modules

If you look at explorer, you have a toolkit folder with a Go module file and you have an app folder with the Go module file. Let`s look at workspaces:

So inside the toolkit-project folder, workspace let me do is work with multiple projects, all referring to local instances of particular modules.

So the command to make this happen

```
go work init toolkit app
```