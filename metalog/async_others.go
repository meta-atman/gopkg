//go:build !(linux && (arm64 || amd64 || mips64 || mips64le || ppc64 || ppc64le || riscv64 || s390x || loong64))

package metalog

func (w *AsyncWriter) writever() {
	panic("not_implemented")
}
