FROM alpine:3.4

WORKDIR /app
ADD builds/minesweeper_api /app
EXPOSE 8080
CMD /app/minesweeper_api
