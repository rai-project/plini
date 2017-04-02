package language

import (
	"github.com/rai-project/linguist"
)

type Language string 
const (
	CPPLanguage = "C++"
	CUDALanguage = "CUDA"
	OpenCLLanguage = "OpenCL"
	PythonLanguage = "Python"
	ShellLanguage = "Shell"
)

func DetectLanguage(prog string) (Language, error) {
	lang := linguist.Detect(prog)
	if lang == "" {
		return "", errors.New("unable to determine language")
	}
	return Language(lang), nil
}