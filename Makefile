all: updateGit

updateGit:
	git add *
	git rm --cached config/config.yaml
	git commit -m "update"
	git push

