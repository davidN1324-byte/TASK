# Makefile

# Image name (local)
IMAGE_NAME := multiplatform-app
VERSION := latest

# Local tags for different platforms
IMAGE_TAG_LINUX := $(IMAGE_NAME):$(VERSION)-linux
IMAGE_TAG_ARM := $(IMAGE_NAME):$(VERSION)-arm
IMAGE_TAG_MAC := $(IMAGE_NAME):$(VERSION)-macos
IMAGE_TAG_WIN := $(IMAGE_NAME):$(VERSION)-windows

# Platforms
PLATFORMS := linux/amd64 linux/arm64 windows/amd64 darwin/arm64

.PHONY: all linux arm macos windows test clean

# Main command: build for all platforms
all: linux arm macos windows

# OS-specific build
linux:
	docker buildx build --platform linux/amd64 -t $(IMAGE_TAG_LINUX) --load .

arm:
	docker buildx build --platform linux/arm64 -t $(IMAGE_TAG_ARM) --load .

macos:
	docker buildx build --platform darwin/arm64 -t $(IMAGE_TAG_MAC) --load .

windows:
	docker buildx build --platform windows/amd64 -t $(IMAGE_TAG_WIN) --load .

# Running Windows tests locally (Git Bash / CMD)
test-windows:
	@echo "Running Windows tests..."
	@cmd //c tests\test_windows.bat

# Running tests in a container (local Linux image only)
test:
	docker run --rm $(IMAGE_TAG_LINUX)

# Cleaning local images
clean:
	docker rmi -f $(IMAGE_TAG_LINUX) $(IMAGE_TAG_ARM) $(IMAGE_TAG_MAC) $(IMAGE_TAG_WIN) || true
