build:
	cd cmd && go build -o EmbedDiscordVideo

run:
	./cmd/EmbedDiscordVideo

dev:
	cd cmd && go build -o EmbedDiscordVideo
	./cmd/EmbedDiscordVideo
