migration_gen:
	@if [ -z "$(MIGRATION_NAME)" ]; then \
		echo "MIGRATION_NAME is required"; \
		exit 1; \
	fi
	atlas migrate diff $(shell date +%Y%m%d%H%M%S)_$(MIGRATION_NAME) \
	--dir "file://ent/migrate/migrations" \
	--to "ent://ent/schema" \
	--dev-url "docker://postgres/15/test?search_path=public"

migration_apply:
	@if [ -z "$(PASSWORD)" ]; then \
		echo "PASSWORD is required"; \
		exit 1; \
	fi
	atlas migrate apply \
		--dir "file://ent/migrate/migrations" \
		--url "postgres://postgres.azrhiuujplcjlpyfccqh:$(PASSWORD)@aws-0-us-east-1.pooler.supabase.com:5432/postgres" \
		--allow-dirty \
		--revisions-schema public

generate_ent:
	go generate ./ent
		