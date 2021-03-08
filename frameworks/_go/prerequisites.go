package _go

import "os/exec"

func Pre(root string) {
	fmt(root)
}

func fmt(root string) {
	exec.Command("go", "fmt", "./...").Output()
	exec.Command("go", "mod", "tidy").Output()

}
