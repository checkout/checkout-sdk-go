package common

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"testing"
)

const (
	modulePath       = "github.com/checkout/checkout-sdk-go/v3"
	commonImportPath = modulePath + "/common"
	httpImportPath   = "net/http"
)

// requestMethods maps each HttpClient method to the index of its path argument.
var requestMethods = map[string]int{
	"GetWithContext": 1, "PostWithContext": 1, "PutWithContext": 1, "PatchWithContext": 1,
	"DeleteWithContext": 1, "UploadWithContext": 1, "PostFormWithContext": 1,
	"Get": 0, "Post": 0, "Put": 0, "Patch": 0, "Delete": 0, "Upload": 0, "PostForm": 0,
}

// Files allowed to pass a request path through unchanged, each with its reason.
var requestPathPassThrough = map[string]string{
	"client/client.go":     "ApiClient methods without a context delegate the path they receive to the WithContext variant",
	"mocks/client_mock.go": "test double, records the path it receives and sends nothing",
}

// BuildPath arguments allowed besides literals, constants and encoded parameters.
var allowedBuildPathCalls = map[string]string{
	"apm/klarna/client.go c.getBaseUrl()": "returns one of two string literals selected by the environment",
}

// net/http functions and methods that create or send a request, and the only files allowed to use them.
var httpRequestFunctions = map[string]bool{
	"NewRequest": true, "NewRequestWithContext": true, "Get": true, "Post": true, "PostForm": true, "Head": true,
}

var allowedHttpRequestFiles = map[string]string{
	"client/client.go":                        "ApiClient joins its base URI with a path built by common.BuildPath",
	"configuration/oauth_keys_credentials.go": "token request to the configured authorization URI, no caller value in the path",
}

// net/http names that build or send requests outside the HttpClient; only the allowed files use them.
var httpRequestNames = map[string]bool{
	"Request": true, "Transport": true, "RoundTripper": true, "DefaultClient": true, "DefaultTransport": true,
}

// Methods that send a request when called on an http.Client, an http.Transport or the ApiClient
// internals; only the allowed files call them.
var sendingMethods = map[string]bool{"Do": true, "RoundTrip": true}
var httpClientMethods = map[string]bool{"Do": true, "Get": true, "Post": true, "PostForm": true, "Head": true}
var apiClientInternals = map[string]bool{
	"invoke": true, "submit": true, "submitForm": true, "buildRequest": true, "doRequest": true,
}

// Files that may create an ApiClient or set its base URI; each base URI must come from the
// configured environment.
var apiClientFactoryFiles = map[string]bool{"client/client.go": true, "nas/checkout_api.go": true, "abc/checkout_api.go": true}
var allowedBaseUri = regexp.MustCompile(`^configuration\.(Environment\.[A-Za-z]+Uri\(\)|EnvironmentSubdomain\.ApiUrl)$`)

// Packages that hold the http.Client configuration; nothing else may reach the http.Client.
var httpClientConfigDirs = map[string]bool{"client": true, "configuration": true, "nas": true, "abc": true, "common": true}

// Imports for raw connections, unsafe memory or reflection, and the files allowed to use them.
var restrictedImports = map[string]map[string]bool{
	"net": {}, "crypto/tls": {}, "unsafe": {}, "syscall": {},
	"reflect": {"client/client.go": true, "common/serializer.go": true},
}

var goosGoarchSuffix = regexp.MustCompile(`_(aix|android|darwin|dragonfly|freebsd|illumos|ios|js|linux|netbsd|openbsd|plan9|solaris|windows|386|amd64|arm|arm64|loong64|mips|mipsle|mips64|mips64le|ppc64|ppc64le|riscv64|s390x|wasm)(_test)?\.go$`)
var buildConstraint = regexp.MustCompile(`^//\s*(\+build|go:build)`)

// allowedBuildPathCalls entries whose function must return only string literals.
var literalOnlyFunctions = map[string]string{"apm/klarna/client.go": "getBaseUrl"}

// commonFunctions are the helpers whose use the guard follows; they may only be called directly.
var commonFunctions = map[string]bool{"BuildPath": true, "BuildQueryPath": true, "EncodePathSegment": true}

