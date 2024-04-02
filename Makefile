SHELL := /bin/bash
DASHBOARD_DIR := dashboards
DASHBOARD_LINTER := $(shell command -v dashboard-linter 2> /dev/null)

lint:
ifndef DASHBOARD_LINTER
	$(error "dashboard-linter is not available, please install it via 'go get github.com/grafana/dashboard-linter'")
endif
	@echo "Linting the dashboard files..."
	@passed=0; \
	failed=0; \
	while IFS= read -r -d '' file; do \
		if $(DASHBOARD_LINTER) lint $$file --verbose --strict > /dev/null 2>&1; then \
			printf "\e[32m✓ %s\e[0m\n" "$$file"; \
			passed=$$((passed+1)); \
		else \
			printf "\e[31m✗ %s\e[0m\n" "$$file"; \
			failed=$$((failed+1)); \
			$(DASHBOARD_LINTER) lint $$file --verbose --strict 2>&1; \
		fi; \
	done < <(find $(DASHBOARD_DIR) -name '*.json' -print0); \
	if [ $$failed -eq 0 ]; then \
		printf "\e[32mSummary: $$passed passed, $$failed failed.\e[0m\n"; \
	else \
		printf "\e[31mSummary: $$passed passed, $$failed failed.\e[0m\n"; \
	fi