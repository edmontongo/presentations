( echo "digraph G {"
    go list -f '{{range .Imports}}{{printf "\t%q -> %q;\n" $.ImportPath .}}{{end}}' \
        		$(go list -f '{{join .Deps " "}}' time) time
    echo "}" ) | dot -Tsvg -o time-deps.svg