type sourceFile struct {
	rel     string
	dir     string
	file    *ast.File
	imports map[string]string // import name -> import path
	consts  map[*ast.ValueSpec]bool
}

type guard struct {
	fset     *token.FileSet
	root     string
	consts   map[string]map[string]bool // package dir -> package-level const names
	problems []string
	calls    int
}

// TestLibraryEncodesEveryCallerValueInRequestPaths reads the library source and fails on any
// request path built from a value that has not gone through EncodePathSegment.
func TestLibraryEncodesEveryCallerValueInRequestPaths(t *testing.T) {
	g, files := runGuard(t, "..")

	// Control: the scan reads the whole library, so a clean result means the calls were checked.
	// Every client.go of the library (found here by glob, apart from the walk) must have been read.
	parsed := map[string]bool{}
	for _, sf := range files {
		parsed[sf.rel] = true
	}
	clients := []string{"client/client.go"}
	for _, pattern := range []string{"../*/client.go", "../*/*/client.go"} {
		matches, err := filepath.Glob(pattern)
		if err != nil {
			t.Fatal(err)
		}
		for _, m := range matches {
			rel := filepath.ToSlash(strings.TrimPrefix(m, ".."+string(filepath.Separator)))
			if !strings.HasPrefix(rel, "common/testdata/") {
				clients = append(clients, rel)
			}
		}
	}
	if len(clients) < 40 {
		t.Fatalf("found only %d client.go files by glob", len(clients))
	}
	for _, rel := range clients {
		if !parsed[rel] {
			t.Fatalf("%s not read; the scan is not reading the library", rel)
		}
	}
	if g.calls == 0 {
		t.Fatal("no BuildPath call checked; the scan is not reading the library")
	}
	for _, p := range g.problems {
		t.Error(p)
	}
}

// TestGuardReportsEachRule runs the guard on the small source trees in testdata/guard: each one
// breaks one rule and must be reported with the expected message; "clean" must report nothing.
func TestGuardReportsEachRule(t *testing.T) {
	cases := map[string]string{
		"clean":                    "",
		"unencoded_value":          "BuildPath argument id is not a constant or a value encoded with common.EncodePathSegment",
		"encoded_after_use":        "BuildPath argument id is not a constant or a value encoded with common.EncodePathSegment",
		"path_not_built":           "is not built by common.BuildPath",
		"helper_as_value":          "EncodePathSegment used as a value",
		"dot_import":               "dot import of " + commonImportPath,
		"build_constraint":         "build constraint in a library file",
		"ignored_file":             "library file ignored by the go tool",
		"platform_file":            "platform-specific library file",
		"linkname":                 "go:linkname in a library file",
		"restricted_import":        "import of unsafe in a library file",
		"goto":                     "goto in a library file",
		"reflection_call":          "method called by reflection",
		"reflection_by_name":       "method looked up by name through reflection",
		"testdata_import":          "import of a testdata package in a library file",
		"http_client_field":        "HttpClient used outside the client configuration",
		"http_transport":           "net/http Transport used outside the HttpClient",
		"http_client_method":       "request sent with HttpClient.Get outside the HttpClient",
		"klarna_base_url_missing":  "getBaseUrl not found; update literalOnlyFunctions",
		"request_method_value":     "request method Get used as a value",
		"http_request_literal":     "net/http Request built by hand instead of the HttpClient",
		"base_uri_access":          "ApiClient base URI used outside client/client.go",
		"api_client_outside_entry": "ApiClient created outside the SDK entry points",
		"http_client":              "net/http Client used outside the client configuration",
		"http_new_request":         "request created with net/http NewRequest instead of the HttpClient",
		"http_do":                  "request sent with Do outside the HttpClient",
		"api_client_literal":       "ApiClient built by hand outside client/client.go",
		"api_client_base_uri":      "ApiClient base URI is not taken from the configured environment",
		"klarna_base_url":          "getBaseUrl must return only string literals",
		"api_client_internal_call": "ApiClient internal invoke called outside client/client.go",
	}
	for name, want := range cases {
		name, want := name, want
		t.Run(name, func(t *testing.T) {
			g, _ := runGuard(t, filepath.Join("testdata", "guard", name))
			if want == "" {
				if len(g.problems) != 0 {
					t.Errorf("expected no problems, got %v", g.problems)
				}
				return
			}
			for _, p := range g.problems {
				if strings.Contains(p, want) {
					return
				}
			}
			t.Errorf("expected a problem containing %q, got %v", want, g.problems)
		})
	}
}

