templ:
    templ generate --watch

tailwindcss:
    tailwindcss -i web/assets/input.css -o web/assets/dist/style.css --watch

rustywind:
    fd -e html -e templ | entr -r rustywind --write .

bun:
    fd -e ts | entr -r bun build  web/assets/main.ts --minify --outdir web/assets/dist/

dev:
    /mnt/c/Program\ Files/Google/Chrome/Application/chrome.exe --app="http://localhost:8080" --user-data-dir='C:\Users\raikr\AppData\Local\Temp\'

watch:
    fd . | entr -r sh -c 'clear && GOOS=js GOARCH=wasm go build -tags client -o web/assets/dist/main.wasm internal/functions/agrows_client_functions.go && (unset GOOS GOARCH && go build -tags server -o server cmd/main.go && ./server -R)'

generate:
    agrows -i internal/functions/functions.go client && agrows -i internal/functions/functions.go server
