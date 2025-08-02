module github.com/vilar95/go-packages/packaging/mod-replace-new/system

go 1.24.3

require github.com/google/uuid v1.6.0

replace github.com/vilar95/go-packages/packaging/mod-replace/math => ../math
