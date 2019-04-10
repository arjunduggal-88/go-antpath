package main

import (
	"C"
	. "github.com/vibrantbyte/go-antpath/antpath"
	"unsafe"
)

//pMatcher
var pMatcher PathMatcher

//init
func init(){
	pMatcher = New()
}

//export Version
func Version() *C.char{
	var cmsg = C.CString("v1.0")
	// 手动释放C分配的存储
	defer C.free(unsafe.Pointer(cmsg))
	return cmsg
}

//export Increment
func Increment(value *int) {
	*value = *value + 1
}

//export IsPattern
func IsPattern(path string) bool {
	return pMatcher.IsPattern(path)
}

//export Match
func Match(pattern,path string) bool{
	return pMatcher.Match(pattern,path)
}

//export MatchStart
func MatchStart(pattern,path string) bool{
	return pMatcher.MatchStart(pattern,path)
}

//export ExtractPathWithinPattern
func ExtractPathWithinPattern(pattern,path string) *C.char {
	result := C.CString(pMatcher.ExtractPathWithinPattern(pattern,path))
	defer C.free(unsafe.Pointer(result))
	return result
}

//export ExtractUriTemplateVariables
func ExtractUriTemplateVariables(pattern,path string) *map[string]string {
	return pMatcher.ExtractUriTemplateVariables(pattern,path)
}

////export GetPatternComparator
//func GetPatternComparator(path string) *AntPatternComparator {
//	return pMatcher.GetPatternComparator(path)
//}

//export Combine
func Combine(pattern1,pattern2 string) *C.char {
	result := C.CString(pMatcher.Combine(pattern1,pattern2))
	defer C.free(unsafe.Pointer(result))
	return result
}

//export SetPathSeparator
func SetPathSeparator(pathSeparator string){
	 pMatcher.SetPathSeparator(pathSeparator)
}

//export SetCaseSensitive
func SetCaseSensitive(caseSensitive bool){
	pMatcher.SetCaseSensitive(caseSensitive)
}

//export SetTrimTokens
func SetTrimTokens(trimTokens bool){
	pMatcher.SetTrimTokens(trimTokens)
}

//export SetCachePatterns
func SetCachePatterns(cachePatterns bool){
	pMatcher.SetCachePatterns(cachePatterns)
}

//main
func main(){}