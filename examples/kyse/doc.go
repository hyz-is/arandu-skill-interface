// Package kyse holds every Kyse snippet the arandu-skill-interface skills show, as compiled
// code.
//
// A snippet in a SKILL.md or a reference is preceded by
//
//	<!-- example: examples/kyse/<file>.go#<Function> -->
//
// and the test in internal/skillcheck fails when the fenced block differs from
// the function declared here. The compiler checks the props; the test in this
// package renders each one and checks the markup carries nothing the CSP would
// refuse.
//
// In a .kyse.go view the same expression is written inside {!! … !!}: a view
// is Go, so what compiles here is what compiles there.
package kyse
