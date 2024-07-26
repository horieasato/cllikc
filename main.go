package main

import (
	"fmt"
	"syscall"
	"unsafe"
)

var (
	user32           = syscall.NewLazyDLL("user32.dll")
	procFindWindow   = user32.NewProc("FindWindowW")
	procFindWindowEx = user32.NewProc("FindWindowExW")
	procSendMessage  = user32.NewProc("SendMessageW")
)

const (
	WM_LBUTTONDOWN = 0x0201
	WM_LBUTTONUP   = 0x0202
	BM_CLICK       = 0x00F5
)

type HWND uintptr

func findWindow(className, windowName *uint16) HWND {
	hwnd, _, _ := procFindWindow.Call(
		uintptr(unsafe.Pointer(className)),
		uintptr(unsafe.Pointer(windowName)),
	)
	return HWND(hwnd)
}

func findWindowEx(parent, child HWND, className, windowName *uint16) HWND {
	hwnd, _, _ := procFindWindowEx.Call(
		uintptr(parent),
		uintptr(child),
		uintptr(unsafe.Pointer(className)),
		uintptr(unsafe.Pointer(windowName)),
	)
	return HWND(hwnd)
}

func sendMessage(hwnd HWND, msg uint32, wparam, lparam uintptr) {
	procSendMessage.Call(
		uintptr(hwnd),
		uintptr(msg),
		wparam,
		lparam,
	)
}

func main() {
	fmt.Printf("hello\n")
	// 找到目标窗口
	windowName, err := syscall.UTF16PtrFromString("電卓")
	if err != nil {
		panic(err)
	}
	hwnd := findWindow(nil, windowName)
	if hwnd == 0 {
		fmt.Println("找不到目标窗口")
		return
	}

	// 找到窗口内的按钮
	buttonClassName, err := syscall.UTF16PtrFromString("Button")
	if err != nil {
		panic(err)
	}
	buttonHwnd := findWindowEx(hwnd, 0, buttonClassName, nil)
	if buttonHwnd == 0 {
		fmt.Println("找不到按钮")
		return
	}

	// 模拟点击按钮
	sendMessage(buttonHwnd, BM_CLICK, 0, 0)
	fmt.Println("按钮点击模拟成功")
}
