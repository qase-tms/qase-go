.DEFAULT_GOAL = compile

compile:
	# TODO:dirty
	@maintainer go vanity build
	@mv docs/CNAME go.qase.io
	@rm -rf docs
	@mv go.qase.io docs
.PHONY: compile
