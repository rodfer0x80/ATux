.PHONY: runserver
runserver: 
	./server/bin/main

.PHONY: buildserver 
buildserver: 
	./scripts/build_server.sh

.PHONY: clean
clean:
	./scripts/clean

.PHONY: runembed
runembed: 
	./server/bin/main

.PHONY: buildembed
buildembed: 
	./scripts/build_embed.sh