func runGuard(t *testing.T, root string) (*guard, []sourceFile) {
	t.Helper()
	g := &guard{fset: token.NewFileSet()}
	files := g.parseLibrary(t, root)
	for _, sf := range files {
		g.checkFile(sf)
	}
	return g, files
}

func (g *guard) report(n ast.Node, format string, args ...interface{}) {
	p := g.fset.Position(n.Pos())
	rel, _ := filepath.Rel(g.root, p.Filename)
	g.problems = append(g.problems, fmt.Sprintf("%s:%d: %s", filepath.ToSlash(rel), p.Line, fmt.Sprintf(format, args...)))
}

func (g *guard) checkFile(sf sourceFile) {
	for name, importPath := range sf.imports {
		if name == "." && (importPath == commonImportPath || importPath == httpImportPath) {
			g.report(sf.file, "dot import of %s", importPath)
		}
	}
	// The scan reads every file, so a file that is only built under some constraint could hide a
	// declaration the compiler sees differently.
	for _, group := range sf.file.Comments {
		if group.Pos() > sf.file.Package {
			break
		}
		for _, c := range group.List {
			if buildConstraint.MatchString(c.Text) {
				g.report(c, "build constraint in a library file")
			}
		}
	}
	if base := filepath.Base(sf.rel); strings.HasPrefix(base, "_") || strings.HasPrefix(base, ".") {
		g.report(sf.file, "library file ignored by the go tool")
	}
	if goosGoarchSuffix.MatchString(sf.rel) {
		g.report(sf.file, "platform-specific library file")
	}
	for _, group := range sf.file.Comments {
		for _, c := range group.List {
			if strings.HasPrefix(c.Text, "//go:linkname") {
				g.report(c, "go:linkname in a library file")
			}
		}
	}
	for _, im := range sf.file.Imports {
		importPath, _ := strconv.Unquote(im.Path.Value)
		if allowed, restricted := restrictedImports[importPath]; restricted && !allowed[sf.rel] {
			g.report(im, "import of %s in a library file", importPath)
		}
		if importPath == "testdata" || strings.HasSuffix(importPath, "/testdata") || strings.Contains(importPath, "/testdata/") {
			g.report(im, "import of a testdata package in a library file")
		}
	}
	_, allowedHttp := allowedHttpRequestFiles[sf.rel]
	ast.Inspect(sf.file, func(n ast.Node) bool {
		switch e := n.(type) {
		case *ast.BranchStmt:
			if e.Tok == token.GOTO {
				g.report(e, "goto in a library file")
			}
		case *ast.SelectorExpr:
			if e.Sel.Name == "MethodByName" {
				g.report(e, "method looked up by name through reflection")
			}
			if _, usesReflect := sf.imports["reflect"]; usesReflect && (e.Sel.Name == "Call" || e.Sel.Name == "Method") {
				g.report(e, "method called by reflection")
			}
			_, _, qualified := sf.importedName(e)
			isClientInterface := qualified && sf.imports[exprString(e.X)] == modulePath+"/client" && e.Sel.Name == "HttpClient"
			if pkg, name, ok := sf.importedName(e); ok && pkg == httpImportPath && name == "Client" && !httpClientConfigDirs[sf.dir] {
				g.report(e, "net/http Client used outside the client configuration")
			}
			if (e.Sel.Name == "HttpClient" || e.Sel.Name == "BuildDefaultClient") && !isClientInterface && !httpClientConfigDirs[sf.dir] {
				g.report(e, "%s used outside the client configuration", e.Sel.Name)
			}
			if pkg, name, ok := sf.importedName(e); ok && pkg == httpImportPath && httpRequestNames[name] && !allowedHttp {
				g.report(e, "net/http %s used outside the HttpClient", name)
			}
		case *ast.CallExpr:
			sel, ok := e.Fun.(*ast.SelectorExpr)
			if !ok || allowedHttp {
				return true
			}
			if sendingMethods[sel.Sel.Name] {
				g.report(e, "request sent with %s outside the HttpClient", sel.Sel.Name)
			}
			if apiClientInternals[sel.Sel.Name] {
				g.report(e, "ApiClient internal %s called outside client/client.go", sel.Sel.Name)
			}
			if httpClientMethods[sel.Sel.Name] && strings.HasSuffix(exprString(sel.X), "HttpClient") {
				g.report(e, "request sent with HttpClient.%s outside the HttpClient", sel.Sel.Name)
			}
		}
		return true
	})
	if name, ok := literalOnlyFunctions[sf.rel]; ok {
		found := false
		for _, decl := range sf.file.Decls {
			fn, ok := decl.(*ast.FuncDecl)
			if !ok || fn.Name.Name != name || fn.Body == nil {
				continue
			}
			found = true
			ast.Inspect(fn.Body, func(n ast.Node) bool {
				if ret, ok := n.(*ast.ReturnStmt); ok {
					for _, r := range ret.Results {
						if lit, ok := r.(*ast.BasicLit); !ok || lit.Kind != token.STRING {
							g.report(r, "%s must return only string literals", name)
						}
					}
				}
				return true
			})
		}
		if !found {
			g.report(sf.file, "%s not found; update literalOnlyFunctions", name)
		}
	}

	// Every reference to the followed helpers and to the net/http request functions must be the
	// function of a call, so that no alias or function value escapes the checks below.
	callees := map[ast.Expr]bool{}
	ast.Inspect(sf.file, func(n ast.Node) bool {
		if call, ok := n.(*ast.CallExpr); ok {
			callees[call.Fun] = true
		}
		if fn, ok := n.(*ast.FuncDecl); ok {
			callees[fn.Name] = true // the declaration itself
		}
		return true
	})
	ast.Inspect(sf.file, func(n ast.Node) bool {
		switch e := n.(type) {
		case *ast.SelectorExpr, *ast.Ident:
			expr := e.(ast.Expr)
			if name, ok := sf.commonFunction(expr); ok && !callees[expr] {
				g.report(expr, "%s used as a value", name)
			}
			if sel, ok := expr.(*ast.SelectorExpr); ok && !callees[expr] {
				if _, isRequest := requestMethods[sel.Sel.Name]; isRequest {
					if _, allowed := requestPathPassThrough[sf.rel]; !allowed {
						g.report(expr, "request method %s used as a value", sel.Sel.Name)
					}
				}
			}
			if pkg, name, ok := sf.importedName(expr); ok && pkg == httpImportPath && httpRequestFunctions[name] {
				if _, allowed := allowedHttpRequestFiles[sf.rel]; !allowed {
					g.report(expr, "request created with net/http %s instead of the HttpClient", name)
				}
			}
		case *ast.CompositeLit:
			if pkg, name, ok := sf.importedName(e.Type); ok && pkg == httpImportPath && name == "Request" {
				if _, allowed := allowedHttpRequestFiles[sf.rel]; !allowed {
					g.report(e, "net/http Request built by hand instead of the HttpClient")
				}
			}
			if strings.HasSuffix(exprString(e.Type), "ApiClient") && sf.rel != "client/client.go" {
				g.report(e, "ApiClient built by hand outside client/client.go")
			}
		}
		if sel, ok := n.(*ast.SelectorExpr); ok && sel.Sel.Name == "BaseUri" && !callees[sel] && sf.rel != "client/client.go" {
			g.report(sel, "ApiClient base URI used outside client/client.go")
		}
		if call, ok := n.(*ast.CallExpr); ok {
			if sel, ok := call.Fun.(*ast.SelectorExpr); ok && sel.Sel.Name == "NewApiClient" {
				if !apiClientFactoryFiles[sf.rel] {
					g.report(call, "ApiClient created outside the SDK entry points")
				} else if len(call.Args) != 2 || !allowedBaseUri.MatchString(exprString(call.Args[1])) {
					g.report(call, "ApiClient base URI is not taken from the configured environment")
				}
			}
		}
		return true
	})

	for _, decl := range sf.file.Decls {
		fn, _ := decl.(*ast.FuncDecl)
		if fn != nil && fn.Body == nil {
			continue
		}
		g.checkDecl(sf, decl, fn)
	}
}

