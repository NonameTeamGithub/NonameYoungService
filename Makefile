all: update

update:
	git add *
	git rm --cached config/config.yaml
	git commit -m "upd"
	git push
