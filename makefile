.PHONY: runserver
runserver: 
	./atux_server/bin/main

.PHONY: buildserver 
buildserver: 
	./scripts/build_server.sh

.PHONY: clean
clean:
	./scripts/clean.sh

.PHONY: runembed
runembed: 
	./atux_embed/bin/main

.PHONY: buildembed
buildembed: 
	./scripts/build_embed.sh


