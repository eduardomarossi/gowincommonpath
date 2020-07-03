package main

import (
	"syscall"
	"unsafe"
	"unicode/utf16"
	"fmt"
)

var CSIDL_APPDATA                 = 0x001A     // Application Data, new for NT4 
var CSIDL_COMMON_APPDATA          = 0x0023     // All Users\Application Data 
var CSIDL_COMMON_DOCUMENTS        = 0x002e     // All Users\Documents 
var CSIDL_DESKTOP                 = 0x0010     // C:\Documents and Settings\username\Desktop 
var CSIDL_FONTS                   = 0x0014     // C:\Windows\Fonts 
var CSIDL_LOCAL_APPDATA           = 0x001C     // non roaming, user\Local Settings\Application Data 
var CSIDL_MYMUSIC                 = 0x000d     // "My Music" folder 
var CSIDL_MYPICTURES              = 0x0027     // My Pictures, new for Win2K 
var CSIDL_PERSONAL                = 0x0005     // My Documents 
var CSIDL_PROGRAM_FILES_COMMON    = 0x002b     // C:\Program Files\Common 
var CSIDL_PROGRAM_FILES           = 0x0026     // C:\Program Files 
var CSIDL_PROGRAMS                = 0x0002     // C:\Documents and Settings\username\Start Menu\Programs 
var CSIDL_RESOURCES               = 0x0038     // %windir%\Resources\, For theme and other windows resources. 
var CSIDL_STARTMENU               = 0x000b     // C:\Documents and Settings\username\Start Menu 
var CSIDL_STARTUP                 = 0x0007     // C:\Documents and Settings\username\Start Menu\Programs\Startup. 
var CSIDL_SYSTEM                  = 0x0025     // GetSystemDirectory() 
var CSIDL_WINDOWS                 = 0x0024     // GetWindowsDirectory() 

type DLL struct {
	*syscall.DLL
}


func (d *DLL) Proc(name string) *syscall.Proc {
	p := d.MustFindProc(name)
	return p
}

var gowincommonpath = syscall.NewLazyDLL("shell32.dll")


func StringToCharPtr(str string) *uint8 {
	chars := append([]byte(str), 0) // null terminated
	return &chars[0]
}

func UTF16PtrToString(p *uint16, max int) string {
if p == nil {
return ""
}
// Find NUL terminator.
end := unsafe.Pointer(p)
n := 0
for *(*uint16)(end) != 0 && n < max {
end = unsafe.Pointer(uintptr(end) + unsafe.Sizeof(*p))
n++
}
s := (*[(1 << 30) - 1]uint16)(unsafe.Pointer(p))[:n:n]
return string(utf16.Decode(s))
}
	

func GoWinCommonPathGet(csidl int) string {
	buf := ""
	bufC, err := syscall.UTF16PtrFromString(buf)
	if err != nil {
		fmt.Println(err)
	}
	gowincommonpath.NewProc("SHGetFolderPathW").Call(uintptr(0), uintptr(csidl), uintptr(0), uintptr(0), uintptr(unsafe.Pointer(bufC)))
	return UTF16PtrToString(bufC, 512)
}


func main() {
	fmt.Println(GoWinCommonPathGet(CSIDL_STARTUP))
}