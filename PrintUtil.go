package main
import (
    "fmt"
    "os/exec"
)
// CopyImageFromDockerRepository return script path
func CopyImageFromDockerRepository(imageSource string) (string) {
    fmt.Println("Started Downloading Image")
    cmd, err := exec.Command("/bin/sh", "copy_image.sh", imageSource, "/Users/mam/Documents/test/sample").Output()
    if err != nil {
    fmt.Printf("error %s", err)
    }
    output := string(cmd)
    fmt.Println("Successfully Downloading Image")
    return output
}
func main() {
    CopyImageFromDockerRepository("docker://quay.io/buildah/stable")
}