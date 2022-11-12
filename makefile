.PHONY: runserver
runserver: 
	./atux-server/bin/main

.PHONY: buildserver 
buildserver: 
	./scripts/build_server.sh

.PHONY: clean
clean:
	./scripts/clean.sh

.PHONY: runembed
runembed: 
	./atux-embed/bin/main

.PHONY: buildembed
buildembed: 
	./scripts/build_embed.sh


