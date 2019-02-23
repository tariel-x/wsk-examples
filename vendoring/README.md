# Vendoring

1. Get dependencies: `cd src && go mod vendor`.
2. Create `exec` from `main.go`: `mv main.go exec`.
3. Move dependencies to current folder: `mv vendor/* && rm -rf vendor`.
4. Make archive with action: `zip -r vendoring.zip *`.
5. Build `wsk -i action create vendoring vendoring.zip --web true --kind go:default`.
6. Invoke action `wsk -i action invoke vendoring -r -p name vasya`.
