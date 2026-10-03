services-start:
	@for dir in */; do \
		(cd "$$dir" && make app-restart) & \
	done; \
	wait