// checkDecl checks every BuildPath call and every HttpClient request inside one declaration. fn
// is nil for declarations that are not functions (package-level variables holding a func literal).
func (g *guard) checkDecl(sf sourceFile, decl ast.Decl, fn *ast.FuncDecl) {
	encoded := map[*ast.Object]token.Pos{}
	var params map[*ast.Object]bool
	if fn != nil {
		params = parameterObjects(fn)
		encoded = encodedParameters(sf, fn, params)
	}

	ast.Inspect(decl, func(n ast.Node) bool {
		call, ok := n.(*ast.CallExpr)
		if !ok {
			return true
		}
		if name, ok := sf.commonFunction(call.Fun); ok && name == "BuildPath" {
			g.calls++
			for _, arg := range call.Args {
				g.checkBuildPathArgument(sf, fn, arg, params, encoded)
			}
			return true
		}
		sel, ok := call.Fun.(*ast.SelectorExpr)
		if !ok {
			return true
		}
		index, ok := requestMethods[sel.Sel.Name]
		if !ok || len(call.Args) <= index+1 {
			return true
		}
		if _, ok := requestPathPassThrough[sf.rel]; ok {
			return true
		}
		pathArg := call.Args[index]
		if sf.isBuiltPath(pathArg) {
			return true
		}
		if id, ok := pathArg.(*ast.Ident); ok && fn != nil && id.Obj != nil && !params[id.Obj] {
			if rhs, ok := singleAssignment(fn, id.Obj); ok && sf.isBuiltPath(rhs) {
				return true
			}
		}
		g.report(pathArg, "request path %s is not built by common.BuildPath", exprString(pathArg))
		return true
	})
}

