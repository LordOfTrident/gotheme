NAME = gotheme
CMD  = cmd/$(NAME)

build:
	go build ./$(CMD)

run:
	go run ./$(CMD)

install: $(NAME)
	cp $(NAME) /usr/bin/$(NAME)

uninstall: /usr/bin/$(NAME)
	rm /usr/bin/$(NAME)

clean: $(NAME)
	rm $(NAME)

all:
	@echo build, run, install, uninstall, clean
