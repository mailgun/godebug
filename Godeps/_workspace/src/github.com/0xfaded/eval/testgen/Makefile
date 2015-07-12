.PHONY: *gen.go clean

*gen.go:
	$(eval OUT:=$(patsubst %.go,%,$@))
	go build -o $(OUT) main.go common.go $@
	./$(OUT) > ../$(patsubst %.go,%_test.go,$@)

default: %gen.go

clean:
	rm *_gen
