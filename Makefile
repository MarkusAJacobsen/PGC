app-name=guarded-island-59755

all:
	build release open

build:
	heroku container:push web -a ${app-name}

release:
	heroku container:release web -a ${app-name}

open:
	heroku open -a ${app-name}

logs:
	heroku logs --tail -a ${app-name}

shell:
	heroku run bash -a ${app-name}