func (g *guard) checkBuildPathArgument(sf sourceFile, fn *ast.FuncDecl, arg ast.Expr, params map[*ast.Object]bool, encoded map[*ast.Object]token.Pos) {
	switch e := arg.(type) {
	case *ast.BasicLit:
		if e.Kind == token.STRING {
			return
		}
	case *ast.Ident:
		if e.Obj == nil {
			// Not declared in this file: a package-level name from another file of the package.
			if g.consts[sf.dir][e.Name] {
				return
			}
		} else if spec, ok := e.Obj.Decl.(*ast.ValueSpec); ok && e.Obj.Kind == ast.Con && sf.consts[spec] {
			return
		} else if params[e.Obj] {
			if pos, ok := encoded[e.Obj]; ok && pos < e.Pos() {
				return
			}
		}
	case *ast.SelectorExpr:
		if pkg, name, ok := sf.importedName(e); ok && strings.HasPrefix(pkg, modulePath+"/") &&
			g.consts[strings.TrimPrefix(pkg, modulePath+"/")][name] {
			return
		}
	case *ast.CallExpr:
		if _, ok := allowedBuildPathCalls[sf.rel+" "+exprString(e)]; ok && isReceiverCall(e, fn) {
			return
		}
	}
	g.report(arg, "BuildPath argument %s is not a constant or a value encoded with common.EncodePathSegment", exprString(arg))
}

// isReceiverCall reports whether call is recv.method() on the receiver of fn, not on a variable
// that shadows it.
func isReceiverCall(call *ast.CallExpr, fn *ast.FuncDecl) bool {
	sel, ok := call.Fun.(*ast.SelectorExpr)
	if !ok || fn == nil || fn.Recv == nil || len(fn.Recv.List) != 1 || len(fn.Recv.List[0].Names) != 1 {
		return false
	}
	x, ok := sel.X.(*ast.Ident)
	return ok && x.Obj != nil && x.Obj == fn.Recv.List[0].Names[0].Obj
}

