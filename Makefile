app-name=guarded-island-59755

all:
	build release open

build:
	GOOS=linux go build -tags seabolt_static -a -installsuffix -o pgc .

buildh:
	heroku container:push web -a ${app-name}

release:
	heroku container:release web -a ${app-name}

open:
	heroku open -a ${app-name}

shell:
	heroku run bash -a ${app-name}
