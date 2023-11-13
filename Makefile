
SHELL=/bin/bash

hub_update:
	@hub_ctrl ${HUB_MODE} ln "$(realpath ./src/watchcat.py)" "${HOME}/.local/bin/watchcat"
