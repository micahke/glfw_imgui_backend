package main

import (
	"fmt"
	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/inkyblackness/imgui-go"
	backend "github.com/micahke/glfw_imgui_backend"
	"runtime"
)

func init() {
	runtime.LockOSThread()
}

func main() {

	// Initialize GLFW through go-gl/glfw
	if err := glfw.Init(); err != nil {
		panic("Error initializing GLFW")
	}
	defer glfw.Terminate()

	// GLFW setup
	glfw.WindowHint(glfw.ContextVersionMajor, 3)
	glfw.WindowHint(glfw.ContextVersionMinor, 3)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	// Initialize window through go-gl/glfw
	window, win_err := glfw.CreateWindow(920, 540, "Hello, world!", nil, nil)
	if win_err != nil {
		panic("Error creating window")
	}

	window.MakeContextCurrent()
	glfw.SwapInterval(1)

	if err := gl.Init(); err != nil {
		panic("Error initializing OpenGL")
	}

	// Initialize imgui
	context := imgui.CreateContext(nil)
	defer context.Destroy()

	// KEY: link imgui context with GLFW window context
	impl := backend.ImguiGlfw3Init(window)
	defer impl.Shutdown()

	showDemoWindow := false
	showAnotherWindow := false
	counter := 0

	for !window.ShouldClose() {
		impl.NewFrame()

		{
			imgui.Text("Hello, world!")

			imgui.Checkbox("Demo Window", &showDemoWindow)
			imgui.Checkbox("Another Window", &showAnotherWindow)

			if imgui.Button("Button") {
				counter++
			}
			imgui.SameLine()
			imgui.Text(fmt.Sprintf("counter = %d", counter))

		}

		if showAnotherWindow {
			imgui.BeginV("Another Window", &showAnotherWindow, 0)
			imgui.Text("Hello from another window!")
			if imgui.Button("Close Me") {
				showAnotherWindow = false
			}
			imgui.End()
		}

		if showDemoWindow {
			imgui.ShowDemoWindow(&showDemoWindow)
		}

		gl.Clear(gl.COLOR_BUFFER_BIT)

		imgui.Render()
		impl.Render(imgui.RenderedDrawData())
		window.SwapBuffers()
		glfw.PollEvents()

	}

}