// isBuiltPath reports whether e is common.BuildPath(...) or common.BuildQueryPath(common.BuildPath(...), q).
func (sf sourceFile) isBuiltPath(e ast.Expr) bool {
	call, ok := e.(*ast.CallExpr)
	if !ok {
		return false
	}
	name, ok := sf.commonFunction(call.Fun)
	if !ok {
		return false
	}
	if name == "BuildPath" {
		return true
	}
	if name == "BuildQueryPath" && len(call.Args) == 2 {
		inner, ok := call.Args[0].(*ast.CallExpr)
		if !ok {
			return false
		}
		innerName, ok := sf.commonFunction(inner.Fun)
		return ok && innerName == "BuildPath"
	}
	return false
}

// commonFunction resolves e to one of the followed helpers of the common package.
func (sf sourceFile) commonFunction(e ast.Expr) (string, bool) {
	if pkg, name, ok := sf.importedName(e); ok && pkg == commonImportPath && commonFunctions[name] {
		return name, true
	}
	if id, ok := e.(*ast.Ident); ok && sf.dir == "common" && commonFunctions[id.Name] &&
		(id.Obj == nil || id.Obj.Kind == ast.Fun) {
		return id.Name, true
	}
	return "", false
}

// importedName resolves pkg.Name, where pkg is an import of the file, to the import path and Name.
func (sf sourceFile) importedName(e ast.Expr) (string, string, bool) {
	sel, ok := e.(*ast.SelectorExpr)
	if !ok {
		return "", "", false
	}
	x, ok := sel.X.(*ast.Ident)
	if !ok || x.Obj != nil {
		return "", "", false
	}
	importPath, ok := sf.imports[x.Name]
	return importPath, sel.Sel.Name, ok
}

func parameterObjects(fn *ast.FuncDecl) map[*ast.Object]bool {
	params := map[*ast.Object]bool{}
	for _, field := range fn.Type.Params.List {
		for _, n := range field.Names {
			if n.Obj != nil {
				params[n.Obj] = true
			}
		}
	}
	return params
}

// encodedParameters returns, for each parameter that the function re-assigns in a top-level
// statement as `p, err := common.EncodePathSegment(p)` (or `=`) followed by
// `if err != nil { return ..., err }`, the position after which p holds the encoded value. A
// parameter written anywhere else in the function (assignment, range, increment, address taken)
// is left out.
func encodedParameters(sf sourceFile, fn *ast.FuncDecl, params map[*ast.Object]bool) map[*ast.Object]token.Pos {
	encoded := map[*ast.Object]token.Pos{}
	stmts := fn.Body.List
	for i := 0; i+1 < len(stmts); i++ {
		as, ok := stmts[i].(*ast.AssignStmt)
		if !ok || len(as.Lhs) != 2 || len(as.Rhs) != 1 {
			continue
		}
		target, ok1 := as.Lhs[0].(*ast.Ident)
		errIdent, ok2 := as.Lhs[1].(*ast.Ident)
		call, ok3 := as.Rhs[0].(*ast.CallExpr)
		if !ok1 || !ok2 || !ok3 || errIdent.Name != "err" || target.Obj == nil || !params[target.Obj] || len(call.Args) != 1 {
			continue
		}
		name, ok := sf.commonFunction(call.Fun)
		argIdent, ok4 := call.Args[0].(*ast.Ident)
		if !ok || name != "EncodePathSegment" || !ok4 || argIdent.Obj != target.Obj {
			continue
		}
		if !isErrReturn(stmts[i+1]) {
			continue
		}
		encoded[target.Obj] = stmts[i+1].End()
	}
	for obj := range encoded {
		if writes(fn, obj) != 1 {
			delete(encoded, obj)
		}
	}
	return encoded
}

// writes counts the statements that may change obj: assignments, range clauses, increments and
// taking its address (which counts twice, as the callee may write through it).
func writes(fn *ast.FuncDecl, obj *ast.Object) int {
	count := 0
	is := func(e ast.Expr) bool {
		id, ok := e.(*ast.Ident)
		return ok && id.Obj == obj
	}
	ast.Inspect(fn.Body, func(n ast.Node) bool {
		switch s := n.(type) {
		case *ast.AssignStmt:
			for _, lhs := range s.Lhs {
				if is(lhs) {
					count++
				}
			}
		case *ast.RangeStmt:
			if (s.Key != nil && is(s.Key)) || (s.Value != nil && is(s.Value)) {
				count++
			}
		case *ast.IncDecStmt:
			if is(s.X) {
				count++
			}
		case *ast.UnaryExpr:
			if s.Op == token.AND && is(s.X) {
				count += 2
			}
		}
		return true
	})
	return count
}

