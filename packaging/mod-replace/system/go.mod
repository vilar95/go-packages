module github.com/vilar95/go-packages/packaging/mod-replace/system

go 1.24.3

replace github.com/vilar95/go-packages/packaging/mod-replace/math => ../math

require github.com/vilar95/go-packages/packaging/mod-replace/math v0.0.0-00010101000000-000000000000
