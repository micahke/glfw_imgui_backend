# GLFW IMGUI Backend

This package is intended for those using GLFW and imgui with Go. It is meant to bridge the gap between `go-gl/glfw` and `imgui-go`. This means you can retain control over your `glfw.Window` object without handing over control to `imgui`. I found this file in an old commit from `go-gl`. Helpful for those building applications with OpenGL in Go.

![](/doc/screenshot.png)

## Usage

For a full example, check out this [example](/example) in the project directory.

To import:

```go
import "github.com/micahke/glfw_imgui_backend"
// OR
import backend "github.com/glfw_imgui_backend"
```

ImGUI can now be used in a similar syntax to C++. To link your `glfw.Window` with `imgui`, call:

```go
// make Window
window, err := glfw.CreateWindow(960, 540, "Hello, world!", nil, nil)
if err != nil {
  panic("Error creating window")
}
window.MakeContextCurrent()

// Create imgui IO context
context := imgui.CreateContext(nil)
defer context.Destroy()

// Link the GLFW and Imgui
impl := gui.ImguiGlfw3Init(window)
defer impl.Shutdown()
```