// singleAssignment returns the right-hand side when obj is a local variable assigned exactly once.
func singleAssignment(fn *ast.FuncDecl, obj *ast.Object) (ast.Expr, bool) {
	as, ok := obj.Decl.(*ast.AssignStmt)
	if !ok || len(as.Rhs) != 1 || writes(fn, obj) != 1 {
		return nil, false
	}
	return as.Rhs[0], true
}

func isErrReturn(stmt ast.Stmt) bool {
	ifs, ok := stmt.(*ast.IfStmt)
	if !ok || ifs.Init != nil || ifs.Else != nil || len(ifs.Body.List) != 1 {
		return false
	}
	cond, ok := ifs.Cond.(*ast.BinaryExpr)
	if !ok || cond.Op != token.NEQ || exprString(cond.X) != "err" || exprString(cond.Y) != "nil" {
		return false
	}
	ret, ok := ifs.Body.List[0].(*ast.ReturnStmt)
	if !ok || len(ret.Results) == 0 {
		return false
	}
	return exprString(ret.Results[len(ret.Results)-1]) == "err"
}

func (g *guard) parseLibrary(t *testing.T, root string) []sourceFile {
	absRoot, err := filepath.Abs(root)
	if err != nil {
		t.Fatal(err)
	}
	g.root = absRoot
	g.consts = map[string]map[string]bool{}
	var files []sourceFile
	err = filepath.Walk(absRoot, func(p string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			if p != absRoot && (strings.HasPrefix(info.Name(), ".") || info.Name() == "vendor" || p == filepath.Join(absRoot, "common", "testdata")) {
				return filepath.SkipDir
			}
			return nil
		}
		if !strings.HasSuffix(p, ".go") || strings.HasSuffix(p, "_test.go") {
			return nil
		}
		f, err := parser.ParseFile(g.fset, p, nil, parser.ParseComments)
		if err != nil {
			return err
		}
		rel, _ := filepath.Rel(absRoot, p)
		sf := sourceFile{
			rel:     filepath.ToSlash(rel),
			dir:     filepath.ToSlash(filepath.Dir(rel)),
			file:    f,
			imports: map[string]string{},
			consts:  map[*ast.ValueSpec]bool{},
		}
		for _, im := range f.Imports {
			importPath, _ := strconv.Unquote(im.Path.Value)
			name := importPath[strings.LastIndex(importPath, "/")+1:]
			if im.Name != nil {
				name = im.Name.Name
			}
			sf.imports[name] = importPath
		}
		if g.consts[sf.dir] == nil {
			g.consts[sf.dir] = map[string]bool{}
		}
		for _, d := range f.Decls {
			if gd, ok := d.(*ast.GenDecl); ok && gd.Tok == token.CONST {
				for _, s := range gd.Specs {
					spec := s.(*ast.ValueSpec)
					sf.consts[spec] = true
					for _, n := range spec.Names {
						g.consts[sf.dir][n.Name] = true
					}
				}
			}
		}
		files = append(files, sf)
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
	return files
}

func exprString(e ast.Expr) string {
	switch x := e.(type) {
	case *ast.Ident:
		return x.Name
	case *ast.BasicLit:
		return x.Value
	case *ast.SelectorExpr:
		return exprString(x.X) + "." + x.Sel.Name
	case *ast.CallExpr:
		args := make([]string, len(x.Args))
		for i, a := range x.Args {
			args[i] = exprString(a)
		}
		return exprString(x.Fun) + "(" + strings.Join(args, ", ") + ")"
	case *ast.StarExpr:
		return "*" + exprString(x.X)
	case *ast.UnaryExpr:
		return x.Op.String() + exprString(x.X)
	case *ast.BinaryExpr:
		return exprString(x.X) + " " + x.Op.String() + " " + exprString(x.Y)
	}
	return fmt.Sprintf("%T", e)
}
