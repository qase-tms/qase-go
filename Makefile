.DEFAULT_GOAL = compile

compile:
	@maintainer go vanity build --host=go.qase.io ./docs
.PHONY: compile

clean:
	@ls client/*.go | xargs -n1 sed -i '' -e 1,9d $$1
.PHONY: clean
