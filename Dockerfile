FROM scratch
ENV TERM=xterm-256color
ENTRYPOINT ["/minesweeper-cli"]
COPY minesweeper-cli /
