# A Self-Documenting Makefile: http://marmelab.com/blog/2016/02/29/auto-documented-makefile.html

.PHONY: update
.DEFAULT_GOAL := help

update: ## Update site
	rm -rf favicon.ico
	rm -rf index.html
	rm -rf inline.*
	rm -rf main.*
	rm -rf styles.*
	rm -rf vendor.*
	cp dist/* .

help:
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
