app-start:
	@for dir in */; do \
		(cd "$$dir" && make services-start) & \
	done; \
	wait