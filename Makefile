.DEFAULT_GOAL = compile

compile:
	@maintainer go vanity build --host=go.qase.io ./docs
.PHONY: compile
