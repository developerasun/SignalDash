# generate swagger doc, should run at root to include all deps
swag init -g ./main.go -o ./docs --parseDependency --parseInternal