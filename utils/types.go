package utils

type CompilerOptions struct {
	Target          string `json:"target"`
	Module          string `json:"module"`
	OutDir          string `json:"outDir"`
	RootDir         string `json:"rootDir"`
	Strict          bool   `json:"strict"`
	EsModuleInterop bool   `json:"esModuleInterop"`
	SkipLibCheck    bool   `json:"skipLibCheck"`
}

type TSConfig struct {
	CompilerOptions CompilerOptions `json:"compilerOptions"`
	Include         []string        `json:"include"`
	Exclude         []string        `json:"exclude"`
}
