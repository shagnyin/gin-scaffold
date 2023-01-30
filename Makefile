.PHONY:model
model:
	@sqlc generate
	@echo "model create success ~~~"

.PHONY:docs
docs:
	@swag init