FROM scratch

ENV PORT 8000
EXPOSE $PORT

COPY rest-api /
CMD ["/rest-api"]
